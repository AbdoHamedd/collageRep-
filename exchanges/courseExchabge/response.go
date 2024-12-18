package courseExchabge

import "time"

type CourseResoForDept struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TotalHours  int       `json:"totalHours"`
}
