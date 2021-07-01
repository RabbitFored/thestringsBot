package main

import (
  tl "github.com/goTelegramBot/telepher"
  "github.com/goTelegramBot/telepher/types"
  cow "github.com/Code-Hex/Neo-cowsay"
  "log"
  "fmt"
)

func cowsay(b tl.Bot,m *types.Message){
  message := "Reply to any text"
  if m.ReplyToMessage != nil{
    message = m.ReplyToMessage.Text
  }

  say, err := cow.Say(
		cow.Phrase(message),
		cow.Type("default"),
		cow.BallonWidth(40),
	)
  	if err != nil {
		log.Println(err)
	}
  text := fmt.Sprintf("`%s`",say)
  b.SendMessage(m.Chat.Id, text ,&tl.Options{ParseMode:"Markdown"})
	
  

}
