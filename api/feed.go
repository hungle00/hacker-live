package api

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type RSSFeed struct {
	XMLName xml.Name  `xml:"rss"`
	Items   []RSSItem `xml:"channel>item"`
}

type RSSItem struct {
	Title       string `xml:"title" json:"title"`
	Description string `xml:"description" json:"description"`
	PubDate     string `xml:"pubDate" json:"pubDate"`
	Creator     string `xml:"creator" json:"creator"`
	Link        string `xml:"link" json:"link"`
}

func formatDate(date string) string {
	t, err := time.Parse(time.RFC1123Z, date)
	if err != nil {
		return date
	}
	return t.Format("Jan 2, 2006 3:04 PM")
}

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://hnrss.org/newcomments")
	if err != nil {
		http.Error(w, "Failed to fetch RSS feed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var feed RSSFeed
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		http.Error(w, "Failed to parse RSS feed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	for _, item := range feed.Items {
		fmt.Fprintf(w, `
			<div class="feed-item">
				<h3><a href="%s">%s</a></h3>
				<div class="meta">
					by %s • %s
				</div>
				<div class="description">%s</div>
			</div>
		`, item.Link,
			template.HTMLEscapeString(item.Title),
			template.HTMLEscapeString(item.Creator),
			formatDate(item.PubDate),
			item.Description)
	}
}
