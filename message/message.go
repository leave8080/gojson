package message

import (
	"github.com/go-playground/validator/v10"
	"github.com/leave8080/gojson/conf"
	"github.com/leave8080/gojson/dao"
)

var (
	srvDao   *dao.Dao
	Validate *validator.Validate
)

type MessgeHandler interface {
	Handle([]byte)
}

// 定义消息处理 handle interface, 使用消息结构体tag校验消息类型
type HttpHandler interface {
	TableNames() string
	Execute()
}

type Handler struct {
}

func InitHandle(c *conf.Config) *Handler {
	//srvDao = dao.New(c)

	Validate = validator.New()
	return &Handler{}
}

//func (* Handler)Handle(data []byte)  {
//	messageHandlers := []HttpHandler{
//		&MxFeedback{},
//	}
//
//	for _, mHandle := range messageHandlers {
//		handle := mHandle.(HttpHandler)
//		// 将json反序列化到结构体
//		if err := json.Unmarshal(data, handle); err != nil {
//			log.Errorf("json.Unmarshal(%s), err:%#v", string(data), err)
//			continue
//		}
//		// 校验数据是否为该类型的消息
//		err := Validate.Struct(handle)
//		if err != nil {
//			// 不符合
//			continue
//		}
//		handle.Execute()
//		//if res := ; res != amqp.Accept {
//		//	return res
//		//}
//	}
//
//
//
//}
