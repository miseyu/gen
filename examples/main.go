package main

import (
	"context"
	"fmt"

	"github.com/miseyu/gen/examples/biz"
	"github.com/miseyu/gen/examples/conf"
	"github.com/miseyu/gen/examples/dal"
	"github.com/miseyu/gen/examples/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
