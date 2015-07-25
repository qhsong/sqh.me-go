package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var localTemple *template.Template

func addLocalTemplateFile() {
	localTemple = template.Must(template.New("index.html").ParseFiles("template/index.html"))
	filepath.Walk("template/", func(path string, f os.FileInfo, err error) error {
		_, fileName := filepath.Split(path)
		if len(fileName) != 0 {
			if (strings.Contains(fileName, ".html") || strings.Contains(fileName, ".htm")) && !(strings.Contains(fileName, ".swp") || strings.Contains(fileName, "index.html")) {
				fmt.Println(fileName)
				localTemple = template.Must(localTemple.ParseFiles(path))
			}
		}
		return nil
	})
}

func main() {
	http.HandleFunc("/", IndexHandle)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	addLocalTemplateFile()

	http.ListenAndServe(":8080", nil)
}
