package workspace

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Workspace struct {
	// ref: https://godoc.org/labix.org/v2/mgo/bson#Marshal
	model.ModelBase `bson:"inline"`
	Name            string `json:"name" bson:"name" binding:"required"`
	Description     string `json:"description" bson:"description" `
	GpuNum          int64  `json:"gpu_num" bson:"gpu_num" `
	CpuNum          int64  `json:"cpu_num" bson:"cpu_num" `
	MemoryNum       int64  `json:"memory_num" bson:"memory_num" `
	Rdma            int64  `json:"rdma" bson:"rdma" `
}

func NewWorkspace() Workspace {
	now := time.Now()
	model := model.ModelBase{
		ID:      primitive.NewObjectID(),
		Ctime:   now,
		Mtime:   now,
		Deleted: false,
	}
	workspace := Workspace{}
	workspace.ModelBase = model
	return workspace
}

type reladateModel struct {
}

type WorkSpaceVo struct {
	Workspace     `bson:"inline"`
	reladateModel `bson:"inline"`
}
