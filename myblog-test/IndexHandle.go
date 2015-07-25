package main

import (
	"./libblog"
	"fmt"
	"net/http"
	"time"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	var data libblog.indexData
	local, _ := time.LoadLocation("")
	data.Articles = make([]ShortPage, 5)
	data.Articles[0] = ShortPage{Title: "11", Link: "111", WriteDate: time.Date(2010, 1, 1, 11, 1, 1, 1, local), Tags: make([]string, 3)}
	data.Articles[1] = ShortPage{Title: "11", Link: "111", WriteDate: time.Date(2010, 1, 1, 11, 1, 1, 1, local), Tags: make([]string, 3)}
	data.Articles[2] = ShortPage{Title: "11", Link: "111", WriteDate: time.Date(2010, 1, 1, 11, 1, 1, 1, local), Tags: make([]string, 3)}
	data.Articles[3] = ShortPage{Title: "11", Link: "111", WriteDate: time.Date(2010, 1, 1, 11, 1, 1, 1, local), Tags: make([]string, 3)}
	data.Articles[4] = ShortPage{Title: "11", Link: "111", WriteDate: time.Date(2010, 1, 1, 11, 1, 1, 1, local), Tags: make([]string, 3)}
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			data.Articles[i].Tags[j] = "tag" + fmt.Sprint(j)
		}
	}
	data.Site.Title = "qhsong "
	data.Title = "Index"
	data.Site.BaseUrl = "http://127.0.0.1:8080"
	data.RSSlink = "http://sqh.me/index.xml"
	err := localTemple.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
