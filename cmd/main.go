package main

import (
	folder "FileSystemTree3/domain"
	config "FileSystemTree3/infrastructure"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()
	_ = conf

	// Print out environment variables
	//fmt.Println(conf.GitHub.Username)
	//fmt.Println(conf.GitHub.APIKey)
	//fmt.Println(conf.DebugMode)
	//fmt.Println(conf.MaxUsers)

	//fmt.Println(conf.TreeCollapseColor, conf.TreeExpandColor, conf.EnabledExtensions, conf.TreeBranchColor, conf.WarningColor,
	//	conf.FoldersColor, conf.FilesColor, conf.SortingIgnoreCase, conf.FileStatsColor)

	//for _, role := range conf.EnabledExtensions {
	//	fmt.Println(role)
	//}
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	RootFolder := folder.NewRootItem(currentDir)
	fmt.Println(RootFolder.FolderName)
	RootFolder.Scan()

	fmt.Println((RootFolder.SubFolders))
	fmt.Println((RootFolder.Files))
}
