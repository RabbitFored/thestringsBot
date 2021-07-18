package main

import (
  "strings"
  "fmt"
  tl "github.com/goTelegramBot/gogram"
  "github.com/goTelegramBot/gogram/types"
)


func count(b tl.Bot,m *types.Message){
    var message string
  if m.ReplyToMessage != nil{

    message = m.ReplyToMessage.Text
  }else{
    b.SendMessage(m.Chat.Id,"_Reply to any message with_ /count",&tl.Options{ParseMode:"Markdown"})
    return
  }

  word_count := len(strings.Fields(message))
  character := len(message)

  var counts string

  for index, element := range repetition(message) {
    li := fmt.Sprintf("*%s*  =  _%d_\n",index, element)
    counts = counts + li
    }

  text := fmt.Sprintf("*Words           :* `%d`\n*Characters  :* `%d`\n\n*Repeatation:*\n%s",word_count,character,counts)
  b.SendMessage(m.Chat.Id, text,&tl.Options{ParseMode:"Markdown"})
}

func repetition(st string) map[string]int {
  
    // using strings.Field Function
    input := strings.Fields(st)
    wc := make(map[string]int)
    for _, word := range input {
        _, matched := wc[word]
        if matched {
            wc[word] += 1
        } else {
            wc[word] = 1
        }
    }
    return wc
}
