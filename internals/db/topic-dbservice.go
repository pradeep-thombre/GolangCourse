package db

import (
	"GolangCourse/commons"
	"GolangCourse/commons/appdb"
	"GolangCourse/commons/apploggers"
	"GolangCourse/configs"
	dbmodel "GolangCourse/internals/db/models"
	"GolangCourse/internals/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type topicDbService struct {
	tcollection appdb.DatabaseCollection
}

type TopicDbService interface {
	GetTopicById(ctx context.Context, id string) (*models.Topic, error)
	DeleteTopicById(ctx context.Context, id string) error
	GetTopics(ctx context.Context) ([]*models.Topic, error)
	SaveTopic(ctx context.Context, topic *dbmodel.TopicSchema) (string, error)
	UpdateTopic(ctx context.Context, topic *dbmodel.TopicSchema, topicId string) error
	HideTopic(ctx context.Context, topicId string) error
}

func NewTopicDbService(dbclient appdb.DatabaseClient) TopicDbService {
	return &topicDbService{
		tcollection: dbclient.Collection(configs.MONGO_TOPICS_COLLECTION),
	}
}

// GetTopicById retrieves a topic by its ID
func (t *topicDbService) GetTopicById(ctx context.Context, topicId string) (*models.Topic, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetTopicById, Id: %s", topicId)

	// Get object ID from the topicId string
	id, err := primitive.ObjectIDFromHex(topicId)
	if err != nil {
		return nil, fmt.Errorf("invalid topicId provided, topicId: %s", topicId)
	}

	var topic *models.Topic
	var filter = bson.M{"_id": id}
	dbError := t.tcollection.FindOne(ctx, filter, &topic)
	if dbError != nil {
		logger.Error(dbError)
		return nil, dbError
	}
	logger.Infof("Executed GetTopicById, topic: %s", commons.PrintStruct(topic))
	return topic, nil
}

// DeleteTopicById deletes a topic by its ID
func (t *topicDbService) DeleteTopicById(ctx context.Context, topicId string) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing DeleteTopicById, Id: %s", topicId)

	// Get object ID from topicId string
	id, err := primitive.ObjectIDFromHex(topicId)
	if err != nil {
		return fmt.Errorf("cannot delete topic, invalid topicId provided, topicId: %s", topicId)
	}
	var filter = bson.M{"_id": id}
	_, dbError := t.tcollection.DeleteOne(ctx, filter)
	if dbError != nil {
		logger.Error(dbError)
		return dbError
	}
	logger.Infof("Executed DeleteTopicById, Id: %s", topicId)
	return nil
}

// GetTopics retrieves all topics
func (t *topicDbService) GetTopics(ctx context.Context) ([]*models.Topic, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetTopics")

	var topics []*models.Topic
	var filter = map[string]interface{}{}
	dbError := t.tcollection.Find(ctx, filter, &options.FindOptions{}, &topics)
	if dbError != nil {
		logger.Error(dbError)
		return nil, dbError
	}
	logger.Infof("Executed GetTopics, topics: %d", len(topics))
	return topics, nil
}

// SaveTopic saves a new topic in the database
func (t *topicDbService) SaveTopic(ctx context.Context, topic *dbmodel.TopicSchema) (string, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing SaveTopic...")

	// Insert topic into db
	result, dbError := t.tcollection.InsertOne(ctx, topic)
	if dbError != nil {
		logger.Error(dbError)
		return "", dbError
	}

	// Extract the inserted ID from the result
	id := result.InsertedID.(primitive.ObjectID).Hex()
	logger.Infof("Executed SaveTopic, topicId: %s", commons.PrintStruct(topic))
	return id, nil
}

// UpdateTopic updates an existing topic
func (t *topicDbService) UpdateTopic(ctx context.Context, topic *dbmodel.TopicSchema, topicId string) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing UpdateTopic...")

	// Get object ID from topicId string
	id, err := primitive.ObjectIDFromHex(topicId)
	if err != nil {
		return fmt.Errorf("cannot update topic, invalid topicId provided, topicId: %s", topicId)
	}
	var filter = bson.M{"_id": id}
	update := bson.M{"$set": topic} // Correct update document
	// Update topic in db
	_, dbError := t.tcollection.UpdateOne(ctx, filter, update)
	if dbError != nil {
		logger.Error(dbError)
		return dbError
	}

	logger.Infof("Executed UpdateTopic, topicId: %s", commons.PrintStruct(topic))
	return nil
}

// HideTopic hides a topic by setting a "hidden" flag
func (t *topicDbService) HideTopic(ctx context.Context, topicId string) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing HideTopic, topicId: %s", topicId)

	// Get object ID from topicId string
	id, err := primitive.ObjectIDFromHex(topicId)
	if err != nil {
		return fmt.Errorf("cannot hide topic, invalid topicId provided, topicId: %s", topicId)
	}
	var filter = bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"hidden": true}} // Set the hidden flag to true
	// Update topic in db
	_, dbError := t.tcollection.UpdateOne(ctx, filter, update)
	if dbError != nil {
		logger.Error(dbError)
		return dbError
	}

	logger.Infof("Executed HideTopic, topicId: %s", topicId)
	return nil
}
