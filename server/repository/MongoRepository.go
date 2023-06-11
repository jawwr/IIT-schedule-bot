package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"server/models"
	"time"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository() *MongoRepository {
	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	return &MongoRepository{
		client: client,
	}
}

func connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}

func (r MongoRepository) GetSchedule() *models.DaySchedule { //todo
	return nil
}

func (r MongoRepository) GetLesson() *models.Lesson {
	return nil
}

func (r MongoRepository) GetToday() *models.DaySchedule {
	return nil
}

func (r MongoRepository) GetGroupSchedule() *models.Group {
	return nil
}

func (r MongoRepository) GetWeekByNumber(number int) *models.WeekSchedule {
	return nil
}