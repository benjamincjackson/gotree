package cmd

import (
	"errors"
	"fmt"
	"github.com/fredericlemoine/gotree/io"
	"github.com/fredericlemoine/gotree/io/utils"
	"github.com/spf13/cobra"
	"os"
)

var difftipsTree1 string
var difftipsTree2 string

// difftipsCmd represents the difftips command
var difftipsCmd = &cobra.Command{
	Use:   "difftips",
	Short: "difftips prints the diff between tip names of two trees",
	Long: `difftips prints the diff between tip names of two trees.

For example:
t1.nh : (t1,t2,(t3,t4));
t2.nh : (t10,t2,(t3,t4));

gotree difftips -i t1.nh -c t2.nh

should produce the following output:
< t1
> t10
= 3

`,
	Run: func(cmd *cobra.Command, args []string) {
		if difftipsTree2 == "none" {
			io.ExitWithMessage(errors.New("Compare tree file must be provided with -c"))
		}
		eq := 0
		if refTree, err := utils.ReadRefTree(difftipsTree1); err != nil {
			io.ExitWithMessage(err)
		} else if compTree, err2 := utils.ReadRefTree(difftipsTree2); err2 != nil {
			io.ExitWithMessage(err2)
		} else {
			for _, t := range refTree.Tips() {
				if ok, err3 := compTree.ExistsTip(t.Name()); err3 != nil {
					io.ExitWithMessage(err)
				} else {
					if !ok {
						fmt.Fprintf(os.Stdout, "< %s\n", t.Name())
					} else {
						eq++
					}
				}
			}
			for _, t := range compTree.Tips() {
				if ok, err4 := refTree.ExistsTip(t.Name()); err4 != nil {
					io.ExitWithMessage(err)
				} else {
					if !ok {
						fmt.Fprintf(os.Stdout, "> %s\n", t.Name())
					}
				}
			}
			fmt.Fprintf(os.Stdout, "= %d\n", eq)
		}
	},
}

func init() {
	RootCmd.AddCommand(difftipsCmd)
	difftipsCmd.Flags().StringVarP(&difftipsTree1, "reftree", "i", "stdin", "Reference tree input file")
	difftipsCmd.Flags().StringVarP(&difftipsTree2, "compared", "c", "none", "Other tree file to compare with")
}
