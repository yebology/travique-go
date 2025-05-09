package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct {

	Id			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`			
	Name		string					`json:"name" validate:"required,min=8"`
	Email		string					`json:"email" validate:"required,email"`
	Avatar		string					`json:"avatar" validate:"required"`

}