package apis

import (
	"GolangCourse/commons"
	"GolangCourse/commons/apploggers"
	"GolangCourse/internals/models"
	"GolangCourse/internals/services"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type topicController struct {
	eservice services.TopicEventService
}

func NewTopicController(eservice services.TopicEventService) topicController {
	return topicController{
		eservice: eservice,
	}
}

// @Tags Topic Management
// @Summary GetAllTopics
// @Description Get a list of all topics
// @Accept json
// @Produce json
// @Success 200 {array} models.Topic
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topics [get]
func (t *topicController) GetAllTopics(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	logger.Info("Executing GetAllTopics")

	// Fetch all topics using the service layer
	topics, err := t.eservice.GetAllTopics(lcontext)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(err.Error(), nil))
	}

	// Return the list of topics
	return c.JSON(http.StatusOK, topics)
}

// @Tags Topic Management
// @Summary CreateTopic
// @Description Creates a new topic with title, description, and order
// @Accept json
// @Produce json
// @Param payload body models.Topic true "Topic data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topics [post]
func (t *topicController) CreateTopic(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	logger.Info("Executing CreateTopic")

	var topic *models.Topic
	err := c.Bind(&topic)
	if err != nil || topic == nil {
		logger.Error("invalid request payload")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("invalid request payload", nil))
	}

	// Basic validation
	if len(strings.TrimSpace(topic.Name)) == 0 {
		logger.Error("'name' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'title' is required", nil))
	}

	// Create topic using the service layer
	topicID, err := t.eservice.CreateTopic(lcontext, topic)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(err.Error(), nil))
	}

	logger.Infof("Created Topic with ID: %s", topicID)
	return c.JSON(http.StatusCreated, map[string]string{"id": topicID})
}

// @Tags Topic Management
// @Summary UpdateTopic
// @Description Update topic details such as title, description, order, and visibility
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Param payload body models.Topic true "Updated topic data"
// @Success 200
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topics/{id} [put]
func (t *topicController) UpdateTopic(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicID := c.Param("id")
	logger.Infof("Executing UpdateTopic, topicId: %s", topicID)

	if len(strings.TrimSpace(topicID)) == 0 {
		logger.Error("'id' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'id' is required", nil))
	}

	var topic *models.Topic
	err := c.Bind(&topic)
	if err != nil || topic == nil {
		logger.Error("invalid request payload")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("invalid request payload", nil))
	}

	// Update the topic using service layer
	err = t.eservice.UpdateTopic(lcontext, topicID, topic)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(err.Error(), nil))
	}

	logger.Infof("Executed UpdateTopic, topicId: %s", topicID)
	return c.NoContent(http.StatusOK)
}

// @Tags Topic Management
// @Summary DeleteTopic
// @Description Delete a topic by ID
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Success 204
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topics/{id} [delete]
func (t *topicController) DeleteTopic(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicID := c.Param("id")
	logger.Infof("Executing DeleteTopic, topicId: %s", topicID)

	if len(strings.TrimSpace(topicID)) == 0 {
		logger.Error("'id' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'id' is required", nil))
	}

	err := t.eservice.DeleteTopic(lcontext, topicID)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(err.Error(), nil))
	}

	logger.Infof("Deleted Topic with ID: %s", topicID)
	return c.NoContent(http.StatusNoContent)
}

// @Tags Topic Management
// @Summary HideTopic
// @Description Hide a topic by ID
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Success 200
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topics/{id}/hide [put]
func (t *topicController) HideTopic(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicID := c.Param("id")
	logger.Infof("Executing HideTopic, topicId: %s", topicID)

	if len(strings.TrimSpace(topicID)) == 0 {
		logger.Error("'id' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'id' is required", nil))
	}

	err := t.eservice.HideTopic(lcontext, topicID)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(err.Error(), nil))
	}

	logger.Infof("Hidden Topic with ID: %s", topicID)
	return c.JSON(http.StatusOK, "Topic hidden successfully")
}
