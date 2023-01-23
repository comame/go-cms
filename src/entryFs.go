package src

import (
	"fmt"
	"io"
	"os"
)

func GetEntriesJson() string {
	fp, err := os.Open("./entries/entries.json")
	if err != nil {
		// entries.json がないときは回復不能
		panic(err)
	}
	defer fp.Close()

	bytes, err := io.ReadAll(fp)
	if err != nil {
		// entries.json を読めないときは回復不能
		panic(err)
	}

	return string(bytes)
}

func GetEntryMarkdown(date Date, entry string, ext string) (string, error) {
	path := fmt.Sprintf("./entries/%d/%s.%s", date.Year, entry, ext)

	fp, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	bytes, err := io.ReadAll(fp)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}
