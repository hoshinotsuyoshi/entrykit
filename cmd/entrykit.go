package main

import (
	"fmt"
	"os"

	"github.com/hoshinotsuyoshi/entrykit"

	// import内のアンスコは「ブランク識別子」
	// http://hogesuke.hateblo.jp/entry/2014/09/12/080005
	// https://golang.org/ref/spec#Import_declarations
	//	"To import a package solely(もっぱら) for its side-effects (initialization),
	//	use the blank identifier as explicit package name:"
	_ "github.com/hoshinotsuyoshi/entrykit/codep"
	_ "github.com/hoshinotsuyoshi/entrykit/prehook"
	_ "github.com/hoshinotsuyoshi/entrykit/render"
	_ "github.com/hoshinotsuyoshi/entrykit/switch"
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
			// $ go run cmd/entrykit.go -v
			// バージョンだよ

			// $ go run cmd/entrykit.go -v
			// したときの
			// len(os.Args)は2,
			// os.Args[0]は "/var/folders/k3/62ynls6s2pz7zhd4fyqzgthc0000gn/T/go-build760712289/command-line-arguments/_obj/exe/entrykit"
			// みたいな文字列になる
			fmt.Println(Version)
			os.Exit(0)
		case "--symlink":
			// {root}/entrykit.go のfunc Symlink()で定義されている
			entrykit.Symlink()
			os.Exit(0)
		}
	}
	entrykit.RunCmd()
}
