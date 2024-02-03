package movie

type Movies struct {
	Movies []Movie `json:"movies"`
}

type Movie struct {
	Title    string  `json:"title"`
	Released int     `json:"realized"`
	Genre    string  `json:"genre"`
	Casts    []Actor `json:"casts"`
	Duration int     `json:"duration"`
	Country  string  `json:"country"`
}

type Actor struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
}
