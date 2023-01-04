package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-files-go/version1"
)

type FilesMockClientV1 struct {
	client  *version1.FilesMockClientV1
	fixture *FilesClientFixtureV1
}

func newEmailTemplatesMockClientV1() *FilesMockClientV1 {
	return &FilesMockClientV1{}
}

func (c *FilesMockClientV1) setup(t *testing.T) *FilesClientFixtureV1 {
	c.client = version1.NewFilesMockClientV1()
	c.fixture = NewFilesClientFixtureV1(c.client)
	return c.fixture
}

func (c *FilesMockClientV1) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newEmailTemplatesMockClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
