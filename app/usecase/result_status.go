package usecase

type ResultStatus struct {
	Status int
	Error  error
}

func NewResultStatus(status int, err error) *ResultStatus {
	resultStatus := new(ResultStatus)
	resultStatus.Status = status
	resultStatus.Error = err
	return resultStatus
}
