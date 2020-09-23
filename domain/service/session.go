package service

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionService interface {
	Get(r *http.Request, sessionID string) (*sessions.Session, error)
	Save(r *http.Request, w http.ResponseWriter, sessionID string, payload map[string]interface{}) error
	Delete(r *http.Request, w http.ResponseWriter, sessionID string) error
}
