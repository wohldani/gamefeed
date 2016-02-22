package gamefeed

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

type Request struct {
	Ctx    context.Context
	Client *http.Client
}

func NewRequest(ctx context.Context) Request {
	return Request{
		Ctx:    ctx,
		Client: urlfetch.Client(ctx),
	}
}

func init() {
	http.HandleFunc("/", handleFeed)
}

func handleFeed(w http.ResponseWriter, r *http.Request) {
	request := NewRequest(appengine.NewContext(r))
	request.parseFeed()
}
