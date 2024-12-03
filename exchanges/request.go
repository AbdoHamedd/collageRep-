package exchanges

type CreateCourseRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TotalHours  int    `json:"totalHours"`
}

type UpdateCourseRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TotalHours  int    `json:"totalHours"`
}

type DeleteCourseRequest struct {
	ID uint `json:"id"`
}

type GetCourseByIdRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TotalHours  int    `json:"totalHours"`
}

type CreateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateDepartmentRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
