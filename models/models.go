package models

import (
	"github.com/kamva/mgm/v3"
	"time"
)

type Company struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	CompanyName      string    `bson:"company_name"`
	BountyUrl        string    `json:"bounty_url"`
	Domains          []Domain  `bson:"domains"`
}

type Domain struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time   `bson:"created_at"`
	UpdatedAt        time.Time   `bson:"updated_at"`
	DomainName       string      `bson:"domain_name"`
	Subdomains       []Subdomain `bson:"subdomains"`
}

type Subdomain struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time         `bson:"created_at"`
	UpdatedAt        time.Time         `bson:"updated_at"`
	SubdomainName    string            `bson:"subdomain_name"`
	Ports            []TcpPort         `bson:"ports"`
	PortScans        []PortScan        `bson:"port_scans"`
	DirScans         []DirScan         `bson:"dir_scans"` //FFUF or dirscanning
	Technologies     []Technology      `bson:"technologies"`
	HTTPDetectedURLs []HTTPDetectedURL `bson:"http_detected_urls"` //HTTPX
	WaybackUris      []WaybackUrl      `bson:"wayback_uris"`
}

type PortScan struct {
	mgm.DefaultModel     `bson:",inline"`
	CreatedAt            time.Time    `bson:"created_at"`
	UpdatedAt            time.Time    `bson:"updated_at"`
	Ports                []TcpPort    `bson:"ports"`
	TechnologiesDetected []Technology `bson:"technologies_detected"` //TODO: Validate if useless
}

type DirScan struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time     `bson:"created_at"`
	UpdatedAt        time.Time     `bson:"updated_at"`
	DirScanPaths     []DirScanPath `bson:"dir_scan_paths"`
}

type DirScanPath struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	URI              string    `bson:"uri"`
	IsDirectory      bool      `bson:"is_directory"`
}

type TcpPort struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	PortNumber       int32     `bson:"port_number"`
	ServiceName      string    `bson:"service_name"`
	Content          string    `bson:"content"`
}

type Technology struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time       `bson:"created_at"`
	UpdatedAt        time.Time       `bson:"updated_at"`
	TechnologyName   string          `bson:"technology_name"`
	Version          string          `bson:"version"`
	PublicExploits   []PublicExploit `bson:"public_exploits"`
}

type PublicExploit struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	TechnologyName   string    `bson:"technology_name"`
}

type WaybackUrl struct {
	mgm.DefaultModel `bson:",inline"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	URL              string    `bson:"url"`
}

type HTTPDetectedURL struct {
	URL   string `bson:"url"`
	Title string `bson:"title"`
	//TODO: Screenshot
}
