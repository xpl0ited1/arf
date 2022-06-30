package scanners

import (
	"activeReconBot/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
