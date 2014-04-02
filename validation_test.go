package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/datadictionary"
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
	"time"
)

var _ = Suite(&DataDictionaryTests{})

type DataDictionaryTests struct{}

func (s *DataDictionaryTests) createFIX40NewOrderSingle() *MessageBuilder {
	msg := NewMessageBuilder()
	msg.Header.Set(NewStringField(tag.MsgType, "D"))
	msg.Header.Set(NewStringField(tag.BeginString, "FIX.4.0"))
	msg.Header.Set(NewStringField(tag.BodyLength, "0"))
	msg.Header.Set(NewStringField(tag.SenderCompID, "0"))
	msg.Header.Set(NewStringField(tag.TargetCompID, "0"))
	msg.Header.Set(NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.Set(NewUTCTimestampField(tag.SendingTime, time.Now()))

	msg.Body.Set(NewStringField(tag.ClOrdID, "A"))
	msg.Body.Set(NewStringField(tag.HandlInst, "1"))
	msg.Body.Set(NewStringField(tag.Symbol, "A"))
	msg.Body.Set(NewStringField(tag.Side, "1"))
	msg.Body.Set(NewStringField(tag.OrdType, "1"))
	msg.Body.Set(NewIntField(tag.OrderQty, 5))

	msg.Trailer.Set(NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *DataDictionaryTests) createFIX43NewOrderSingle() *MessageBuilder {
	msg := NewMessageBuilder()
	msg.Header.Set(NewStringField(tag.MsgType, "D"))
	msg.Header.Set(NewStringField(tag.BeginString, "FIX.4.3"))
	msg.Header.Set(NewStringField(tag.BodyLength, "0"))
	msg.Header.Set(NewStringField(tag.SenderCompID, "0"))
	msg.Header.Set(NewStringField(tag.TargetCompID, "0"))
	msg.Header.Set(NewStringField(tag.MsgSeqNum, "0"))
	msg.Header.Set(NewUTCTimestampField(tag.SendingTime, time.Now()))

	msg.Body.Set(NewStringField(tag.ClOrdID, "A"))
	msg.Body.Set(NewStringField(tag.HandlInst, "1"))
	msg.Body.Set(NewStringField(tag.Symbol, "A"))
	msg.Body.Set(NewStringField(tag.Side, "1"))
	msg.Body.Set(NewIntField(tag.OrderQty, 5))
	msg.Body.Set(NewStringField(tag.OrdType, "1"))
	msg.Body.Set(NewUTCTimestampField(tag.TransactTime, time.Now()))

	msg.Trailer.Set(NewStringField(tag.CheckSum, "000"))

	return msg
}

func (s *DataDictionaryTests) TestValidateInvalidTagNumber(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.Set(NewStringField(9999, "hello"))
	msg, _ := builder.Build()
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Trailer.Set(NewStringField(9999, "hello"))
	msg, _ = builder.Build()
	reject = Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))

	builder = s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(9999, "hello"))
	msg, _ = builder.Build()
	reject = Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidTagNumber)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(9999))
}

func (s *DataDictionaryTests) TestValidateTagNotDefinedForMessage(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(41, "hello"))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagNotDefinedForThisMessageType)
	c.Check(*reject.RefTagID(), Equals, tag.Tag(41))
}

func (s *DataDictionaryTests) TestValidateTagNotDefinedForMessageComponent(c *C) {
	dict, err := datadictionary.Parse("spec/FIX43.xml")
	c.Check(err, IsNil)
	builder := s.createFIX43NewOrderSingle()
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, IsNil)
}

func (s *DataDictionaryTests) TestValidateFieldNotFound(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Body.init(normalFieldOrder)
	builder.Body.Set(NewStringField(tag.ClOrdID, "A"))
	builder.Body.Set(NewStringField(tag.HandlInst, "A"))
	builder.Body.Set(NewStringField(tag.Symbol, "A"))
	builder.Body.Set(NewStringField(tag.Side, "A"))
	builder.Body.Set(NewStringField(tag.OrderQty, "A"))

	//ord type is required
	//builder.Body.Set(NewStringField(tag.OrdType, "A"))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.OrdType)

	builder = s.createFIX40NewOrderSingle()
	builder.Header.init()
	builder.Header.Set(NewStringField(tag.MsgType, "D"))
	builder.Header.Set(NewStringField(tag.BeginString, "FIX.4.0"))
	builder.Header.Set(NewStringField(tag.BodyLength, "0"))
	builder.Header.Set(NewStringField(tag.SenderCompID, "0"))
	builder.Header.Set(NewStringField(tag.TargetCompID, "0"))
	builder.Header.Set(NewStringField(tag.MsgSeqNum, "0"))
	//sending time is required
	//msg.Header.FieldMap.Set(NewStringField(tag.SendingTime, "0"))

	msg, _ = builder.Build()

	reject = Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, RequiredTagMissing)
	c.Check(*reject.RefTagID(), Equals, tag.SendingTime)
}

func (s *DataDictionaryTests) TestValidateTagSpecifiedWithoutAValue(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(tag.ClientID, ""))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagSpecifiedWithoutAValue)
	c.Check(*reject.RefTagID(), Equals, tag.ClientID)
}

func (s *DataDictionaryTests) TestValidateInvalidMsgType(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")

	builder := s.createFIX40NewOrderSingle()
	builder.Header.Set(NewStringField(tag.MsgType, "z"))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, InvalidMsgType)
}

func (s *DataDictionaryTests) TestValidateValueIsIncorrect(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(tag.HandlInst, "4"))
	msg, _ := builder.Build()

	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, ValueIsIncorrect)
	c.Check(*reject.RefTagID(), Equals, tag.HandlInst)
}

func (s *DataDictionaryTests) TestValidateIncorrectDataFormatForValue(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	builder.Body.Set(NewStringField(tag.OrderQty, "+200.00"))
	msg, _ := builder.Build()
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, IncorrectDataFormatForValue)
	c.Check(*reject.RefTagID(), Equals, tag.OrderQty)
}

func (s *DataDictionaryTests) TestValidateTagSpecifiedOutOfRequiredOrder(c *C) {
	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	builder := s.createFIX40NewOrderSingle()
	//should be in header
	builder.Body.Set(NewStringField(tag.OnBehalfOfCompID, "CWB"))
	msg, _ := builder.Build()
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagSpecifiedOutOfRequiredOrder)
	c.Check(*reject.RefTagID(), Equals, tag.OnBehalfOfCompID)
}

func (s *DataDictionaryTests) TestValidateTagAppearsMoreThanOnce(c *C) {

	msg, err := MessageFromParsedBytes([]byte("8=FIX.4.09=10735=D34=249=TW52=20060102-15:04:0556=ISLD11=ID21=140=140=254=138=20055=INTC60=20060102-15:04:0510=234"))
	c.Check(err, IsNil)

	dict, _ := datadictionary.Parse("spec/FIX40.xml")
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, TagAppearsMoreThanOnce)
	c.Check(*reject.RefTagID(), Equals, tag.OrdType)
}

func (s *DataDictionaryTests) TestFloatValidation(c *C) {
	msg, err := MessageFromParsedBytes([]byte("8=FIX.4.29=10635=D34=249=TW52=20140329-22:38:4556=ISLD11=ID21=140=154=138=+200.0055=INTC60=20140329-22:38:4510=178"))
	c.Check(err, IsNil)

	dict, _ := datadictionary.Parse("spec/FIX42.xml")
	reject := Validate(dict, *msg)
	c.Check(reject, NotNil)
	c.Check(reject.RejectReason(), Equals, IncorrectDataFormatForValue)
}