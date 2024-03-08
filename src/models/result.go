package models

type SearchBookResult struct {
	Result     []Book `json:"books"`
	ResultCode int    `json:"resultCode"`
}
