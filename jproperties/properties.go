package jproperties

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Properties struct {
	dict        map[string]string
	description string
}

func (p *Properties) Load(filename string) {
	p = &Properties{dict: make(map[string]string)}
	if len(filename) == 0 {
		log.Fatal("Invalid file name when loading properties file")
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// p.readProps(file)
}

func (p *Properties) Store(filename, description string) {

}

func readProps(reader io.Reader) (p *Properties, err error) {
	p = &Properties{dict: make(map[string]string)}
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				p.dict[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Properties) writeProps(writer io.Writer, description string) {

}
