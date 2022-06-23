package entity

type Ipcheck struct {
	Id     uint8
	Office string `json:"office"`
	Ip     string `json:"ip"`
	City   string `json:"city"`
	Server string `json:"server"`
}

type SshLog struct {
	Id     uint8
	Time   string `json:"time"`
	Ip     string `json:"ip"`
	Text   string `json:"text"`
	Server string `json:"server"`
}
