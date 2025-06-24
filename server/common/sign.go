package common

import (
	stdpath "path"

	"github.com/kinyokun/OpenList/internal/conf"
	"github.com/kinyokun/OpenList/internal/model"
	"github.com/kinyokun/OpenList/internal/setting"
	"github.com/kinyokun/OpenList/internal/sign"
)

func Sign(obj model.Obj, parent string, encrypt bool) string {
	if obj.IsDir() || (!encrypt && !setting.GetBool(conf.SignAll)) {
		return ""
	}
	return sign.Sign(stdpath.Join(parent, obj.GetName()))
}
