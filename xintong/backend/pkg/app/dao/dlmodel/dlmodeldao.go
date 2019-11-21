package dlmodel

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/dao"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dlmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

type DlmodelDao struct {
	dao.Dao
	Results []dlmodel.Dlmodel
	Model   dlmodel.Dlmodel
}

func NewDlmodelDao() *DlmodelDao {
	return &DlmodelDao{
		dao.Dao{}, make([]dlmodel.Dlmodel, 0, 1),
		dlmodel.Dlmodel{},
	}
}

func (dao *DlmodelDao) GetModeName() string {
	return "dlmodel"
}

func (dao *DlmodelDao) PushData(cur *mongo.Cursor) {
	m := dlmodel.Dlmodel{}
	cur.Decode(&m)
	dao.Results = append(dao.Results, m)
}

func (dao *DlmodelDao) SetData(singleReult *mongo.SingleResult) {
	singleReult.Decode(&dao.Model)
}
