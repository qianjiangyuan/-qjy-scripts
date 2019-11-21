package group

import (
  "github.com/xxmyjk/xintong/backend/pkg/app/dao"
  "github.com/xxmyjk/xintong/backend/pkg/app/model/group"
  "go.mongodb.org/mongo-driver/mongo"
)
type GroupDao struct {
  dao.Dao
  Results [] group.Group
  Model  group.Group
}

func NewGroupDao() * GroupDao {
  return &GroupDao {
    dao.Dao{}, make([]group.Group,0,1),
    group.Group{},

  }
}

func (dao *GroupDao)GetModeName() string{
  return "group"
}

func (dao *GroupDao)PushData(cur *mongo.Cursor) {
  m :=group.Group{}
  cur.Decode(&m);
  dao.Results = append(dao.Results,m)
}
func (dao *GroupDao)SetData(singleReult *mongo.SingleResult) {
  singleReult.Decode(&dao.Model)
}

