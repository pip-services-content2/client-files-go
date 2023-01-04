package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IFilesClientV1 interface {
	GetGroups(ctx context.Context, correlationId string, paging *data.PagingParams) (data.DataPage[string], error)
	GetFilesByFilter(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*FileV1], error)
	GetFilesByIds(ctx context.Context, correlationId string, fileIds []string) ([]*FileV1, error)
	GetFileById(ctx context.Context, correlationId string, fileId string) (*FileV1, error)
	CreateFile(ctx context.Context, correlationId string, file *FileV1) (*FileV1, error)
	UpdateFile(ctx context.Context, correlationId string, file *FileV1) (*FileV1, error)
	DeleteFileById(ctx context.Context, correlationId string, fileId string) (*FileV1, error)
}
