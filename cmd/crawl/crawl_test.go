package crawl

import (
	"testing"
)

func TestCrawl(t *testing.T) {
	// https://github.com/govcms-tests/govcms-cli/blob/3216b7334c184e2cd52f013f588d2af9529b4929/docs/testing_cobra.md
	url := "https://ryuucafe.com"
	if url == "" {
		t.Errorf("url is empty")
	}
}
