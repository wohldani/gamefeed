package gamefeed

import (
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"
)

func (r Request) parseFeed() {
	links := r.getLinks("http://bluesnews.com", "https://www.youtube.com")
	var wg sync.WaitGroup
	for _, link := range links {
		if n, _ := link.Count(r.Ctx); n == 0 {
			wg.Add(1)
			go func(link Link) {
				defer wg.Done()
				link.Put(r.Ctx)
				title, err := r.getYoutubeTitle(link.Href)
				if err != nil {
					return
				}
				r.sendEmailMessage(title, link.Href)
			}(link)
		}
	}
	wg.Wait()
}

func (r Request) getLinks(fromUrl string, toDomain string) []Link {
	var links []Link
	res, _ := r.Client.Get(fromUrl)
	nodes := findNodes(res.Body, "a")

	for _, link := range nodes {
		for _, a := range link.Attr {
			if a.Key == "href" {
				if strings.Contains(link.String(), toDomain) {
					links = append(links, NewLink(a.Val))
				}
				break
			}
		}
	}
	return links
}

func (r Request) getYoutubeTitle(url string) (string, error) {
	res, err := r.Client.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return regexp.MustCompile("<title>(.*?)</title>").FindStringSubmatch(string(str))[1], nil
}

func (r Request) sendEmailMessage(subject string, message string) {
	msg := &mail.Message{
		Sender:  "Bluesnews Youtube Watcher wohldani@gmail.com",
		To:      []string{"wohldani@gmail.com"},
		Subject: subject,
		Body:    message,
	}
	if err := mail.Send(r.Ctx, msg); err != nil {
		log.Errorf(r.Ctx, "Couldn't send email: %v", err)
	}
}
