package main

import (
	"math/rand"
	"time"
	"bufio"
	"os"
	"strings"
	"unicode"
	"fmt"
	"flag"
)

var tlds = []string{"com", "net"}
const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

var tld *string

func init() {
	tld = flag.String("tld", "com", "tld(ex com, net jp)")
	flag.Parse()
}
func main() {


	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r){
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + *tld)
	}

}


