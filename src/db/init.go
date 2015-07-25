package db

import (
	"container/list"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"libblog"
)

var DBSqlite *sql.DB

func init() {
	var err error
	DBSqlite, err = sql.Open("sqlite3", "blog.db")

	if err != nil {
		libblog.Logger.Fatal(err)
	}
}

func GetCategory() *list.List {
	categoryList := list.New()
	rows, _ := DBSqlite.Query("Select distinct(category) from Articles")
	defer rows.Close()

	for rows.Next() {
		var cate string
		rows.Scan(&cate)
		categoryList.PushBack(cate)
	}

	return categoryList
}
