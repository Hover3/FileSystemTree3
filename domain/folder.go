package folder

type FolderInfo struct {
	Parent     *FolderInfo
	SubFolders []FolderInfo
	Files      []FileInfo
	IsScanned  bool
	CantAccess bool
	FolderName string
}

type FileInfo struct {
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
