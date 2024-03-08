package models

import "net/http"

const (
	DB_INSERT_FAIL = "데이터 저장 오류"
	DB_UPDATE_FAIL = "데이터 갱신 오류"
	DB_DELETE_FAIL = "데이터 삭제 오류"

	DB_NO_CONTENT = "조회 내역 없음"
	DB_CONFLICT   = "상태를 변경할 수 없음"
)

type MultiBookResult struct {
	Data          []Book `json:"result"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

func (r *MultiBookResult) SetResult(data []Book, resultCode int, resultMessage string) {
	r.Data = data
	r.ResultCode = resultCode
	r.ResultMessage = resultMessage
}

type SingleBookResult struct {
	Data          Book   `json:"result"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

func (r *SingleBookResult) SetResult(data Book, resultCode int, resultMessage string) {
	r.Data = data
	r.ResultCode = resultCode
	r.ResultMessage = resultMessage
}

type OpResult struct {
	Data          BookOp `json:"result"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

func (r *OpResult) SetResult(data BookOp, resultCode int, resultMessage string) {
	r.Data = data
	r.ResultCode = resultCode
	r.ResultMessage = resultMessage
}

func CheckErr(err error) int {
	if err != nil {
		// todo -  err 값에 따른 에러 코드 세분화  필요
		return http.StatusInternalServerError
	} else {
		return http.StatusOK
	}
}

func CheckResult(nRow int64, dbResultCode int) int {
	if nRow != int64(0) {
		return http.StatusOK
	} else {
		return dbResultCode
	}
}
