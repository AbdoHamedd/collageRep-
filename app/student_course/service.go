package student_course

import (
	"basicCrudoperations/DataBase"
	"basicCrudoperations/exchanges/studenr_course"
	"basicCrudoperations/models"
)

func studentJoinCourseService(req studenr_course.StudentJoinCourseRequest) models.StusentCourse {
	studentJoinCourse := setJoinHelper(req)
	DataBase.DB.Create(&studentJoinCourse)
	return studentJoinCourse
}

func JoinDrobService(req studenr_course.StudentJoinCourseRequest) {
	DropJoin(req)
}
