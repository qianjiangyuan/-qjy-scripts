package group
import (
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
type Group struct {
    model.ModelBase `bson:"inline"`
    Name string  `json:"name" bson:"name"`
}

type GroupVo struct {
	Group `bson:"inline"`
	reladateModel  `bson:"inline"`
}
type reladateModel struct {
}
func NewGroup() Group {
	now:= time.Now()
	model := model.ModelBase{
		ID: primitive.NewObjectID(),
		Ctime: now,
		Mtime: now,
	}
	group :=Group{}
	group.ModelBase = model;
	return group
}
