package main

//Config - root struct for config
type Config struct {
	Host        string   `json:"host"`
	Port        string   `json:"port"`
	Compression bool     `json:"compression"`
	Security    Security `json:"security"`
	Cache       Cache    `json:"caching"`
	Routes      []Route  `json:"routes"`
}

//Security - struct for security details
type Security struct {
	Enabled       string `json:"enabled"`
	CertPath      string `json:"certPath"`
	KeyPath       string `json:"keyPath"`
	MinTLSVersion string `json:"minTLSVersion"`
	MaxTLSVersion string `json:"maxTLSVersion"`
}

//Route - struct for routing details
type Route struct {
	Path        string  `json:"path"`
	Headers     Headers `json:"headers"`
	Target      string  `json:"target"`
	TargetType  string  `json:"targetType"`
	CacheExpiry int32   `json:"cacheExpiry"`
}

//Headers - struct for headers
type Headers struct {
	Request  map[string]string `json:"request"`
	Response map[string]string `json:"response"`
}

//Cache - struct for caching details
type Cache struct {
	Enabled  bool   `json:"enabled"`
	EndPoint string `json:"endPoint"`
}
