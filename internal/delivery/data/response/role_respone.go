package response

type RoleResponse struct {
	Id    int    `json:"id"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Value string `json:"value"`
}
