package response

import "github.com/awahids/belajar-go/common"

type Response struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data,omitempty"`
	Meta    *common.Meta `json:"meta,omitempty"`
}
