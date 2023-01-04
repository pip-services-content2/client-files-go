package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type FilesNullClientV1 struct {
}

func NewFilesNullClientV1() *FilesNullClientV1 {
	return &FilesNullClientV1{}
}

func (c *FilesNullClientV1) GetGroups(ctx context.Context, correlationId string, paging *data.PagingParams) (data.DataPage[string], error) {
	return *data.NewEmptyDataPage[string](), nil
}

func (c *FilesNullClientV1) GetFilesByFilter(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*FileV1], error) {
	return *data.NewEmptyDataPage[*FileV1](), nil
}

func (c *FilesNullClientV1) GetFilesByIds(ctx context.Context, correlationId string, fileIds []string) ([]*FileV1, error) {
	return make([]*FileV1, 0), nil
}

func (c *FilesNullClientV1) GetFileById(ctx context.Context, correlationId string, fileId string) (*FileV1, error) {
	return nil, nil
}

func (c *FilesNullClientV1) CreateFile(ctx context.Context, correlationId string, file *FileV1) (*FileV1, error) {
	return file, nil
}

func (c *FilesNullClientV1) UpdateFile(ctx context.Context, correlationId string, file *FileV1) (*FileV1, error) {
	return file, nil
}

func (c *FilesNullClientV1) DeleteFileById(ctx context.Context, correlationId string, fileId string) (*FileV1, error) {
	return nil, nil
}
