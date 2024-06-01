package web

type NewsResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Content string `json:"content"`
}
