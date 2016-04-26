package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jhh/puka"
)

// LegacyBookmark is the previous format for bookmarks.
type LegacyBookmark struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Title   string
	URL     string
	Comment string
	Date    time.Time `bson:"add_date"`
	Tags    []string
}

// NewBookmark converts to current Bookmark.
func (b *LegacyBookmark) NewBookmark() puka.Bookmark {
	return puka.Bookmark{
		Title:       b.Title,
		URL:         b.URL,
		Description: b.Comment,
		Timestamp:   b.Date,
		Tags:        b.Tags,
	}
}
