package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBConfig struct {
	DBName  string
	Timeout time.Duration
	DBHost  string
	DBPort  string
}

const SWEAT_TABLE string = "sweat" // the collection name

type Sweat struct {
	// Database specific fields.
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userid,omitempty" json:"UserID"`
	CreatedAt time.Time          `bson:"created_at"`

	// Potential disease Diagnosis
	Glucose  float32 `bson:"glucose" json:"Glucose,omitempty"`   // excess indicates diabetes
	Chloride float32 `bson:"chloride" json:"Chloride,omitempty"` // excess indicates cystic fibrosis

	// Electrolytes
	Sodium    float32 `bson:"sodium" json:"Sodium,omitempty"`
	Potassium float32 `bson:"potassium" json:"Potassium,omitempty"` // excess indicates exercise / workout
	Magnesium float32 `bson:"magnesium" json:"Magnesium,omitempty"` // excess indicates exercise / workout
	Calcium   float32 `bson:"calcium" json:"Calcium,omitempty"`     // excess indicates exercise / workout

	// Environmental conditions and determining criteria
	Humidity         float32 `bson:"humidity" json:"Humidity,omitempty"`                // high humidity increseas sweating
	RoomTemperatureF float32 `bson:"room_temperature" json:"RoomTemperature,omitempty"` // cooler room temperature with sweat indicates hyperdidrosis
	BodyTemperatureF float32 `bson:"body_temperature" json:"BodyTemperature,omitempty"` // high body temperature with sweat indicates fever
	HeartBeat        int32   `bson:"heartbeat" json:"HeartBeat,omitempty"`              // sweating without apparent reason is an alarming condition!
}

type Store interface {
	Store(ctx context.Context, collection string, data any) error
	Update(ctx context.Context, collection string, filter any, update any) error
	//FindAll(context.Context, string, any, any) error
	//FindOne(context.Context, string, any, any, any) error
	//StartSession() (Session, error)
}
