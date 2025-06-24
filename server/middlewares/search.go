package middlewares

import (
	"github.com/kinyokun/OpenList/internal/conf"
	"github.com/kinyokun/OpenList/internal/errs"
	"github.com/kinyokun/OpenList/internal/setting"
	"github.com/kinyokun/OpenList/server/common"
	"github.com/gin-gonic/gin"
)

func SearchIndex(c *gin.Context) {
	mode := setting.GetStr(conf.SearchIndex)
	if mode == "none" {
		common.ErrorResp(c, errs.SearchNotAvailable, 500)
		c.Abort()
	} else {
		c.Next()
	}
}
