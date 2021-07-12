package folder

import "os"

type TreeFolderInfo struct {
	Parent     *TreeFolderInfo
	SubFolders []TreeFolderInfo
	Files      []FileInfo
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
		SubFolders: make([]TreeFolderInfo, 0),
		Files:      make([]FileInfo, 0),
		IsScanned:  false,
		FolderName: absolutePath,
		CantAccess: false,
	}
}

func newSubFolderItem(folderName string, parent *TreeFolderInfo) *TreeFolderInfo {
	return &TreeFolderInfo{
		Parent:     parent,
		SubFolders: make([]TreeFolderInfo, 0),
		Files:      make([]FileInfo, 0),
		IsScanned:  false,
		CantAccess: false,
		FolderName: folderName,
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
			f.SubFolders = append(f.SubFolders, *tempFolder)
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
			f.Files = append(f.Files, tempFile)
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
