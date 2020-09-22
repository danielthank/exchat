package infra

import (
	"net/http"

	"github.com/danielthank/exchat-server/domain/service"
	"github.com/gorilla/sessions"
)

type sessionService struct {
	*RedisHandler
}

func NewSessionService(redisHandler *RedisHandler) service.SessionService {
	sessionService := &sessionService{redisHandler}
	return sessionService
}

func (t *sessionService) Save(r *http.Request, w http.ResponseWriter, sessionID string, payload map[string]interface{}) error {
	session, err := t.Get(r, sessionID)
	if err != nil {
		return err
	}
	for key, value := range payload {
		session.Values[key] = value
	}
	if err = session.Save(r, w); err != nil {
		return err
	}
	return nil
}

func (t *sessionService) Get(r *http.Request, sessionID string) (*sessions.Session, error) {
	session, err := t.RedisHandler.Store.Get(r, sessionID)
	if err != nil {
		return nil, err
	}
	return session, nil
}
