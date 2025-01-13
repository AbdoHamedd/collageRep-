package studenr_course

type StudentJoinCourseRequest struct {
	UserID   uint `json:"user_id" `
	CourseID uint `json:"course_id"`
}
