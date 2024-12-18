package deptexchange

type CreateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateDepartmentRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetByIdRequest struct {
	ID uint `json:"id"`
}
