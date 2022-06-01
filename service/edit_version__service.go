package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"
)

type EditVersionService struct {
	Version string `json:"version"`
	URL     string `json:"url"`
	ID      int64  `json:"id"`
}

func (v *EditVersionService) EditVersion() response.Response {
	version, err := mysql.EditVersion(v.Version, v.URL, v.ID)
	if err != nil {
		util.Log().Error("edit version err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: version,
	}
}
