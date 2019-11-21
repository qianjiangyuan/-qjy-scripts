package image

import (
	"github.com/xxmyjk/xintong/backend/pkg/app/dao"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/image"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImageDao struct {
	dao.Dao
	Results []image.ImageVo
	Model   image.Image
}

func NewImageDao() *ImageDao {
	return &ImageDao{
		dao.Dao{}, make([]image.ImageVo, 0, 1),
		image.Image{},
	}
}

func (dao *ImageDao) GetModeName() string {
	return "image"
}

func (dao *ImageDao) PushData(cur *mongo.Cursor) {
	m := image.ImageVo{}
	cur.Decode(&m)
	m.Cuser.Name = "管理员"
	dao.Results = append(dao.Results, m)
}

func (dao *ImageDao) SetData(singleReult *mongo.SingleResult) {
	singleReult.Decode(&dao.Model)
}
