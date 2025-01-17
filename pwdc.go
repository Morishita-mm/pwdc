package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// フラグを格納する構造体
type flags struct {
	logical  bool
	physical bool
	colorize bool
	space    bool
	tree     bool
}

func parseFlags() flags {
	// 大文字小文字を区別せずにフラグを処理
	logical := flag.Bool("L", true, "Use logical path (default)")
	physical := flag.Bool("P", false, "Use physical path")
	colorize := flag.Bool("C", false, "Colorize each directory level")
	space := flag.Bool("S", false, "Add spaces between directory names")
	tree := flag.Bool("T", false, "Display the path in a tree-like format")

	// 小文字のエイリアスも設定
	flag.BoolVar(logical, "l", true, "Use logical path (default)")
	flag.BoolVar(physical, "p", false, "Use physical path")
	flag.BoolVar(colorize, "c", false, "Colorize each directory level")
	flag.BoolVar(space, "s", false, "Add spaces between directory names")
	flag.BoolVar(tree, "t", false, "Display the path in a tree-like format")

	flag.Parse()

	return flags{
		logical:  *logical,
		physical: *physical,
		colorize: *colorize,
		space:    *space,
		tree:     *tree,
	}
}

func main() {
	// フラグの解析
	opts := parseFlags()

	// 論理パスと物理パスの優先順位（pwdと同じ動作）
	if opts.physical {
		opts.logical = false
	}

	// 現在のパスを取得
	var pwd string
	var err error
	if opts.logical {
		pwd, err = os.Getwd()
	} else {
		pwd, err = filepath.EvalSymlinks(os.Getenv("PWD"))
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// 階段状出力が有効な場合
	if opts.tree {
		printTreeFormat(pwd, opts)
		return
	}

	// 通常のフォーマット出力
	printStandardFormat(pwd, opts)
}

func printTreeFormat(pwd string, opts flags) {
	segments := strings.Split(pwd, string(os.PathSeparator))
	for i, segment := range segments {
		if segment == "" {
			fmt.Println("/")
			continue
		}

		prefix := strings.Repeat("  ", i) // 階段状のインデント

		if opts.colorize {
			// 階層ごとに異なる色を設定
			c := color.New(color.Attribute(30 + (i % 8))) // 8色をループ
			segment = c.Sprint(segment)
		}

		// スペースを加えるオプションの処理
		if opts.space {
			segment = strings.ReplaceAll(segment, "/", " / ")
		}

		fmt.Printf("%s%s\n", prefix, segment)
	}
}

func printStandardFormat(pwd string, opts flags) {
	segments := strings.Split(pwd, string(os.PathSeparator))
	var result string
	for i, segment := range segments {
		if segment == "" {
			continue // ルート("/")の処理
		}

		if opts.colorize {
			// 階層ごとに異なる色を設定
			c := color.New(color.Attribute(30 + (i % 8))) // 8色をループ
			segment = c.Sprint(segment)
		}

		if opts.space {
			result += segment + " / "
		} else {
			result += string(os.PathSeparator) + segment
		}
	}

	// 出力の整形
	if opts.space {
		result = strings.TrimSuffix(result, " / ")
	}
	fmt.Println(result)
}
