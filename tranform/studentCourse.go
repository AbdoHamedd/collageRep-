package tranform

import (
	"basicCrudoperations/exchanges/studenr_course"
	"basicCrudoperations/models"
)

func TransformJoin(studentCourse *models.StusentCourse) studenr_course.JoinResponse {
	return studenr_course.JoinResponse{
		CourseID: studentCourse.CourseId,
		UserID:   studentCourse.UserId,
	}

}
