package main

import (
	"github.com/kam2yar/user-service/internal"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	zap.L().Info("Running app")

	internal.Bootstrap()
}
