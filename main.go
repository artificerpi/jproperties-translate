package main

import (
	"flag"
	"log"

	"github.com/artificerpi/gotranslate"
	"github.com/artificerpi/jproperties-translate/jproperties"
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

	var srcProps jproperties.Properties
	dstProps := jproperties.Properties{}
	srcProps.Load(srcFile)
	for _, name := range srcProps.Keys() {
		log.Println("Translating prop:", name, srcProps.Get(name))
		translated := gotranslate.QuickTranslation(srcProps.Get(name), gotranslate.English, gotranslate.ChineseS)
		dstProps.Put(name, translated)
	}
	dstProps.Store(dstFile, "Translated Properties")
}
