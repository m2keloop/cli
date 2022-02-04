package cmd

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"time"
)

const (
	UNIX = iota + 1
	ADD
)

var timeMode int8
var content string

func init() {
	timeCmd.Flags().Int8VarP(&timeMode, "mode", "m", 0, "请输入模式")
	timeCmd.Flags().StringVarP(&content, "content", "c", "", "时间内容")
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间相关命令",
	Long:  "时间相关命令集",
	Run: func(cmd *cobra.Command, args []string) {
		switch timeMode {
		case UNIX:
			FormatPrint(time.Unix(cast.ToInt64(content), 0))
		case ADD:
			FormatPrint(time.Now().Add(cast.ToDuration(content) * time.Hour))
		default:
			fmt.Printf("未知模式")
		}
	},
}

func FormatPrint(time2 time.Time) {
	format := "2006-01-02 15:04:05"
	parse := time2.Format(format)
	fmt.Printf("time:%v\n", parse)
}
