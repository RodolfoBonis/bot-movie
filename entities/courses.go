package entities

type CoursesData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Links       []struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"links"`
}
