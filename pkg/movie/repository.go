package movie

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Database *mongo.Database
	CTX      context.Context
}

func (movie *MongoRepository) GetMoviesExcludeHistory(userID string) ([]Movie, error) {
	collection := movie.Database.Collection("movies")
	ctx := movie.CTX
	var movies []Movie

	history, err := movie.getHistory(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": bson.M{"$nin": history}}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sample", Value: bson.D{
			{Key: "size", Value: 15},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &movies); err != nil {
		return nil, err
	}
	return movies, nil
}

func (movie *MongoRepository) GetRadomMovies() ([]Movie, error) {
	collection := movie.Database.Collection("movies")
	ctx := movie.CTX
	var movies []Movie

	pipeline := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{
			{Key: "size", Value: 15},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (movie *MongoRepository) getHistory(userID string) ([]primitive.ObjectID, error) {
	collection := movie.Database.Collection("history")
	ctx := movie.CTX
	var history []primitive.ObjectID

	filter := bson.M{"user_id": userID}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc MovieHistory
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		objectID, err := primitive.ObjectIDFromHex(doc.MovieID)
		if err != nil {
			return nil, err
		}

		history = append(history, objectID)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return history, nil
}
