package domain

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UsersResult struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Result  []Users `json:"result"`
}
