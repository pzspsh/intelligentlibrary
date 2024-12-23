/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:30:17
*/
package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

func Match(exprRule string, sourceData map[string]interface{}) (bool, error) {
	if len(exprRule) == 0 {
		return false, errors.New("empty exprRule")
	}

	// 解析表达式
	exprAst, err := parser.ParseExpr(exprRule)
	if err != nil {
		return false, fmt.Errorf("parse exprRule err: %w", err)
	}
	//fset := token.NewFileSet()
	//ast.Print(fset, exprAst)

	res, err := judge(exprAst, sourceData)
	if err != nil {
		return false, fmt.Errorf("judge err: %w", err)
	}
	return res, nil
}

// 递归解析ast
func judge(expr ast.Expr, sourceData map[string]interface{}) (bool, error) {

	switch t := expr.(type) {
	case *ast.BinaryExpr:
		if isBinaryLeaf(t) {
			return cmpBinary(t, sourceData)
		}
		// 递归比较
		lRes, err := judge(t.X, sourceData)
		// fmt.Printf("-- lr: %v, op: %v \n", lRes, t.Op)
		if err != nil {
			return false, err
		}
		rRes, err := judge(t.Y, sourceData)
		//fmt.Printf("## lr: %v, rr: %v, op: %v \n", lRes, rRes, t.Op)
		if err != nil {
			return false, err
		}

		switch t.Op {
		case token.LAND:
			return lRes && rRes, nil
		case token.LOR:
			return lRes || rRes, nil
		}
		return false, fmt.Errorf("not support op xx")
	case *ast.CallExpr: // 匹配到函数
		return matchFunc(t, sourceData)
	case *ast.ParenExpr:
		return judge(t.X, sourceData)
	default:
		return false, errors.New(fmt.Sprintf("%#v type is not support", expr))
	}
}

func isBinaryLeaf(expr *ast.BinaryExpr) bool {
	// 二元表达式的最小单位，左节点是Ident，右节点是BasicLit
	_, lType := expr.X.(*ast.Ident)
	_, rType := expr.Y.(*ast.BasicLit)
	return lType && rType
}

const (
	funcInData  = "in_data"
	funcInArray = "in_array"
)

func matchFunc(expr *ast.CallExpr, sourceData map[string]interface{}) (bool, error) {
	funIdent, ok := expr.Fun.(*ast.Ident)
	if !ok {
		return false, fmt.Errorf("CallExpr node error, node info: %v", expr)
	}

	switch funIdent.Name {
	case funcInData:
		return InOp(expr.Args, sourceData)
	default:
		return false, fmt.Errorf("not support func: %s", funIdent.Name)
	}
}

func InOp(args []ast.Expr, sourceData map[string]interface{}) (bool, error) {
	if len(args) != 2 {
		return false, fmt.Errorf("in_data args length error")
	}
	key, ok := args[0].(*ast.BasicLit)
	if !ok || key.Kind != token.STRING {
		return false, fmt.Errorf("in_data args key type error, node: %v", args)
	}
	data, ok := sourceData[trimQuotes(key.Value)]
	if !ok {
		return false, fmt.Errorf("in_data args key not exists in source data, node: %v", args)
	}
	dataString, err := convToString(data)
	if err != nil {
		return false, err
	}

	valBasicLit, ok := args[1].(*ast.BasicLit)
	if !ok {
		return false, fmt.Errorf("in_data args val type error, node: %#v", args)
	}
	valList := strings.Split(trimQuotes(valBasicLit.Value), "|")
	for _, val := range valList {
		if val == dataString {
			return true, nil
		}
	}

	return false, nil
}
func cmpBinary(expr *ast.BinaryExpr, sourceData map[string]interface{}) (bool, error) {
	var (
		xName string
		xVal  interface{}
	)

	xName = expr.X.(*ast.Ident).Name
	xVal, ok := sourceData[xName]
	if !ok {
		return false, fmt.Errorf("source data key: %s not exists", xName)
	}
	y := expr.Y.(*ast.BasicLit)

	// 根据不同类型比较
	switch y.Kind {
	case token.INT: // 转换为int64比较
		xValInt64, err := convToInt64(xVal)
		if err != nil {
			return false, err
		}
		yInt64, err := strconv.ParseInt(y.Value, 10, 64)
		if err != nil {
			return false, err
		}
		return cmpInt64(xValInt64, yInt64, expr.Op)
	case token.STRING:
		xValString, err := convToString(xVal)
		if err != nil {
			return false, err
		}
		return cmpString(xValString, strings.Trim(y.Value, "\""), expr.Op)
	}

	return false, fmt.Errorf("cmpBinary error")
}

func cmpString(x, y string, op token.Token) (bool, error) {
	switch op {
	case token.EQL:
		return x == y, nil
	default:
		return false, fmt.Errorf("not support string op: %v", op)
	}
}

func cmpInt64(x, y int64, op token.Token) (bool, error) {
	switch op {
	case token.EQL:
		return x == y, nil
	case token.LSS:
		return x < y, nil
	case token.GTR:
		return x > y, nil
	case token.NEQ:
		return x != y, nil
	case token.LEQ:
		return x <= y, nil
	case token.GEQ:
		return x >= y, nil
	default:
		return false, fmt.Errorf("not support number op: %v", op)
	}
	// return false, fmt.Errorf("not support number op: %v", op)
}

func convToInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case int64:
		return v, nil
	case string:
		vI64, e := strconv.ParseInt(v, 10, 64)
		return vI64, e
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	}
	return int64(0), fmt.Errorf("data: %v convert to int64 err", v)
}

func convToString(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case int32:
		return strconv.Itoa(int(v)), nil
	case int64:
		return strconv.Itoa(int(v)), nil
	default:
		return "", fmt.Errorf("data: %v convert to string err", v)
	}
}

func trimQuotes(str string) string {
	return strings.Trim(str, "\"")
}

type rule struct {
	data map[string]interface{}
	rule string
	res  bool
}

func main() {
	   ruleList := []rule{
      {
         data: map[string]interface{}{
            "name": "cooper",
            "age":  11,
            "sex":  2,
         },
         rule: `(age > 10 && in_data("name", "cooper|jack")) || sex == 1`,
         res:  true,
      },
      {
         data: map[string]interface{}{
            "sex":  11,
            "addr": "aaa",
         },
         rule: `addr == "aaa" && sex == 1`,
         res:  false,
      },
   }

   for _, ca := range ruleList {
      res, err := Match(ca.rule, ca.data)
      if err != nil && res != ca.res {
         fmt.Printf("res: %v, err: %v", ca.res, err)
      }
   }
}