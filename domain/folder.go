package folder

type FolderInfo struct {
	Parent     *FolderInfo
	Subfolders []FolderInfo
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
		Subfolders: make([]FolderInfo, 0),
		Files:      make([]FileInfo, 0),
		IsScanned:  false,
		FolderName: absolutePath,
		CantAccess: false,
	}
}

func new(foldername string, parent *FolderInfo) *FolderInfo {
	return &FolderInfo{
		Parent:     parent,
		Subfolders: make([]FolderInfo, 0),
		Files:      make([]FileInfo, 0),
		IsScanned:  false,
		CantAccess: false,
		FolderName: foldername,
	}
}
