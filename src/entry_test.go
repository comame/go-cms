package src

import (
	"testing"
)

func TestParseEntries(t *testing.T) {
	json := `[{
        "entry": "safari-bookmark-import",
        "title": "Safari のブックマークを Chrome にインポートできない問題",
        "date": "2019-04-22",
        "tags": [
            "troubleshoot",
            "Chrome"
        ],
        "type": "html"
    },
    {
        "entry": "improve-web-performance",
        "title": "Web ページの読み込みパフォーマンスを改善する",
        "date": "2019-04-23",
        "tags": [
            "Web"
        ],
        "type": "md"
    }]`

	_, err := ParseEntries(json)
	if err != nil {
		t.FailNow()
	}
}
