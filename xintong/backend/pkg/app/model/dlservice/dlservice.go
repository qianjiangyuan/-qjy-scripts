package dlservice

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dlmodel"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/image"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Dlservice struct {
	model.ModelBase `bson:"inline"`
	Name            string             `json:"name" bson:"name"`
	Version         string             `json:"version" bson:"version"`
	Label           string             `json:"label" bson:"label"`
	Platform        string             `json:"platform" bson:"platform"`
	Path            string             `json:"path" bson:"path"`
	Description     string             `json:"description" bson:"description"`
	ImageID         primitive.ObjectID `json:"image_id" bson:"image_id"`
	DlmodelID       primitive.ObjectID `json:"dlmodel_id" bson:"dlmodel_id"`
}

type DlserviceVo struct {
	Dlservice    `bson:"inline"`
	relatedModel `bson:"inline"`
}
type relatedModel struct {
	Cuser   user.User       `json:"cuser"`
	Image   image.Image     `json:"image"`
	Dlmodel dlmodel.Dlmodel `json:"dlmodel"`
}

func NewDlservice() Dlservice {
	now := time.Now()
	model := model.ModelBase{
		ID:    primitive.NewObjectID(),
		Ctime: now,
		Mtime: now,
	}
	dlservice := Dlservice{}
	dlservice.ModelBase = model
	return dlservice
}
