package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo-studyroom/kadai1/uesho/imgconv"
)

const CodeFailure = 1

// 引数の解析をする
func setArgs() map[string]string {

	from := flag.String("from", "jpg", "変更前の拡張子")
	to := flag.String("to", "png", "変更後の拡張子")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		return nil
	}
	dir := args[0]

	return map[string]string{"from": *from, "to": *to, "dir": dir}
}

// このアプリの使い方を表示する
func printUsage() {
	fmt.Println("使用方法:")
	fmt.Println("  imgconv [-from=<ext>] [-to=<ext>] target_directory")
	fmt.Println("引数:")
	fmt.Println("  -from=<ext> 変換前の拡張子", imgconv.ValidImageExts, "(default: jpg)")
	fmt.Println("  -to=<ext>   変換後の拡張子", imgconv.ValidImageExts, "(default: png)")
}

func main() {
	args := setArgs()
	if args == nil {
		printUsage()
		os.Exit(CodeFailure)
	}
	fmt.Println(args)

	err := imgconv.Do(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(CodeFailure)
	}
}
