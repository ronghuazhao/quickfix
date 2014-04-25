package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MassQuote msg type = i.
type MassQuote struct {
	message.Message
}

//MassQuoteBuilder builds MassQuote messages.
type MassQuoteBuilder struct {
	message.MessageBuilder
}

//CreateMassQuoteBuilder returns an initialized MassQuoteBuilder with specified required fields.
func CreateMassQuoteBuilder(
	quoteid field.QuoteID,
	noquotesets field.NoQuoteSets) MassQuoteBuilder {
	var builder MassQuoteBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("i"))
	builder.Body.Set(quoteid)
	builder.Body.Set(noquotesets)
	return builder
}

//QuoteReqID is a non-required field for MassQuote.
func (m MassQuote) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from MassQuote.
func (m MassQuote) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for MassQuote.
func (m MassQuote) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MassQuote.
func (m MassQuote) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for MassQuote.
func (m MassQuote) QuoteResponseLevel() (*field.QuoteResponseLevel, errors.MessageRejectError) {
	f := new(field.QuoteResponseLevel)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from MassQuote.
func (m MassQuote) GetQuoteResponseLevel(f *field.QuoteResponseLevel) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DefBidSize is a non-required field for MassQuote.
func (m MassQuote) DefBidSize() (*field.DefBidSize, errors.MessageRejectError) {
	f := new(field.DefBidSize)
	err := m.Body.Get(f)
	return f, err
}

//GetDefBidSize reads a DefBidSize from MassQuote.
func (m MassQuote) GetDefBidSize(f *field.DefBidSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DefOfferSize is a non-required field for MassQuote.
func (m MassQuote) DefOfferSize() (*field.DefOfferSize, errors.MessageRejectError) {
	f := new(field.DefOfferSize)
	err := m.Body.Get(f)
	return f, err
}

//GetDefOfferSize reads a DefOfferSize from MassQuote.
func (m MassQuote) GetDefOfferSize(f *field.DefOfferSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a required field for MassQuote.
func (m MassQuote) NoQuoteSets() (*field.NoQuoteSets, errors.MessageRejectError) {
	f := new(field.NoQuoteSets)
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from MassQuote.
func (m MassQuote) GetNoQuoteSets(f *field.NoQuoteSets) errors.MessageRejectError {
	return m.Body.Get(f)
}
