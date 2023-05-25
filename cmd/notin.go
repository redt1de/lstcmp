package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// notinCmd represents the notin command
var notinCmd = &cobra.Command{
	Use:                   "notin [FileA] [FileB]",
	DisableFlagParsing:    true,
	Args:                  cobra.ExactArgs(2),
	Short:                 "list lines in file A that are NOT IN file B",
	ArgAliases:            []string{"fileA", "fileB"},
	Long:                  ``,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fileA := args[0]
		fileB := args[1]
		listA, err := readLines(fileA)
		if err != nil {
			log.Fatal("Failed to open File A:", fileA)
		}
		listB, err := readLines(fileB)
		if err != nil {
			log.Fatal("Failed to open File B:", fileB)
		}

		for _, a := range listA {
			if !listContains(listB, a) {
				fmt.Println(a)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(notinCmd)

}


