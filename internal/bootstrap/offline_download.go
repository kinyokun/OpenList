package bootstrap

import (
	"github.com/kinyokun/OpenList/internal/offline_download/tool"
	"github.com/kinyokun/OpenList/pkg/utils"
)

func InitOfflineDownloadTools() {
	for k, v := range tool.Tools {
		res, err := v.Init()
		if err != nil {
			utils.Log.Warnf("init tool %s failed: %s", k, err)
		} else {
			utils.Log.Infof("init tool %s success: %s", k, res)
		}
	}
}
