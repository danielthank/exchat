package service

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionService interface {
	Save(r *http.Request, w http.ResponseWriter, sessionID string, payload map[string]interface{}) error
	Get(r *http.Request, sessionID string) (*sessions.Session, error)
}
