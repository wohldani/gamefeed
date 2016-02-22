package gamefeed

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Link struct {
	Href string `datastore:"href"`
}

func NewLink(href string) Link {
	return Link{
		Href: href,
	}
}

func (link Link) Count(ctx context.Context) (int, error) {
	return datastore.NewQuery("links").
		Filter("href =", link.Href).Count(ctx)
}

func (link Link) Put(ctx context.Context) error {
	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "links", nil), &link)
	return err
}
