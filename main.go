package main

import (
	"fmt"

	"github.com/artificerpi/gotranslate"
)

func main() {
	fmt.Println(gotranslate.QuickTranslation("Hello", gotranslate.English, gotranslate.ChineseS))
}
