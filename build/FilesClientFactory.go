package build

import (
	clients1 "github.com/pip-services-content2/client-files-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type FilesClientFactory struct {
	*cbuild.Factory
}

func NewFilesClientFactory() *FilesClientFactory {
	c := &FilesClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-files", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-files", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-files", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewFilesNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewFilesMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewFilesCommandableHttpClientV1)

	return c
}
