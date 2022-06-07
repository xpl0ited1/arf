package models

import (
	"github.com/kamva/mgm/v3"
)

type Company struct {
	mgm.DefaultModel `bson:",inline"`
	CompanyName      string   `json:"company_name"`
	BountyUrl        string   `json:"bounty_url"`
	Domains          []Domain `json:"domains"`
}

type Domain struct {
	mgm.DefaultModel `bson:",inline"`
	CompanyID        string      `json:"company_id"`
	DomainName       string      `json:"domain_name"`
	Subdomains       []Subdomain `json:"subdomains"`
}

type Subdomain struct {
	mgm.DefaultModel `bson:",inline"`
	DomainID         string            `json:"domain_id"`
	SubdomainName    string            `json:"subdomain_name"`
	Ports            []TcpPort         `json:"ports"`
	PortScans        []PortScan        `json:"port_scans"`
	DirScans         []DirScan         `json:"dir_scans"` //FFUF or dirscanning
	Technologies     []Technology      `json:"technologies"`
	HTTPDetectedURLs []HTTPDetectedURL `json:"http_detected_urls"` //HTTPX
	WaybackUris      []WaybackUrl      `json:"wayback_uris"`
	HTTPTitle        string            `json:"http_title"`
}

type PortScan struct {
	mgm.DefaultModel     `bson:",inline"`
	Ports                []TcpPort    `json:"ports"`
	TechnologiesDetected []Technology `json:"technologies_detected"` //TODO: Validate if useless
}

type DirScan struct {
	mgm.DefaultModel `bson:",inline"`
	DirScanPaths     []DirScanPath `json:"dir_scan_paths"`
}

type DirScanPath struct {
	mgm.DefaultModel `bson:",inline"`
	URI              string `json:"uri"`
	IsDirectory      bool   `json:"is_directory"`
}

type TcpPort struct {
	PortNumber  int32  `json:"port_number"`
	ServiceName string `json:"service_name"`
	Content     string `json:"content"`
}

type Technology struct {
	TechnologyName string          `json:"technology_name"`
	Version        string          `json:"version"`
	PublicExploits []PublicExploit `json:"public_exploits"`
}

type PublicExploit struct {
	mgm.DefaultModel `bson:",inline"`
	TechnologyName   string `json:"technology_name"`
}

type WaybackUrl struct {
	mgm.DefaultModel `bson:",inline"`
	URL              string `json:"url"`
}

type HTTPDetectedURL struct {
	mgm.DefaultModel `bson:",inline"`
	URL              string `json:"url"`
	Title            string `json:"title"`
	//TODO: Screenshot
}

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username"`
	Password         string `json:"password"`
}

type ApiKey struct {
	mgm.DefaultModel `bson:",inline"`
	ApiKey           string `json:"api_key"`
}
