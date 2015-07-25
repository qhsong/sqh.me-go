package db

import (
	"container/list"
	"libblog"
)

func GetIndexData(page int) *libblog.IndexData {
	data := new(libblog.IndexData)
	//data.Articles = make([]libblog.ShortPage, 5)
	AricleList := list.New()

	//get Article short info
	rows, err := DBSqlite.Query("Select id,short_title,title,create_time,category from Articles order by create_time desc limit ?,?", page*5, 5)
	defer rows.Close()
	if err != nil {
		libblog.Logger.Fatal(err)
	}
	for rows.Next() {
		var category string
		var short_title string
		var id int
		var Articles libblog.ShortPage
		rows.Scan(&id, &short_title, &Articles.Title, &Articles.WriteDate, &category)

		//Permalink
		Articles.Permalink = libblog.BaseUrl + "/" + category + "/" + short_title

		//tags
		tagRows, _ := DBSqlite.Query("select a.name from tag a inner join tagArticleRelation b on a.id=b.tagid where b.articleid=?", id)
		tagList := list.New()
		for tagRows.Next() {
			var tag string
			tagRows.Scan(&tag)
			tagList.PushBack(tag)
		}
		Articles.Tags = make([]string, tagList.Len())
		tagIndex := 0
		for e := tagList.Front(); e != nil; e = e.Next() {
			Articles.Tags[tagIndex] = e.Value.(string)
			tagIndex++
		}
		AricleList.PushBack(Articles)
	}

	data.Articles = make([]libblog.ShortPage, AricleList.Len())

	index := 0
	for e := AricleList.Front(); e != nil; e = e.Next() {
		data.Articles[index] = e.Value.(libblog.ShortPage)
		index++
	}

	data.PageNum = page
	data.Site.Title = libblog.Title
	data.Site.BaseUrl = libblog.BaseUrl
	data.Site.Copyright = libblog.Copyright
	data.Title = "Index-" + data.Site.Title
	data.RSSlink = "http://sqh.me/index.xml"
	return data
}

func GetMaxIndexPage() int {
	rows, err := DBSqlite.Query("Select count(*) as count from Articles ")
	defer rows.Close()
	if err != nil {
		libblog.Logger.Panic(err)
	}
	var count int
	rows.Next()
	rows.Scan(&count)

	TotalPage := count/5 - 1
	if count%5 != 0 {
		TotalPage = TotalPage + 1
	}

	return TotalPage
}

func HasNextPage(page int) bool {
	TotalPage := GetMaxIndexPage()
	if page < TotalPage {
		return true
	} else {
		return false
	}
}
