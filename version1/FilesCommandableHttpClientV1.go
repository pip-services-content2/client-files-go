package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type FilesCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewFilesCommandableHttpClientV1() *FilesCommandableHttpClientV1 {
	return &FilesCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/files"),
	}
}

func (c *FilesCommandableHttpClientV1) GetGroups(ctx context.Context, correlationId string, paging *data.PagingParams) (data.DataPage[string], error) {
	params := data.NewAnyValueMapFromTuples(
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_groups", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[string](), err
	}

	return clients.HandleHttpResponse[data.DataPage[string]](res, correlationId)
}

func (c *FilesCommandableHttpClientV1) GetFilesByFilter(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*FileV1], error) {
	params := data.NewAnyValueMapFromTuples(
		"paging", paging,
		"filter", filter,
	)

	res, err := c.CallCommand(ctx, "get_files_by_filter", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[*FileV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*FileV1]](res, correlationId)
}

func (c *FilesCommandableHttpClientV1) GetFilesByIds(ctx context.Context, correlationId string, fileIds []string) ([]*FileV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"file_ids", fileIds,
	)

	res, err := c.CallCommand(ctx, "get_files_by_ids", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]*FileV1](res, correlationId)
}

func (c *FilesCommandableHttpClientV1) GetFileById(ctx context.Context, correlationId string, fileId string) (*FileV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"file_id", fileId,
	)

	res, err := c.CallCommand(ctx, "get_file_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*FileV1](res, correlationId)
}

func (c *FilesCommandableHttpClientV1) CreateFile(ctx context.Context, correlationId string, file *FileV1) (*FileV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"file", file,
	)

	res, err := c.CallCommand(ctx, "create_file", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*FileV1](res, correlationId)
}

func (c *FilesCommandableHttpClientV1) UpdateFile(ctx context.Context, correlationId string, file *FileV1) (*FileV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"file", file,
	)

	res, err := c.CallCommand(ctx, "update_file", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*FileV1](res, correlationId)
}

func (c *FilesCommandableHttpClientV1) DeleteFileById(ctx context.Context, correlationId string, fileId string) (*FileV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"file_id", fileId,
	)

	res, err := c.CallCommand(ctx, "delete_file_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*FileV1](res, correlationId)
}
