package grifts

import (
	"firstproject/actions"
	
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
