package version1

import (
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type FileV1 struct {
	// Identification
	Id    string
	Group string
	Name  string

	// Content
	Description  string
	ContentId    string
	ContentUri   string
	ThumbnailId  string
	ThumbnailUri string
	CreateTime   time.Time
	ExpireTime   time.Time
	Attributes   *data.StringValueMap

	// Custom fields
	CustomHdr any
	CustomDat any
}

func NewFileV1(id, group, name, description, contentId, contentUri string, expireTime time.Time, attributes *data.StringValueMap) *FileV1 {
	c := &FileV1{
		Id:          id,
		Group:       group,
		Name:        name,
		Description: description,
		ContentId:   contentId,
		ContentUri:  contentUri,
		CreateTime:  time.Now(),
		ExpireTime:  expireTime,
	}

	if attributes != nil {
		c.Attributes = attributes
	} else {
		c.Attributes = data.NewEmptyStringValueMap()
	}

	return c
}
