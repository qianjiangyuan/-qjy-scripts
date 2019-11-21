package image

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Image struct {
	// ref: https://godoc.org/labix.org/v2/mgo/bson#Marshal
	model.ModelBase `bson:"inline"`
	Name            string `json:"name" bson:"name" binding:"required"`
	Label           string `json:"label" bson:"label" `
	Visibility      string `json:"visibility" bson:"visibility" `
	Url             string `json:"url" bson:"url" `
	Platform        string `json:"platform" bson:"platform" `
	Description     string `json:"description" bson:"description"`
	Version         string `json:"version" bson:"version"`
	FileName        string `json:"filename"" bson:"filename"`
	Command         string `json:"command"" bson:"command"`
}
type reladateModel struct {
	Cuser user.User `json:"cuser"`
}
type ImageVo struct {
	Image         `bson:"inline"`
	reladateModel `bson:"inline"`
	Files         []model.File `json:files`
	FileTotal     int          `json:file_total`
}

func NewImage() Image {
	now := time.Now()
	model := model.ModelBase{
		ID:      primitive.NewObjectID(),
		Ctime:   now,
		Mtime:   now,
		Deleted: false,
	}
	image := Image{}
	image.ModelBase = model
	return image
}


type HanborImage struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	ProjectID    int           `json:"project_id"`
	Description  string        `json:"description"`
	PullCount    int           `json:"pull_count"`
	StarCount    int           `json:"star_count"`
	TagsCount    int           `json:"tags_count"`
	Labels       []interface{} `json:"labels"`
	CreationTime time.Time     `json:"creation_time"`
	UpdateTime   time.Time     `json:"update_time"`
	Version   	 string 		`json:"version"`
	ImageUrl	string			`json:"image_url"`
}

type HarborImageVersion struct {
	Name          string        `json:"name"`
}


