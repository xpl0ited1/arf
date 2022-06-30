package scanners

import (
	"crypto/tls"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"time"
)

func GetPageTitle(hostname string) (string, error) {
	title := ""
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := http.Client{
		Transport: http.DefaultTransport.(*http.Transport),
		Timeout:   2 * time.Second,
	}
	url := "http://" + hostname

	res, err := client.Get(url)

	if err != nil {
		log.Println(err)
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err)
			}
		}(res.Body)
		title, _ = GetHtmlTitle(res.Body)
		if title != "" {
			url := "https://" + hostname

			res, err := client.Get(url)

			if err != nil {
				log.Println(err)
			} else {
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						log.Println(err)
					}
				}(res.Body)
				title, _ = GetHtmlTitle(res.Body)
			}

			if err != nil {
				log.Println(err)
			}

		}
	}

	if err != nil {
		log.Println(err)
	}

	return title, nil
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func GetHtmlTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Println("Fail to parse html")
	}
	return traverse(doc)
}
