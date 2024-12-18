package courseExchabge

type CreateCourseRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	TotalHours   int    `json:"totalHours"`
	DepartmentID uint   `json:"departmentID"`
}

type UpdateCourseRequest struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	TotalHours   int    `json:"totalHours"`
	DepartmentID uint   `json:"departmentID"`
}

type IdRequest struct {
	ID uint `json:"id"`
}
