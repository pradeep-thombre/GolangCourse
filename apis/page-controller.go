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

type pcontroller struct {
	eservice services.PageEventService
}

func NewPageController(eservice services.PageEventService) pcontroller {
	return pcontroller{
		eservice: eservice,
	}
}

// @Tags Page Management
// @Summary GetPageById
// @Description Gets page details by page id such as title, content, topicID, etc.
// @Accept json
// @Produce json
// @Param topicId path string true "Topic id"
// @Param id path string true "Page id"
// @Success 200 {object} models.Page
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topic/{topicId}/pages/{id} [Get]
func (p *pcontroller) GetPageById(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicId := c.Param("topicId")
	pageId := c.Param("id")
	logger.Infof("Executing GetPageById, topicId: %s, pageId: %s", topicId, pageId)
	if len(strings.TrimSpace(topicId)) == 0 || len(strings.TrimSpace(pageId)) == 0 {
		logger.Error("'topicId' and 'id' are required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'topicId' and 'id' are required", nil))
	}
	page, serror := p.eservice.GetPageById(lcontext, topicId, pageId)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Infof("Executed GetPageById, topicId: %s, pageId: %s, page: %v", topicId, pageId, page)
	return c.JSON(http.StatusOK, page)
}

// @Tags Page Management
// @Summary DeletePageById
// @Description Delete page details by page id for a specific topic
// @Accept json
// @Produce json
// @Param topicId path string true "Topic id"
// @Param id path string true "Page id"
// @Success 204
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topic/{topicId}/pages/{id} [Delete]
func (p *pcontroller) DeletePageById(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicId := c.Param("topicId")
	pageId := c.Param("id")
	logger.Infof("Executing DeletePageById, topicId: %s, pageId: %s", topicId, pageId)
	if len(strings.TrimSpace(topicId)) == 0 || len(strings.TrimSpace(pageId)) == 0 {
		logger.Error("'topicId' and 'id' are required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'topicId' and 'id' are required", nil))
	}
	serror := p.eservice.DeletePage(lcontext, topicId, pageId)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Infof("Executed DeletePageById, topicId: %s, pageId: %s", topicId, pageId)
	return c.NoContent(http.StatusNoContent)
}

// @Tags Page Management
// @Summary GetPagesByTopicId
// @Description Get details of all pages for a given topicId
// @Accept json
// @Produce json
// @Param topicId path string true "Topic id"
// @Success 200 {object} []models.Page
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topic/{topicId}/pages [Get]
func (p *pcontroller) GetPagesByTopicId(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicId := c.Param("topicId")
	logger.Infof("Executing GetPagesByTopicId, topicId: %s", topicId)
	if len(strings.TrimSpace(topicId)) == 0 {
		logger.Error("'topicId' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'topicId' is required", nil))
	}
	pages, serror := p.eservice.GetPagesByTopicId(lcontext, topicId)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Infof("Executed GetPagesByTopicId, topicId: %s, pages: %v", topicId, pages)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": len(pages),
		"pages": pages,
	})
}

// @Tags Page Management
// @Summary CreatePage
// @Description Save a new page with details like title, content, topicID
// @Accept json
// @Produce json
// @Param topicId path string true "Topic id"
// @Param payload body models.Page true "Page data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topic/{topicId}/pages [post]
func (p *pcontroller) CreatePage(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicId := c.Param("topicId")
	logger.Infof("Executing CreatePage, topicId: %s", topicId)
	var page *models.Page
	err := c.Bind(&page)
	if err != nil || page == nil {
		logger.Error("invalid request payload")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("invalid request payload", nil))
	}
	if len(strings.TrimSpace(page.Title)) == 0 {
		logger.Error("'title' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'title' is required", nil))
	}
	if len(strings.TrimSpace(page.Content)) == 0 {
		logger.Error("'content' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'content' is required", nil))
	}
	pageID, serror := p.eservice.CreatePage(lcontext, topicId, page)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Infof("Executed CreatePage, pageID: %s", pageID)
	return c.JSON(http.StatusCreated, map[string]string{
		"id": pageID,
	})
}

// @Tags Page Management
// @Summary UpdatePage
// @Description Update existing page details like title, content, and topicID
// @Accept json
// @Produce json
// @Param topicId path string true "Topic id"
// @Param id path string true "Page id"
// @Param payload body models.Page true "Updated page data"
// @Success 200
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topic/{topicId}/pages/{id} [patch]
func (p *pcontroller) UpdatePage(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicId := c.Param("topicId")
	pageId := c.Param("id")
	logger.Infof("Executing UpdatePage, topicId: %s, pageId: %s", topicId, pageId)
	if len(strings.TrimSpace(topicId)) == 0 || len(strings.TrimSpace(pageId)) == 0 {
		logger.Error("'topicId' and 'id' are required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'topicId' and 'id' are required", nil))
	}
	var page *models.Page
	err := c.Bind(&page)
	if err != nil || page == nil {
		logger.Error("invalid request payload")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("invalid request payload", nil))
	}
	if len(strings.TrimSpace(page.Title)) == 0 {
		logger.Error("'title' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'title' is required", nil))
	}
	if len(strings.TrimSpace(page.Content)) == 0 {
		logger.Error("'content' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'content' is required", nil))
	}
	serror := p.eservice.UpdatePage(lcontext, topicId, pageId, page)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Infof("Executed UpdatePage, topicId: %s, pageId: %s", topicId, pageId)
	return c.NoContent(http.StatusOK)
}

// @Tags Page Management
// @Summary HidePage
// @Description Hide a page by page id (marking it as hidden) for a specific topic
// @Accept json
// @Produce json
// @Param topicId path string true "Topic id"
// @Param id path string true "Page id"
// @Success 204
// @Failure 400 {object} commons.ApiErrorResponsePayload
// @Router /topic/{topicId}/pages/{id}/hide [patch]
func (p *pcontroller) HidePage(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	topicId := c.Param("topicId")
	pageId := c.Param("id")
	logger.Infof("Executing HidePage, topicId: %s, pageId: %s", topicId, pageId)
	if len(strings.TrimSpace(topicId)) == 0 || len(strings.TrimSpace(pageId)) == 0 {
		logger.Error("'topicId' and 'id' are required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'topicId' and 'id' are required", nil))
	}
	serror := p.eservice.HidePage(lcontext, topicId, pageId)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Infof("Executed HidePage, topicId: %s, pageId: %s", topicId, pageId)
	return c.NoContent(http.StatusNoContent)
}
