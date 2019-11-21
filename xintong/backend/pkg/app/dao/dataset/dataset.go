package dataset

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/dao"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dataset"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSetDao struct {
	dao.Dao
	Results []dataset.DataSetVo
	Model   dataset.DataSet
}

func NewDataSetDao() *DataSetDao {
	return &DataSetDao{
		dao.Dao{}, make([]dataset.DataSetVo, 0, 1),
		dataset.DataSet{},
	}
}

func (dsDao *DataSetDao) GetModeName() string {
	return "dataset"
}

func (dsDao *DataSetDao) PushData(cur *mongo.Cursor) {
	m := dataset.DataSetVo{}
	cur.Decode(&m)
	m.Cuser.Name = "管理员"
	dsDao.Results = append(dsDao.Results, m)
}

func (dsDao *DataSetDao) SetData(singleReult *mongo.SingleResult) {
	singleReult.Decode(&dsDao.Model)
}
