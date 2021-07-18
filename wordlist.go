package main

import (
  tl "github.com/goTelegramBot/gogram"
  "github.com/goTelegramBot/gogram/types"
)


func wordlist(b tl.Bot,m *types.Message){
   args := m.Args()
   if len(args) < 2 {
      b.SendMessage(m.Chat.Id,"*Specify some keyword to get wordlist from available collection*",&tl.Options{ParseMode:"Markdown"})
      return
  }
   if args[1] == "profanity" {
  document := types.InputFile{
    FilePath:"wordlists/profanity-en.txt",  
  }
     _,err := b.SendDocument(m.Chat.Id, document,nil)
     if err != nil{
        b.SendMessage(m.Chat.Id,"*Wordlist not found*",&tl.Options{ParseMode:"Markdown"})
     }
     
   }else{
  b.SendMessage(m.Chat.Id,"*Wordlist not found*",&tl.Options{ParseMode:"Markdown"})
   }

}
