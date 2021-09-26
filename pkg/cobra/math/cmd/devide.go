package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

const (
	ContinueOnParseError = 1 // 解析错误尝试继续处理

	PanicOnDividedByZero = 5 // 除0 painc
)

var (
	dividedByZeroHanding int // 除 0 如何处理
)
var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divide subcommand divide all passed args.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}
		var r float64
		r, _ = strconv.ParseFloat(args[0], 64)
		for _, v := range args[1:] {
			f, _ := strconv.ParseFloat(v, 64)
			r /= f
		}
		fmt.Printf("%s = %.2f\n", strings.Join(args, "/"), r)
	},
}

func init() {
	divideCmd.Flags().IntVarP(&dividedByZeroHanding, "divide_by_zero", "d", int(PanicOnDividedByZero), "do what when divided by zero")

	rootCmd.AddCommand(divideCmd)
}
