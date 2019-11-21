package dlmodel

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Dlmodel struct {
	model.ModelBase `bson:"inline"`
	Name            string             `json:"name" bson:"name"`
	Version         string             `json:"version" bson:"version"`
	Url             string             `json:"url" bson:"url"`
	TrainTaskId     primitive.ObjectID `json:"train_task_id" bson:"train_task_id"`
	Label           string             `json:"label" bson:"label"`
	Platform        string             `json:"platform" bson:"platform"`
	ModelName       string              `json:"model_name" bson:"model_name"`
}

type DlmodelVo struct {
	Dlmodel      `bson:"inline"`
	relatedModel `bson:"inline"`
	Files        []model.File `json:files`
	FileTotal    int          `json:file_total`
}
type relatedModel struct {
	Cuser user.User `json:"cuser"`
}

func NewDlmodel() Dlmodel {
	now := time.Now()
	model := model.ModelBase{
		ID:    primitive.NewObjectID(),
		Ctime: now,
		Mtime: now,
	}
	dlmodel := Dlmodel{}
	dlmodel.ModelBase = model
	return dlmodel
}
