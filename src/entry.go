package src

import (
	"encoding/json"
	"errors"
	"regexp"
	"sort"
)

type Entry struct {
	Entry string   `json:"entry"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
	Type  string   `json:"type"`
}

func ParseEntries(s string) ([]Entry, error) {
	var entries []Entry
	err := json.Unmarshal([]byte(s), &entries)
	if err != nil {
		return nil, err
	}

	validated := true
	for _, entry := range entries {
		ok := ValidateEntry(entry)
		if !ok {
			validated = false
			break
		}
	}
	if !validated {
		return nil, errors.New("Invalid Format")
	}

	return entries, nil
}

func EntriesToString(entries []Entry) string {
	sort.SliceStable(entries, func(i, j int) bool {
		aDate, err := ParseDate(entries[i].Date)
		if err != nil {
			panic("invalid date format")
		}
		bDate, err := ParseDate(entries[j].Date)
		if err != nil {
			panic("invalid date format")
		}

		if aDate.Year < bDate.Year {
			return true
		}
		if aDate.Month < bDate.Month {
			return true
		}
		if aDate.Date < bDate.Date {
			return true
		}
		return false
	})

	str, _ := json.Marshal(entries)
	return string(str)
}

var acceptableTypes = []string{
	"html",
	"md",
}

func ValidateEntry(entry Entry) bool {
	if !isAcceptableType(entry.Type) {
		return false
	}

	if !isAcceptableEntry(entry.Entry) {
		return false
	}

	if !isAcceptableDate(entry.Date) {
		return false
	}

	return true
}

func isAcceptableType(s string) bool {
	for _, v := range acceptableTypes {
		if s == v {
			return true
		}
	}
	return false
}

func isAcceptableEntry(s string) bool {
	reg := regexp.MustCompile(`^[\w-_]+$`)
	return reg.MatchString(s)
}

func isAcceptableDate(s string) bool {
	_, err := ParseDate(s)
	if err != nil {
		return false
	}
	return true
}
