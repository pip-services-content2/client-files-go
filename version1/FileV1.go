package version1

import (
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type FileV1 struct {
	// Identification
	Id    string `json:"id"`
	Group string `json:"group"`
	Name  string `json:"name"`

	// Content
	Description  string               `json:"description"`
	ContentId    string               `json:"content_id"`
	ContentUri   string               `json:"content_uri"`
	ThumbnailId  string               `json:"thumbnail_id"`
	ThumbnailUri string               `json:"thumbnail_uri"`
	CreateTime   time.Time            `json:"create_time"`
	ExpireTime   time.Time            `json:"expire_time"`
	Attributes   *data.StringValueMap `json:"attributes"`

	// Custom fields
	CustomHdr any `json:"custom_hdr"`
	CustomDat any `json:"custom_dat"`
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
