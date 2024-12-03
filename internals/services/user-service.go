package services

import (
	"GolangCourse/commons/apploggers"
	"GolangCourse/internals/db"
	dbmodel "GolangCourse/internals/db/models"
	"GolangCourse/internals/models"
	"context"
	"encoding/json"
)

type UserEventService interface {
	GetUserById(context context.Context, userId string) (*models.User, error)
	DeleteUserById(context context.Context, userId string) error
	GetUsers(context context.Context) ([]*models.User, error)
	CreateUser(context context.Context, user *models.User) (string, error)
	UpdateUser(context context.Context, user *models.User, userId string) error
}

type eservice struct {
	dbservice db.UserDbService
}

func NewUserEventService(userDbService db.UserDbService) UserEventService {
	return &eservice{
		dbservice: userDbService,
	}
}

func (e *eservice) GetUserById(context context.Context, userId string) (*models.User, error) {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing GetUserById, userId: %s", userId)
	user, dberror := e.dbservice.GetUserById(context, userId)
	if dberror != nil {
		logger.Error(dberror)
		return nil, dberror
	}
	logger.Infof("Executed GetUserById, userId: %s", userId)
	return user, nil
}

func (e *eservice) DeleteUserById(context context.Context, userId string) error {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing DeleteUserById, userId: %s", userId)
	dberror := e.dbservice.DeleteUserById(context, userId)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}
	logger.Infof("Executed DeleteUserById, userId: %s", userId)
	return nil
}

func (e *eservice) GetUsers(context context.Context) ([]*models.User, error) {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing GetUsers...")
	users, dberror := e.dbservice.GetUsers(context)
	if dberror != nil {
		logger.Error(dberror)
		return nil, dberror
	}
	logger.Infof("Executed GetUsers, users: %d", len(users))
	return users, nil
}

func (e *eservice) CreateUser(context context.Context, user *models.User) (string, error) {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing CreateUser...")
	var userSchema *dbmodel.UserSchema
	pbyes, _ := json.Marshal(user)
	uerror := json.Unmarshal(pbyes, &userSchema)
	if uerror != nil {
		logger.Error(uerror.Error())
		return "", uerror
	}
	userId, dberror := e.dbservice.SaveUser(context, userSchema)
	if dberror != nil {
		logger.Error(dberror)
		return "", dberror
	}
	logger.Infof("Executed CreateUser, userId: %v", userId)
	return userId, nil
}

func (e *eservice) UpdateUser(context context.Context, user *models.User, userId string) error {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing UpdateUser...")
	var userSchema *dbmodel.UserSchema
	pbyes, _ := json.Marshal(user)
	uerror := json.Unmarshal(pbyes, &userSchema)
	if uerror != nil {
		logger.Error(uerror.Error())
		return uerror
	}
	dberror := e.dbservice.UpdateUser(context, userSchema, userId)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}
	logger.Infof("Executed UpdateUser, userId: %v", userId)
	return nil
}
