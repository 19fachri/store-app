package servicemodels

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Pagination struct {
	Data      interface{} `json:"data"`
	TotalPage uint        `json:"total_page"`
	NextPage  bool        `json:"next_page"`
}
