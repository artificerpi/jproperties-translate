package jproperties


// These replacements allow escaped character not to be translated
var replacementTable = [...]rune{
	'\\',
	':',
	',',
}

func Escape(message string) (text string, args []rune){
	var escapedRunes []rune
	i:=0
	for _, c := range message {
		// just use as an empty value
		var foundRune rune = '`'
		for _, r := range replacementTable {
			if c == r {
				foundRune = r
				break
			}
		}

		if foundRune != '`' {
			args = append(args, foundRune)
			escapedRunes = append(escapedRunes, '`')
			escapedRunes = append(escapedRunes, rune(i))
			escapedRunes = append(escapedRunes, '`')
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
		if runeValues[i] == '`' &&  i+2 < len(runeValues) && runeValues[i+2] == '`' {
			d := int(runeValues[i+1])
			if d < len(args){
				unescapedRunes = append(unescapedRunes, args[d])
				// escape 2 rune after current one
				i = i + 2
			}
		}
		unescapedRunes = append(unescapedRunes, runeValues[i])
	}


	return string(unescapedRunes)
}