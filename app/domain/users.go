package domain

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required,min=1,max=100,alphanum"`
	Password string `json:"-" binding:"required,min=1,max=10,alphanum"`
}

type UsersResult struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Result  []Users `json:"result"`
}
