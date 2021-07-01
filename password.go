package main

import (
  "strings"
  "math/rand"
  "time"
  "fmt"
  "strconv"
  zxcvbn "github.com/nbutton23/zxcvbn-go"
  tl "github.com/goTelegramBot/telepher"
  "github.com/goTelegramBot/telepher/types"
)

var (
    lowerCharSet   = "abcdedfghijklmnopqrst"
    upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialCharSet = "!@#$%&*"
    numberSet      = "0123456789"
    allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
    var password strings.Builder

    //Set special character
    for i := 0; i < minSpecialChar; i++ {
        random := rand.Intn(len(specialCharSet))
        password.WriteString(string(specialCharSet[random]))
    }

    //Set numeric
    for i := 0; i < minNum; i++ {
        random := rand.Intn(len(numberSet))
        password.WriteString(string(numberSet[random]))
    }

    //Set uppercase
    for i := 0; i < minUpperCase; i++ {
        random := rand.Intn(len(upperCharSet))
        password.WriteString(string(upperCharSet[random]))
    }

    remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
    for i := 0; i < remainingLength; i++ {
        random := rand.Intn(len(allCharSet))
        password.WriteString(string(allCharSet[random]))
    }
    inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
func password(b tl.Bot,m *types.Message){
   
    args :=  m.Args()
    
    rand.Seed(time.Now().Unix())
    minSpecialChar := 1
    minNum := 1
    minUpperCase := 1
    passwordLength := 8
    if len(args) != 1 {
    i, err := strconv.Atoi(args[1])
    if err != nil {
b.SendMessage(m.Chat.Id, "*Provide an integer as password length*",&tl.Options{ParseMode:"Markdown"})
      return
    }
      passwordLength = i
    }
    if  (3 < passwordLength) && (passwordLength <  4000) {
    password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
    
    b.SendMessage(m.Chat.Id, password,nil)
    }else{
      b.SendMessage(m.Chat.Id, "*Password length should be between 3 and 4000*",&tl.Options{ParseMode:"Markdown"})
    }

}


func strength(b tl.Bot,m *types.Message){
    var message string
  if m.ReplyToMessage != nil{

    message = m.ReplyToMessage.Text
  }else{
    b.SendMessage(m.Chat.Id,"_Reply to any password text with_ /strength _to check its strength_",&tl.Options{ParseMode:"Markdown"})
    return
  }
  if len(message) > 150 {
    
        b.SendMessage(m.Chat.Id,"*Password length is too long. This maybe a strong password.*",&tl.Options{ParseMode:"Markdown"})
  }
  
  
  strength := zxcvbn.PasswordStrength(message, []string{})
  text := fmt.Sprintf("*Password :* `%s`\n*Entropy     :* `%f`\n*CrackTime :* `%f`\n*CrackTimeDisplay :* `%s`\n*Score :* `%d`\n*CalcTime  :* `%f`",strength.Password,strength.Entropy,strength.CrackTime,strength.CrackTimeDisplay,strength.Score,strength.CalcTime)

    b.SendMessage(m.Chat.Id,text,&tl.Options{ParseMode:"Markdown"})
  
}
