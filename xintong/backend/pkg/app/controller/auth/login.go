package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller"
	userDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/admin/user"
	"github.com/xxmyjk/xintong/backend/pkg/app/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"time"
)

type Auth struct {
	Username string             `json:"username"`
	Password string             `json:"password"`
	ID       primitive.ObjectID `json:"id"`
	Token    string             `json:"token"`
	Administrator string         `json:"administrator"`
	Code     string               `json:"code"`
}

func Login(ctx *gin.Context) {

	auth := Auth{}
	err := ctx.Bind(&auth)
	if controller.ErrCatch(err, ctx) {
		return
	}
	uDao := userDao.NewUserDao()
	user, err := uDao.Login(auth.Username, auth.Password)
	if controller.ErrCatch(err, ctx) {
		return
	}
	reAuth := Auth{}
	reAuth.ID = user.ID
	reAuth.Token, err = util.GenerateToken(user)
	reAuth.Username = user.Nickname
	reAuth.Administrator=user.Administrator
	if controller.ErrCatch(err, ctx) {
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": reAuth,
	})
	return
}


//LoginGithub github OAuth 登陆
func GitHubLogin(c *gin.Context) {


	auth := Auth{}
	err := c.Bind(&auth)
	if controller.ErrCatch(err, c) {
		return
	}
	//code 通过 github.com OAuth API 换取 token
	// token 根据GitHub 开发API接口获取用户信息 githubUser
	gu, err := fetchGithubUser(auth.Code)
	if controller.ErrCatch(err, c) {
		return
	}

	//比对或者插入GitHub User 到数据库
	//同时参数自己的jwt token
	uDao := userDao.NewUserDao()
	user, err := uDao.GithubLogin(*gu.Email, gu.Name)
	if controller.ErrCatch(err, c) {
		return
	}
	reAuth := Auth{}
	reAuth.ID = user.ID
	reAuth.Token, err = util.GenerateToken(user)
	reAuth.Username = user.Nickname
	reAuth.Administrator=user.Administrator
	if controller.ErrCatch(err, c) {
		return
	}
	c.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": reAuth,
	})
}

//fetchGithubUser 获取github 用户信息
func fetchGithubUser(code string) (*githubUser, error) {
	client := http.Client{}
	githubClinetID := connect.Conf.GetString("github.client_id")
	githubClientSecret := connect.Conf.GetString("github.client_secret")
	params := fmt.Sprintf(`{"client_id":"%s","client_secret":"%s","code":"%s"}`, githubClinetID, githubClientSecret, code)
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBufferString(params))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	gt := githubToken{}
	err = json.Unmarshal(bs, &gt)
	if err != nil {
		return nil, err
	}

	//开始获取用户信息
	fmt.Println("begin get user token"+gt.AccessToken)
	req, err = http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Add("Authorization", "Bearer "+gt.AccessToken)

	res, err = client.Do(req)
	fmt.Println("end get user token"+gt.AccessToken)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("using github token to fetch User Info failed with not 200 error")
	}
	defer res.Body.Close()
	bs, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	gu := &githubUser{}
	err = json.Unmarshal(bs, gu)
	if err != nil {
		return nil, err
	}
	if gu.Email == nil {
		tEmail := fmt.Sprintf("%d@github.com", gu.ID)
		gu.Email = &tEmail
	}

	gu.Token = gt.AccessToken
	return gu, nil
}

type githubToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
type githubUser struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             *string   `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Token             string    `json:"-"`
}

