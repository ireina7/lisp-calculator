package service

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/result"
	"github.com/ireina7/lisp-calculator/model"
)

type Parser[A any] interface {
	Parse(string) result.Result[A]
}

type lispParser struct {
	interfaces.Logger
}

func NewLispParser() *lispParser {
	return &lispParser{}
}

func (parser *lispParser) Parse(src string) result.Result[model.LispExpr] {
	err := result.FromErr[model.LispExpr](nil)
	if len(src) == 0 {
		return err
	}
	if src[0] == '(' {
		_, list, err := parser.parseList(src)
		if err != nil {
			return result.FromErr[model.LispExpr](err)
		}
		return result.From[model.LispExpr](list)
	} else if unicode.IsDigit(rune(src[0])) {
		_, num, err := parser.parseNumber(src)
		if err != nil {
			return result.FromErr[model.LispExpr](err)
		}
		return result.From[model.LispExpr](model.Integer(num))
	} else {
		_, idem, err := parser.parseIdem(src)
		if err != nil {
			return result.FromErr[model.LispExpr](err)
		}
		return result.From[model.LispExpr](model.Name(idem))
	}
}

func (parser *lispParser) parseNumber(src string) (string, int, error) {
	src = strings.Trim(src, " ")
	if len(src) == 0 {
		return "", 0, fmt.Errorf("empty src")
	}
	numBuffer := make([]rune, 0)
	index := 0
	for i, c := range src {
		if !unicode.IsDigit(c) {
			index = i
			break
		}
		numBuffer = append(numBuffer, c)
	}
	num, err := strconv.ParseInt(string(numBuffer), 10, 64)
	if err != nil {
		return "", 0, err
	}
	return string([]rune(src)[index:]), int(num), nil
}

func (parser *lispParser) parseIdem(src string) (rest string, idem string, err error) {
	src = strings.Trim(src, " ")
	if len(src) == 0 {
		return "", "", fmt.Errorf("empty src")
	}
	strBuffer := make([]rune, 0)
	index := 0
	for i, c := range src {
		if c == '(' || c == ')' || unicode.IsSpace(c) {
			index = i
			break
		}
		strBuffer = append(strBuffer, c)
	}
	return string([]rune(src)[index:]), string(strBuffer), nil
}

func (parser *lispParser) parseList(src string) (rest string, list *model.List, err error) {
	src = strings.Trim(src, " ")
	if len(src) == 0 {
		return "", nil, fmt.Errorf("empty src")
	}
	if src[0] != '(' {
		return "", nil, fmt.Errorf("not a list")
	}
	ans := model.List{
		Op:   "",
		Args: []model.LispExpr{},
	}
	src, op, err := parser.parseIdem(src[1:])
	if err != nil {
		return "", nil, err
	}
	ans.Op = model.Name(op)
	src = strings.Trim(src, " ")
	for len(src) > 0 && src[0] != ')' {
		if src[0] == '(' {
			rest, list, err := parser.parseList(src)
			if err != nil {
				return "", nil, err
			}
			ans.Args = append(ans.Args, list)
			src = strings.Trim(rest, " ")
		} else if unicode.IsDigit(rune(src[0])) {
			rest, num, err := parser.parseNumber(src)
			if err != nil {
				return "", nil, err
			}
			ans.Args = append(ans.Args, model.Integer(num))
			src = strings.Trim(rest, " ")
		} else {
			rest, idem, err := parser.parseIdem(src)
			if err != nil {
				return "", nil, err
			}
			ans.Args = append(ans.Args, model.Name(idem))
			src = strings.Trim(rest, " ")
		}
	}
	if len(src) == 0 {
		return "", &ans, nil
	}
	return src[1:], &ans, nil
}
