package folder

import "os"

type FolderInfo struct {
	Parent     *FolderInfo
	SubFolders []FolderInfo
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

func New(absolutePath string) *FolderInfo {
	return &FolderInfo{
		Parent:     nil,
		SubFolders: make([]FolderInfo, 0),
		Files:      make([]FileInfo, 0),
		IsScanned:  false,
		FolderName: absolutePath,
		CantAccess: false,
	}
}

func new(folderName string, parent *FolderInfo) *FolderInfo {
	return &FolderInfo{
		Parent:     parent,
		SubFolders: make([]FolderInfo, 0),
		Files:      make([]FileInfo, 0),
		IsScanned:  false,
		CantAccess: false,
		FolderName: folderName,
	}
}

func (f FolderInfo) Scan() error {
	FileSystemObjects, err := os.ReadDir(f.GetAbsolutePath())
	if err != nil {
		f.CantAccess = true
		return err
	}
	for _, fso := range FileSystemObjects {
		if fso.IsDir() {
			tempFolder := FolderInfo{FolderName: fso.Name()}
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
			f.Files = append(f.Files, tempFile)
		}
		return nil
	}

}

func (f FolderInfo) GetAbsolutePath() string {
	if f.Parent == nil {
		return f.FolderName
	} else {
		return f.Parent.FolderName + "\\" + f.FolderName
	}
}
