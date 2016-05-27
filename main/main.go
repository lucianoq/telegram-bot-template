package main

import (
	"net/http"
	"bytes"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"encoding/json"
)

const TOKEN = "" //TODO
const SERVER = "https://api.telegram.org/bot" + TOKEN
const ROUTE = "" //TODO
const PORT = "" //TODO
const PEM_FILE = "" //TODO
const KEY_FILE = "" //TODO

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	r := gin.Default()
	r.POST(ROUTE, func(c *gin.Context) {
		var update Update
		if c.BindJSON(&update) == nil {
			//TODO
			msg := update.Message.Text
			SendMessage(update.Message.Chat.ID, msg)
		}
	})
	r.RunTLS("0.0.0.0:" + PORT, PEM_FILE, KEY_FILE)
}

func SendMessage(chat_id int64, text string) {
	msg := MessageToSend{chat_id, text}
	jsonS, _ := json.Marshal(msg)
	req, _ := http.NewRequest("POST", SERVER + "/sendMessage", bytes.NewBuffer(jsonS))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(req)
}