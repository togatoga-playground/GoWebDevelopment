package main

import (
	"math/rand"
	"time"
	"bufio"
	"os"
	"fmt"
	"strings"
)

const otherWorld = "*"

var transforms = []string{
	otherWorld,
	otherWorld,
	otherWorld,
	otherWorld,
	otherWorld + "app",
	otherWorld + "site",
	otherWorld + "time",
	"get" + otherWorld,
	"go" + otherWorld,
	"lets" + otherWorld,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWorld, s.Text(), -1))
	}
}
