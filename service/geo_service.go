package service

import (
	"accelerator/entity/db"
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"

	"github.com/gin-gonic/gin"
)

type GeoService struct {
}

func (m *GeoService) GetGeos(c *gin.Context) (res *response.Response) {
	var geos []*table.Geo
	err := db.DB.Find(&geos).Error
	if err != nil {
		return &response.Response{
			Code: errcode.CodeDBError,
			Msg:  errcode.Text(errcode.CodeDBError),
			Data: nil,
		}
	} else {
		return &response.Response{
			Code: errcode.CodeSuccess,
			Msg:  errcode.Text(errcode.CodeSuccess),
			Data: geos,
		}
	}
}
