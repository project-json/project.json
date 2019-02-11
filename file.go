package projectcfg

type FileAction string

const (
	FileCopy    string = "copy"
	FileDelete  string = "delete"
	FileKeep    string = "keep"
	FileExclude string = "exclude"
)

type Filepath string
type Filepaths []string

type FileMap map[FileAction]Filepaths

type FilePathTemplate struct {
	Template string
	Params   []string
}
