package dataset

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DataSet struct {
	// ref: https://godoc.org/labix.org/v2/mgo/bson#Marshal
	model.ModelBase `bson:"inline"`
	Name            string `json:"name" bson:"name" binding:"required"`
	Label           string `json:"label" bson:"label" `
	Url             string `json:"url" bson:"url" `
	Platform        string `json:"platform" bson:"platform" `
	Description     string `json:"description" bson:"description" `
	Visibility      string `json:"visibility" bson:"visibility"`
}
type DataSetVo struct {
	DataSet       `bson:"inline"`
	reladateModel `bson:"inline"`
	Files         []model.File `json:files`
	FileTotal     int          `json:file_total`
}

type reladateModel struct {
	Cuser user.User `json:"cuser"`
}

func NewDataSet() DataSet {
	now := time.Now()
	model := model.ModelBase{
		ID:    primitive.NewObjectID(),
		Ctime: now,
		Mtime: now,
	}
	dataset := DataSet{}
	dataset.ModelBase = model
	return dataset
}

