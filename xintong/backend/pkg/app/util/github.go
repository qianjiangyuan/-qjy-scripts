package util

import (
	"fmt"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
)

func Clone(url string,path string) ( error) {
	fmt.Println(url)
	fmt.Println(path)
	path=connect.Conf.GetString("filepath.path")+"/"+path
	conf := connect.Conf
	username := conf.GetString("github.username")
	token := conf.GetString("github.token")
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth: &http.BasicAuth{
			Username: username, // yes, this can be anything except an empty string
			Password: token,
		},
		Progress: os.Stdout,
	})
	if( err != nil ){
		fmt.Println(err)
		return err
	}
	return nil

}