package repositories

import (
	"context"
	"errors"
	"hrms/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// import "go.mongodb.org/mongo-driver/mongo"

type userRepository struct{
	database *mongo.Database
	collection string
}

func NewUserRepository(database *mongo.Database, collection string) models.UserRepository{
return &userRepository{
database: database,
collection: collection,
}
}
// CreateUser(ctx context.Context, user *User) error
// GetUsers(ctx context.Context) ([]*User, error)
// GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
// GetUserByUsername(ctx context.Context, username string) (*User, error)
// GetUserByEmail(ctx context.Context, email string) (*User, error)
// UpdateUser(ctx context.Context, user *User) error
// DeleteUser(ctx context.Context, id primitive.ObjectID) error
func (ur *userRepository)CreateUser(ctx context.Context, user *models.User) (primitive.ObjectID, error){
	res,  err := ur.database.Collection(ur.collection).InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}
	return oid, nil
}

func (ur *userRepository) GetUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	cursor, err := ur.database.Collection(ur.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user *models.User
	err := ur.database.Collection(ur.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil{
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user *models.User
	err := ur.database.Collection(ur.collection).FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil{
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User
	err := ur.database.Collection(ur.collection).FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil{
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	res, err := ur.database.Collection(ur.collection).UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0{
		return errors.New("no user was matched")
	} else if res.ModifiedCount == 0 {
		return errors.New("no fields were updated")
	}
	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	res, err := ur.database.Collection(ur.collection).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no user was matched")
	}
	return nil
}