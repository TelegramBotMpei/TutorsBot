/*
 * Телеграм Бот для работы с научными группами
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Преподователь тип
type Tutor struct {
	// Идентификатор преподавателя
	TutorId int32 `json:"tutor_id,omitempty"`
	// Имя преподавателя
	Name string `json:"name"`
	// Фамилия преподавателя
	Surname string `json:"surname"`
	// Телефон преподавателя для того, чтобы ему можно было позвонить
	Phone string `json:"phone,omitempty"`
	// Почта преподавателя
	Email string `json:"email,omitempty"`
	// Есть ли свободные места
	HasFreePlaces bool `json:"has_free_places,omitempty"`
	// Научные интересы преподавателя
	ScientificInterests []string `json:"scientific_interests,omitempty"`
}
