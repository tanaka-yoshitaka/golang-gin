package usecase

type ResultStatus struct {
	StatusCode int
	Error      error
}

func NewResultStatus(status int, err error) *ResultStatus {
	rs := new(ResultStatus)
	rs.StatusCode = status
	rs.Error = err
	return rs
}
