package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var fatalError error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalError = e
}

func main() {
	defer func() {
		if fatalError != nil {
			os.Exit(1)
		}
	}()

	log.Println(" データベースに接続します ...")
	db, err := mgo.Dial("localhost")
	if err != nil {
		fatal(err)
		return
	}

	defer func() {
		log.Println(" データベース接続を閉じます ...")
		db.Close()
	}()
	pollData := db.DB("ballots").C("polls")
}
