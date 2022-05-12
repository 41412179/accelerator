package errcode

import (
	"accelerator/entity/response"

	"github.com/gin-gonic/gin"
)

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	// CodeSuccess 成功
	CodeSuccess = 0

	// CodeNodeNotExist 节点不存在
	CodeNodeNotExist = 20001

	// CodeGoodNotExist 商品不存在
	CodeGoodNotExist = 20002
)

// errMap 错误码文案映射
var errMap = map[int]string{
	CodeDBError:      "数据库操作失败",
	CodeEncryptError: "加密失败",
	CodeParamErr:     "参数错误",
	CodeSuccess:      "成功",
	CodeCheckLogin:   "未登录",
	CodeNoRightErr:   "未授权访问",
	CodeNodeNotExist: "节点不存在",
	CodeGoodNotExist: "商品不存在",
}

//Text 错误码转换文案
func Text(code int) string {
	errMsg, ok := errMap[code]
	if !ok {
		return "未定义错误"
	}
	return errMsg
}

// NewErrorByCode 根据错误码生成error
// func NewErrorByCode(code int) error {
// 	return errs.New(code, Text(code))
// }

// Err 通用错误处理
// func Err(errCode int, msg string, err error) response.Response {
// 	res := response.Response{
// 		Code: errCode,
// 		Msg:  msg,
// 	}
// 	// 生产环境隐藏底层报错
// 	if err != nil && gin.Mode() != gin.ReleaseMode {
// 		res.Error = err.Error()
// 	}
// 	return res
// }

// NewErr 通用错误处理
func NewErr(errCode int, err error) response.Response {
	res := response.Response{
		Code: errCode,
		Msg:  errMap[errCode],
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr 数据库操作失败
// func DBErr(msg string, err error) response.Response {
// 	if msg == "" {
// 		msg = "数据库操作失败"
// 	}
// 	return Err(CodeDBError, msg, err)
// }

// ParamErr 各种参数错误
// func ParamErr(msg string, err error) response.Response {
// 	if msg == "" {
// 		msg = "参数错误"
// 	}
// 	return Err(CodeParamErr, msg, err)
// }
