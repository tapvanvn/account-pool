package entity

type ConfigDocDB struct {
	Provider         string `json:"Provider"`
	ConnectionString string `json:"ConnectionString"`
	DatabaseName     string `json:"DatabaseName"`
}

type Config struct {
	Environment string       `json:"Environment"`
	DocDB       *ConfigDocDB `json:"DocDB"`
}
