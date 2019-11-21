package code

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/dao"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/code_house"
	"go.mongodb.org/mongo-driver/mongo"
)
type CodeHouseDao struct {
	dao.Dao
	Results [] code.CodeHouse
	Model  code.CodeHouse
}

func NewCodeHouseDao() * CodeHouseDao {
	return &CodeHouseDao {
		dao.Dao{}, make([]code.CodeHouse,0,1),
		code.CodeHouse{},

	}
}

func (dao *CodeHouseDao)GetModeName() string{
	return "codeHouse"
}

func (dao *CodeHouseDao)PushData(cur *mongo.Cursor) {
	m :=code.CodeHouse{}
	cur.Decode(&m)
	dao.Results = append(dao.Results,m)
}

func (dao *CodeHouseDao)SetData(singleReult *mongo.SingleResult) {
	singleReult.Decode(&dao.Model)
}


