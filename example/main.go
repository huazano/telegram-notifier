package main

import (
	"log"
	"os"

	telegramnotifier "github.com/huazano/telegram-notifier"
)

func main() {
	// Get bot token from environment variables
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not configured")
	}

	// Create the notifier
	notifier := telegramnotifier.New(botToken)

	// Send a notification
	chatID := int64(123456789) // Replace with actual chat ID
	err := notifier.Send(chatID, "¡Hola! Esta es una notificación de prueba.")
	if err != nil {
		log.Printf("Error sending notification: %v", err)
		return
	}

	log.Println("Notification sent successfully")

	// Send another notification
	err = notifier.Send(chatID, "Esta es otra notificación de prueba.")
	if err != nil {
		log.Printf("Error sending notification: %v", err)
		return
	}

	log.Println("Second notification sent successfully")
}