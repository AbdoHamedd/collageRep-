package shared

import (
	"basicCrudoperations/exchanges/courseExchabge"
	"basicCrudoperations/exchanges/deptexchange"
	"time"
)

type CourseResponse struct {
	ID          uint                            `json:"id"`
	CreatedAt   time.Time                       `json:"createdAt"`
	UpdatedAt   time.Time                       `json:"updatedAt"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	TotalHours  int                             `json:"totalHours"`
	Department  deptexchange.DepartmentResponse `json:"department"`
}

type DepartmentResponseForDept struct {
	ID          uint                               `json:"id"`
	CreatedAt   time.Time                          `json:"createdAt"`
	UpdatedAt   time.Time                          `json:"updatedAt"`
	Name        string                             `json:"name"`
	Description string                             `json:"description"`
	Courses     []courseExchabge.CourseResoForDept `json:"courses"`
}
