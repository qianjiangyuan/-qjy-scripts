package user

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/group"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/workspace"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	// ref: https://godoc.org/labix.org/v2/mgo/bson#Marshal
	model.ModelBase `bson:"inline"`
	Name            string             `json:"name" bson:"name" binding:"required"`
	Nickname        string             `json:"nickname" bson:"nickname" `
	Passwd          string             `json:"passwd" bson:"passwd" `
	Phone           string             `json:"phone" bson:"phone" `
	Email           string             `json:"email" bson:"email" `
	WorkspaceID     primitive.ObjectID `json:"workspace_id" bson:"workspace_id" `
	GroupId			primitive.ObjectID  `json:"group_id" bson:"group_id" `
	Administrator   string 	`json:"administrator" bson:"administrator" `
	Oauth2            string             `json:"oauth2" bson:"oauth2"`
}
type UserVo struct {
	User          `bson:"inline"`
	reladateModel `bson:"inline"`
}

type reladateModel struct {
	Workspace workspace.Workspace `json:"workspace"`
	Group group.Group `json:"group"`
}

func NewUser() User {
	now := time.Now()
	model := model.ModelBase{
		ID:    primitive.NewObjectID(),
		Ctime: now,
		Mtime: now,
	}
	user := User{}
	user.ModelBase = model
	return user
}
