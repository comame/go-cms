package src

import (
	"errors"
	"sort"

	"github.com/russross/blackfriday/v2"
)

var entriesCache []Entry

func GetAllEntry() ([]Entry, error) {
	if entriesCache != nil {
		return entriesCache, nil
	}

	json := GetEntriesJson()
	entries, err := ParseEntries(json)
	if err != nil {
		return nil, err
	}

	entriesCache = entries
	return entries, nil
}

func GetEntriesByTag(tag string) ([]Entry, error) {
	entries, err := GetAllEntry()
	if err != nil {
		return nil, err
	}

	filtered := make([]Entry, 0)
	for _, v := range entries {
		if containsTag(tag, v) {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}

func GetEntryByYear(year uint) ([]Entry, error) {
	entries, err := GetAllEntry()
	if err != nil {
		return nil, err
	}

	filtered := make([]Entry, 0)
	for _, entry := range entries {
		// バリデーション済みなので OK
		date, _ := ParseDate(entry.Date)
		if date.Year == year {
			filtered = append(filtered, entry)
		}
	}

	return filtered, nil
}

func GetYears() ([]uint, error) {
	entries, err := GetAllEntry()
	if err != nil {
		return nil, err
	}

	yearsSet := make(map[uint]struct{}, 0)
	for _, entry := range entries {
		// バリデーション済みなので OK
		date, _ := ParseDate(entry.Date)
		yearsSet[date.Year] = struct{}{}
	}

	years := make([]uint, 0)
	for k := range yearsSet {
		years = append(years, k)
	}

	sort.Slice(years, func(i, j int) bool {
		return years[i] < years[j]
	})

	return years, nil
}

func GetTags() ([]string, error) {
	entries, err := GetAllEntry()
	if err != nil {
		return nil, err
	}

	tagsSet := make(map[string]struct{}, 0)
	for _, entry := range entries {
		for _, tag := range entry.Tags {
			tagsSet[tag] = struct{}{}
		}
	}

	tags := make([]string, 0)
	for k := range tagsSet {
		tags = append(tags, k)
	}

	sort.SliceStable(tags, func(i, j int) bool {
		return tags[i] < tags[j]
	})

	return tags, nil
}

func GetEntry(dateStr string, entry string, ext string) (*Entry, string, error) {
	entries, err := GetAllEntry()
	if err != nil {
		return nil, "", err
	}

	var target Entry
	targetFound := false
	for _, v := range entries {
		if v.Date == dateStr && v.Entry == entry && v.Type == ext {
			target = v
			targetFound = true
			break
		}
	}
	if !targetFound {
		return nil, "", errors.New("Not Found")
	}

	// バリデーション済みなので OK
	date, _ := ParseDate(target.Date)
	text, err := GetEntryMarkdown(*date, target.Entry, target.Type)
	if err != nil {
		return nil, "", err
	}

	return &target, text, err
}

func GetEntryHtml(dateStr string, entry string, ext string) (string, error) {
	_, content, err := GetEntry(dateStr, entry, ext)
	if err != nil {
		return "", err
	}

	if ext == "html" {
		return content, nil
	}

	if ext != "md" {
		return "", errors.New("Unknown type")
	}

	html := blackfriday.Run([]byte(content))
	return string(html), nil
}

func containsTag(tag string, entry Entry) bool {
	for _, v := range entry.Tags {
		if tag == v {
			return true
		}
	}
	return false
}
