package main

import (
	_ "db"
	"fmt"
	"libblog"
	"net/http"
	"strings"
)

func ArticlesHandle(w http.ResponseWriter, r *http.Request) {

	libblog.Logger.Println("Access " + r.URL.Path)
	urlsp := strings.Split(r.URL.Path, "/")
	if len(urlsp[2]) != 0 { //load articles page
		fmt.Println(urlsp[2])
	} else { //load List page
		fmt.Println("aaa")
	}
}
