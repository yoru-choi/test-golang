package repository

import (
	"context"

	"test-golang/src/models" // 수정된 경로

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository defines the methods to interact with the user data
type UserRepository interface {
	GetUsers(ctx context.Context) ([]*models.User, error) // 수정: 반환 타입 변경
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, id string, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

// userRepository implements UserRepository
type userRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) GetUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User // 사용자 정보를 담을 슬라이스 생성

	cursor, err := r.collection.Find(ctx, bson.D{}) // 모든 사용자 문서 찾기
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) // 커서 닫기

	// 커서에서 문서를 간단히 반복하여 디코딩
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil // 사용자 슬라이스 반환
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) UpdateUser(ctx context.Context, id string, user *models.User) error {
	// 업데이트할 필드들을 bson.M에 담아 전달
	update := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
		},
	}
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update, options.Update())
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
