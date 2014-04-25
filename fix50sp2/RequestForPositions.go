package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//RequestForPositions msg type = AN.
type RequestForPositions struct {
	message.Message
}

//RequestForPositionsBuilder builds RequestForPositions messages.
type RequestForPositionsBuilder struct {
	message.MessageBuilder
}

//CreateRequestForPositionsBuilder returns an initialized RequestForPositionsBuilder with specified required fields.
func CreateRequestForPositionsBuilder(
	posreqid field.PosReqID,
	posreqtype field.PosReqType,
	clearingbusinessdate field.ClearingBusinessDate,
	transacttime field.TransactTime) RequestForPositionsBuilder {
	var builder RequestForPositionsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AN"))
	builder.Body.Set(posreqid)
	builder.Body.Set(posreqtype)
	builder.Body.Set(clearingbusinessdate)
	builder.Body.Set(transacttime)
	return builder
}

//PosReqID is a required field for RequestForPositions.
func (m RequestForPositions) PosReqID() (*field.PosReqID, errors.MessageRejectError) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from RequestForPositions.
func (m RequestForPositions) GetPosReqID(f *field.PosReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqType is a required field for RequestForPositions.
func (m RequestForPositions) PosReqType() (*field.PosReqType, errors.MessageRejectError) {
	f := new(field.PosReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqType reads a PosReqType from RequestForPositions.
func (m RequestForPositions) GetPosReqType(f *field.PosReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for RequestForPositions.
func (m RequestForPositions) MatchStatus() (*field.MatchStatus, errors.MessageRejectError) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from RequestForPositions.
func (m RequestForPositions) GetMatchStatus(f *field.MatchStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for RequestForPositions.
func (m RequestForPositions) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from RequestForPositions.
func (m RequestForPositions) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RequestForPositions.
func (m RequestForPositions) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RequestForPositions.
func (m RequestForPositions) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RequestForPositions.
func (m RequestForPositions) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RequestForPositions.
func (m RequestForPositions) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RequestForPositions.
func (m RequestForPositions) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RequestForPositions.
func (m RequestForPositions) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for RequestForPositions.
func (m RequestForPositions) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from RequestForPositions.
func (m RequestForPositions) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for RequestForPositions.
func (m RequestForPositions) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from RequestForPositions.
func (m RequestForPositions) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for RequestForPositions.
func (m RequestForPositions) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from RequestForPositions.
func (m RequestForPositions) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from RequestForPositions.
func (m RequestForPositions) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from RequestForPositions.
func (m RequestForPositions) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for RequestForPositions.
func (m RequestForPositions) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from RequestForPositions.
func (m RequestForPositions) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for RequestForPositions.
func (m RequestForPositions) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from RequestForPositions.
func (m RequestForPositions) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for RequestForPositions.
func (m RequestForPositions) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from RequestForPositions.
func (m RequestForPositions) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from RequestForPositions.
func (m RequestForPositions) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for RequestForPositions.
func (m RequestForPositions) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from RequestForPositions.
func (m RequestForPositions) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for RequestForPositions.
func (m RequestForPositions) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from RequestForPositions.
func (m RequestForPositions) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for RequestForPositions.
func (m RequestForPositions) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from RequestForPositions.
func (m RequestForPositions) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for RequestForPositions.
func (m RequestForPositions) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from RequestForPositions.
func (m RequestForPositions) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for RequestForPositions.
func (m RequestForPositions) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from RequestForPositions.
func (m RequestForPositions) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for RequestForPositions.
func (m RequestForPositions) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from RequestForPositions.
func (m RequestForPositions) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for RequestForPositions.
func (m RequestForPositions) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from RequestForPositions.
func (m RequestForPositions) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for RequestForPositions.
func (m RequestForPositions) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from RequestForPositions.
func (m RequestForPositions) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for RequestForPositions.
func (m RequestForPositions) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from RequestForPositions.
func (m RequestForPositions) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for RequestForPositions.
func (m RequestForPositions) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from RequestForPositions.
func (m RequestForPositions) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for RequestForPositions.
func (m RequestForPositions) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from RequestForPositions.
func (m RequestForPositions) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for RequestForPositions.
func (m RequestForPositions) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from RequestForPositions.
func (m RequestForPositions) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for RequestForPositions.
func (m RequestForPositions) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from RequestForPositions.
func (m RequestForPositions) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for RequestForPositions.
func (m RequestForPositions) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from RequestForPositions.
func (m RequestForPositions) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for RequestForPositions.
func (m RequestForPositions) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from RequestForPositions.
func (m RequestForPositions) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from RequestForPositions.
func (m RequestForPositions) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from RequestForPositions.
func (m RequestForPositions) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for RequestForPositions.
func (m RequestForPositions) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from RequestForPositions.
func (m RequestForPositions) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for RequestForPositions.
func (m RequestForPositions) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from RequestForPositions.
func (m RequestForPositions) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for RequestForPositions.
func (m RequestForPositions) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from RequestForPositions.
func (m RequestForPositions) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from RequestForPositions.
func (m RequestForPositions) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for RequestForPositions.
func (m RequestForPositions) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from RequestForPositions.
func (m RequestForPositions) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from RequestForPositions.
func (m RequestForPositions) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from RequestForPositions.
func (m RequestForPositions) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from RequestForPositions.
func (m RequestForPositions) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from RequestForPositions.
func (m RequestForPositions) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from RequestForPositions.
func (m RequestForPositions) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for RequestForPositions.
func (m RequestForPositions) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from RequestForPositions.
func (m RequestForPositions) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for RequestForPositions.
func (m RequestForPositions) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from RequestForPositions.
func (m RequestForPositions) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for RequestForPositions.
func (m RequestForPositions) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from RequestForPositions.
func (m RequestForPositions) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for RequestForPositions.
func (m RequestForPositions) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from RequestForPositions.
func (m RequestForPositions) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for RequestForPositions.
func (m RequestForPositions) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from RequestForPositions.
func (m RequestForPositions) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for RequestForPositions.
func (m RequestForPositions) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from RequestForPositions.
func (m RequestForPositions) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for RequestForPositions.
func (m RequestForPositions) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from RequestForPositions.
func (m RequestForPositions) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from RequestForPositions.
func (m RequestForPositions) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for RequestForPositions.
func (m RequestForPositions) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from RequestForPositions.
func (m RequestForPositions) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from RequestForPositions.
func (m RequestForPositions) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from RequestForPositions.
func (m RequestForPositions) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from RequestForPositions.
func (m RequestForPositions) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for RequestForPositions.
func (m RequestForPositions) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from RequestForPositions.
func (m RequestForPositions) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for RequestForPositions.
func (m RequestForPositions) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from RequestForPositions.
func (m RequestForPositions) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for RequestForPositions.
func (m RequestForPositions) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from RequestForPositions.
func (m RequestForPositions) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for RequestForPositions.
func (m RequestForPositions) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from RequestForPositions.
func (m RequestForPositions) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for RequestForPositions.
func (m RequestForPositions) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from RequestForPositions.
func (m RequestForPositions) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for RequestForPositions.
func (m RequestForPositions) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from RequestForPositions.
func (m RequestForPositions) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for RequestForPositions.
func (m RequestForPositions) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from RequestForPositions.
func (m RequestForPositions) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from RequestForPositions.
func (m RequestForPositions) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for RequestForPositions.
func (m RequestForPositions) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from RequestForPositions.
func (m RequestForPositions) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for RequestForPositions.
func (m RequestForPositions) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from RequestForPositions.
func (m RequestForPositions) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from RequestForPositions.
func (m RequestForPositions) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from RequestForPositions.
func (m RequestForPositions) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from RequestForPositions.
func (m RequestForPositions) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for RequestForPositions.
func (m RequestForPositions) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from RequestForPositions.
func (m RequestForPositions) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for RequestForPositions.
func (m RequestForPositions) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from RequestForPositions.
func (m RequestForPositions) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for RequestForPositions.
func (m RequestForPositions) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from RequestForPositions.
func (m RequestForPositions) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from RequestForPositions.
func (m RequestForPositions) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for RequestForPositions.
func (m RequestForPositions) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from RequestForPositions.
func (m RequestForPositions) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for RequestForPositions.
func (m RequestForPositions) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from RequestForPositions.
func (m RequestForPositions) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from RequestForPositions.
func (m RequestForPositions) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from RequestForPositions.
func (m RequestForPositions) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for RequestForPositions.
func (m RequestForPositions) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from RequestForPositions.
func (m RequestForPositions) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for RequestForPositions.
func (m RequestForPositions) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from RequestForPositions.
func (m RequestForPositions) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for RequestForPositions.
func (m RequestForPositions) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from RequestForPositions.
func (m RequestForPositions) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for RequestForPositions.
func (m RequestForPositions) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from RequestForPositions.
func (m RequestForPositions) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for RequestForPositions.
func (m RequestForPositions) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from RequestForPositions.
func (m RequestForPositions) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from RequestForPositions.
func (m RequestForPositions) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for RequestForPositions.
func (m RequestForPositions) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from RequestForPositions.
func (m RequestForPositions) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for RequestForPositions.
func (m RequestForPositions) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from RequestForPositions.
func (m RequestForPositions) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for RequestForPositions.
func (m RequestForPositions) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from RequestForPositions.
func (m RequestForPositions) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for RequestForPositions.
func (m RequestForPositions) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from RequestForPositions.
func (m RequestForPositions) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for RequestForPositions.
func (m RequestForPositions) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from RequestForPositions.
func (m RequestForPositions) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for RequestForPositions.
func (m RequestForPositions) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from RequestForPositions.
func (m RequestForPositions) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for RequestForPositions.
func (m RequestForPositions) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from RequestForPositions.
func (m RequestForPositions) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for RequestForPositions.
func (m RequestForPositions) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from RequestForPositions.
func (m RequestForPositions) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from RequestForPositions.
func (m RequestForPositions) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from RequestForPositions.
func (m RequestForPositions) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from RequestForPositions.
func (m RequestForPositions) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from RequestForPositions.
func (m RequestForPositions) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for RequestForPositions.
func (m RequestForPositions) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from RequestForPositions.
func (m RequestForPositions) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for RequestForPositions.
func (m RequestForPositions) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from RequestForPositions.
func (m RequestForPositions) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for RequestForPositions.
func (m RequestForPositions) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from RequestForPositions.
func (m RequestForPositions) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for RequestForPositions.
func (m RequestForPositions) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from RequestForPositions.
func (m RequestForPositions) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for RequestForPositions.
func (m RequestForPositions) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from RequestForPositions.
func (m RequestForPositions) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for RequestForPositions.
func (m RequestForPositions) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from RequestForPositions.
func (m RequestForPositions) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for RequestForPositions.
func (m RequestForPositions) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from RequestForPositions.
func (m RequestForPositions) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for RequestForPositions.
func (m RequestForPositions) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from RequestForPositions.
func (m RequestForPositions) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for RequestForPositions.
func (m RequestForPositions) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from RequestForPositions.
func (m RequestForPositions) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for RequestForPositions.
func (m RequestForPositions) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from RequestForPositions.
func (m RequestForPositions) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for RequestForPositions.
func (m RequestForPositions) ResponseTransportType() (*field.ResponseTransportType, errors.MessageRejectError) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from RequestForPositions.
func (m RequestForPositions) GetResponseTransportType(f *field.ResponseTransportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for RequestForPositions.
func (m RequestForPositions) ResponseDestination() (*field.ResponseDestination, errors.MessageRejectError) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from RequestForPositions.
func (m RequestForPositions) GetResponseDestination(f *field.ResponseDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for RequestForPositions.
func (m RequestForPositions) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from RequestForPositions.
func (m RequestForPositions) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from RequestForPositions.
func (m RequestForPositions) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from RequestForPositions.
func (m RequestForPositions) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for RequestForPositions.
func (m RequestForPositions) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from RequestForPositions.
func (m RequestForPositions) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}
