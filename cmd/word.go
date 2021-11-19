package cmd

import (
	"cli/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换,模式如下：",
	"1:全部单词转为大写",
	"2:全部单词转为小写",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help命令查看帮助文档")
		}
		log.Printf("输出结果：%v", content)
	},
}
var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
