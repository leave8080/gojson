package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leave8080/gojson/message"
	"gofer/pkg/log"
)

func getDass(c *gin.Context) {

	req := new(message.MessInfo)
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error("bindJson", err)
		return
	}
	//log.Debug("ShouldBindJSON:", req.Table, req.Opt, req.BeforeValues, req.Values)
	// 校验数据是否为该类型的消息
	err = message.Validate.Struct(message.Mx[req.Table])
	if err != nil {
		// 不符合
		log.Error("Validate err:", err)
		return
	}

	message.Mx[req.Table].Execute(req)

}
