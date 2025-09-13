package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const SessionKey = "user_id"

func SetSession(c *gin.Context, userID uuid.UUID) {
	session := sessions.Default(c)
	session.Set(SessionKey, userID)
	session.Save()
}

func GetSessionUserID(c *gin.Context) (uuid.UUID, bool) {
	session := sessions.Default(c)
	id, ok := session.Get(SessionKey).(uuid.UUID)
	return id, ok
}

func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
