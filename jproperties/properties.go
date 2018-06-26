package jproperties

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Properties provides the way to load and store properties to file.
type Properties struct {
	dict        map[string]string
	description string
}

// Load reads properties from file with specified name
func (p *Properties) Load(filename string) {
	if len(filename) == 0 {
		log.Fatal("Invalid file name when loading properties file")
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	props, err := readProps(file)
	if err != nil {
		log.Fatal(err)
	}

	p.dict = props.dict
	p.description = props.description
}

// Store writes properties to file
func (p *Properties) Store(filename, description string) {
	if len(filename) == 0 {
		log.Fatal("Invalid file name when storing properties file")
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = writeProps(file, p)
	if err != nil {
		log.Fatal(err)
	}
}

// Get retrives value from Properties by specified key
func (p *Properties) Get(key string) string {
	if p == nil || p.dict == nil {
		log.Fatal("Try to get value from null Properties")
	}

	return p.dict[key]
}

// Put save a property
func (p *Properties) Put(key, value string) {
	if p == nil {
		log.Fatal("Try to save value from null Properties")
	}

	if p.dict == nil {
		p.dict = make(map[string]string)
	}
	p.dict[key] = value
}

// Keys return all keys of properties
func (p *Properties) Keys() []string {
	keys := []string{}
	for k := range p.dict {
		keys = append(keys, k)
	}

	return keys
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

func writeProps(writer io.Writer, p *Properties) error {
	if p == nil {
		return errors.New("Writing null Properties")
	}

	if len(p.description) > 0 {
		fmt.Fprintln(writer, p.description)
	}

	for k, v := range p.dict {
		prop := k + "=" + v
		fmt.Fprintln(writer, prop)
	}

	return nil
}
