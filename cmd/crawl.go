package cmd

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

var (
	c        *colly.Collector
	crawlCmd = &cobra.Command{
		Use:   "crawl",
		Short: "Crawl a website",
		Long:  "Crawl a website and generate JSON file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Start crawling...")
			if url == "" {
				fmt.Println("URL is required. use --url or -u flag")
				return
			}
			c = colly.NewCollector()
			c.OnHTML("html", func(e *colly.HTMLElement) {
				fmt.Println(e.Text)
			})

			c.OnRequest(func(r *colly.Request) {
				fmt.Println("Visiting", r.URL.String())
			})

			c.Visit(url)
		},
	}

	bulkCrawlCmd = &cobra.Command{
		Use:   "bulk-crawl",
		Short: "Bulk crawl websites",
		Long:  "Bulk crawl websites and generate JSON file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(urls) == 0 {
				fmt.Println("URLs are required. use --urls or -U flag")
				return
			}

			count := len(urls)
			crawlWebsite := func(url string, resp chan string, w *sync.WaitGroup) {
				c = colly.NewCollector()
				c.OnHTML("h1", func(e *colly.HTMLElement) {
					defer w.Done()
					resp <- e.Text
				})

				c.OnRequest(func(r *colly.Request) {
					fmt.Println("Visiting", count)
					count--
				})

				c.Visit(url)
			}

			wg := sync.WaitGroup{}
			resp := make(chan string, len(urls))

			for _, u := range urls {
				wg.Add(1)
				go crawlWebsite(u, resp, &wg)
			}

			// Wait for all goroutines to finish
			wg.Wait()
			// Close the channel to avoid deadlock
			close(resp)

			for r := range resp {
				fmt.Println(r)
			}
		},
	}
)
