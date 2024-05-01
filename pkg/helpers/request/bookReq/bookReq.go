package bookReq

type CreateBookReq struct {
	Title  string `validate:"required,min=1,max=200" json:"name"`
	Author string `validate:"required,min=1,max=200" json:"author"`
	Year   int    `validate:"required" json:"year"`
}

type UpdateBookReq struct {
	UUID   string `validate:"required" json:"uuid"`
	Title  string `validate:"required,min=1,max=200" json:"name"`
	Author string `validate:"required,min=1,max=200" json:"author"`
	Year   int    `validate:"required" json:"year"`
}
