package vo

type ConnectVO struct {
	Url      string `json:"url"`
	Db       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
}
