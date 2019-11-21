package code
import (
"github.com/xxmyjk/xintong/backend/pkg/app/model"
"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
"go.mongodb.org/mongo-driver/bson/primitive"
"time"
)
type CodeHouse struct {
	model.ModelBase `bson:"inline"`
	Name string  `json:"name" bson:"name"`
	Git string  `json:"git" bson:"git"`
	Platform string  `json:"platform" bson:"platform"`
	Label string  `json:"label" bson:"label"`
	Describe string  `json:"describe" bson:"describe"`
	Cmd string  `json:"cmd" bson:"cmd"`
	Visibility      string `json:"visibility" bson:"visibility"`

}

type CodeHouseVo struct {
	CodeHouse `bson:"inline"`
	reladateModel  `bson:"inline"`
}
type reladateModel struct {
	Cuser      user.User `json:"cuser"`
}
func NewCodeHouse() CodeHouse {
	now:= time.Now()
	model := model.ModelBase{
		ID: primitive.NewObjectID(),
		Ctime: now,
		Mtime: now,
	}
	codeHouse :=CodeHouse{}
	codeHouse.ModelBase = model;
	return codeHouse
}

