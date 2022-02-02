package usecase

type ResultStatus struct {
	StatusCode int
	Error      error
}

func NewResultStatus(status int, err error) *ResultStatus {
	resultStatus := new(ResultStatus)
	resultStatus.StatusCode = status
	resultStatus.Error = err
	return resultStatus
}
