package jproperties

import "strconv"

// TRANSPLATE_SEPARATOR is used with arguments to format message
// this can be useful to make the message text not to be escaped during translation
const TRANSPLATE_SEPARATOR = '-'

// These replacements allow escaped character not to be translated
var replacementTable = [...]rune{
	'\\',
	':',
	',',
	'\n',
	'\r',
	'\t',
}

func Escape(message string) (text string, args []rune){
	var escapedRunes []rune
	i:=0
	for _, c := range message {
		// just use it as an "empty" value
		var foundRune rune = TRANSPLATE_SEPARATOR
		for _, r := range replacementTable {
			if c == r {
				foundRune = r
				break
			}
		}

		if foundRune != TRANSPLATE_SEPARATOR {
			args = append(args, foundRune)
			escapedRunes = append(escapedRunes, TRANSPLATE_SEPARATOR)
			// TODO rune int to char
			escapedRunes = append(escapedRunes, rune('0' + i))
			escapedRunes = append(escapedRunes, TRANSPLATE_SEPARATOR)
			i++
		}else{
			escapedRunes = append(escapedRunes, c)
		}
	}

	return string(escapedRunes), args
}


func Format(text string, args []rune) (message string){
	runeValues := []rune(text)
	var unescapedRunes []rune
	for i :=0; i <len(runeValues); i++{
		if runeValues[i] == TRANSPLATE_SEPARATOR &&  i+2 < len(runeValues) && runeValues[i+2] == TRANSPLATE_SEPARATOR {
			d, err:= strconv.Atoi(string(runeValues[i+1]))
			if err==nil && d < len(args){
				unescapedRunes = append(unescapedRunes, args[d])
				// escape 2 rune after current one
				i = i + 2
				continue
			}
		}
		unescapedRunes = append(unescapedRunes, runeValues[i])
	}


	return string(unescapedRunes)
}