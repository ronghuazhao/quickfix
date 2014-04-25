package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//IOI msg type = 6.
type IOI struct {
	message.Message
}

//IOIBuilder builds IOI messages.
type IOIBuilder struct {
	message.MessageBuilder
}

//CreateIOIBuilder returns an initialized IOIBuilder with specified required fields.
func CreateIOIBuilder(
	ioiid field.IOIID,
	ioitranstype field.IOITransType,
	side field.Side,
	ioiqty field.IOIQty) IOIBuilder {
	var builder IOIBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("6"))
	builder.Body.Set(ioiid)
	builder.Body.Set(ioitranstype)
	builder.Body.Set(side)
	builder.Body.Set(ioiqty)
	return builder
}

//IOIID is a required field for IOI.
func (m IOI) IOIID() (*field.IOIID, errors.MessageRejectError) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from IOI.
func (m IOI) GetIOIID(f *field.IOIID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOITransType is a required field for IOI.
func (m IOI) IOITransType() (*field.IOITransType, errors.MessageRejectError) {
	f := new(field.IOITransType)
	err := m.Body.Get(f)
	return f, err
}

//GetIOITransType reads a IOITransType from IOI.
func (m IOI) GetIOITransType(f *field.IOITransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIRefID is a non-required field for IOI.
func (m IOI) IOIRefID() (*field.IOIRefID, errors.MessageRejectError) {
	f := new(field.IOIRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIRefID reads a IOIRefID from IOI.
func (m IOI) GetIOIRefID(f *field.IOIRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for IOI.
func (m IOI) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from IOI.
func (m IOI) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for IOI.
func (m IOI) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from IOI.
func (m IOI) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for IOI.
func (m IOI) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from IOI.
func (m IOI) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for IOI.
func (m IOI) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from IOI.
func (m IOI) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for IOI.
func (m IOI) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from IOI.
func (m IOI) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for IOI.
func (m IOI) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from IOI.
func (m IOI) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for IOI.
func (m IOI) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from IOI.
func (m IOI) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for IOI.
func (m IOI) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from IOI.
func (m IOI) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for IOI.
func (m IOI) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from IOI.
func (m IOI) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for IOI.
func (m IOI) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from IOI.
func (m IOI) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for IOI.
func (m IOI) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from IOI.
func (m IOI) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for IOI.
func (m IOI) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from IOI.
func (m IOI) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for IOI.
func (m IOI) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from IOI.
func (m IOI) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for IOI.
func (m IOI) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from IOI.
func (m IOI) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for IOI.
func (m IOI) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from IOI.
func (m IOI) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for IOI.
func (m IOI) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from IOI.
func (m IOI) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for IOI.
func (m IOI) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from IOI.
func (m IOI) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for IOI.
func (m IOI) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from IOI.
func (m IOI) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for IOI.
func (m IOI) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from IOI.
func (m IOI) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for IOI.
func (m IOI) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from IOI.
func (m IOI) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for IOI.
func (m IOI) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from IOI.
func (m IOI) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for IOI.
func (m IOI) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from IOI.
func (m IOI) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for IOI.
func (m IOI) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from IOI.
func (m IOI) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for IOI.
func (m IOI) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from IOI.
func (m IOI) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for IOI.
func (m IOI) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from IOI.
func (m IOI) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for IOI.
func (m IOI) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from IOI.
func (m IOI) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for IOI.
func (m IOI) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from IOI.
func (m IOI) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for IOI.
func (m IOI) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from IOI.
func (m IOI) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for IOI.
func (m IOI) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from IOI.
func (m IOI) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for IOI.
func (m IOI) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from IOI.
func (m IOI) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for IOI.
func (m IOI) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from IOI.
func (m IOI) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for IOI.
func (m IOI) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from IOI.
func (m IOI) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for IOI.
func (m IOI) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from IOI.
func (m IOI) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for IOI.
func (m IOI) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from IOI.
func (m IOI) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for IOI.
func (m IOI) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from IOI.
func (m IOI) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for IOI.
func (m IOI) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from IOI.
func (m IOI) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for IOI.
func (m IOI) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from IOI.
func (m IOI) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for IOI.
func (m IOI) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from IOI.
func (m IOI) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for IOI.
func (m IOI) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from IOI.
func (m IOI) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for IOI.
func (m IOI) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from IOI.
func (m IOI) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for IOI.
func (m IOI) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from IOI.
func (m IOI) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for IOI.
func (m IOI) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from IOI.
func (m IOI) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for IOI.
func (m IOI) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from IOI.
func (m IOI) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for IOI.
func (m IOI) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from IOI.
func (m IOI) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for IOI.
func (m IOI) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from IOI.
func (m IOI) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for IOI.
func (m IOI) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from IOI.
func (m IOI) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for IOI.
func (m IOI) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from IOI.
func (m IOI) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for IOI.
func (m IOI) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from IOI.
func (m IOI) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for IOI.
func (m IOI) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from IOI.
func (m IOI) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for IOI.
func (m IOI) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from IOI.
func (m IOI) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for IOI.
func (m IOI) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from IOI.
func (m IOI) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for IOI.
func (m IOI) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from IOI.
func (m IOI) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for IOI.
func (m IOI) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from IOI.
func (m IOI) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for IOI.
func (m IOI) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from IOI.
func (m IOI) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for IOI.
func (m IOI) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from IOI.
func (m IOI) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for IOI.
func (m IOI) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from IOI.
func (m IOI) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for IOI.
func (m IOI) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from IOI.
func (m IOI) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for IOI.
func (m IOI) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from IOI.
func (m IOI) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for IOI.
func (m IOI) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from IOI.
func (m IOI) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQty is a required field for IOI.
func (m IOI) IOIQty() (*field.IOIQty, errors.MessageRejectError) {
	f := new(field.IOIQty)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQty reads a IOIQty from IOI.
func (m IOI) GetIOIQty(f *field.IOIQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for IOI.
func (m IOI) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from IOI.
func (m IOI) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for IOI.
func (m IOI) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from IOI.
func (m IOI) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for IOI.
func (m IOI) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from IOI.
func (m IOI) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for IOI.
func (m IOI) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from IOI.
func (m IOI) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for IOI.
func (m IOI) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from IOI.
func (m IOI) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for IOI.
func (m IOI) ValidUntilTime() (*field.ValidUntilTime, errors.MessageRejectError) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from IOI.
func (m IOI) GetValidUntilTime(f *field.ValidUntilTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQltyInd is a non-required field for IOI.
func (m IOI) IOIQltyInd() (*field.IOIQltyInd, errors.MessageRejectError) {
	f := new(field.IOIQltyInd)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQltyInd reads a IOIQltyInd from IOI.
func (m IOI) GetIOIQltyInd(f *field.IOIQltyInd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOINaturalFlag is a non-required field for IOI.
func (m IOI) IOINaturalFlag() (*field.IOINaturalFlag, errors.MessageRejectError) {
	f := new(field.IOINaturalFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetIOINaturalFlag reads a IOINaturalFlag from IOI.
func (m IOI) GetIOINaturalFlag(f *field.IOINaturalFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoIOIQualifiers is a non-required field for IOI.
func (m IOI) NoIOIQualifiers() (*field.NoIOIQualifiers, errors.MessageRejectError) {
	f := new(field.NoIOIQualifiers)
	err := m.Body.Get(f)
	return f, err
}

//GetNoIOIQualifiers reads a NoIOIQualifiers from IOI.
func (m IOI) GetNoIOIQualifiers(f *field.NoIOIQualifiers) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for IOI.
func (m IOI) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from IOI.
func (m IOI) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for IOI.
func (m IOI) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from IOI.
func (m IOI) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for IOI.
func (m IOI) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from IOI.
func (m IOI) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for IOI.
func (m IOI) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from IOI.
func (m IOI) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for IOI.
func (m IOI) URLLink() (*field.URLLink, errors.MessageRejectError) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from IOI.
func (m IOI) GetURLLink(f *field.URLLink) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for IOI.
func (m IOI) NoRoutingIDs() (*field.NoRoutingIDs, errors.MessageRejectError) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from IOI.
func (m IOI) GetNoRoutingIDs(f *field.NoRoutingIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for IOI.
func (m IOI) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from IOI.
func (m IOI) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for IOI.
func (m IOI) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from IOI.
func (m IOI) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for IOI.
func (m IOI) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from IOI.
func (m IOI) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for IOI.
func (m IOI) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from IOI.
func (m IOI) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for IOI.
func (m IOI) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from IOI.
func (m IOI) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for IOI.
func (m IOI) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from IOI.
func (m IOI) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for IOI.
func (m IOI) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from IOI.
func (m IOI) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for IOI.
func (m IOI) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from IOI.
func (m IOI) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for IOI.
func (m IOI) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from IOI.
func (m IOI) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for IOI.
func (m IOI) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from IOI.
func (m IOI) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for IOI.
func (m IOI) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from IOI.
func (m IOI) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for IOI.
func (m IOI) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from IOI.
func (m IOI) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for IOI.
func (m IOI) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from IOI.
func (m IOI) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for IOI.
func (m IOI) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from IOI.
func (m IOI) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}
