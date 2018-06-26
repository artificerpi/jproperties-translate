package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/artificerpi/gotranslate"
	"github.com/artificerpi/jproperties-translate/jproperties"
	"golang.org/x/text/language"
)

const extension = ".properties"

var (
	srcFile string
	dstFile string
	lang    string
	dirPath string
)

func translateDir(path string, langTag language.Tag) {
	var wg sync.WaitGroup

	for _, propFile := range listPropFiles(path) {
		wg.Add(1)

		go func(filename string) {
			defer wg.Done()

			var srcProps jproperties.Properties
			dstProps := jproperties.Properties{}
			srcProps.Load(filename)

			for _, name := range srcProps.Keys() {
				translated := gotranslate.QuickTranslate(srcProps.Get(name), langTag)
				dstProps.Put(name, translated)
			}
			dstProps.Store(filename, "Translated Properties")
			log.Println("Translated", filename, "into", filename, "with language", langTag.String())
		}(propFile)

	}

	wg.Wait()
}

func listPropFiles(dir string) []string {
	propFiles := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}

		if !info.IsDir() && filepath.Ext(info.Name()) == extension {
			propFiles = append(propFiles, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return propFiles
}

func main() {
	flag.StringVar(&srcFile, "s", "source.properties", "specify the source properties file you want to translate")
	flag.StringVar(&dstFile, "t", srcFile, "specify the target properties file to save")
	flag.StringVar(&lang, "lang", "zh-CN", "specify the language wanted")
	flag.StringVar(&dirPath, "dir", "", "specify the dir store properties files to be translated")
	flag.Parse()

	langTag, err := language.Parse(lang)
	if err != nil {
		log.Fatal(err)
	}

	if len(dirPath) > 0 {
		translateDir(dirPath, langTag)
		return
	}

	var srcProps jproperties.Properties
	dstProps := jproperties.Properties{}
	srcProps.Load(srcFile)
	for _, name := range srcProps.Keys() {
		message := srcProps.Get(name)
		text, args := jproperties.Escape(message)
		log.Println(string(args))
		translatedText := gotranslate.QuickTranslate(text, langTag)
		translatedMessage := jproperties.Format(translatedText, args)
		log.Println(message ,translatedMessage)
		dstProps.Put(name, translatedMessage)
	}
	dstProps.Store(dstFile, "Translated Properties")
	log.Println("Translated", srcFile, "into", dstFile, "with language", langTag.String())
}
