package view_test

// import (
// 	"testing"

// 	"github.com/derailed/k9s/internal/dao"
// 	"github.com/derailed/k9s/internal/ui"
// 	"github.com/derailed/k9s/internal/view"
// 	"github.com/stretchr/testify/assert"
// )

// BOZO!!
// func TestHelpNew(t *testing.T) {
// 	ctx := makeCtx()

// 	app := ctx.Value(ui.KeyApp).(*view.App)
// 	po := view.NewPod(dao.GVR("v1/pods"))
// 	po.Init(ctx)
// 	app.Content.Push(po)

// 	v := view.NewHelp()
// 	v.Init(ctx)

// 	assert.Equal(t, 32, v.GetRowCount())
// 	assert.Equal(t, 10, v.GetColumnCount())
// 	assert.Equal(t, "<backspace>", v.GetCell(1, 0).Text)
// 	assert.Equal(t, "Erase", v.GetCell(1, 1).Text)
// }
