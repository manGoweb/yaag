package gin

import (
	"log"
	"net/http/httptest"
	"strings"

	"github.com/manGoweb/yaag/middleware"
	"github.com/manGoweb/yaag/yaag"
	"github.com/manGoweb/yaag/yaag/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func Document() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !yaag.IsOn() {
			return
		}
		writer := httptest.NewRecorder()
		apiCall := models.ApiCall{}
		middleware.Before(&apiCall, c.Request)
		c.Next()
		if writer.Code != 404 {
			apiCall.MethodType = c.Request.Method
			apiCall.CurrentPath = strings.Split(c.Request.RequestURI, "?")[0]
			apiCall.ResponseBody = ""
			apiCall.ResponseCode = c.Writer.Status()
			headers := map[string]string{}
			for k, v := range c.Writer.Header() {
				log.Println(k, v)
				headers[k] = strings.Join(v, " ")
			}
			apiCall.ResponseHeader = headers
			go yaag.GenerateHtml(&apiCall)
		}
	}
}
