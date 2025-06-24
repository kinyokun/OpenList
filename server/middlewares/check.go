package middlewares

import (
	"strings"

	"github.com/kinyokun/OpenList/internal/conf"
	"github.com/kinyokun/OpenList/pkg/utils"
	"github.com/kinyokun/OpenList/server/common"
	"github.com/gin-gonic/gin"
)

func StoragesLoaded(c *gin.Context) {
	if conf.StoragesLoaded {
		c.Next()
	} else {
		if utils.SliceContains([]string{"", "/", "/favicon.ico"}, c.Request.URL.Path) {
			c.Next()
			return
		}
		paths := []string{"/assets", "/images", "/streamer", "/static"}
		for _, path := range paths {
			if strings.HasPrefix(c.Request.URL.Path, path) {
				c.Next()
				return
			}
		}
		common.ErrorStrResp(c, "Loading storage, please wait", 500)
		c.Abort()
	}
}
