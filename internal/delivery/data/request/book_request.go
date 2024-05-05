package request

type CreateBookReq struct {
	Title  string `validate:"required" json:"title"`
	Author string `validate:"required" json:"author"`
	Year   int    `validate:"required" json:"year"`
}

type UpdateBookReq struct {
	UUID   string `validate:"required" json:"uuid"`
	Title  string `validate:"required" json:"title"`
	Author string `validate:"required" json:"author"`
	Year   int    `validate:"required" json:"year"`
}
