package cmd

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

var countBytes bool
var countLines bool
var countWords bool
var countChars bool

var rootCmd = &cobra.Command{
	Use:   "ccwc [flags] [input text]",
	Short: "A unix like wc binary written in golang",
	RunE:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	var file *os.File
	var err error
	var fileName string

	if len(args) > 0 {
		fileName = args[0]
		file, err = os.OpenFile(fileName, os.O_RDONLY, 0444)
		if err != nil {
			fmt.Printf("failed to open file: %s \n", err)
			os.Exit(1)
		}
		defer file.Close()
		calcOutput(file, fileName)
	} else {
		calcOutput(bufio.NewReader(os.Stdin), "")
	}
	return nil
}

func calcOutput(r io.Reader, fileName string) {
	b, err := io.ReadAll(r)
	if err != nil {
		fmt.Printf("failed to read file: %s \n", err)
		os.Exit(1)
	}

	var byteSizes string
	var lines string
	var words string

	if !countBytes && !countChars && !countLines && !countWords {
		lines = calcLineCounts(b)
		words = calcWordCounts(b)
		byteSizes = calcByteCounts(b)
		fmt.Println(lines, words, byteSizes, fileName)
		return
	}

	if countLines {
		lines = calcLineCounts(b)
	}
	if countWords {
		words = calcWordCounts(b)
	}
	if countBytes {
		byteSizes = calcByteCounts(b)
	}
	if countChars {
		byteSizes = calcCharCounts(b)
	}
	fmt.Println(lines, words, byteSizes, fileName)
}

func calcByteCounts(b []byte) string {
	return strconv.Itoa(binary.Size(b))
}

func calcLineCounts(b []byte) string {
	r := bytes.NewReader(b)
	var count int
	var err error
	var target []byte = []byte("\n")
	buffer := make([]byte, 32*1024)

	for {
		read, err := r.Read(buffer)
		if err != nil {
			break
		}
		count += bytes.Count(buffer[:read], target)
	}

	if err != nil {
		fmt.Printf("failed to count line: %s \n", err)
		os.Exit(1)
	}

	return strconv.Itoa(count)
}

func calcWordCounts(b []byte) string {
	str := string(b)
	return strconv.Itoa(len(strings.Fields(str)))
}

func calcCharCounts(b []byte) string {
	str := string(b)
	return strconv.Itoa(utf8.RuneCountInString(str))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&countBytes, "countbytes", "c", false, "The number of bytes in each input file is written to the standard outputs")
	rootCmd.PersistentFlags().BoolVarP(&countLines, "countlines", "l", false, "The number of lines in each input file is written to the standard outputs")
	rootCmd.PersistentFlags().BoolVarP(&countWords, "countwords", "w", false, "The number of words in each input file is written to the standard outputs")
	rootCmd.PersistentFlags().BoolVarP(&countChars, "countchars", "m", false, "The number of chars in each input file is written to the standard outputs")
}
