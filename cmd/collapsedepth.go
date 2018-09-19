package cmd

import (
	"github.com/fredericlemoine/gotree/io"
	"github.com/spf13/cobra"
)

var maxdepthThreshold int
var mindepthThreshold int

// collapsedepthCmd represents the collapsedepth command
var collapsedepthCmd = &cobra.Command{
	Use:   "depth",
	Short: "Collapse branches having a given depth",
	Long: `Collapse branches having a given depth.

Branches having depth (number of taxa on the lightest side of 
the bipartition) d such that:

min-depth<=d<=max-depth

will be collapsed.

`,
	Run: func(cmd *cobra.Command, args []string) {
		f := openWriteFile(outtreefile)
		defer closeWriteFile(f, outtreefile)

		treefile, treechan := readTrees(intreefile)
		defer treefile.Close()

		for t := range treechan {
			t.Tree.ReinitIndexes()
			if t.Err != nil {
				io.ExitWithMessage(t.Err)
			}
			t.Tree.CollapseTopoDepth(mindepthThreshold, maxdepthThreshold)
			f.WriteString(t.Tree.Newick() + "\n")
		}
	},
}

func init() {
	collapseCmd.AddCommand(collapsedepthCmd)

	collapsedepthCmd.Flags().IntVarP(&mindepthThreshold, "min-depth", "m", 0, "Min depth cutoff to collapse branches")
	collapsedepthCmd.Flags().IntVarP(&maxdepthThreshold, "max-depth", "M", 0, "Max Depth cutoff to collapse branches")
}
