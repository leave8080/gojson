package message

import (
	"errors"
	"gofer/pkg/log"
)

type MessInfo struct {
	Opt          string      `json:"opt"  binding:"required"`
	Table        string      `json:"table"  binding:"required"`
	BeforeValues interface{} `json:"before_values"`
	Values       interface{} `json:"values"  binding:"required"`
}

func Task(req *MessInfo) error {
	var err error
	if req.Opt == "update" {

	} else {
		log.Debug(req.Values)
	}
	switch req.Opt {
	case "update":
		//todo
		if req.BeforeValues == nil {
			err = errors.New("BeforeValues = nil")
			return err
		}

	case "add":
	//todo
	case "delete":
		//todo
	default:
		err = errors.New("Opt err")
		return err
	}
	return nil
}
