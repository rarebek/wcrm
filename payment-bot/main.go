package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {
	botToken := "6950271943:AAH-E-aI0O6bB625qzLm-yv5Nh4P539xk58"
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update.Message)
		} else if update.PreCheckoutQuery != nil {
			handlePreCheckoutQuery(bot, update.PreCheckoutQuery)
		} else if update.Message != nil && update.Message.SuccessfulPayment != nil {
			handleSuccessfulPayment(bot, update.Message)
		}
	}
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	case "/start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to our bot! Use /buy to purchase.")
		bot.Send(msg)
	case "/buy":
		sendInvoice(bot, message.Chat.ID)
	}
}

func sendInvoice(bot *tgbotapi.BotAPI, chatID int64) {
	// Fetch product details from API
	products, err := getProductsFromAPI()
	if err != nil {
		log.Println("Failed to fetch product details:", err)
		return
	}

	// Create invoice
	var prices []tgbotapi.LabeledPrice
	var productNames []string
	for _, product := range products {
		prices = append(prices, tgbotapi.LabeledPrice{Label: product.Name, Amount: product.Price})
		productNames = append(productNames, product.Name)
	}

	invoice := tgbotapi.NewInvoice(chatID,
		strings.Join(productNames, ", "),
		"Description of products",
		"payload", // Unique payload identifier
		"398062629:TEST:999999999_F91D8F69C042267444B74CC0B3C747757EB0E065", // Replace with your actual provider token from @BotFather
		"StartParam", // Unique deep-linking parameter
		"UZS",        // Currency code
		prices)

	// Add empty array for suggested tip amounts
	invoice.SuggestedTipAmounts = []int{}

	if _, err := bot.Send(invoice); err != nil {
		log.Panic(err)
	}
}

func getProductsFromAPI() ([]Product, error) {
	// Replace "your_api_endpoint" with the actual endpoint of your API
	resp, err := http.Get("http://localhost:8080/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func handlePreCheckoutQuery(bot *tgbotapi.BotAPI, query *tgbotapi.PreCheckoutQuery) {
	preCheckoutConfig := tgbotapi.PreCheckoutConfig{
		PreCheckoutQueryID: query.ID,
		OK:                 true,
		ErrorMessage:       "",
	}
	preCheckoutConfig.OK = true
	if _, err := bot.Request(preCheckoutConfig); err != nil {
		log.Panic(err)
	}
}

func handleSuccessfulPayment(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Thank you for your purchase! Your order has been processed.")
	bot.Send(msg)
}
