package model

type Employee struct {
	ID       int64   `json:"id" example:"1"`
	Name     string  `json:"name" example:"malik"`
	Email    string  `json:"email" example:"malik@gmail.com"`
	Position string  `json:"position" example:"intern full stack developer"`
	Salary   float64 `json:"salary" example:"3000000"`
	// CreatedAt time.Time	`json:"createdAt"`
}
type UpdatedEmployee struct {
	ID       int64   `json:"id" example:"1"`
	Name     string  `json:"name" example:"malik"`
	Email    string  `json:"email" example:"malik@gmail.com"`
	Position string  `json:"position" example:"intern backend developer"`
	Salary   float64 `json:"salary" example:"2000000"`
	// CreatedAt time.Time	`json:"createdAt"`
}