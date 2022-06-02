package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"
)

type EditVersionService struct {
	Version string `json:"version" form:"version" binding:"required"`
	URL     string `json:"url" form:"url" binding:"required"`
	Desc    string `json:"desc" form:"desc" binding:"required"`
	ID      int64  `json:"id" form:"id" binding:"required"`
	CodeID  string `json:"code_id" form:"code_id" binding:"required"`
}

func (v *EditVersionService) EditVersion() response.Response {
	version, err := mysql.EditVersion(v.Version, v.URL, v.Desc, v.CodeID, v.ID)
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
