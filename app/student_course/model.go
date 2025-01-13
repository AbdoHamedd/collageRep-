package student_course

import (
	"basicCrudoperations/DataBase"
	"basicCrudoperations/app/user"
	"basicCrudoperations/exchanges/studenr_course"
	"basicCrudoperations/models"
)

func userIsStudent(id uint) bool {
	User := user.GetUserById(id)
	if User.Role != "student" {
		return false
	}
	return true
}

func joinValidate(req studenr_course.StudentJoinCourseRequest) bool {
	var studentCourse models.StusentCourse
	DataBase.DB.Where("course_id = ? AND user_id = ?", req.CourseID, req.UserID).First(&studentCourse)
	if studentCourse.ID != 0 {
		return false
	}
	return true
}

func DropJoin(req studenr_course.StudentJoinCourseRequest) {
	DataBase.DB.Where("user_id = ? AND course_id = ?", req.UserID, req.CourseID).Unscoped().Delete(&models.StusentCourse{})
}
