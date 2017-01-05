package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"net/url"
	"strings"
	"github.com/bitly/go-nsq"
)

type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

type tweet struct {
	Text string
}

func readFromTwitter(votes chan <- string) {
	options, err := loadOptions()
	if err != nil {
		log.Println("選択肢の読み込みに失敗いたしました: ", err)
		return
	}
	u, err := url.Parse("https://stream.twitter.com/1.1/statuses/filter.json")
	if err != nil {
		log.Println("URLの解析に失敗いたしました: ", err)
		return
	}
	query := make(url.Values)
	query.Set("track", strings.Join(options, ","))
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(query.Encode()))
	if err != nil {
		log.Println("検索のリクエストの作成に失敗しました: ", err)
		return
	}
	resp, err := makeRequest(req, query)
	if err != nil {
		log.Println("検索のリクエストに失敗しました: ", err)
		return
	}
	reader = resp.Body
	decoder := json.NewDecoder(reader)
	for {
		var tweet tweet
		if err := decoder.Decode(&tweet); err != nil {
			break
		}
		for _, option := range options {
			if strings.Contains(strings.ToLower(tweet.Text), strings.ToLower(option)) {
				log.Println("投票: ", option)
				votes <- option
			}
		}
	}
}

func publishVotes(votes <- chan string) <-chan struct{}{
	stopchan := make(chan struct{}, 1)
	pub, _ := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	go func(){
		for vote := range votes{
			pub.Publish("votes", []byte(vote))
		}
		log.Println("Publisher: 停止中です")
		pub.Stop()
		log.Println("Publisher: 停止しました")
		stopchan <- struct{}{}
	}()
	return stopchan
}


func main() {

}

var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("Connected...MongoDB: localhost")
	db, err = mgo.Dial("localhost")
	return err
}

func closedb() {
	db.Close()
	log.Println("データベース接続が閉じられました。")

}
