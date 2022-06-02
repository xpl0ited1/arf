package scanners

type ShodanDNSResolveResult struct {
	Items []ShodanDNSResolveItem
}

type ShodanDNSResolveItem struct {
	Hostname string
	IP       string
}

type ShodanHostResult struct {
	RegionCode  string
	CountryCode string
	City        string
	LastUpdate  string
	Latitude    float64
	CountryName string
	Org         string
	Data        []ShodanHostResultData
	ASN         string
	ISP         string
	Longitude   float64
	OS          string
	Ports       []int32
}

type ShodanHostResultData struct {
	OS        string
	ISP       string
	Port      int32
	Data      string
	Org       string
	ASN       string
	Transport string
	Product   string
	Tags      []string
	Version   string
	SSH       ShodanHostResultDataSSHItem
	CPE32     []string
	Info      string
	CPE       []string
	HTTP      ShodanHostResultDataHTTPItem
}

type ShodanHostResultDataHTTPItem struct {
	Status      int32
	SecurityTxt string
	Title       string
	Robots      string
	Server      string
	HeadersHash int32
	HTML        string
	Location    string
	WAF         string
	Components  map[string]map[string][]string //"CFML": {	"categories": [	"Programming languages"]}
}

type ShodanHostResultDataSSHItem struct {
	Hassh       string
	Fingerprint string
	Mac         string
	Cipher      string
	Key         string
	Kex         ShodanHostResultDataSSHItemKex
	Type        string
}

type ShodanHostResultDataSSHItemKex struct {
	Languages               []string
	ServerHostKeyAlgorithms []string
	EncryptionAlgorithms    []string
	KexFollows              bool
	Unused                  int32
	KexAlgorithms           []string
	CompressionAlgorithms   []string
	MacAlgorithms           []string
}
