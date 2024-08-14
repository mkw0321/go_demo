package main

import (
	"fmt"
)

type student struct {
	id    int
	name  string
	class string
}

// 是student的构造函数
func newStuedent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudents []*student
}

// 是studentMgr的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 添加学生 使用指针接收来增加学生切片
func (s *studentMgr) addStudent(newStudent *student) {
	s.allStudents = append(s.allStudents, newStudent)
}

// 编辑学生
func (s *studentMgr) editStudent(newStudent *student) {
	for i, v := range s.allStudents {
		if v.id == newStudent.id {
			s.allStudents[i] = newStudent
			return
		}
	}
	fmt.Println("newStudent not exist")
}

// 展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Print("学号:%d 姓名:%s \n", v.id, v.name)
	}
}
