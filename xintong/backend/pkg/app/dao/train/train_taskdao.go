package train_task

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/dao"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/train"
	"go.mongodb.org/mongo-driver/mongo"
)

type Train_taskDao struct {
	dao.Dao
	Results []train_task.Train_task
	Model   train_task.Train_task
}

func NewTrain_taskDao() *Train_taskDao {
	return &Train_taskDao{
		dao.Dao{}, make([]train_task.Train_task, 0, 1),
		train_task.Train_task{},
	}
}

func (dao *Train_taskDao) GetModeName() string {
	return "train_task"
}

func (dao *Train_taskDao) PushData(cur *mongo.Cursor) {
	m := train_task.Train_task{}
	cur.Decode(&m)
	dao.Results = append(dao.Results, m)
}

func (dao *Train_taskDao) SetData(singleReult *mongo.SingleResult) {
	singleReult.Decode(&dao.Model)
}
