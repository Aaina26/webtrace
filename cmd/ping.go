/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

var flag_set bool

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping command is used to crawl a website and print a sitemap for it.",
	Long:  `ping first check if a website is accessible or not and then goes on to generate a sitemap for it.
	Example Usage: webtrace ping <url>
	
	You can also use it with a flag "--json" to export sitemap in json file.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//check for response
		input_url := args[0]
		response, response_time := get_request(input_url)
		fmt.Println("Status:", response.Status)
		fmt.Println("Response Time:", response_time)

		if response.StatusCode == 200 {
			base, err := url.Parse(input_url)
			if err != nil {
				fmt.Println("Invalid url:", err)
			}

			root := &Node{Name: input_url, Children: make(map[string]*Node)}
			crawl(input_url, base, root)
			print_site_map(root, 0)

			if flag_set {
				err = writeJSONToFile(root, "sitemap.json")
				if err != nil {
					fmt.Println("Error writing sitemap to file:", err)
				} else {
					fmt.Println("Sitemap saved to sitemap.json")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pingCmd.Flags().BoolVar(&flag_set, "json", false, "To export sitemap as a json file")
}

// this function is main function that return htt
func get_request(input_url string) (resp *http.Response, t time.Duration) {
	start := time.Now()
	resp, err := http.Get(input_url)
	if err != nil {
		fmt.Println("Invalid URL:", err)
	}
	duration := time.Since(start)
	return resp, duration
}

type Node struct {
	Name     string
	Children map[string]*Node
}

var visited = make(map[string]bool)

func crawl(link string, base *url.URL, root *Node) {
	if visited[link] {
		return
	}
	visited[link] = true

	resp, err := http.Get(link)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	links := extract_links(doc, base)
	for _, href := range links {
		if visited[href] {
			continue
		}
		u, _ := url.Parse(href)
		add_to_tree(root, u.Path)
		crawl(href, base, root)
	}
}

func extract_links(n *html.Node, base *url.URL) []string {
	var links []string
	var visit func(*html.Node)

	visit = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := attr.Val
					u, err := base.Parse(href)
					if err != nil {
						continue
					}
					if u.Host == base.Host {
						normalized := u.Scheme + "://" + u.Host + u.Path
						links = append(links, normalized)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(n)
	return links
}

// Add URL path segments to tree
func add_to_tree(root *Node, path string) {
	if path == "" || path == "/" {
		return
	}
	segments := strings.Split(strings.Trim(path, "/"), "/")
	current := root
	for _, seg := range segments {
		if current.Children[seg] == nil {
			current.Children[seg] = &Node{Name: seg, Children: make(map[string]*Node)}
		}
		current = current.Children[seg]
	}
}

// Print tree-style site map
func print_site_map(n *Node, indent int) {
	prefix := strings.Repeat("│   ", indent)
	if indent == 0 {
		fmt.Println("Site Map:")
		fmt.Printf("- %s\n", n.Name)
	} else {
		fmt.Printf("%s├── %s\n", prefix, n.Name)
	}

	i := 0
	for _, child := range sorted_keys(n.Children) {
		print_site_map(n.Children[child], indent+1)
		i++
	}
}

// Keep consistent ordering of children (alphabetical)
func sorted_keys(m map[string]*Node) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func writeJSONToFile(n *Node, filename string) error {
	sitemapJSON, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(sitemapJSON)
	if err != nil {
		return err
	}

	return nil
}
