package user

import (
	"context"
	"fmt"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/dao"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
	"github.com/xxmyjk/xintong/backend/pkg/app/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDao struct {
	dao.Dao
	Results []user.User
	Model   user.User
}

func NewUserDao() *UserDao {
	return &UserDao{
		dao.Dao{}, make([]user.User, 0, 1),
		user.User{},
	}
}

func (dao *UserDao) GetModeName() string {
	return "user"
}

func (dao *UserDao) PushData(cur *mongo.Cursor) {
	m := user.User{}
	cur.Decode(&m)
	dao.Results = append(dao.Results, m)
}

func (dao *UserDao) SetData(singleReult *mongo.SingleResult) {
	singleReult.Decode(&dao.Model)
}

func (dao *UserDao) Login(userName string, passWord string) (user.User, error) {

	filter := bson.D{
		{"name", userName},
	}
	projection := bson.D{
		// useless stub for now
		{"runner", 0},
	}
	user := user.User{}
	collection := connect.Conn.Mongo.Collection("user")
	rs := collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection))

	err := rs.Decode(&user)

	if err != nil {
		return user, err
	}

	passWordMd5 := util.MD5(passWord)

	fmt.Print("data="+passWordMd5+","+user.Passwd )
	if user.Passwd != passWordMd5 {
		//return user, errors.New("username or password  is incorrect")
	}
	return user, nil
}
func (dao *UserDao) FindOne(userName string) (user.User, error) {
	filter := bson.D{
		{"name", userName},
	}
	projection := bson.D{
		// useless stub for now
		{"runner", 0},
	}
	user := user.User{}
	collection := connect.Conn.Mongo.Collection("user")
	rs := collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection))

	err := rs.Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (dao *UserDao) UpdatePassWd( m user.UserVo ) error {
	//collection := connect.Conn.Mongo.Collection(baseResult.GetModeName())
	//filter := bson.D{{"_id", bson.D{{"$in", ids}}}}
	//_, err := collection.DeleteMany(context.TODO(), filter) {
	collection := connect.Conn.Mongo.Collection(dao.GetModeName())
	m.Passwd = util.MD5(m.Passwd)
	updateStr := bson.M{"$set":bson.M{"passwd":m.Passwd} }
	filter := bson.D{
		{"_id", m.ID},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, updateStr)
	if err != nil {
		return  err
	}
	return nil

}

func (dao *UserDao) GithubLogin(email string, name string) ( user.User, error) {

	filter := bson.D{
		{"name", email},
		{"oauth2", "github"},
	}
	projection := bson.D{
		// useless stub for now
		{"runner", 0},
	}
	user := user.User{}

	collection := connect.Conn.Mongo.Collection("user")
	rs := collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection))
	err := rs.Decode(&user)
	if err != nil {
		user.Name=email
		user.Oauth2="github"
		user.Nickname=name;
		_,err :=dao.Insert(user,dao)
		if(err != nil){
			return user,err;
		}
	}
	return user, nil
}