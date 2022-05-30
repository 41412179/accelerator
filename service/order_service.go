package service

import (
	"accelerator/conf"
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"

	// "github.com/go-pay/pkg/util"
	payutil "github.com/go-pay/gopay/pkg/util"
)

// OrderService 订单服务
type OrderService struct {
	GoodID    int64  `form:"good_id" json:"good_id" binding:"required"`
	PayType   string `form:"pay_type" json:"pay_type" binding:"required"`
	ChannelID int64  `form:"channel_id" json:"channel_id" binding:"required"`
	user      *table.User
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

// GetExpireTimeByUserId 获取过期时间
func (o *OrderService) GetExpireTimeByUserId(userID int64, remainingTime int64) int64 {
	return time.Now().Unix() + remainingTime
}

// GetRemainingTime 获取剩余时间
func (o *OrderService) GetRemainingTimeByUserId(userId int64) (int64, error) {
	orders, err := mysql.GetOrdersByUserID(userId)
	if err != nil {
		util.Log().Error("get orders by user id err: %+v", err)
		return 0, err
	}
	var remainingTime int64
	for _, order := range orders {
		if order.Status == mysql.OrderStatusPaid {
			if order.EndTime.Unix() > time.Now().Unix() {
				remainingTime += order.EndTime.Unix() - int64(time.Now().Unix())
			}
		}
	}

	return remainingTime, nil
}

// CreateOrder 创建订单
func (o *OrderService) CreateOrder(c *gin.Context) response.Response {
	// 获取用户
	user := util.GetUserByCtx(c)
	if user == nil {
		util.Log().Error("user not logined, user: %v", user)
		return errcode.NewErr(errcode.CodeCheckLogin, nil)
	}
	o.user = user
	// var order table.Order
	order := o.generateOrder()
	if order == nil {
		util.Log().Error("generate order err, user: %+v, orderInfo: %+v", user, o)
		return errcode.NewErr(errcode.CodeDBError, nil)
	}

	// 创建订单
	id, err := mysql.InsertOrder(order)
	if err != nil {
		util.Log().Error("insert order err: %+v", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}
	// 计算佣金
	go o.computerCommission(id, order)
	return o.setRsponse()

}

// computerCommission 计算佣金
func (o *OrderService) computerCommission(orderId int64, order *table.Order) {
	// 如果没有被邀请人，则不计算佣金
	if o.user.InviterId == 0 {
		return
	}

	// 计算佣金
	c := new(table.Commission)
	c.UserId = o.user.InviterId
	c.OrderId = orderId
	c.Type = table.AddCommissionType
	c.Change = float64(order.PayActualPrice) * float64(0.3) / 100

	err := mysql.InsertCommission(c)
	if err != nil {
		util.Log().Error("insert commission err: %+v", err)
		return
	}

}

func (o *OrderService) setRsponse() response.Response {
	str, err := o.createOrderStr()
	if err != nil {
		util.Log().Error("create order str err: %+v", err)
		return errcode.NewErr(errcode.CodeCreateOrderFailed, err)
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: str,
	}
}

// generateOrder 生成订单
func (o *OrderService) generateOrder() *table.Order {
	good, err := mysql.GetGoodByID(o.GoodID)
	if err != nil {
		util.Log().Error("get good by id err: %+v", err)
		return nil
	}
	order := new(table.Order)
	order.GoodId = int(o.GoodID)
	order.PayType = o.PayType
	order.UserId = o.user.ID
	order.StartTime = time.Now()
	order.EndTime = time.Now().Add(time.Minute * time.Duration(good.Duration))
	order.Status = mysql.OrderStatusPaid
	order.ChannelId = o.ChannelID
	order.PayActualPrice = float64(good.Price) / 100
	return order
}

func (o *OrderService) createOrderStr() (string, error) {

	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	// privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient(conf.PayConf.AppID, conf.PayConf.AppPrivateKey, conf.PayConf.Pro)
	if err != nil {
		util.Log().Error("初始化支付宝客户端失败", err)
		return "", err
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl(conf.PayConf.NotifyUrl)

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "加速器支付")
	bm.Set("out_trade_no", payutil.RandomString(32))
	bm.Set("total_amount", "1.00")
	//手机APP支付参数请求
	payParam, err := client.TradeAppPay(context.Background(), bm)
	if err != nil {
		util.Log().Error("err:", err)
		return "", err
	}
	return payParam, nil
}
