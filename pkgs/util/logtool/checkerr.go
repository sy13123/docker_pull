package logtool

import "go.uber.org/zap"



func Fatalerror( _e error ) {
	if _e != nil {
		Logc.Fatal("",zap.Error(_e))
	}
  }

func Errorerror( _e error ) {
	if _e != nil {
		Logc.Error("",zap.Error(_e))
	}
}