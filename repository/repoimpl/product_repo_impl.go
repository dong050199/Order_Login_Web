package repoimpl

import (
	models "SQLite_JWT/model/product"
	repo "SQLite_JWT/repository"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductREpoImpl struct {
	DB *mongo.Database
}

func NewProductRepo(db *mongo.Database) repo.ProductRepo {
	return &ProductREpoImpl{
		DB: db,
	}
}

func (mongo *ProductREpoImpl) Update(u models.Product, id int) error {
	var product models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongo.DB.Collection("product").UpdateOne(ctx, models.Product{Id_product: id}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}

func (mongo *ProductREpoImpl) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongo.DB.Collection("product").DeleteOne(ctx, models.Product{Id_product: id})
	if err != nil {
		return err
	}
	return nil
}

func (mongo *ProductREpoImpl) Select() ([]models.Product, error) {
	var products []models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	curso, err := mongo.DB.Collection("product").Find(ctx, bson.M{})
	if err != nil {
		return products, err
	}
	for curso.Next(ctx) {
		var prodct models.Product
		curso.Decode(prodct)
		products = append(products, prodct)

	}
	return products, nil
}

func (mongo *ProductREpoImpl) SelectId(id int) (models.Product, error) {
	var product models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := mongo.DB.Collection("product").FindOne(ctx, models.Product{Id_product: id}).Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (mongo *ProductREpoImpl) Insert(u models.Product) error {
	bbytes, _ := bson.Marshal(u)
	_, err := mongo.DB.Collection("product").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}
	return nil
}
