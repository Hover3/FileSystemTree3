package folder

import (
	ColorPrinting "FileSystemTree3/app"
	"fmt"
	"os"
)

type TreeFolderInfo struct {
	Parent     *TreeFolderInfo
	SubFolders []*TreeFolderInfo
	Files      []*FileInfo
	IsScanned  bool
	CantAccess bool
	FolderName string
}

type FileInfo struct {
	FileName   string
	FileSize   int64
	FileExt    string
	CantAccess bool
}

func (f FileInfo) GetFileExtension() string {
	return ""
}

func NewRootItem(absolutePath string) *TreeFolderInfo {
	return &TreeFolderInfo{
		Parent:     nil,
		SubFolders: make([]*TreeFolderInfo, 0),
		Files:      make([]*FileInfo, 0),
		IsScanned:  false,
		FolderName: absolutePath,
		CantAccess: false,
	}
}

func newSubFolderItem(folderName string, parent *TreeFolderInfo) *TreeFolderInfo {
	return &TreeFolderInfo{
		Parent:     parent,
		SubFolders: make([]*TreeFolderInfo, 0),
		Files:      make([]*FileInfo, 0),
		IsScanned:  false,
		CantAccess: false,
		FolderName: folderName,
	}
}

func (f *TreeFolderInfo) ScanRecurrent() {
	f.Scan()
	for _, el := range f.SubFolders {
		el.ScanRecurrent()
	}
}

func (f *TreeFolderInfo) PrintRecurrent(prefix string, last bool) {

	f.Print(prefix, last)
	for i := 0; i < len(f.SubFolders); i++ {
		if i == len(f.SubFolders)-1 {
			if f.Parent == nil {
				f.SubFolders[i].PrintRecurrent(prefix+" ", true)
			} else {
				var s string
				if last == true {
					s = "  "
				} else {
					s = "\u2502" + " "
				}
				f.SubFolders[i].PrintRecurrent(prefix+s, true)
			}
		} else {

			if f.Parent == nil {
				f.SubFolders[i].PrintRecurrent(prefix+" ", false)
			} else {
				var s string
				if last == true {
					s = "  "
				} else {
					s = "\u2502" + " "
				}
				f.SubFolders[i].PrintRecurrent(prefix+s, false)
			}
		}
	}
	f.PrintFiles(prefix + "\u2502" + " ")
}

func (f *TreeFolderInfo) Print(prefix string, last bool) {
	var s string
	if last == true {
		s = "\u2514"
	} else {
		s = "\u251C"
	}
	fmt.Print(prefix, s, "\u2500")
	ColorPrinting.PrintFolderName(f.FolderName)
	if f.CantAccess == true {
		ColorPrinting.PrintError("\t No access!")
	}
	fmt.Print("\n")
}

func (f *TreeFolderInfo) PrintFiles(prefix string) {
	for i := range f.Files {
		var s string
		if i == len(f.Files)-1 {
			s = "\u2514"

		} else {
			s = "\u251C"

		}
		fmt.Print(prefix, s, "\u2500")
		ColorPrinting.PrintFileName(f.Files[i].FileName)
		ColorPrinting.PrintFileSize(fmt.Sprint("\t", "size:", f.Files[i].FileSize, "\n"))
	}
}

func (f *TreeFolderInfo) Scan() error {
	FileSystemObjects, err := os.ReadDir(f.GetAbsolutePath())
	if err != nil {
		f.CantAccess = true
		return err
	}
	for _, fso := range FileSystemObjects {
		if fso.IsDir() {
			tempFolder := newSubFolderItem(fso.Name(), f)
			f.SubFolders = append(f.SubFolders, tempFolder)
			//Adding directory
		} else {
			//adding file
			tempFile := FileInfo{FileName: fso.Name()}
			fileInformation, err := fso.Info()
			if err != nil {
				tempFile.CantAccess = true
			} else {
				tempFile.FileSize = fileInformation.Size()
			}
			f.Files = append(f.Files, &tempFile)
		}

	}
	f.IsScanned = true
	return nil
}

func (f *TreeFolderInfo) GetAbsolutePath() string {
	if f.Parent == nil {
		return f.FolderName
	} else {
		return f.Parent.FolderName + "\\" + f.FolderName
	}
}
