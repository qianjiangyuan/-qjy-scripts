package upload

import "github.com/xxmyjk/xintong/backend/pkg/app/model"

type Files struct {
	Files         []model.File `json:files`
	Total     int          `json:total`
}
