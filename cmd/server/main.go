package main

import (
	"context"
	worker "github.com/littlebugger/tinode4chat/internal/tinode"
	tinode "github.com/littlebugger/tinode4chat/pkg/tinode"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/littlebugger/tinode4chat/internal/service/handlers"
	"github.com/littlebugger/tinode4chat/internal/service/repository"
	"github.com/littlebugger/tinode4chat/internal/service/usecase"
	chatroom "github.com/littlebugger/tinode4chat/pkg/server/chatroom"
	message "github.com/littlebugger/tinode4chat/pkg/server/message"
	user "github.com/littlebugger/tinode4chat/pkg/server/user"
)

func main() {
	ctx := context.Background()

	// Get the MongoDB URL from environment variables
	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_URL environment variable is required")
	}

	dbName := "chat_app"

	// Create a MongoRepository instance by connecting to MongoDB
	repo, err := repository.ConnectToMongo(ctx, mongoURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ensure MongoDB connection is closed on exit
	defer func() {
		if err := repo.CloseMongoConnection(ctx); err != nil {
			log.Fatalf("Error closing MongoDB connection: %v", err)
		}
	}()

	// Get the Tinode server URL from environment variables or configuration
	tinodeURL := os.Getenv("TINODE_URL")
	if tinodeURL == "" {
		tinodeURL = "tinode:6060" // Default to Docker Compose service name
	}

	// Initialize the Tinode client
	tinodeClient := tinode.NewTinodeClient(tinodeURL)

	// Establish the WebSocket connection
	if err := tinodeClient.Connect(); err != nil {
		log.Fatalf("Failed to connect to Tinode server: %v", err)
	}
	defer tinodeClient.Close()

	// Set Up UseCases
	userUC := usecase.NewUserUseCase(repo, tinodeClient)
	chatroomUC := usecase.NewChatRoomUseCase(repo, tinodeClient)
	messageUC := usecase.NewMessageUseCase(repo, tinodeClient)

	// Create the worker
	tinodeWorker := worker.NewTinodeWorker(tinodeClient)

	// Start the worker
	tinodeWorker.Start()

	// Handle events from the worker
	go eventLoop(tinodeWorker, userUC, chatroomUC, messageUC)

	// Wait for interrupt signal to gracefully shutdown the worker
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down worker...")
	tinodeWorker.Stop()
	log.Println("Worker stopped.")

	// Login as admin
	adminEmail := os.Getenv("TINODE_ADMIN_EMAIL")
	adminPassword := os.Getenv("TINODE_ADMIN_PASSWORD")

	if err := tinodeClient.Login(adminEmail, adminPassword); err != nil {
		log.Fatalf("Failed to login to Tinode as admin: %v", err)
	}

	// Set up Echo
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize handlers and register routes
	userHandler := handlers.NewUserHandler(userUC)
	user.RegisterHandlers(e, userHandler) // This binds the generated routes with our handlers

	chatroomHandler := handlers.NewChatRoomHandler(chatroomUC)
	chatroom.RegisterHandlers(e, chatroomHandler)

	messageHandler := handlers.NewMessageHandler(messageUC)
	message.RegisterHandlers(e, messageHandler)

	// Start the HTTP server
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func eventLoop(
	tinodeWorker *worker.TinodeWorker,
	_ *usecase.UserService,
	roomUC *usecase.ChatRoomService,
	messageUC *usecase.MessageService,
) {
	for event := range tinodeWorker.Events() {
		ctx := context.Background()
		switch event.Type {
		case "data":
			data := event.Payload.(map[string]interface{})
			err := messageUC.HandleDataEvent(ctx, data)
			if err != nil {
				log.Printf("Error handling data event: %v", err)
			}
		case "meta":
			meta := event.Payload.(map[string]interface{})
			err := roomUC.HandleMetaEvent(ctx, meta)
			if err != nil {
				log.Printf("Error handling meta event: %v", err)
			}
		case "pres":
			// Here should be user meta events but i do not have support for them in db
		default:
			log.Printf("Unknown event type: %s", event.Type)
		}
	}
}
