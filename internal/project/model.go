package project

type Project struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	ApiKeys   []string `json:"api_keys"`
	UserId    string   `json:"user_id"`
	CreatedAt string   `json:"created_at"`
}
