package db

import (
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

const (
	INSERT_TEMPLATE = `func (c *{{.model}}) {{.model}Create(ctx context.Context, model *model.{{.model}}) (err error) {
	count, err := GetMassWithContext(ctx).Insert(model)
	if err != nil {
		return
	}

	if count <= 0 {
		err = errors.New("insert corporateLog error")
	}
	return
}

`
	UPDATE_TEMPLATE = `func (c *{{.model}}Dao) {{.model}}UpdateByItem(ctx context.Context, model *model.{{.model}}) (err error) {
	if model.Id <= 0 {
		err = errors.New("Update id can't not nil")
		return
	}
	_, err = getMassDB().Context(ctx).ID(model.Id).Update(model)
	return
}

`
	GET_TEMPLATE = `func (c *{{.model}}Dao) {{.model}}GetByItem(ctx context.Context, model *model.{{.model}}) (has bool, err error) {
	has, err = getMassDB().Context(ctx).OrderBy("id desc").Get(model)
	if err != nil {
		return
	}
	return
}

`
	ALL = INSERT_TEMPLATE + UPDATE_TEMPLATE + GET_TEMPLATE
)

var tableName string

var t2dCmd = &cobra.Command{
	Use:   "t2d",
	Short: "表结构dao生成",
	Long:  "表结构dao生成",
	Run: func(cmd *cobra.Command, args []string) {
		tmpl, err := template.New("t2d").Parse(ALL)
		if err != nil {
			return
		}
		data := map[string]string{
			"model": tableName,
		}
		err = tmpl.Execute(os.Stdout, data)
		if err != nil {
			return
		}
	},
}

func GetCmd() *cobra.Command {
	return t2dCmd
}
func init() {
	t2dCmd.Flags().StringVarP(&tableName, "table", "t", "", "请输入表明")
}
