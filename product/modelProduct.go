package product

import (
	"time"
)

type Products struct {
	ID 			int			`json:"id" bson:"_id"`
	Name 		string		`json:"name" bson:"name"`
	Status		int			`json:"status" bson:"status"`
	CreatedAt 	time.Time 	`json:"created_at" bson:"created_at"`
	CreatedBy 	string		`json:"created_by" bson:"created_by"`
	UpdatedAt 	time.Time 	`json:"created_at" bson:"updated_at"`
	UpdatedBy 	string		`json:"created_by" bson:"updated_by"`
}