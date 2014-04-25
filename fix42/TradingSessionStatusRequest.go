package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradingSessionStatusRequest msg type = g.
type TradingSessionStatusRequest struct {
	message.Message
}

//TradingSessionStatusRequestBuilder builds TradingSessionStatusRequest messages.
type TradingSessionStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateTradingSessionStatusRequestBuilder returns an initialized TradingSessionStatusRequestBuilder with specified required fields.
func CreateTradingSessionStatusRequestBuilder(
	tradsesreqid field.TradSesReqID,
	subscriptionrequesttype field.SubscriptionRequestType) TradingSessionStatusRequestBuilder {
	var builder TradingSessionStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("g"))
	builder.Body.Set(tradsesreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//TradSesReqID is a required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesReqID() (*field.TradSesReqID, errors.MessageRejectError) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradSesReqID(f *field.TradSesReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMethod is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesMethod() (*field.TradSesMethod, errors.MessageRejectError) {
	f := new(field.TradSesMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMethod reads a TradSesMethod from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradSesMethod(f *field.TradSesMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMode is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesMode() (*field.TradSesMode, errors.MessageRejectError) {
	f := new(field.TradSesMode)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMode reads a TradSesMode from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradSesMode(f *field.TradSesMode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}
