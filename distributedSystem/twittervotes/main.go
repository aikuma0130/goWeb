package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

func main() {}

var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("MongoDBにダイヤル中: localhost")
	db, err = mgo.Dial("localhost")
	return err
}

func closedb() {
	db.Close()
	log.Println("データベース接続が閉じられました")

}
