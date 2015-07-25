package main

import (
	"db"
	"libblog"
	"net/http"
	"strconv"
	"strings"
	_ "time"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	libblog.Logger.Println("Access " + r.URL.Path)
	data := db.GetIndexData(0)
	data.HasNextPage = true
	err := localTemple.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func IndexOtherPageHandler(w http.ResponseWriter, r *http.Request) {
	libblog.Logger.Println("Access " + r.URL.Path)

	requestNumString := strings.Split(r.URL.Path, "/")[2]
	libblog.Logger.Println(requestNumString)
	requestNum, err := strconv.Atoi(requestNumString)
	if err != nil || requestNum > db.GetMaxIndexPage() {
		libblog.Logger.Println("Convert string " + requestNumString + " to int error")
		requestNum = 0
	}

	data := db.GetIndexData(requestNum)
	data.HasNextPage = db.HasNextPage(requestNum)

	err = localTemple.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
