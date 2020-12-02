package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leave8080/gojson/message"
	"gofer/pkg/log"
)

func getDass(c *gin.Context) {
	messageHandlers := []message.HttpHandler{

		&message.MxProject{},

		&message.MxFeedback{},
		&message.MxTest{},
	}

	for i, mHandle := range messageHandlers {
		handle := mHandle.(message.HttpHandler)
		//err := c.ShouldBindJSON(handle)
		//if err != nil {
		//	log.Error("bindJson", err)
		//	continue
		//}
		name2 := handle.TableNames()
		// 校验数据是否为该类型的消息
		err := message.Validate.Struct(handle)
		if err != nil {
			// 不符合
			continue
		}
		//names := mHandle.TableNames()
		log.Debug("@@@", i, name2)
		//enable := true
		//for _, key := range names {
		//	_, ok := c.Items[key]
		//	if !ok {
		//		enable = false
		//		break
		//	}
		//}
		handle.Execute()
	}

}
