package db

import (
	_ "embed"
	"errors"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/template"
)

const (
	INSERT_TEMPLATE = `func {{.model}}Create(ctx context.Context, model *model.{{.model}}) (err error) {
	count, err := getDB().Context(ctx).Insert(model)
	if err != nil {
		return
	}

	if count <= 0 {
		err = errors.New("insert {{.model}} error")
	}
	return
}

`
	UPDATE_TEMPLATE = `func (c *{{.model}}) {{.model}}UpdateByItem(ctx context.Context, model *model.{{.model}}) (err error) {
	if model.Id <= 0 {
		err = errors.New("Update id can't not nil")
		return
	}

	count, err = getDB().Context(ctx).ID(model.Id).Update(model)
	if err != nil {
		return
	}

	if count <= 0 {
		err = errors.New("Update by it return 0")
		return
	}
	return
}

`
	GET_TEMPLATE = `func (c *{{.model}}) {{.model}}GetByItem(ctx context.Context, model *model.{{.model}}) (has bool, err error) {
	has, err = getDB().Context(ctx).OrderBy("id desc").Get(model)
	return
}

`
	ALL = INSERT_TEMPLATE + UPDATE_TEMPLATE + GET_TEMPLATE
)

//go:embed sql.tmpl
var templateStr string

var tableName string

func init() {
	t2dCmd.Flags().StringVarP(&tableName, "table", "t", "", "请输入表明")
}

var t2dCmd = &cobra.Command{
	Use:   "t2d",
	Short: "表结构dao生成",
	Long:  "表结构dao生成",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err error
		)
		defer func() {
			if err != nil {
				log.Fatal(err)
			}
		}()
		if len(tableName) <= 0 {
			err = errors.New("tableName not null")
			return
		}
		tmp := template.New("tmp")
		parse, err := tmp.Parse(templateStr)
		if err != nil {
			return
		}
		data := map[string]string{
			"model": tableName,
		}
		err = parse.Execute(os.Stdout, data)
		if err != nil {
			return
		}
	},
}

func GetCmd() *cobra.Command {
	return t2dCmd
}
