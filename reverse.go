package main

import (
  tl "github.com/goTelegramBot/gogram"
     "github.com/goTelegramBot/gogram/types"
)

func reverse(b tl.Bot,m *types.Message){

  var message string
  if m.ReplyToMessage != nil{

    message = m.ReplyToMessage.Text
  }else{
    b.SendMessage(m.Chat.Id,"_Reply to any message with_ /reverse _to reverse texts_",&tl.Options{ParseMode:"Markdown"})
    return
  }
  
  reverse := reverseString(message)
  
  b.SendMessage(m.Chat.Id, reverse,nil)

}

func reverseString(str string) (result string) {
    for _, v := range str {
        result = string(v) + result
    }
    return
}
