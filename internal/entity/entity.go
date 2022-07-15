package entity

type Ipcheck struct {
	Id     uint8
	Office string `json:"office"`
	Ip     string `json:"ip"`
	City   string `json:"city"`
	Server string `json:"server"`
}

type SshLog struct {
	Id   uint32
	Time string `json:"time"`
	Ip   string `json:"ip"`
	Text string `json:"text"`
}

type CheckNetLog struct {
	Id     uint64
	Time   string `json:"time"`
	Office string `json:"office"`
	Ip     string `json:"ip"`
	City   string `json:"city"`
	Server string `json:"server"`
	Text   string `json:"text"`
}

const (
	UserSql = "usersql"
	PassSql = "Nomu8@RAmBat"
	HostSql = "10.101.2.194:3306"
)
