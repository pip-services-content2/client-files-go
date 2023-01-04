package test_version1

import (
	"context"
	"testing"
	"time"

	"github.com/pip-services-content2/client-files-go/version1"
	"github.com/stretchr/testify/assert"
)

type FilesClientFixtureV1 struct {
	Client version1.IFilesClientV1
}

func NewFilesClientFixtureV1(client version1.IFilesClientV1) *FilesClientFixtureV1 {
	return &FilesClientFixtureV1{
		Client: client,
	}
}

func (c *FilesClientFixtureV1) clear() {
	c.Client = nil
}

func (c *FilesClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create file
	file := version1.NewFileV1("", "test", "file-1.dat", "Test file", "111", "", time.Time{}, nil)

	file1, err := c.Client.CreateFile(context.Background(), "123", file)
	assert.Nil(t, err)
	assert.NotNil(t, file1)

	// Update file
	file.Name = "new_file.dat"

	file, err = c.Client.UpdateFile(context.Background(), "123", file)
	assert.Nil(t, err)

	assert.NotNil(t, file)
	assert.Equal(t, file.Name, "new_file.dat")

	// Get files
	page, err := c.Client.GetFilesByFilter(context.Background(), "123", nil, nil)
	assert.Nil(t, err)

	assert.Len(t, page.Data, 1)

	// Get groups
	groupsPage, err := c.Client.GetGroups(context.Background(), "123", nil)
	assert.Nil(t, err)

	assert.Len(t, groupsPage.Data, 1)

	// Delete file
	_, err = c.Client.DeleteFileById(context.Background(), "123", file1.Id)
	assert.Nil(t, err)

	// Try to get deleted file
	file, err = c.Client.GetFileById(context.Background(), "123", file1.Id)
	assert.Nil(t, err)
	assert.Nil(t, file)
}
