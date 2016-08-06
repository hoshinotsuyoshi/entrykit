package main

import (
	"github.com/hoshinotsuyoshi/entrykit"
	_ "github.com/hoshinotsuyoshi/entrykit/codep"
)

var cmd = "codep"

func main() {
	entrykit.Cmds[cmd](
		entrykit.NewConfig(cmd, true))
}
