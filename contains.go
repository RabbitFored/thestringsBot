package main

import (
  "regexp"
  tl "github.com/goTelegramBot/gogram"
  "github.com/goTelegramBot/gogram/types"
)

func contains(b tl.Bot,m *types.Message){
    var message string
  if m.ReplyToMessage != nil{

    message = m.ReplyToMessage.Text
  }else{
    b.SendMessage(m.Chat.Id,"_Reply to any message with_ /contains <keyword> _to match texts_",&tl.Options{ParseMode:"Markdown"})
    return
  }
  args := m.Args()
  
  if len(args) < 1 {
      b.SendMessage(m.Chat.Id,"*Specify some keyword to match text*",&tl.Options{ParseMode:"Markdown"})
      return
  }
  rx := regexp.MustCompile(args[1])
  if rx.MatchString(message){
  b.SendMessage(m.Chat.Id,"Contains",&tl.Options{ParseMode:"Markdown"})
  }else{
    b.SendMessage(m.Chat.Id, "*Message dont match*",&tl.Options{ParseMode:"Markdown"})
  }

}
