package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-files-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type FilesCommandableHttpClientV1 struct {
	client  *version1.FilesCommandableHttpClientV1
	fixture *FilesClientFixtureV1
}

func newFilesCommandableHttpClientV1() *FilesCommandableHttpClientV1 {
	return &FilesCommandableHttpClientV1{}
}

func (c *FilesCommandableHttpClientV1) setup(t *testing.T) *FilesClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewFilesCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewFilesClientFixtureV1(c.client)

	return c.fixture
}

func (c *FilesCommandableHttpClientV1) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newFilesCommandableHttpClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
