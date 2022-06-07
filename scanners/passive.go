package scanners

import (
	"activeReconBot/utils"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type HTMLParseError struct{}

func (m *HTMLParseError) Error() string {
	return "Can't parse html body"
}

const (
	SHODAN_API_RESOLVE_URL = "https://api.shodan.io/dns/resolve?hostnames=%s&key=%s"
	SHODAN_API_HOST_URL    = "https://api.shodan.io/shodan/host/%s?key=%s"
	SHODAN_APIKEY          = ""
)

func ResolveIpFromShodan(hostnames []string) []ShodanDNSResolveItem {
	var ips []ShodanDNSResolveItem
	var result map[string]string
	hostnamesChunks := utils.ChunkSlice(hostnames, 10)

	for _, chunk := range hostnamesChunks {
		hosts := strings.Join(chunk[:], ",")
		url := fmt.Sprintf(SHODAN_API_RESOLVE_URL, hosts, SHODAN_APIKEY)

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Accept", "application/json")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		result := map[string]string{}
		if err := json.Unmarshal(body, &result); err != nil {
			panic(err)
		}
	}

	for k, v := range result {
		item := ShodanDNSResolveItem{
			Hostname: k,
			IP:       v,
		}
		ips = append(ips, item)
	}

	return ips
}

func ResolveIpFromHostnameShodan(hostname string) []ShodanDNSResolveItem {
	var ips []ShodanDNSResolveItem
	var result map[string]string
	url := fmt.Sprintf(SHODAN_API_RESOLVE_URL, hostname, SHODAN_APIKEY)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	result = map[string]string{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}

	for k, v := range result {
		item := ShodanDNSResolveItem{
			Hostname: k,
			IP:       v,
		}
		ips = append(ips, item)
	}

	return ips
}

func GetShodanDetailsFromIPAddress(ip string) ShodanHostResult {
	url := fmt.Sprintf(SHODAN_API_HOST_URL, ip, SHODAN_APIKEY)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	result := ShodanHostResult{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}
	return result
}

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
