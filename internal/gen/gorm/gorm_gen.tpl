// generated by sqlgen, do not edit.

package model

import (
    "context"
    "fmt"
    "time"

    "gorm.io/gorm"

    "github.com/shopspring/decimal"
)

// {{UpperCamel $.Table.Name}}Model represents a {{$.Table.Name}} model.
type {{UpperCamel $.Table.Name}}Model struct {
    db gorm.DB
}

// {{UpperCamel $.Table.Name}} represents a {{$.Table.Name}} struct data.
type {{UpperCamel $.Table.Name}} struct { {{range $.Table.Columns}}
    {{UpperCamel .Name}} {{.GoType}} `gorm:"{{if IsPrimary .Name}}primaryKey;{{end}}{{if .AutoIncrement}}autoIncrement;{{end}}column:{{.Name}}" json:"{{LowerCamel .Name}}"`{{end}}
}
{{range $stmt := .SelectStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Having.IsValid}}{{$stmt.Having.ParameterStructure "Having"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}{{$stmt.ReceiverStructure}}
{{end}}

{{range $stmt := .UpdateStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}
{{end}}

{{range $stmt := .DeleteStmt}}{{if $stmt.Where.IsValid}}{{$stmt.Where.ParameterStructure "Where"}}
{{end}}{{if $stmt.Limit.Multiple}}{{$stmt.Limit.ParameterStructure}}
{{end}}
{{end}}

// TableName returns the table name. it implemented by gorm.Tabler.
func ({{UpperCamel $.Table.Name}}) TableName() string {
    return "{{$.Table.Name}}"
}

// Create creates  {{$.Table.Name}} data.
func (m *{{UpperCamel $.Table.Name}}Model) Create(ctx context.Context, data ...*{{UpperCamel $.Table.Name}}) error {
    if len(data)==0{
        return fmt.Errorf("data is empty")
    }

    db:=m.db.WithContext(ctx)
    return db.Create(&data).Error
}
{{range $stmt := .SelectStmt}}
// {{.FuncName}} is generated from sql:
// {{$stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Having.IsValid}}, having {{$stmt.Having.ParameterStructureName "Having"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}})({{if $stmt.Limit.One}}*{{$stmt.ReceiverName}}, {{else}}[]*{{$stmt.ReceiverName}}, {{end}} error){
    var result {{if $stmt.Limit.One}} = new({{$stmt.ReceiverName}}){{else}}[]*{{$stmt.ReceiverName}}{{end}}
    var db = m.db.WithContext(ctx)
    db.Select({{$stmt.SelectSQL}})
    {{if $stmt.Where.IsValid}}db.Where({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}})
    {{end }}{{if $stmt.GroupBy.IsValid}}db.Group({{$stmt.GroupBy.SQL}})
    {{end}}{{if $stmt.Having.IsValid}}db.Having({{$stmt.Having.SQL}}, {{$stmt.Having.Parameters "having"}})
    {{end}}{{if $stmt.OrderBy.IsValid}}db.Order({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}db{{if gt $stmt.Limit.Offset 0}}.Offset({{$stmt.Limit.OffsetParameter "limit"}}){{end}}.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}})
    {{end}}db.Find(&result)
    return result, db.Error
}
{{end}}

{{range $stmt := .UpdateStmt}}
// {{.FuncName}} is generated from sql:
// {{$stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context, data *{{UpperCamel $.Table.Name}}{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}}) error {
    var db = m.db.WithContext(ctx)
    db.Model(&{{UpperCamel $.Table.Name}}{})
    {{if $stmt.Where.IsValid}}db.Where({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}})
    {{end}}{{if $stmt.OrderBy.IsValid}}db.Order({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}db{{if gt $stmt.Limit.Offset 0}}.Offset({{$stmt.Limit.OffsetParameter "limit"}}){{end}}.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}})
    {{end}}db.Updates(map[string]interface{}{
        {{range $name := $stmt.Columns}}"{{$name}}": data.{{UpperCamel $name}},
        {{end}}
    })
    return db.Error
}
{{end}}

{{range $stmt := .DeleteStmt}}
// {{.FuncName}} is generated from sql:
// {{$stmt.SQL}}
func (m *{{UpperCamel $.Table.Name}}Model){{.FuncName}}(ctx context.Context{{if $stmt.Where.IsValid}}, where {{$stmt.Where.ParameterStructureName "Where"}}{{end}}{{if $stmt.Limit.Multiple}}, limit {{$stmt.Limit.ParameterStructureName}}{{end}}) error {
    var db = m.db.WithContext(ctx)
    {{if $stmt.Where.IsValid}}db.Where({{$stmt.Where.SQL}}, {{$stmt.Where.Parameters "where"}})
    {{end}}{{if $stmt.OrderBy.IsValid}}db.Order({{$stmt.OrderBy.SQL}})
    {{end}}{{if $stmt.Limit.IsValid}}db{{if gt $stmt.Limit.Offset 0}}.Offset({{$stmt.Limit.OffsetParameter "limit"}}){{end}}.Limit({{if $stmt.Limit.One}}1{{else}}{{$stmt.Limit.LimitParameter "limit"}}{{end}})
    {{end}}db.Delete(&{{UpperCamel $.Table.Name}}{})
    return db.Error
}
{{end}}