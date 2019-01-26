package utils

import (
	"github.com/common-nighthawk/go-figure"
)

func ShowBanner() {

  myFigure := figure.NewFigure("Hackalist-CLI", "doom", true)
  myFigure.Print()
}