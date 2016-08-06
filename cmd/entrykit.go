package main

import (
	"fmt"
	"os"

	"github.com/progrium/entrykit"

	_ "github.com/progrium/entrykit/codep"
	_ "github.com/progrium/entrykit/prehook"
	_ "github.com/progrium/entrykit/render"
	_ "github.com/progrium/entrykit/switch"
)

var Version string

// $ go run cmd/entrykit.goするとここを通るよ
func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			// ここで
			// Version = "バージョンだよ"
			// とすると以下のように出力されるよ
			// [/Users/berlin/go/src/github.com/hoshinotsuyoshi/entrykit]$ go run cmd/entrykit.go -v
			// バージョンだよ
			fmt.Println(Version)
			os.Exit(0)
		case "--symlink":
			entrykit.Symlink()
			os.Exit(0)
		}
	}
	entrykit.RunCmd()
}
