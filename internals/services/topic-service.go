package services

import (
	"GolangCourse/commons/apploggers"
	"GolangCourse/internals/db"
	dbmodel "GolangCourse/internals/db/models"
	"GolangCourse/internals/models"
	"context"
	"encoding/json"
	"errors"
)

type TopicEventService interface {
	GetAllTopics(ctx context.Context) ([]models.Topic, error)
	CreateTopic(context context.Context, topic *models.Topic) (string, error)
	UpdateTopic(context context.Context, topicID string, topic *models.Topic) error
	DeleteTopic(context context.Context, topicID string) error
	HideTopic(context context.Context, topicID string) error
}

type topicService struct {
	topicDbService db.TopicDbService
}

func NewTopicEventService(topicDbService db.TopicDbService) TopicEventService {
	return &topicService{
		topicDbService: topicDbService,
	}
}

// CreateTopic creates a new topic in the database
func (t *topicService) CreateTopic(context context.Context, topic *models.Topic) (string, error) {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Info("Executing CreateTopic...")

	// Mapping the models.Topic to dbmodel.TopicSchema for database interaction
	var topicSchema *dbmodel.TopicSchema
	pbyes, _ := json.Marshal(topic)
	uerror := json.Unmarshal(pbyes, &topicSchema)
	if uerror != nil {
		logger.Error(uerror.Error())
		return "", uerror
	}

	// Save the topic in the database
	topicID, dberror := t.topicDbService.SaveTopic(context, topicSchema)
	if dberror != nil {
		logger.Error(dberror)
		return "", dberror
	}

	logger.Infof("Executed CreateTopic, topicID: %v", topicID)
	return topicID, nil
}

// UpdateTopic updates the details of an existing topic
func (t *topicService) UpdateTopic(context context.Context, topicID string, topic *models.Topic) error {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing UpdateTopic, topicID: %s", topicID)

	// Ensure the topicID is not empty
	if len(topicID) == 0 {
		return errors.New("topicID is required")
	}

	// Mapping the models.Topic to dbmodel.TopicSchema for database interaction
	var topicSchema *dbmodel.TopicSchema
	pbyes, _ := json.Marshal(topic)
	uerror := json.Unmarshal(pbyes, &topicSchema)
	if uerror != nil {
		logger.Error(uerror.Error())
		return uerror
	}

	// Update the topic in the database
	dberror := t.topicDbService.UpdateTopic(context, topicSchema, topicID)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}

	logger.Infof("Executed UpdateTopic, topicID: %s", topicID)
	return nil
}

// DeleteTopic deletes a topic by its ID
func (t *topicService) DeleteTopic(context context.Context, topicID string) error {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing DeleteTopic, topicID: %s", topicID)

	// Ensure the topicID is not empty
	if len(topicID) == 0 {
		return errors.New("topicID is required")
	}

	// Delete the topic from the database
	dberror := t.topicDbService.DeleteTopicById(context, topicID)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}

	logger.Infof("Executed DeleteTopic, topicID: %s", topicID)
	return nil
}

// HideTopic hides a topic by setting a "hidden" flag in the database
func (t *topicService) HideTopic(context context.Context, topicID string) error {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing HideTopic, topicID: %s", topicID)

	// Ensure the topicID is not empty
	if len(topicID) == 0 {
		return errors.New("topicID is required")
	}

	// Hide the topic by updating the hidden flag in the database
	dberror := t.topicDbService.HideTopic(context, topicID)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}

	logger.Infof("Executed HideTopic, topicID: %s", topicID)
	return nil
}

// GetAllTopics retrieves all topics from the database
func (t *topicService) GetAllTopics(ctx context.Context) ([]models.Topic, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Info("Executing GetAllTopics")

	// Fetch all topics from the database using the topicDbService
	topics, err := t.topicDbService.GetAllTopics(ctx)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Infof("Executed GetAllTopics, found %d topics", len(topics))
	return topics, nil
}
