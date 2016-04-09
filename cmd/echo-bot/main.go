package main

import (
	"fmt"
	"github.com/rockneurotiko/go-tgbot"
	"log"
	"os"
)

const tokenString string = "botToken"
const hookURLString string = "hookURL"

func echoHandler(bot tgbot.TgBot, msg tgbot.Message, vals []string, kvals map[string]string) *string {
	newmsg := fmt.Sprintf("[Echoed]: %s", vals[1])
	log.Println(newmsg)
	return &newmsg
}

func main() {
	token := os.Getenv(tokenString)
	bot := tgbot.NewTgBot(token).
		CommandFn(`echo (.+)`, echoHandler)
	
	hookURL := os.Getenv(hookURLString)
	if hookURL == "" {
		log.Println("No hook URL. Pooling")
		bot.SimpleStart()
	} else {
		log.Println("Working with web hook");
		bot.ServerStart(hookURL, "/")
	}
}
