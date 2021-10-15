package main

import (
	tl "github.com/goTelegramBot/gogram"
	"github.com/goTelegramBot/gogram/types"
	"strings"
)

func flip(b tl.Bot, m *types.Message) {

	var message string
	if m.ReplyToMessage != nil {

		message = m.ReplyToMessage.Text
	} else {
		b.SendMessage(m.Chat.Id, "_Reply to any message with_ /reverse _to reverse texts_", &tl.Options{ParseMode: "Markdown"})
		return
	}

	reverse := reverseString(message)
	flip := flipString(reverse)

	b.SendMessage(m.Chat.Id, flip, nil)

}

func flipString(str string) string {

	normal := "abcdefghijklmnopqrstuvwxyz_,;.?!/\\'"
	inverse := "ɐqɔpǝɟbɥıظʞןɯuodbɹsʇnʌʍxʎz‾'؛˙¿¡/\\,"

	normal += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	inverse += "∀qϽᗡƎℲƃHIſʞ˥WNOԀὉᴚS⊥∩ΛMXʎZ"

	normal += "0123456789 "
	inverse += "0ƖᄅƐㄣϛ9ㄥ86 "

	strArr := []rune(inverse)
	strArr2 := []rune(normal)
	var flipped string
	for _, s := range str {

		letter := string(s)
		var a int
		b := rune(1)
		if strings.Contains(normal, letter) {
			a = strings.Index(normal, letter)
			b = strArr[a]

		} else if strings.Contains(inverse, letter) {
			a = strings.Index(inverse, letter)
			b = strArr2[a]

		} else {
			b = rune('\t')
		}
		flipped += string(b)

	}
	return flipped
}
