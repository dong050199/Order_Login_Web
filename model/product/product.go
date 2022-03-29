package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Id_product  int                `json:"id_product" bson:"id_product,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Price       float32            `json:"price" bson:"price,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Category    string             `json:"category" bson:"category,omitempty"`
	Image       string             `json:"image" bson:"image,omitempty"`
	Rate        float32            `json:"rate" bson:"rate,omitempty"`
	Count       int                `json:"count" bson:"count,omitempty"`
}
