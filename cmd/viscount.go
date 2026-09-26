package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(viscountCmd)
}

var viscountCmd = &cobra.Command{
	Use:   "viscount",
	Short: "count files in subfolders",
	Long:  `count files in every direct subfolder of working directory`,
	Run:   Viscount,
}

func Viscount(cmd *cobra.Command, args []string) {
	root, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	err = viscountFiles(root)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

func viscountFiles(root string) (err error) {
	dir, err := os.ReadDir(root)
	if err != nil {
		return err
	}
	fmt.Printf("root: %s\n", root)
	for _, d := range dir {
		if d.IsDir() {
			n := root + "/" + d.Name()
			fileCount, _, err := countFiles(n)
			if err != nil {
				return err
			}
			s := fmt.Sprintf("%d                   ", fileCount)
			fmt.Printf("   %s   %s\n", s[:7], n)
		}
	}
	return err
}
