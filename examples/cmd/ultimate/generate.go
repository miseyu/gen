package main

import (
	"github.com/miseyu/gen"
	"github.com/miseyu/gen/examples/conf"
	"github.com/miseyu/gen/examples/dal"
	"github.com/miseyu/gen/examples/dal/model"
	"gorm.io/gorm"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()

	prepare(dal.DB) // prepare table for generate
}

var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	"int":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	"json": func(columnType gorm.ColumnType) string { return "json.RawMessage" },
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,

		WithUnitTest: true,

		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
	})

	g.UseDB(dal.DB)

	g.WithDataTypeMap(dataMap)
	g.WithJSONTagNameStrategy(func(c string) string { return "-" })

	g.ApplyBasic(model.Customer{})
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
