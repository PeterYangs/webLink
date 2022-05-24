package main

import (
	"context"
	"github.com/PeterYangs/webLink"
	"time"
)

func main() {

	c, _ := context.WithTimeout(context.Background(), 50*time.Minute)

	w := webLink.NewWebLink(c, "url.txt")

	w.Link("https://www.925g.com/")

	w.Regular(`\.html$`)

	w.Run()

}
