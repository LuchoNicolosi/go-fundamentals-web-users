package domain

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func GetAll() ([]User, error) {
	return []User{
		{
			ID:        1,
			FirstName: "Lucho",
			LastName:  "Nicolosi",
			Email:     "nicolosi@gmail.com",
		},
		{
			ID:        2,
			FirstName: "tomcat",
			LastName:  "perito",
			Email:     "perito@gmail.com",
		},
		{
			ID:        3,
			FirstName: "zaraza",
			LastName:  "sumario",
			Email:     "sumario@gmail.com",
		},
	}, nil
}
