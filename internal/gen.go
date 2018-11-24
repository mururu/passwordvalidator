package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"strings"
)

const passwordFile = "common-passwords.txt.gz"
const listFile = "../common.go"

func main() {
	passwords := make(map[string]struct{})
	f, err := os.Open(passwordFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		return
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		password := scanner.Text()
		if strings.Contains(password, `\`) {
			password = strings.Replace(password, `\`, `\\`, -1)
		}
		if strings.Contains(password, `"`) {
			password = strings.Replace(password, `"`, `\"`, -1)
		}
		passwords[password] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("num: %d\n", len(passwords))

	f, err = os.Create(listFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	if _, err = fmt.Fprintf(w, "package passwordvalidator\n\nvar commonPasswords = map[string]struct{}{\n"); err != nil {
		log.Fatal(err)
	}
	for key := range passwords {
		if _, err = fmt.Fprintf(w, "	\"%s\": struct{}{},\n", key); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := fmt.Fprintf(w, "}\n"); err != nil {
		log.Fatal(err)
	}
	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}
}
