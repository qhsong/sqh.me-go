package libblog

import (
	"time"
)

type ShortPage struct {
	Title     string
	Link      string
	WriteDate time.Time
	Tags      []string
	Permalink string
}

type Page struct {
	Main    ShortPage
	Content string
}
