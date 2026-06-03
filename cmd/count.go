package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(countCmd)
}

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "count all files in current folder and subfolders",
	Long:  `Count all files in current folder and subfolders`,
	Run:   Count,
}

func Count(cmd *cobra.Command, args []string) {
	root, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fileCount, folderCount, err := countFiles(root)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%d files in %d subfolders\n", fileCount, folderCount)
}

func countFiles(root string) (files, folders int, err error) {
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			folders++ // Directory
		} else {
			files++ // Regular file
		}
		return nil
	})
	return files, folders - 1, err // Minus one folder so we don't count where we started
}
