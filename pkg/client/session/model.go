package session

type Session struct {
	Id       string           `json:"id"`
	Image    string           `json:"image"`
	Location string           `json:"location"`
	Host     string           `json:"host"`
	Created  int64            `json:"created"`
	Port     map[string]int64 `json:"port"`
}
