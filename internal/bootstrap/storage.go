package bootstrap

import (
	"context"

	"github.com/kinyokun/OpenList/internal/conf"
	"github.com/kinyokun/OpenList/internal/db"
	"github.com/kinyokun/OpenList/internal/model"
	"github.com/kinyokun/OpenList/internal/op"
	"github.com/kinyokun/OpenList/pkg/utils"
)

func LoadStorages() {
	storages, err := db.GetEnabledStorages()
	if err != nil {
		utils.Log.Fatalf("failed get enabled storages: %+v", err)
	}
	go func(storages []model.Storage) {
		for i := range storages {
			err := op.LoadStorage(context.Background(), storages[i])
			if err != nil {
				utils.Log.Errorf("failed get enabled storages: %+v", err)
			} else {
				utils.Log.Infof("success load storage: [%s], driver: [%s], order: [%d]",
					storages[i].MountPath, storages[i].Driver, storages[i].Order)
			}
		}
		conf.StoragesLoaded = true
	}(storages)
}
