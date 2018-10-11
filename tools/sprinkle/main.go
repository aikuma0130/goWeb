package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = make([]string, 0)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	exe, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}
	dirPath := filepath.Dir(exe)
	f, err := os.Open(dirPath + "/transforms.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	t := bufio.NewScanner(f)
	for t.Scan() {
		transforms = append(transforms, t.Text())
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		transformedWord := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(transformedWord, otherWord, s.Text(), -1))
	}
}
