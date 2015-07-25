package libblog

type indexData struct {
	Articles []ShortPage
	Site     Siteinfo
	Title    string
	RSSlink  string
}

type Siteinfo struct {
	Title   string
	BaseUrl string
}
