package main

import (
	"errors"

	"github.com/overred/xout"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithError(errors.New("wow err")).Info("TEST TEST")
	logrus.Infoln("TEST TEST")

	x := xout.NewPresetText()
	x.WithError(errors.New("OOOPS")).Info("YAY!")
}
