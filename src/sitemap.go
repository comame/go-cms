package src

import (
	"fmt"
	"strings"
)

func GetSitemap(hostname string) (string, error) {
	entries, err := GetAllEntry()
	if err != nil {
		return "", nil
	}

	entryUrls := make([]string, 0)
	for _, entry := range entries {
		entryUrls = append(
			entryUrls,
			fmt.Sprintf("https://%s/entries/%s/%s.%s", hostname, entry.Date, entry.Entry, entry.Type),
		)
	}

	tags, err := GetTags()
	if err != nil {
		return "", nil
	}

	tagUrls := make([]string, 0)
	for _, tag := range tags {
		tagUrls = append(
			tagUrls,
			fmt.Sprintf("https://%s/entries/tags/%s", hostname, tag),
		)
	}

	entryUrlsStr := strings.Join(entryUrls, "\n")
	tagUrlsStr := strings.Join(tagUrls, "\n")
	return entryUrlsStr + "\n" + tagUrlsStr, nil
}
