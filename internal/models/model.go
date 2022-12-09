package models

type Student struct {
	ID                 int    `json:"id"`
	TutorID            int    `json:"tutor_id"`
	Name               string `json:"name"`
	Surname            string `json:"surname"`
	PasswordHash       string `json:"password_hash"`
	Salt               string `json:"salt"`
	ScientificDirector string `json:"scientific_director"`
	CourseWorkSubject  string `json:"course_work_subject"`
	DiplomaWorkSubject string `json:"diploma_work_subject"`
	Course             int    `json:"course"`
}

type Tutor struct {
	ID                  int      `json:"id"`
	Name                string   `json:"name"`
	Surname             string   `json:"surname"`
	Salt                string   `json:"salt"`
	PasswordHash        string   `json:"password_hash"`
	Email               string   `json:"email"`
	Phone               string   `json:"phone"`
	IsFreePlaces        bool     `json:"is_free_places"`
	ScientificInterests []string `json:"scientific_interests"`
}
