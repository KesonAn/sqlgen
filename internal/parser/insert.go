package parser

import (
	"github.com/pingcap/parser/ast"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func parseInsert(stmt *ast.InsertStmt) (*spec.InsertStmt, error) {
	var text = stmt.Text()
	var ret spec.InsertStmt

	tableName, err := parseTableRefsClause(stmt.Table)
	if err != nil {
		return nil, errorNearBy(err, text)
	}

	columns, err := parseColumns(stmt.Columns)
	if err != nil {
		return nil, err
	}

	ret.Table = tableName
	ret.Action = spec.ActionCreate
	ret.SQL = text
	ret.Columns = columns

	return &ret, nil
}