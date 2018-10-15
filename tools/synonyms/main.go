package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aikuma0130/goWeb/tools/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%q の類語検索に失敗しました: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("%q に類語はありませんでした \n", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
