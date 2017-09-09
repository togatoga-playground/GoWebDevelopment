package main

import (
	"math/rand"
	"time"
	"bufio"
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
)

const otherWorld = "*"

func readTransforms() []string {
	b, err := ioutil.ReadFile("config/transforms.conf")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\n")
}

func main() {
	transforms := readTransforms()

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWorld, s.Text(), -1))
	}
}
