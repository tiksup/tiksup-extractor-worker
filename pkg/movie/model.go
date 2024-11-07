package movie

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL         string             `bson:"url" json:"url"`
	Title       string             `bson:"title" json:"title"`
	Genre       []string           `bson:"genre" json:"genre"`
	Protagonist string             `bson:"protagonist" json:"protagonist"`
	Director    string             `bson:"director" json:"director"`
}

type MovieHistory struct {
	UserID  string `bson:"user_id" json:"user_id"`
	MovieID string `bson:"movie_id" json:"movie_id"`
}
