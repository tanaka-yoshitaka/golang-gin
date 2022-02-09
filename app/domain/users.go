package domain

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string
}

type UsersResult struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Result  []Users `json:"result"`
}
