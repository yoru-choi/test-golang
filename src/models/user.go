// models/user.go
package models

// User 사용자 모델 구조체
type User struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
