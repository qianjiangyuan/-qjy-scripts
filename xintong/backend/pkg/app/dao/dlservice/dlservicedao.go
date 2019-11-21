package dlservice

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/dao"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dlservice"
	"go.mongodb.org/mongo-driver/mongo"
)

type DlserviceDao struct {
	dao.Dao
	Results []dlservice.Dlservice
	Model   dlservice.Dlservice
}

func NewDlserviceDao() *DlserviceDao {
	return &DlserviceDao{
		dao.Dao{}, make([]dlservice.Dlservice, 0, 1),
		dlservice.Dlservice{},
	}
}

func (dao *DlserviceDao) GetModeName() string {
	return "dlservice"
}

func (dao *DlserviceDao) PushData(cur *mongo.Cursor) {
	m := dlservice.Dlservice{}
	cur.Decode(&m)
	dao.Results = append(dao.Results, m)
}

func (dao *DlserviceDao) SetData(singleReult *mongo.SingleResult) {
	singleReult.Decode(&dao.Model)
}
