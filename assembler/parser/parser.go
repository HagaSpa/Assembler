package parser

import (
	"fmt"
	"strings"
)

type CommandType string

type Parser struct {
	lines       []string
	commandType CommandType
	symbol      string
	dest        string
	comp        string
	jump        string
}

/*
Parserはファイルの全内容を受け取る必要がある。hasMoreCommands()やadvanceのため。
そのためNew内部で以下のことを行う

byteの配列から 文字列化して改行文字でパース。
linesに1行ごとにいれる.
行頭が//になってるのはコメントとして、削除する。
（行末にコメントも考えられるので、//が出てきたら改行までをトリムする処理を作って共通化）
*/
func New(f []byte) *Parser {
	p := &Parser{}

	// Only Unix
	scontent := strings.Split(string(f), "\n")

	fmt.Printf("File contents: %s", scontent)
	return p
}

// func (p *Parser) hasMoreCommands() bool {
// }
