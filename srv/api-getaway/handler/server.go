package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	__ "github.com/yuhang-jieke/orderai/srv/proto"

	"github.com/yuhang-jieke/orderai/srv/api-getaway/basic/config"
)

func OrderAdd(c *gin.Context) {
	var form __.AddOrdersReq
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	_, err := config.OrderClient.AddOrders(c, &__.AddOrdersReq{
		Name:  form.Name,
		Num:   form.Num,
		Price: form.Price,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
	return
}
func DelOrder(c *gin.Context) {
	var form __.DelOrdersReq
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	_, err := config.OrderClient.DelOrders(c, &__.DelOrdersReq{
		Id: form.Id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
	return
}
func GetId(c *gin.Context) {
	var form __.GetOrdersByIdReq
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	r, err := config.OrderClient.GetOrdersById(c, &__.GetOrdersByIdReq{
		Id: form.Id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": r,
	})
	return
}
func UpdateId(c *gin.Context) {
	var form __.UpdateOrdersReq
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	r, err := config.OrderClient.UpdateOrders(c, &__.UpdateOrdersReq{
		Price: form.Price,
		Id:    form.Id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": r,
	})
	return
}
