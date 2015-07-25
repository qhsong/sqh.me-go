package libblog

type IndexData struct {
	Articles    []ShortPage
	Site        Siteinfo
	Title       string
	RSSlink     string
	PageNum     int
	HasNextPage bool
}

type Siteinfo struct {
	Title     string
	BaseUrl   string
	Copyright string
}

func (p *IndexData) GetNextPageNum() int {
	return p.PageNum + 1
}

func (p *IndexData) GetPrevPageNum() int {
	return p.PageNum - 1
}
