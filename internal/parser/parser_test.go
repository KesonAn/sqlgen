package parser

import (
	_ "embed"
	"fmt"
	"log"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/anqiansong/sqlgen/internal/spec"
)

//go:embed test.sql
var testSql string

func TestParse(t *testing.T) {
	dxl, err := Parse(testSql)
	if err != nil {
		log.Fatal(err)
	}

	ctx, err := spec.From(dxl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ctx[0].SelectStmt[0].OrderBy.ParameterStructure("OrderBy"))
}

func TestFrom(t *testing.T) {
	logx.Disable()
	dxl, err := From("root:mysqlpw@tcp(127.0.0.1:55000)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dxl)
}
