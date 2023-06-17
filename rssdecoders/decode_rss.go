package rssdecoders

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/sobhanhsa/simpleblog/db"
)

func Decoder(address string, category string) {
	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return
	}
	defer resp.Body.Close()

	rss := Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return
	}
	var data Item
	for _, item := range rss.Channel.Items {
		data = item
		db.CreateArticle("sobhanhsa1", category, data.Title, data.Desc, "")
	}
}
