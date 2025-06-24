package doubao_share

import (
	"github.com/kinyokun/OpenList/internal/driver"
	"github.com/kinyokun/OpenList/internal/op"
)

type Addition struct {
	driver.RootPath
	Cookie   string `json:"cookie" type:"text"`
	ShareIds string `json:"share_ids" type:"text" required:"true"`
}

var config = driver.Config{
	Name:              "DoubaoShare",
	LocalSort:         true,
	OnlyLocal:         false,
	OnlyProxy:         false,
	NoCache:           false,
	NoUpload:          true,
	NeedMs:            false,
	DefaultRoot:       "/",
	CheckStatus:       false,
	Alert:             "",
	NoOverwriteUpload: false,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &DoubaoShare{}
	})
}
