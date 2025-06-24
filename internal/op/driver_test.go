package op_test

import (
	"testing"

	_ "github.com/kinyokun/OpenList/drivers"
	"github.com/kinyokun/OpenList/internal/op"
)

func TestDriverItemsMap(t *testing.T) {
	itemsMap := op.GetDriverInfoMap()
	if len(itemsMap) != 0 {
		t.Logf("driverInfoMap: %v", itemsMap)
	} else {
		t.Errorf("expected driverInfoMap not empty, but got empty")
	}
}
