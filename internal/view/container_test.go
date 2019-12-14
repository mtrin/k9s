package view_test

import (
	"testing"

	"github.com/derailed/k9s/internal/dao"
	"github.com/derailed/k9s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestContainerNew(t *testing.T) {
	po := view.NewContainer(dao.GVR("containers"))
	po.Init(makeCtx())

	assert.Equal(t, "Containers", po.Name())
	assert.Equal(t, 19, len(po.Hints()))
}
