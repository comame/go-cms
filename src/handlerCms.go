package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func HandleCms(w http.ResponseWriter, r *http.Request) bool {
	if reg := regexp.MustCompile(`^/api/entries/?$`); reg.MatchString(r.URL.Path) {
		data, err := GetAllEntry()
		if err != nil {
			res, _ := json.Marshal(ErrorResponse{
				Message: "InternalError",
				Code:    InternalError.String(),
			})
			http.Error(w, string(res), http.StatusInternalServerError)
			return true
		}
		res, _ := json.Marshal(Response{
			Data: data,
		})
		fmt.Fprint(w, string(res))
		return true
	}

	if reg := regexp.MustCompile(`^/api/entries/tag/(\w+)/?$`); reg.MatchString(r.URL.Path) {
		match := reg.FindAllStringSubmatch(r.URL.Path, -1)
		tag := match[0][1]

		data, err := GetEntriesByTag(tag)
		if err != nil {
			res, _ := json.Marshal(ErrorResponse{
				Message: "InternalError",
				Code:    InternalError.String(),
			})
			http.Error(w, string(res), http.StatusInternalServerError)
			return true
		}
		res, _ := json.Marshal(Response{
			Data: data,
		})
		fmt.Fprint(w, string(res))
		return true
	}

	if reg := regexp.MustCompile(`^/api/entries/year/(\d+)/?$`); reg.MatchString(r.URL.Path) {
		match := reg.FindAllStringSubmatch(r.URL.Path, -1)
		yearStr, _ := strconv.ParseUint(match[0][1], 10, 64)

		data, err := GetEntryByYear(uint(yearStr))
		if err != nil {
			res, _ := json.Marshal(ErrorResponse{
				Message: "InternalError",
				Code:    InternalError.String(),
			})
			http.Error(w, string(res), http.StatusInternalServerError)
			return true
		}
		res, _ := json.Marshal(Response{
			Data: data,
		})
		fmt.Fprint(w, string(res))
		return true
	}

	if reg := regexp.MustCompile(`^/api/entries/(\d{4}-\d{2}-\d{2})/([\w-_]+)\.(\w+)$`); reg.MatchString(r.URL.Path) {
		match := reg.FindAllStringSubmatch(r.URL.Path, -1)
		date := match[0][1]
		entryName := match[0][2]
		ext := match[0][3]

		entry, text, err := GetEntry(date, entryName, ext)
		if err != nil {
			res, _ := json.Marshal(ErrorResponse{
				Message: "Not Found",
				Code:    NotFound.String(),
			})
			http.Error(w, string(res), http.StatusNotFound)
			return true
		}
		res, _ := json.Marshal(Response{
			Data: map[string]any{
				"entry": entry,
				"text":  text,
			},
		})
		fmt.Fprint(w, string(res))
		return true
	}

	if reg := regexp.MustCompile(`^/api/entries/years/?$`); reg.MatchString(r.URL.Path) {
		data, err := GetYears()
		if err != nil {
			res, _ := json.Marshal(ErrorResponse{
				Message: "InternalError",
				Code:    InternalError.String(),
			})
			http.Error(w, string(res), http.StatusInternalServerError)
			return true
		}
		res, _ := json.Marshal(Response{
			Data: data,
		})
		fmt.Fprint(w, string(res))
		return true
	}

	if reg := regexp.MustCompile(`^/api/entries/tags/?$`); reg.MatchString(r.URL.Path) {
		data, err := GetTags()
		if err != nil {
			res, _ := json.Marshal(ErrorResponse{
				Message: "InternalError",
				Code:    InternalError.String(),
			})
			http.Error(w, string(res), http.StatusInternalServerError)
			return true
		}
		res, _ := json.Marshal(Response{
			Data: data,
		})
		fmt.Fprint(w, string(res))
		return true
	}

	if reg := regexp.MustCompile(`^/api/feed.xml$`); reg.MatchString(r.URL.Path) {
		r.Header.Set("content-type", "application/xml")
		data, err := GetFeed(HOSTNAME)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return true
		}
		fmt.Fprint(w, data)
		return true
	}

	if reg := regexp.MustCompile(`^/api/sitemap.txt$`); reg.MatchString(r.URL.Path) {
		r.Header.Set("content-type", "text/plain")
		data, err := GetSitemap(HOSTNAME)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return true
		}
		fmt.Fprint(w, data)
		return true
	}

	return false
}
