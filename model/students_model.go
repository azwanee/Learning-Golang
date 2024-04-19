package model

type Student struct {
	StudentID int64  `json:"student_id"`
	Name      string `json:"name" validate:"required"`
	Age       string `json:"age" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Hoby      string `json:"hoby" validate:"required"`
}
