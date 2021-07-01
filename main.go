package main

import (
  tl "github.com/goTelegramBot/telepher"
  "github.com/goTelegramBot/telepher/types"
  "log"
  "os"
  "fmt"
)


func start(bot tl.Bot,message *types.Message){
    text := "*Hi there!*\n\nI am theStringsbot.\nYou can do different tasks using me.\n\n*Hit* /help *to know more on how to use me.*"
    
    markup := tl.InlineKeyboardMarkup()
but1 := types.InlineKeyboardButton{Text:"Channel",Url: "https://t.me/theostrich"}
    row1 := markup.Row(but1)

  keyboard := markup.Parse(row1)
     
     bot.SendMessage(message.Chat.Id, text,&tl.Options{ReplyMarkup:&keyboard,ParseMode:"Markdown"})
}

func help(bot tl.Bot,message *types.Message){
    text := fmt.Sprintf(`*Hi %s!* 
    Here is a detailed guide on using me.
    
    *Helpful commands:*
    - /start : Starts me! You've probably already used this.
    - /help  : I'll tell you more about myself!
    - /password : Generates a random password. Default password length is 8.
    - /strength: Check strength of your password.
    - /wordlist <keyword> : Provides wordlist of <keyword> [Under development]
    - /symbols : Get a list of symbols.
    - /reverse : .txet ruoy esreveR
    - /case <style> : Change sentence case of your text.
        Style maybe :
                  `+"`"+`upper`+"`"+`, `+"`"+`lower`+"`"+`, `+"`"+`title`+"`"+`, `+"`"+`camel`+"`"+`, `+"`"+`camel-lower`+"`"+`, `+"`"+`title-proper`+"`"+`, `+"`"+`mixed`+"`"+`, `+"`"+`mixed-inverse`+"`"+`
        Converts to Upper case if <style> not provided.
    - /count : Counts words and length of the replied to message.
    - /write : Make me write in a notebook.
    - /about  : Know about me.
    `, message.From.FirstName)
    
    markup := tl.InlineKeyboardMarkup()
but1 := types.InlineKeyboardButton{Text:"Get Help",Url: "https://t.me/ostrichdiscussion"}
but2 := types.InlineKeyboardButton{Text:"🔖Add me in group",Url: "https://t.me/thestringsBot?startgroup=new"}
    row1 := markup.Row(but1,but2)

  keyboard := markup.Parse(row1)
     
     bot.SendMessage(message.Chat.Id, text,&tl.Options{ReplyMarkup:&keyboard,ParseMode:"Markdown"})
}
func about(bot tl.Bot,message *types.Message){
    text :="<b>About Me :</b>\n\n" +
    "  - <b>Name        :</b> theStrings\n" +
    "  - <b>Creator     :</b> @theostrich\n" +
    "  - <b>Language    :</b> Golang\n" +
    "  - <b>Library     :</b> <a href=\"https://github.com/goTelegramBot/telepher\">Telepher</a>"

    markup := tl.InlineKeyboardMarkup()
but1 := types.InlineKeyboardButton{Text:"Channel",Url: "https://t.me/theostrich"}
but2 := types.InlineKeyboardButton{Text:"Support Group",Url: "https://t.me/ostrichdiscussion"}
    row1 := markup.Row(but1,but2)

  keyboard := markup.Parse(row1)

     bot.SendMessage(message.Chat.Id, text,&tl.Options{ReplyMarkup:&keyboard,ParseMode:"html",DisableWebPagePreview:true})
}

func main() {
 b,err := tl.NewBot(os.Getenv("TOKEN"),nil)

  if err != nil{
  log.Println(err)
  return
  }

b.Command("start",start)
b.Command("help",help)
b.Command("about",about)
b.Command("password",password)
b.Command("wordlist",wordlist)
b.Command("count",count)
b.Command("cowsay",cowsay)
b.Command("upper",upper)
b.Command("lower",lower)
b.Command("case",change_case)
b.Command("reverse",reverse)
b.Command("symbols",symbols)
b.Command("contains",contains)
b.Command("strength",strength)
// b.Command("write",write)
b.Start()
}
