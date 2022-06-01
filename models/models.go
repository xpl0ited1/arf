package models

import (
	"github.com/kamva/mgm/v3"
	"time"
)

type Company struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	CompanyName      string    `json:"company_name"`
	BountyUrl        string    `json:"bounty_url"`
	Domains          []Domain  `json:"domains"`
}

type Domain struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time   `bson:"created_at"`
	UpdatedAt        time.Time   `bson:"updated_at"`
	DomainName       string      `json:"domain_name"`
	Subdomains       []Subdomain `json:"subdomains"`
}

type Subdomain struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time         `bson:"created_at"`
	UpdatedAt        time.Time         `bson:"updated_at"`
	SubdomainName    string            `json:"subdomain_name"`
	Ports            []TcpPort         `json:"ports"`
	PortScans        []PortScan        `json:"port_scans"`
	DirScans         []DirScan         `json:"dir_scans"` //FFUF or dirscanning
	Technologies     []Technology      `json:"technologies"`
	HTTPDetectedURLs []HTTPDetectedURL `json:"http_detected_urls"` //HTTPX
	WaybackUris      []WaybackUrl      `json:"wayback_uris"`
}

type PortScan struct {
	mgm.DefaultModel     `bson:",inline"`
	CreatedAt            time.Time    `bson:"created_at"`
	UpdatedAt            time.Time    `bson:"updated_at"`
	Ports                []TcpPort    `json:"ports"`
	TechnologiesDetected []Technology `json:"technologies_detected"` //TODO: Validate if useless
}

type DirScan struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time     `bson:"created_at"`
	UpdatedAt        time.Time     `bson:"updated_at"`
	DirScanPaths     []DirScanPath `json:"dir_scan_paths"`
}

type DirScanPath struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	URI              string    `json:"uri"`
	IsDirectory      bool      `json:"is_directory"`
}

type TcpPort struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	PortNumber       int32     `json:"port_number"`
	ServiceName      string    `json:"service_name"`
	Content          string    `json:"content"`
}

type Technology struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time       `bson:"created_at"`
	UpdatedAt        time.Time       `bson:"updated_at"`
	TechnologyName   string          `json:"technology_name"`
	Version          string          `json:"version"`
	PublicExploits   []PublicExploit `json:"public_exploits"`
}

type PublicExploit struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	TechnologyName   string    `json:"technology_name"`
}

type WaybackUrl struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	URL              string    `json:"url"`
}

type HTTPDetectedURL struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	URL              string    `json:"url"`
	Title            string    `json:"title"`
	//TODO: Screenshot
}
