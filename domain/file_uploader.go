package domain

import "bytes"

type FileModel struct {
	data     *bytes.Buffer
	fileName string
}

func NewFileModel(data *bytes.Buffer, fileName string) *FileModel {
	return &FileModel{data: data, fileName: fileName}
}

func (f FileModel) Data() *bytes.Buffer {
	return f.data
}

func (f FileModel) FileName() string {
	return f.fileName
}

type FileUploader interface {
	Upload(fileModel *FileModel) error
}
