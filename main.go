package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jaswdr/faker"
	"github.com/spf13/cobra"
)

func main() {
	var rootDir string
	var depth int
	var numFiles int
	var fileSize int

	rand.Seed(time.Now().UnixNano())

	rootCmd := &cobra.Command{
		Use:   "fauxdir",
		Short: "Create deep nested directory structure with random names and random text files",
		RunE: func(cmd *cobra.Command, args []string) error {
			return exec(rootDir, depth, numFiles, fileSize)
		},
	}

	rootCmd.Flags().StringVarP(&rootDir, "root", "r", "", "root directory for the faux directory structure")
	rootCmd.MarkFlagRequired("root")
	rootCmd.Flags().IntVarP(&depth, "depth", "d", 5, "depth of the faux directory structure")
	rootCmd.Flags().IntVarP(&numFiles, "num-files", "n", 5, "number of random text files to create in each directory")
	rootCmd.Flags().IntVarP(&fileSize, "size", "s", 20, "size of the files to create in bytes")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func exec(rootDir string, depth int, numFiles int, fileSize int) error {
	err := os.Mkdir(rootDir, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	return createDirTree(rootDir, depth, numFiles, fileSize, faker.New())
}

func createDirTree(rootDir string, depth int, numFiles int, fileSize int, f faker.Faker) error {
	if depth <= 0 {
		return nil
	}

	dirName := sanitize(f.Address().Country())
	dirPath := filepath.Join(rootDir, dirName)

	err := os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	for i := 0; i < numFiles; i++ {
		fileName := sanitize(f.Address().Country()) + ".txt"
		filePath := filepath.Join(dirPath, fileName)

		err := ioutil.WriteFile(filePath, f.Lorem().Bytes(fileSize), 0644)
		if err != nil {
			return err
		}
	}

	return createDirTree(dirPath, depth-1, numFiles, fileSize, f)
}

func sanitize(input string) string {
	return strings.ToLower(strings.Join(strings.Split(input, " "), "-"))
}
