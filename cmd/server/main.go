package main

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/repository"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/littlebugger/tinode4chat/internal/service/handlers"
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

	// Set Up UseCases
	userUC := usecase.NewUserUseCase(repo)
	chatroomUC := usecase.NewChatRoomUseCase(repo)
	messageUC := usecase.NewMessageUseCase(repo)

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
