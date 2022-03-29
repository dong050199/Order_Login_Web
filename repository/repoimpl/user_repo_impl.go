package repoimpl

import (
	models "SQLite_JWT/model/user"
	repo "SQLite_JWT/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserREpoImpl struct {
	DB *mongo.Database
}

func NewUserRepo(db *mongo.Database) repo.UserRepo {
	return &UserREpoImpl{
		DB: db,
	}
}

func (mongo *UserREpoImpl) FindUserByEmail(email string) (models.User, error) {
	user := models.User{}

	result := mongo.DB.Collection("user").FindOne(context.Background(), bson.M{"email": email})
	err := result.Decode(&user)
	if user == (models.User{}) {
		return user, models.ERR_USER_NOT_FOUND
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func (mongo *UserREpoImpl) CheckLoginInfo(email string, password string) (models.User, error) {
	user := models.User{}
	result := mongo.DB.Collection("user").FindOne(context.Background(), bson.M{"email": email, "password": password})
	err := result.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (mongo *UserREpoImpl) Insert(user models.User) error {
	bbytes, _ := bson.Marshal(user)
	_, err := mongo.DB.Collection("user").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}
	return nil
}
