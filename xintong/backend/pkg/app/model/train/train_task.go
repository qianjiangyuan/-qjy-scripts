package train_task

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/model"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/code_house"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dataset"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/image"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Train_task struct {
	model.ModelBase `bson:"inline"`
	Name            string             `json:"name" bson:"name"`
	Phone           string             `json:"phone" bson:"phone"`
	Contacts        string             `json:"contacts" bson:"contacts"`
	Email           string             `json:"email" bson:"email"`
	ModelUrl        string             `json:"model_url" bson:"model_url"`
	DatasetId       primitive.ObjectID `json:"dataset_id" bson:"dataset_id"`
	ImageUrl        string			   `json:"image_url" bson:"image_url"`
	PodNum          int32              `json:"pod_num" bson:"pod_num"`
	CpuNum          int64              `json:"cpu_num" bson:"cpu_num"`
	GpuNum          int64              `json:"gpu_num" bson:"gpu_num"`
	MemoryNum       int64              `json:"memory_num" bson:"memory_num"`
	Rdma            int64              `json:"rdma" bson:"rdma"`
	Cmd             string             `json:"cmd" bson:"cmd"`
	CodeHouseId     primitive.ObjectID `json:"codehouse_id" bson:"codehouse_id"`
	Type            string              `json:"type" bson:"type"`

}


type Train_taskVo struct {
	ModelUrl   string  `json:"model_url"`
	CodePath   string
	Status     string  `json:"status"`
	Train_task   `bson:"inline"`
	RelatedModel `bson:"inline"`
}
type RelatedModel struct {
	Cuser   user.User       `json:"cuser"`
	DataSet dataset.DataSet `json:"dataset"`
	CodeHouse code.CodeHouse `json:"codehouse"`
	Image   image.Image		`json:"image"`
}

func NewTrain_task() Train_task {
	now := time.Now()
	model := model.ModelBase{
		ID:    primitive.NewObjectID(),
		Ctime: now,
		Mtime: now,
	}
	train_task := Train_task{}
	train_task.ModelBase = model
	return train_task
}
