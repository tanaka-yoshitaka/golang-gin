package domain

type Users struct {
	ID   int
	Name string
}

type UsersForGet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *Users) BuildForGet() UsersForGet {
	user := UsersForGet{}
	user.ID = u.ID
	user.Name = u.Name

	return user
}
