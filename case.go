package main

import (
  "strings"
  "fmt"
  "unicode"
  tl "github.com/goTelegramBot/gogram"
  "github.com/goTelegramBot/gogram/types"
)

func upper(b tl.Bot,m *types.Message){
    var message string
  if m.ReplyToMessage != nil{

    message = m.ReplyToMessage.Text
  }else{
    b.SendMessage(m.Chat.Id,"_Reply to any message with_ /upper",&tl.Options{ParseMode:"Markdown"})
    return
  }
  upperCase := strings.ToUpper(message)
  b.SendMessage(m.Chat.Id,upperCase,nil)

}

func lower(b tl.Bot,m *types.Message){
    var message string
  if m.ReplyToMessage != nil{

    message = m.ReplyToMessage.Text
  }else{
    b.SendMessage(m.Chat.Id,"_Reply to any message with_ /lower",&tl.Options{ParseMode:"Markdown"})
    return
  }
  lowerCase := strings.ToLower(message)
  b.SendMessage(m.Chat.Id,lowerCase,nil)

}

func change_case(b tl.Bot,m *types.Message){
    var message string
    var text string
    args := m.Args()

  if m.ReplyToMessage != nil{

    message = m.ReplyToMessage.Text
  }else{
    b.SendMessage(m.Chat.Id,"_Reply to any message with_ /case <keyword>",&tl.Options{ParseMode:"Markdown"})
    return
  }
      change_case := "upper"
   if len(args) != 1 {
     change_case = strings.ToLower(args[1])
  }
  switch change_case{
    case "upper":
      text = strings.ToUpper(message)
    case "lower":
      text = strings.ToLower(message)
    case "title":
      text = strings.Title(strings.ToLower(message))
    case "camel":
      text = ToCamel(message)
    case "camel-lower":
      text = ToLowerCamel(message)
    case "title-proper":
      text = properTitle(message)
    case "mixed":
      text = AlternateCase(strings.ToLower(message))   
    case "mixed-inverse":
      text = AlternateCaseInvert(strings.ToLower(message))     
    default:
      text = fmt.Sprintf("Unknown case type : %s",change_case)

  }
  b.SendMessage(m.Chat.Id,text,nil)

}

var uppercaseAcronym = map[string]string{
	"ID": "id",
}

func toCamelInitCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	if a, ok := uppercaseAcronym[s]; ok {
		s = a
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}

// ToCamel converts a string to CamelCase
func ToCamel(s string) string {
	return toCamelInitCase(s, true)
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	return toCamelInitCase(s, false)
}

func properTitle(input string) string {
    words := strings.Split(input, " ")
    smallwords := " a an on the to "

    for index, word := range words {
        if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
            words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func AlternateCase(s string) string {
    rs, upper := []rune(s), false
    for i, r := range rs {
        if unicode.IsLetter(r) {
            if upper = !upper; upper {
                rs[i] = unicode.ToUpper(r)
            }
        }
    }
    return string(rs)
}

func AlternateCaseInvert(s string) string {
    rs, upper := []rune(s), true
    for i, r := range rs {
        if unicode.IsLetter(r) {
            if upper = !upper; upper {
                rs[i] = unicode.ToUpper(r)
            }
        }
    }
    return string(rs)
}
