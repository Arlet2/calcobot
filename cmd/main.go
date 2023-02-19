package main

import (
	"calcobot/internal/database"
	"calcobot/internal/network"
	"fmt"
	"os"
)

func main() {
	db, err := database.NewPostgresDatabase(os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), false)
	if err != nil {
		fmt.Println("Found error: "+err.Error())
		return
	}
	defer db.Close()

	bot, err := network.NewBot(os.Getenv("BOT_TOKEN"), true)

	if err != nil {
		fmt.Println("Found error: "+err.Error())
		return
	}
	go network.StartHttpServer("/lol", db)
	bot.StartWorking(db)
}