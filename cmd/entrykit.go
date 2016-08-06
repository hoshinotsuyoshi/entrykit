package main

import (
	"fmt"
	"os"

	"github.com/progrium/entrykit"

	// import内のアンスコは「ブランク識別子」
	// http://hogesuke.hateblo.jp/entry/2014/09/12/080005
	// https://golang.org/ref/spec#Import_declarations
	//	"To import a package solely(もっぱら) for its side-effects (initialization),
	//	use the blank identifier as explicit package name:"
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
