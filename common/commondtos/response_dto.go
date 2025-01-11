package commondtos

type ResponseDTO[T any] struct {
	Success    bool           `json:"success"`
	Data       []T            `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
}
