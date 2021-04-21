package parser

import (
	"fmt"
	"strconv"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/scanner"
	"github.com/nkanaev/numb/pkg/value"
)

type parser struct {
	s *scanner.Scanner
}

func (p *parser) parseUnaryExpr() ast.Node {
	switch p.s.Token {
	case scanner.NUM:
		num, err := strconv.Atoi(p.s.Value)
		if err != nil {
			panic(err)
		}
		p.s.Scan()
		// TODO: pass string directly
		return value.NewInt(int64(num))
	}	
	fmt.Println(p.s.Token)
	panic("die")
}

func (p *parser) parseBinaryExpr(prec1 int) ast.Node {
	lhs := p.parseUnaryExpr()
	for {
		prec := p.s.Token.Precedence()
		if prec < prec1 {
			break
		}
		tok := p.s.Token
		p.s.Scan()
		rhs := p.parseBinaryExpr(prec + 1)
		lhs = &ast.BinOP{Lhs: lhs, Rhs: rhs, Op: tok}
	}
	return lhs
}

func Parse(line string) ast.Node {
	s := scanner.New(line)
	p := &parser{s: s}
	s.Scan()
	return p.parseBinaryExpr(scanner.LowestPrec+1)
}