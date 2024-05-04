package bookRes

type BookRes struct {
	Id     int    `json:"id"`
	UUID   string `json:"uuid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
