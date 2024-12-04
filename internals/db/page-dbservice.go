package db

import (
	"GolangCourse/commons"
	"GolangCourse/commons/appdb"
	"GolangCourse/commons/apploggers"
	"GolangCourse/configs"
	"GolangCourse/internals/db/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PageDbService interface {
	GetPageById(ctx context.Context, pageID string) (*models.PageSchema, error)
	DeletePageById(ctx context.Context, pageID string) error
	GetPagesByTopicId(ctx context.Context, topicID string) ([]*models.PageSchema, error)
	SavePage(ctx context.Context, page *models.PageSchema) (string, error)
	UpdatePage(ctx context.Context, page *models.PageSchema, pageID string) error
	HidePage(ctx context.Context, pageID string) error
}

type pageDbService struct {
	pcollection appdb.DatabaseCollection
}

func NewPageDbService(dbclient appdb.DatabaseClient) PageDbService {
	return &pageDbService{
		pcollection: dbclient.Collection(configs.MONGO_PAGES_COLLECTION),
	}
}

// GetPageById retrieves a page by its ID
func (p *pageDbService) GetPageById(ctx context.Context, pageID string) (*models.PageSchema, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetPageById, pageID: %s", pageID)

	// Convert pageID to ObjectID
	id, err := primitive.ObjectIDFromHex(pageID)
	if err != nil {
		return nil, fmt.Errorf("invalid pageID provided, pageID: %s", pageID)
	}

	var page models.PageSchema
	filter := bson.M{"_id": id}
	// Using FindOne with proper options (nil in this case is fine for simple queries)
	err = p.pcollection.FindOne(ctx, filter, &page)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Infof("Executed GetPageById, page: %s", commons.PrintStruct(page))
	return &page, nil
}

// DeletePageById deletes a page by its ID
func (p *pageDbService) DeletePageById(ctx context.Context, pageID string) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing DeletePageById, pageID: %s", pageID)

	// Convert pageID to ObjectID
	id, err := primitive.ObjectIDFromHex(pageID)
	if err != nil {
		return fmt.Errorf("invalid pageID provided, pageID: %s", pageID)
	}

	filter := bson.M{"_id": id}
	_, err = p.pcollection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Infof("Executed DeletePageById, pageID: %s", pageID)
	return nil
}

// GetPagesByTopicId retrieves all pages associated with a given topic ID
func (p *pageDbService) GetPagesByTopicId(ctx context.Context, topicID string) ([]*models.PageSchema, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetPagesByTopicId, topicID: %s", topicID)

	var pages []*models.PageSchema
	filter := bson.M{"topic_id": topicID}

	// Using Find with proper options (nil in this case is fine for simple queries)
	err := p.pcollection.Find(ctx, filter, &options.FindOptions{}, &pages)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Infof("Executed GetPagesByTopicId, pages found: %d", len(pages))
	return pages, nil
}

// SavePage saves a new page to the database
func (p *pageDbService) SavePage(ctx context.Context, page *models.PageSchema) (string, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing SavePage...")

	// Insert page into the database
	result, err := p.pcollection.InsertOne(ctx, page)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	// Get the inserted ID and return as hex string
	pageID := result.InsertedID.(primitive.ObjectID).Hex()
	logger.Infof("Executed SavePage, pageID: %s", pageID)
	return pageID, nil
}

// UpdatePage updates an existing page in the database
func (p *pageDbService) UpdatePage(ctx context.Context, page *models.PageSchema, pageID string) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing UpdatePage, pageID: %s", pageID)

	// Convert pageID to ObjectID
	id, err := primitive.ObjectIDFromHex(pageID)
	if err != nil {
		return fmt.Errorf("invalid pageID provided, pageID: %s", pageID)
	}

	// Update page in the database
	filter := bson.M{"_id": id}
	update := bson.M{"$set": page}
	_, err = p.pcollection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Infof("Executed UpdatePage, pageID: %s", pageID)
	return nil
}

// HidePage sets the hidden flag of a page to true
func (p *pageDbService) HidePage(ctx context.Context, pageID string) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing HidePage, pageID: %s", pageID)

	// Convert pageID to ObjectID
	id, err := primitive.ObjectIDFromHex(pageID)
	if err != nil {
		return fmt.Errorf("invalid pageID provided, pageID: %s", pageID)
	}

	// Update page to set "hidden" flag to true
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isHidden": true}}
	_, err = p.pcollection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Infof("Executed HidePage, pageID: %s", pageID)
	return nil
}
