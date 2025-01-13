package student_course

import (
	"basicCrudoperations/exchanges/studenr_course"
	"basicCrudoperations/models"
)

func setJoinHelper(req studenr_course.StudentJoinCourseRequest) models.StusentCourse {
	return models.StusentCourse{
		CourseId: req.CourseID,
		UserId:   req.UserID,
	}
}
