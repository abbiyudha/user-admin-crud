package admin

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"sejutaCita/services/admin-service/entities"
)

type AdminRepository struct {
	database *mongo.Database
}

func NewAdminRepository(db *mongo.Database) *AdminRepository {
	return &AdminRepository{
		database: db,
	}
}

type AdminRepositoryInferface interface {
	CreateAdmin(adminParam entities.Admin) error
	LoginAdmin(email string) (entities.Admin, error)
	GetAdminById(id string) (entities.Admin, error)
	UpdateAdmin(id string, adminParam entities.Admin) error
	DeleteAdmin(id string) error
}

func (ar *AdminRepository) CreateAdmin(adminParam entities.Admin) error {
	var ctx context.Context

	adminParam.Id = primitive.NewObjectID()
	encrypt, _ := bcrypt.GenerateFromPassword([]byte(adminParam.Password), bcrypt.DefaultCost)
	adminParam.Password = string(encrypt)

	_, err := ar.database.Collection("Admin").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	_, err = ar.database.Collection("Admin").InsertOne(ctx, adminParam)
	if err != nil {
		return errors.New("Email already used")
	}

	return nil
}

func (ar *AdminRepository) LoginAdmin(email string) (entities.Admin, error) {
	var ctx context.Context
	var admin entities.Admin

	result := ar.database.Collection("Admin").FindOne(ctx, bson.M{"email": email})
	if result.Err() != nil {
		return admin, result.Err()
	}

	err := result.Decode(&admin)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (ar *AdminRepository) GetAdminById(id string) (entities.Admin, error) {
	var ctx context.Context
	var adminParam entities.Admin

	adminParam.Id, _ = primitive.ObjectIDFromHex(id)

	result := ar.database.Collection("Admin").FindOne(ctx, bson.M{"_id": adminParam.Id})
	if result.Err() != nil {
		return adminParam, errors.New("User not found")
	}

	err := result.Decode(&adminParam)
	if err != nil {
		return entities.Admin{}, err
	}

	return adminParam, nil
}

func (ar *AdminRepository) UpdateAdmin(id string, adminParam entities.Admin) error {
	var ctx context.Context

	idParam, _ := primitive.ObjectIDFromHex(id)
	adminParam.Id = idParam
	encrypt, _ := bcrypt.GenerateFromPassword([]byte(adminParam.Password), bcrypt.DefaultCost)
	adminParam.Password = string(encrypt)

	_, err := ar.database.Collection("Admin").UpdateOne(ctx, bson.M{"_id": idParam}, bson.M{"$set": adminParam})
	if err != nil {
		return err
	}

	return err
}

func (ar *AdminRepository) DeleteAdmin(id string) error {
	var ctx context.Context
	var adminParam entities.Admin

	idParam, _ := primitive.ObjectIDFromHex(id)
	adminParam.Id = idParam

	_, err := ar.database.Collection("Admin").DeleteOne(ctx, bson.M{"_id": adminParam.Id})
	if err != nil {
		return err
	}

	return err
}
