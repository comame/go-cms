package src

import (
	"fmt"
	"strings"
	"time"
)

var feedCache *string = nil

func GetFeed(hostname string) (string, error) {
	if feedCache != nil {
		return *feedCache, nil
	}

	entries, err := GetAllEntry()
	if err != nil {
		return "", nil
	}

	items := make([]string, 0)
	for _, entry := range entries {
		_, content, err := GetEntry(entry.Date, entry.Entry, entry.Type)
		if err != nil {
			return "", err
		}
		// TODO: パースする必要あり
		items = append(items, item(
			entry.Title,
			fmt.Sprintf("https://%s/entries/%s/%s.%s", hostname, entry.Date, entry.Date, entry.Type),
			entry.Date,
			content,
		))
	}

	now := time.Now().Format("1970-01-01")
	return base(hostname, now, items), nil
}

func base(hostname, updated string, items []string) string {
	joined := strings.Join(items, "")
	return fmt.Sprintf(`<?xml version='1.0'?>
<feed xmlns='http://www.w3.org/2005/Atom'>
<id>https://%s/</id>
<title>%s</title>
<link rel='alternate' href='https://%s/' />
<link rel='self' href='https://%s/feed.xml' />
<author><name>comame</name></author>
<updated>%sT00:00:00Z</updated>
%s
</feed>`, hostname, hostname, hostname, hostname, updated, joined)
}

func item(title, link, date, content string) string {
	return fmt.Sprintf(`<entry>
<title>%s</title>
<link rel='alternate' href='%s' />
<id>%s</id>
<updated>%sT00:00:00Z</updated>
<summary>%s</summary>
</entry>`, title, link, link, date, escape(content))
}

func escape(text string) string {
	text = strings.ReplaceAll(text, "&", "&amp;")
	text = strings.ReplaceAll(text, "'", "&apos;")
	text = strings.ReplaceAll(text, "&", "&amp;")
	text = strings.ReplaceAll(text, "<", "&lt;")
	text = strings.ReplaceAll(text, ">", "&gt;")
	return text
}
