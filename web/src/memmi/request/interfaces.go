package request

import (
	"github.com/golang/protobuf/proto"
	"memmi/pbuf"
	"net/http"
)

var BODY_READ_ERROR = &pbuf.RequestError{Reason: "Could not read request body."}
var USER_HISTORY_UPDATE_ERROR = &pbuf.RequestError{Reason: "Could not update user history."}

type HttpAuthenticator interface {
	AuthenticateUser(r *http.Request) pbuf.User
}

type HandleResult struct {
	Continue        bool
	ResponseWritten bool
}

type RequestHandler interface {
	ShouldHandle(r *http.Request, u pbuf.User, responseWritten bool) bool
	Handle(w http.ResponseWriter, r *http.Request, u pbuf.User) HandleResult
}

type HttpRouter interface {
	AddHandler(RequestHandler)
	GetHandleFunc() func(http.ResponseWriter, *http.Request)
}

type ProtoIO interface {
	WriteProtoResponse(w http.ResponseWriter, message proto.Message) error
	WriteCodedProtoResponse(w http.ResponseWriter, message proto.Message, statusCode int) error
	ReadCardScoreReport(r *http.Request) (pbuf.CardScoreReport, error)
	ReadCardSetRequest(r *http.Request) (pbuf.CardSetRequest, error)
	ReadCardRequest(r *http.Request) (pbuf.CardRequest, error)
}
