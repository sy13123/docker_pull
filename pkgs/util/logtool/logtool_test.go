package logtool

import (
	"os"
	"testing"

	"go.uber.org/zap"
)

func handleError(_e string) {
	SugLog.Desugar()
}

func TestInitEvent(t *testing.T) {
	InitEvent()
	a := 1
	SugLog.Infof("aa%v", a)
	SugLog.Infow("bb", zap.String("msg","aaac"))
	handleError("666")
	_, b := os.Open("/tmp")
	//Logc.Error("",zap.Error(b))
	Fatalerror(b)
}
