package logging

import (
	"net/http"

	"github.com/gin-gonic/gin"

	cfg "github.com/nannigalaxy/go-rest-api/app/config"

	"github.com/rollbar/rollbar-go"
)

func SetRollbarConfig() {

	rollbar.SetToken(cfg.Config.RollbarToken)
	rollbar.SetEnvironment(cfg.Config.Env)
	rollbar.SetCodeVersion(cfg.Config.RollbarCodeVer)
	rollbar.SetServerHost(cfg.Config.RollbarHost)
	rollbar.SetServerRoot(cfg.Config.RollbarRoot)

}

func RollbarMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				rollbar.Critical(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				c.Abort()
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			for _, ginErr := range c.Errors {
				rollbar.Error(ginErr.Err)
			}
		}
	}
}
