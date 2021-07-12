package ColorPrinting

import (
	config "FileSystemTree3/infrastructure"
	"fmt"
)

var (
	colorReset = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"

	useColors     bool
	filesColor    string
	foldersColor  string
	filesizeColor string
	errorColor    string
)

func PrintTree(s string) {
	fmt.Print()
}
func ColorReset() {
	fmt.Print(colorReset)
}

func PrintFolderName(s string) {
	fmt.Print(foldersColor)
	fmt.Print(s)
	fmt.Print(colorReset)
}

func PrintFileName(s string) {
	fmt.Print(filesColor)
	fmt.Print(s)
	fmt.Print(colorReset)
}
func PrintFileSize(s string) {
	fmt.Print(filesizeColor)
	fmt.Print(s)
	fmt.Print(colorReset)
}

func PrintError(s string) {
	fmt.Print(errorColor)
	fmt.Print(s)
	fmt.Print(colorReset)
}

func InitPrinting(c *config.Config) {
	useColors = c.EnableColorText
	filesColor = GetColorByName(c.FilesColor)
	foldersColor = GetColorByName(c.FoldersColor)
	filesizeColor = GetColorByName(c.FileStatsColor)
	errorColor = GetColorByName((c.WarningColor))
}

func GetColorByName(s string) string {
	switch s {
	case "red":
		return colorRed
	case "green":
		return colorGreen
	case "blue":
		return colorBlue
	case "cyan":
		return colorCyan
	case "purple":
		return colorPurple
	case "yellow":
		return colorYellow
	case "white":
		return colorWhite
	default:
		return colorReset

	}
}

//
//fmt.Println(string(colorRed), "test")
//fmt.Println(string(colorGreen), "test")
//fmt.Println(string(colorYellow), "test")
//fmt.Println(string(colorBlue), "test")
//fmt.Println(string(colorPurple), "test")
//fmt.Println(string(colorWhite), "test")
//fmt.Println(string(colorCyan), "test", string(colorReset))
//fmt.Println("next")
