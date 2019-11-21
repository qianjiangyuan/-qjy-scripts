package workspace

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/dao"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/workspace"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkSpaceDao struct {
	dao.Dao
	Results []workspace.WorkSpaceVo
	Model   workspace.Workspace
}

func NewWorkSpaceDao() *WorkSpaceDao {
	return &WorkSpaceDao{
		dao.Dao{}, make([]workspace.WorkSpaceVo, 0, 1),
		workspace.Workspace{},
	}
}

func (dao *WorkSpaceDao) GetModeName() string {
	return "workspace"
}

func (dao *WorkSpaceDao) PushData(cur *mongo.Cursor) {
	m := workspace.WorkSpaceVo{}
	cur.Decode(&m)
	dao.Results = append(dao.Results, m)
}

func (dao *WorkSpaceDao) SetData(singleReult *mongo.SingleResult) {
	singleReult.Decode(&dao.Model)
}
