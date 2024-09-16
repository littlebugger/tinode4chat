package worker

import (
	"fmt"
	"log"
	"time"

	tinodeclient "github.com/littlebugger/tinode4chat/pkg/tinode"
)

type TinodeWorker struct {
	client    *tinodeclient.TinodeClient
	stopChan  chan struct{}
	eventChan chan TinodeEvent
}

type TinodeEvent struct {
	Type    string
	Payload interface{}
}

// NewTinodeWorker creates a new TinodeWorker instance.
func NewTinodeWorker(client *tinodeclient.TinodeClient) *TinodeWorker {
	return &TinodeWorker{
		client:    client,
		stopChan:  make(chan struct{}),
		eventChan: make(chan TinodeEvent),
	}
}

// Start runs the worker in a separate goroutine.
func (w *TinodeWorker) Start() {
	go w.run()
}

// Stop signals the worker to stop.
func (w *TinodeWorker) Stop() {
	close(w.stopChan)
}

// Events returns the channel for receiving events.
func (w *TinodeWorker) Events() <-chan TinodeEvent {
	return w.eventChan
}

func (w *TinodeWorker) run() {
	defer w.client.Close()

	for {
		// Establish connection
		if err := w.client.Connect(); err != nil {
			log.Printf("Worker failed to connect to Tinode server: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Listen for messages
		if err := w.listenForMessages(); err != nil {
			log.Printf("Worker encountered error: %v", err)
			// Reconnect after a delay
			time.Sleep(5 * time.Second)
			continue
		}

		select {
		case <-w.stopChan:
			log.Println("Worker stopping...")
			return
		default:
			// Continue listening
		}
	}
}

func (w *TinodeWorker) listenForMessages() error {
	for {
		select {
		case <-w.stopChan:
			return nil
		default:
			msg, err := w.client.ReadMessage()
			if err != nil {
				return fmt.Errorf("error reading message: %w", err)
			}
			w.handleMessage(msg)
		}
	}
}

func (w *TinodeWorker) handleMessage(msg map[string]interface{}) {
	if meta, ok := msg["meta"].(map[string]interface{}); ok {
		// Handle meta messages
		w.handleMeta(meta)
	} else if data, ok := msg["data"].(map[string]interface{}); ok {
		// Handle data messages
		w.handleData(data)
	} else if pres, ok := msg["pres"].(map[string]interface{}); ok {
		// Handle presence messages
		w.handlePres(pres)
	} else {
		log.Printf("Unhandled message type: %v", msg)
	}
}

func (w *TinodeWorker) handleMeta(meta map[string]interface{}) {
	// Example: Handle new topics or subscriptions
	if contacts, ok := meta["sub"].([]interface{}); ok {
		for _, c := range contacts {
			contact := c.(map[string]interface{})
			topic := contact["topic"].(string)
			// Process the topic or subscription
			log.Printf("New subscription to topic: %s", topic)
			// Send event
			w.eventChan <- TinodeEvent{
				Type:    "meta",
				Payload: contact,
			}
		}
	}
}

func (w *TinodeWorker) handleData(data map[string]interface{}) {
	// Example: Handle new messages in topics
	topic := data["topic"].(string)
	content := data["content"].(string)
	from := data["from"].(string)
	// Process the message
	log.Printf("New message in topic %s from %s: %s", topic, from, content)
	// Send event
	w.eventChan <- TinodeEvent{
		Type:    "data",
		Payload: data,
	}
}

func (w *TinodeWorker) handlePres(pres map[string]interface{}) {
	// Example: Handle user online/offline status, subscriptions, etc.
	what := pres["what"].(string)
	topic := pres["topic"].(string)
	src := pres["src"].(string)

	switch what {
	case "on":
		log.Printf("User %s is online", src)
	case "off":
		log.Printf("User %s is offline", src)
	case "acs":
		// Access mode changed
		log.Printf("Access mode changed for %s in topic %s", src, topic)
	case "upd":
		// Topic or user info updated
		log.Printf("Update in topic %s", topic)
	case "gone":
		// Topic deleted or user unsubscribed
		log.Printf("Topic %s is gone", topic)
	default:
		log.Printf("Unhandled presence event: %v", pres)
	}

	// Send event
	w.eventChan <- TinodeEvent{
		Type:    "pres",
		Payload: pres,
	}
}
