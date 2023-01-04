package version1

import (
	"context"
	"strings"
	"time"

	blobClients "github.com/pip-services-infrastructure2/client-blobs-go/version1"
	facetClients "github.com/pip-services-infrastructure2/client-facets-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type FilesMockClientV1 struct {
	files        []*FileV1
	facetsClient facetClients.IFacetsClientV1
	blobsClient  blobClients.IBlobsClientV1
	facetsGroup  string
}

func NewFilesMockClientV1() *FilesMockClientV1 {
	return &FilesMockClientV1{
		files:        make([]*FileV1, 0),
		facetsClient: facetClients.NewFacetsMockClientV1(),
		blobsClient:  blobClients.NewBlobsMockClientV1(),
		facetsGroup:  "files",
	}
}

func (c *FilesMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *FilesMockClientV1) matchSearch(item *FileV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Name, search) {
		return true
	}
	if c.matchString(item.Description, search) {
		return true
	}
	return false
}

func (c *FilesMockClientV1) composeFilter(filter *data.FilterParams) func(item *FileV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	name, nameOk := filter.GetAsNullableString("name")
	group, groupOk := filter.GetAsNullableString("group")
	expired, expiredOk := filter.GetAsNullableBoolean("expired")
	fromCreateTime, fromCreateTimeOk := filter.GetAsNullableDateTime("from_create_time")
	toCreateTime, toCreateTimeOk := filter.GetAsNullableDateTime("to_create_time")

	now := time.Now()

	return func(item *FileV1) bool {
		if searchOk && !c.matchSearch(item, search) {
			return false
		}
		if idOk && id != item.Id {
			return false
		}
		if nameOk && name != item.Name {
			return false
		}
		if groupOk && group != item.Group {
			return false
		}
		if expiredOk && expired && item.ExpireTime.Unix() > now.Unix() {
			return false
		}
		if expiredOk && !expired && item.ExpireTime.Unix() <= now.Unix() {
			return false
		}
		if fromCreateTimeOk && item.CreateTime.Unix() >= fromCreateTime.Unix() {
			return false
		}
		if toCreateTimeOk && item.CreateTime.Unix() < toCreateTime.Unix() {
			return false
		}
		return true
	}
}

func (c *FilesMockClientV1) GetGroups(ctx context.Context, correlationId string, paging *data.PagingParams) (data.DataPage[string], error) {
	page, err := c.facetsClient.GetFacetsByGroup(context.Background(), correlationId, c.facetsGroup, paging)
	if err != nil {
		return *data.NewEmptyDataPage[string](), err
	}

	groups := make([]string, 0)
	for _, item := range page.Data {
		groups = append(groups, item.Group)
	}

	return *data.NewDataPage(groups, page.Total), nil
}

func (c *FilesMockClientV1) GetFilesByFilter(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*FileV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*FileV1, 0)
	for _, v := range c.files {
		item := *v
		if filterFunc(&item) {
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.files)), nil
}

func (c *FilesMockClientV1) GetFilesByIds(ctx context.Context, correlationId string, fileIds []string) ([]*FileV1, error) {
	files := make([]*FileV1, 0)

	for _, file := range c.files {
		for _, id := range fileIds {
			if file.Id == id {
				files = append(files, file)
				break
			}
		}
	}

	return files, nil
}

func (c *FilesMockClientV1) GetFileById(ctx context.Context, correlationId string, fileId string) (*FileV1, error) {
	var file *FileV1

	for _, el := range c.files {
		if file.Id == fileId {
			buf := *el
			file = &buf
			break
		}
	}

	return file, nil
}
func (c *FilesMockClientV1) normalizeName(name string) string {
	if name == "" {
		return ""
	}

	name = strings.ReplaceAll(name, "\\", "/")
	pos := strings.LastIndex(name, "/")
	if pos >= 0 {
		name = name[pos+1:]
	}

	return name
}

func (c *FilesMockClientV1) CreateFile(ctx context.Context, correlationId string, file *FileV1) (*FileV1, error) {
	if file.Id != "" {
		file.Id = data.IdGenerator.NextLong()
	}

	file.Name = c.normalizeName(file.Name)
	file.CreateTime = time.Now()

	// Create file
	buf := *file
	c.files = append(c.files, &buf)

	// Add group to facet search
	if file.Group != "" {
		_, err := c.facetsClient.AddFacet(context.Background(), correlationId, c.facetsGroup, file.Group)
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

func (c *FilesMockClientV1) UpdateFile(ctx context.Context, correlationId string, file *FileV1) (*FileV1, error) {
	var newFile *FileV1

	file.Name = c.normalizeName(file.Name)

	// Update file
	for i, f := range c.files {
		if file.Id == f.Id {
			buf := *file
			c.files[i] = &buf
			newFile = file
			break
		}
	}

	// Remove old group from facet search
	if file.Group != "" && file.Group != newFile.Group {
		_, err := c.facetsClient.RemoveFacet(ctx, correlationId, c.facetsGroup, file.Group)
		if err != nil {
			return nil, err
		}
	}

	// Add new group from facet search
	if newFile.Group != "" && file.Group != newFile.Group {
		_, err := c.facetsClient.AddFacet(ctx, correlationId, c.facetsGroup, newFile.Group)
		if err != nil {
			return nil, err
		}
	}

	return newFile, nil
}

func (c *FilesMockClientV1) DeleteFileById(ctx context.Context, correlationId string, fileId string) (*FileV1, error) {
	var file *FileV1

	// Delete file
	for i, f := range c.files {
		if fileId == f.Id {
			file = f
			if i < len(c.files) {
				c.files = append(c.files[:i], c.files[i+1:]...)
			} else {
				c.files = c.files[:i]
			}
			break
		}
	}

	// Delete content blob
	if file.ContentId != "" {
		err := c.blobsClient.DeleteBlobById(ctx, correlationId, file.ContentId)
		if err != nil {
			return nil, err
		}
	}

	// Delete thumbnail blob
	if file.ThumbnailId != "" {
		err := c.blobsClient.DeleteBlobById(ctx, correlationId, file.ThumbnailId)
		if err != nil {
			return nil, err
		}
	}

	// Remove group from facet search
	if file.Group != "" {
		_, err := c.facetsClient.RemoveFacet(ctx, correlationId, c.facetsGroup, file.Group)
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}
