package webLink

import (
	"context"
	"errors"
	"fmt"
	"github.com/PeterYangs/request/v2"
	"github.com/PuerkitoBio/goquery"
	urls "net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type webLink struct {
	cxt     context.Context
	client  *request.Client
	link    string
	list    sync.Map
	host    string
	scheme  string
	lock    sync.Mutex
	file    *os.File
	wait    sync.WaitGroup
	regular string
}

func NewWebLink(cxt context.Context, filePath string) *webLink {

	c := request.NewClient()

	c.Timeout(15 * time.Second)

	f, _ := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)

	return &webLink{cxt: cxt, client: c, list: sync.Map{}, lock: sync.Mutex{}, file: f, wait: sync.WaitGroup{}}
}

func (w *webLink) Link(link string) *webLink {

	w.link = link

	u, err := urls.Parse(link)

	if err == nil && u.Host != "" {

		w.host = u.Host

		w.scheme = u.Scheme
	}

	return w

}

func (w *webLink) Regular(regular string) *webLink {

	w.regular = regular

	return w
}

func (w *webLink) Run() error {

	defer w.file.Close()

	if w.file == nil {

		return errors.New("未配置文件地址")
	}

	for i := 0; i < 8; i++ {

		w.wait.Add(1)

		go w.getLink(w.link, 0)
	}

	w.wait.Wait()

	return nil

}

func (w *webLink) getLink(url string, flag int) {

	defer func() {

		if flag == 0 {

			w.wait.Done()
		}

	}()

	select {

	case <-w.cxt.Done():

		return

	default:

	}

	rsp, err := w.client.R().Get(url)

	if err != nil {

		fmt.Println(err)

		return
	}

	//text/html
	if !strings.Contains(rsp.Header().Get("Content-Type"), "text/html") {

		h := w.getUrl(url)

		w.putUrl(h)

		return
	}

	ct, err := rsp.Content()

	if err != nil {

		fmt.Println(err)

		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(ct.ToString()))

	if err != nil {

		fmt.Println(err)

		return
	}

	doc.Find("a").Each(func(i int, selection *goquery.Selection) {

		href, ok := selection.Attr("href")

		if !ok {

			return
		}

		href = w.getUrl(href)

		u, err := urls.Parse(href)

		if err != nil || u.Host == "" || u.Host != w.host {

			return
		}

		b := w.putUrl(href)

		if b {

			w.getLink(href, flag+1)

		}

	})

}

//获取完整链接
func (w *webLink) getUrl(u string) string {

	s, _ := regexp.MatchString(`^/.*$`, u)

	if s {

		u = w.scheme + "://" + w.host + u

		return u
	}

	return u
}

func (w *webLink) putUrl(u string) bool {

	w.lock.Lock()

	defer w.lock.Unlock()

	_, ok := w.list.Load(u)

	if !ok {

		w.list.Store(u, u)

		if w.regular != "" {

			re1, _ := regexp.MatchString(w.regular, strings.Replace(u, w.scheme+"://"+w.host, "", 1))

			if re1 {

				w.file.Write([]byte(u + "\n"))
			}

		} else {

			w.file.Write([]byte(u + "\n"))

		}

		fmt.Println(u)

		return true

	}

	return false

}
