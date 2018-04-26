package main

import (
	"flag"
	"log"

	"github.com/artificerpi/gotranslate"
	"github.com/artificerpi/jproperties-translate/jproperties"
	"golang.org/x/text/language"
)

var (
	srcFile string
	dstFile string
	lang    string
)

func main() {
	flag.StringVar(&srcFile, "s", "source.properties", "specify the source properties file you want to translate")
	flag.StringVar(&dstFile, "t", srcFile, "specify the target properties file to save")
	flag.StringVar(&lang, "lang", "zh-CN", "specify the language you want")
	flag.Parse()

	langTag, err := language.Parse(lang)
	if err != nil {
		log.Fatal(err)
	}

	var srcProps jproperties.Properties
	dstProps := jproperties.Properties{}
	srcProps.Load(srcFile)
	for _, name := range srcProps.Keys() {
		translated := gotranslate.QuickTranslate(srcProps.Get(name), langTag)
		dstProps.Put(name, translated)
	}
	dstProps.Store(dstFile, "Translated Properties")
}
