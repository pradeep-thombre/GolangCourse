package services

import (
	"GolangCourse/commons/apploggers"
	"GolangCourse/internals/db"
	dbmodel "GolangCourse/internals/db/models"
	"GolangCourse/internals/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PageEventService interface {
	CreatePage(context context.Context, topicID string, page *models.Page) (string, error)
	UpdatePage(context context.Context, topicID string, pageID string, page *models.Page) error
	DeletePage(context context.Context, topicID string, pageID string) error
	HidePage(context context.Context, topicID string, pageID string) error
	GetPageById(context context.Context, topicID string, pageID string) (*models.Page, error)
	GetPagesByTopicId(context context.Context, topicID string) ([]*models.Page, error)
}

type pageService struct {
	pageDbService db.PageDbService
}

func NewPageEventService(pageDbService db.PageDbService) PageEventService {
	return &pageService{
		pageDbService: pageDbService,
	}
}

// CreatePage creates a new page in the database under a specific topic
func (p *pageService) CreatePage(ctx context.Context, topicID string, page *models.Page) (string, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Info("Executing CreatePage...")

	// Ensure the topicID is provided
	if len(topicID) == 0 {
		return "", errors.New("topicID is required")
	}

	// Mapping the models.Page to dbmodel.PageSchema for database interaction
	var pageSchema *dbmodel.PageSchema
	pbyes, _ := json.Marshal(page)
	uerror := json.Unmarshal(pbyes, &pageSchema)
	if uerror != nil {
		logger.Error(uerror.Error())
		return "", uerror
	}

	// Convert the string topicID to primitive.ObjectID
	topicObjID, err := primitive.ObjectIDFromHex(topicID)
	if err != nil {
		return "", fmt.Errorf("invalid topicID provided, topicID: %s", topicID)
	}

	// Set the topicID in the page schema
	pageSchema.TopicID = topicObjID

	// Save the page in the database
	pageID, dberror := p.pageDbService.SavePage(ctx, pageSchema)
	if dberror != nil {
		logger.Error(dberror)
		return "", dberror
	}

	logger.Infof("Executed CreatePage, pageID: %v", pageID)
	return pageID, nil
}

// UpdatePage updates the details of an existing page under a specific topic
func (p *pageService) UpdatePage(ctx context.Context, topicID string, pageID string, page *models.Page) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing UpdatePage, topicID: %s, pageID: %s", topicID, pageID)

	// Ensure the pageID and topicID are provided
	if len(pageID) == 0 || len(topicID) == 0 {
		return errors.New("topicID and pageID are required")
	}

	// Mapping the models.Page to dbmodel.PageSchema for database interaction
	var pageSchema *dbmodel.PageSchema
	pbyes, _ := json.Marshal(page)
	uerror := json.Unmarshal(pbyes, &pageSchema)
	if uerror != nil {
		logger.Error(uerror.Error())
		return uerror
	}
	// Convert the string topicID to primitive.ObjectID
	topicObjID, err := primitive.ObjectIDFromHex(topicID)
	if err != nil {
		return fmt.Errorf("invalid topicID provided, topicID: %s", topicID)
	}
	// Set the topicID in the page schema
	pageSchema.TopicID = topicObjID

	// Update the page in the database
	dberror := p.pageDbService.UpdatePage(ctx, pageSchema, pageID)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}

	logger.Infof("Executed UpdatePage, topicID: %s, pageID: %s", topicID, pageID)
	return nil
}

// DeletePage deletes a page by its ID under a specific topic
func (p *pageService) DeletePage(ctx context.Context, topicID string, pageID string) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing DeletePage, topicID: %s, pageID: %s", topicID, pageID)

	// Ensure the pageID and topicID are provided
	if len(pageID) == 0 || len(topicID) == 0 {
		return errors.New("topicID and pageID are required")
	}

	// Delete the page from the database
	dberror := p.pageDbService.DeletePageById(ctx, pageID)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}

	logger.Infof("Executed DeletePage, topicID: %s, pageID: %s", topicID, pageID)
	return nil
}

// HidePage hides a page by setting a "hidden" flag in the database under a specific topic
func (p *pageService) HidePage(ctx context.Context, topicID string, pageID string) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing HidePage, topicID: %s, pageID: %s", topicID, pageID)

	// Ensure the pageID and topicID are provided
	if len(pageID) == 0 || len(topicID) == 0 {
		return errors.New("topicID and pageID are required")
	}

	// Hide the page by updating the hidden flag in the database
	dberror := p.pageDbService.HidePage(ctx, pageID)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}

	logger.Infof("Executed HidePage, topicID: %s, pageID: %s", topicID, pageID)
	return nil
}

// GetPageById retrieves a page by its ID under a specific topic
func (p *pageService) GetPageById(ctx context.Context, topicID string, pageID string) (*models.Page, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetPageById, topicID: %s, pageID: %s", topicID, pageID)

	// Ensure the pageID and topicID are provided
	if len(pageID) == 0 || len(topicID) == 0 {
		return nil, errors.New("topicID and pageID are required")
	}

	// Retrieve the page from the database
	pageSchema, err := p.pageDbService.GetPageById(ctx, pageID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// Convert dbmodel.PageSchema to models.Page
	page := &models.Page{
		ID:        pageSchema.ID,
		Title:     pageSchema.Title,
		Content:   pageSchema.Content,
		TopicID:   pageSchema.TopicID.Hex(),
		CreatedAt: pageSchema.CreatedAt,
		UpdatedAt: pageSchema.UpdatedAt,
		DeletedAt: pageSchema.DeletedAt,
		Hidden:    pageSchema.Hidden,
	}

	logger.Infof("Executed GetPageById, page: %s", pageID)
	return page, nil
}

// GetPagesByTopicId retrieves all pages associated with a given topic ID
func (p *pageService) GetPagesByTopicId(ctx context.Context, topicID string) ([]*models.Page, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetPagesByTopicId, topicID: %s", topicID)

	// Ensure the topicID is provided
	if len(topicID) == 0 {
		return nil, errors.New("topicID is required")
	}

	// Retrieve the pages associated with the topic from the database
	pageSchemas, err := p.pageDbService.GetPagesByTopicId(ctx, topicID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// Convert dbmodel.PageSchemas to models.Pages
	var pages []*models.Page
	for _, pageSchema := range pageSchemas {
		page := &models.Page{
			ID:        pageSchema.ID,
			Title:     pageSchema.Title,
			Content:   pageSchema.Content,
			TopicID:   pageSchema.TopicID.Hex(),
			CreatedAt: pageSchema.CreatedAt,
			UpdatedAt: pageSchema.UpdatedAt,
			DeletedAt: pageSchema.DeletedAt,
			Hidden:    pageSchema.Hidden,
		}
		pages = append(pages, page)
	}

	logger.Infof("Executed GetPagesByTopicId, pages found: %d", len(pages))
	return pages, nil
}
