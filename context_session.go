package xingyun

import (
	"fmt"
)

var (
	SessionKey   string = "ZQSESSID"
	sessionIDLen int    = 36
)

func newSessionID() string {
	return GenRandString(sessionIDLen)
}

func (ctx *Context) SetSession(key string, data []byte) {
	ctx.Server.SessionStorage.SetSession(ctx.GetSessionID(), key, data)
}

func (ctx *Context) GetSession(key string) []byte {
	return ctx.Server.SessionStorage.GetSession(ctx.GetSessionID(), key)
}

func (ctx *Context) ClearSession(key string) {
	ctx.Server.SessionStorage.ClearSession(ctx.GetSessionID(), key)
}

func (ctx *Context) setNewSessionID() (sessionID string) {
	sessionID = newSessionID()
	ctx.SetCookie(SessionKey, sessionID)
	return
}

// SetCookie adds a cookie header to the response.
func (ctx *Context) GetSessionID() (sessionID string) {
	var cookieVal string
	ctx.GetCookie(SessionKey,&cookieVal)
	
	fmt.Println("GetSessionID",cookieVal)
	
	if cookieVal == "" || len(cookieVal) != sessionIDLen {
		return ctx.setNewSessionID()
	}
	return cookieVal
}
