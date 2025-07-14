# Telegram Notifier

A simple Go package for sending notifications via Telegram. This package is specifically designed for sending notification messages and does not handle polling or user responses.

## Installation

```bash
go get github.com/huazano/telegram-notifier
```

## Usage

### Basic Setup

```go
package main

import (
    "log"
    telegramnotifier "github.com/huazano/telegram-notifier"
)

func main() {
    // Create the notifier with your bot token
    notifier := telegramnotifier.New("YOUR_BOT_TOKEN")
    
    // Send a notification
    chatID := int64(123456789) // Chat or user ID
    err := notifier.Send(chatID, "Something important happened!")
    if err != nil {
        log.Printf("Error: %v", err)
    }
}
```

### Available Methods

- `New(botToken string) *Notifier`: Creates a new notifier instance
- `Send(chatID int64, message string) error`: Sends a message to the specified chat

## Use Cases

- System alerts
- Error notifications
- Automated reports
- Monitoring alerts
- Any notification that doesn't require user response

## Requirements

1. A Telegram bot (create one with @BotFather)
2. The bot token
3. The chat ID where you want to send messages

### How to get the bot token

1. Talk to @BotFather on Telegram
2. Use the `/newbot` command
3. Follow the instructions to create your bot
4. Save the token provided

### How to get the Chat ID

1. Send a message to your bot
2. Visit: `https://api.telegram.org/bot<YOUR_TOKEN>/getUpdates`
3. Look for the `chat.id` in the JSON response

## Complete Example

```go
package main

import (
    "log"
    "os"
    telegramnotifier "github.com/huazano/telegram-notifier"
)

func main() {
    // Configure from environment variables
    botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
    if botToken == "" {
        log.Fatal("TELEGRAM_BOT_TOKEN is not configured")
    }

    notifier := telegramnotifier.New(botToken)
    
    // Send different types of notifications
    chatID := int64(123456789)
    
    // Error notification
    err := notifier.Send(chatID, "🚨 Critical system error")
    if err != nil {
        log.Printf("Error: %v", err)
    }
    
    // Success notification
    err = notifier.Send(chatID, "✅ Process completed successfully")
    if err != nil {
        log.Printf("Error: %v", err)
    }
}
```

## Note

This package is designed solely for sending messages. It does not include polling functionality, command handling, or user response management. It's perfect for alert systems and automated notifications.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.