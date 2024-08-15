package cmd

import (
	"fmt"

	"github.com/anwam/spider/cmd/crawl"
	"github.com/spf13/cobra"
)

var (
	url     string
	urls    []string
	rootCmd = &cobra.Command{
		Use:   "go-scraper",
		Short: "Go Scraper is a CLI tool to crawl a website",
		Long:  `Go Scraper is a CLI tool to crawl a website and generate a JSON file with the results.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "URL to crawl")
	rootCmd.PersistentFlags().StringArrayVarP(&urls, "urls", "U", []string{}, "URLs to crawl")
	rootCmd.AddCommand(crawl.CrawlCmd)
	rootCmd.AddCommand(crawl.BulkCrawlCmd)
}

func initConfig() {
	fmt.Println("initConfig called")
}
