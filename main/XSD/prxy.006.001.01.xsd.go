// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
	"time"
)

// AccountIdentification4Choice is x
type AccountIdentification4Choice struct {
	Othr *GenericAccountIdentification1 `xml:"Othr"`
}

// AddressType2Code ...
type AddressType2Code string

// BICFIIdentifier ...
type BICFIIdentifier string

// BranchAndFinancialInstitutionIdentification5 is x
type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId *FinancialInstitutionIdentification8 `xml:"FinInstnId"`
	BrnchId    *BranchData2                         `xml:"BrnchId"`
}

// BranchData2 ...
type BranchData2 struct {
	Id      string          `xml:"Id"`
	Nm      string          `xml:"Nm"`
	PstlAdr *PostalAddress6 `xml:"PstlAdr"`
}

// CashAccount40 is x
type CashAccount40 struct {
	Id *AccountIdentification4Choice `xml:"Id"`
	Tp *CashAccountType2ChoiceProxy  `xml:"Tp"`
	Nm string                        `xml:"Nm"`
}

// CashAccountType2ChoiceProxy is x
type CashAccountType2ChoiceProxy struct {
	Prtry string `xml:"Prtry"`
}

// ClearingSystemMemberIdentification2 ...
type ClearingSystemMemberIdentification2 struct {
	BICFI string `xml:"BICFI"`
}

// CountryCode ...
type CountryCode string

// Document is x
type Document struct {
	PrxyNqryRspn *ProxyEnquiryResponseV01 `xml:"PrxyNqryRspn"`
}

// Exact4AlphaNumericText ...
type Exact4AlphaNumericText string

// ExternalFinancialInstitutionIdentification1Code ...
type ExternalFinancialInstitutionIdentification1Code string

// ExternalStatusReason1Code ...
type ExternalStatusReason1Code string

// FinancialIdentificationSchemeName1Choice ...
type FinancialIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// FinancialInstitutionIdentification8 is x
type FinancialInstitutionIdentification8 struct {
	BICFI       string                               `xml:"BICFI"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	Nm          string                               `xml:"Nm"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr"`
}

// GenericAccountIdentification1 is x
type GenericAccountIdentification1 struct {
	Id string `xml:"Id"`
}

// GenericFinancialIdentification1 is x
type GenericFinancialIdentification1 struct {
	Id      string                                    `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                    `xml:"Issr"`
}

// GenericIdentification47 ...
type GenericIdentification47 struct {
	Id   string `xml:"Id"`
	Issr string `xml:"Issr"`
}

// GroupHeader60 is x
type GroupHeader60 struct {
	MsgId   string         `xml:"MsgId"`
	CreDtTm time.Time      `xml:"CreDtTm"`
	MsgRcpt *Party12Choice `xml:"MsgRcpt"`
}

// IndividualPerson28 ...
type IndividualPerson28 struct {
	GvnNm  string `xml:"GvnNm"`
	MddlNm string `xml:"MddlNm"`
	Nm     string `xml:"Nm"`
}

// ISODate ...
type ISODate time.Time

// ISODateTime ...
type ISODateTime time.Time

// Max12Text ...
type Max12Text string

// Max140Text ...
type Max140Text string

// Max16Text ...
type Max16Text string

// Max34Text ...
type Max34Text string

// Max350Text ...
type Max350Text string

// Max35Text ...
type Max35Text string

// Max4AlphaNumericText ...
type Max4AlphaNumericText string

// Max70Text ...
type Max70Text string

// Min11Max11Text ...
type Min11Max11Text string

// Organisation22 ...
type Organisation22 struct {
	Nm      string                   `xml:"Nm"`
	RegnDt  time.Time                `xml:"RegnDt"`
	TpOfOrg *OrganisationType1Choice `xml:"TpOfOrg"`
}

// OrganisationType1Choice ...
type OrganisationType1Choice struct {
	Prtry *GenericIdentification47 `xml:"Prtry"`
}

// OriginalGroupInformation3 is x
type OriginalGroupInformation3 struct {
	OrgnlMsgId   string    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId string    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm time.Time `xml:"OrgnlCreDtTm"`
}

// Party12Choice is x
type Party12Choice struct {
	Agt *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

// Party30Choice ...
type Party30Choice struct {
	Org      *Organisation22     `xml:"Org"`
	IndvPrsn *IndividualPerson28 `xml:"IndvPrsn"`
}

// PostalAddress6 ...
type PostalAddress6 struct {
	AdrTp       string   `xml:"AdrTp"`
	Dept        string   `xml:"Dept"`
	SubDept     string   `xml:"SubDept"`
	StrtNm      string   `xml:"StrtNm"`
	BldgNb      string   `xml:"BldgNb"`
	PstCd       string   `xml:"PstCd"`
	TwnNm       string   `xml:"TwnNm"`
	CtrySubDvsn string   `xml:"CtrySubDvsn"`
	Ctry        string   `xml:"Ctry"`
	AdrLine     []string `xml:"AdrLine"`
}

// ProxyAccountStatusCode ...
type ProxyAccountStatusCode string

// ProxyAccountType ...
type ProxyAccountType string

// ProxyEnquiryAccount1 is x
type ProxyEnquiryAccount1 struct {
	Agt      *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
	Acct     *CashAccount40                                `xml:"Acct"`
	Sts      string                                        `xml:"Sts"`
	AcctHldr *Party30Choice                                `xml:"AcctHldr"`
}

// ProxyEnquiryDefinition1 is x
type ProxyEnquiryDefinition1 struct {
	Tp  string `xml:"Tp"`
	Val string `xml:"Val"`
	Sts string `xml:"Sts"`
}

// ProxyEnquiryInformation1 is x
type ProxyEnquiryInformation1 struct {
	RegnId      string                                        `xml:"RegnId"`
	DsplNm      string                                        `xml:"DsplNm"`
	Ptcpt       *BranchAndFinancialInstitutionIdentification5 `xml:"Ptcpt"`
	PrxyInf     *ProxyEnquiryDefinition1                      `xml:"PrxyInf"`
	AcctInf     *ProxyEnquiryAccount1                         `xml:"AcctInf"`
	RegnSts     string                                        `xml:"RegnSts"`
	RegnDtTm    time.Time                                     `xml:"RegnDtTm"`
	PreAuthrsd  bool                                          `xml:"PreAuthrsd"`
	SplmtryData *BISupplementaryData1                         `xml:"SplmtryData"`
}

// ProxyEnquiryResponse1 is x
type ProxyEnquiryResponse1 struct {
	PrxRspnSts string                      `xml:"PrxRspnSts"`
	StsRsnInf  *ProxyStatusChoice          `xml:"StsRsnInf"`
	Rspn       []*ProxyEnquiryInformation1 `xml:"Rspn"`
}

// ProxyEnquiryResponseV01 is x
type ProxyEnquiryResponseV01 struct {
	GrpHdr      *GroupHeader60             `xml:"GrpHdr"`
	OrgnlGrpInf *OriginalGroupInformation3 `xml:"OrgnlGrpInf"`
	NqryRspn    *ProxyEnquiryResponse1     `xml:"NqryRspn"`
}

// ProxyEnquiryStatusCode ...
type ProxyEnquiryStatusCode string

// ProxyRegistrationStatusCode ...
type ProxyRegistrationStatusCode string

// ProxyStatusChoice is x
type ProxyStatusChoice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ProxyStatusCode ...
type ProxyStatusCode string

// YesNoIndicator ...
type YesNoIndicator bool

// BISupplementaryData1 is x
type BISupplementaryData1 struct {
	XMLName  xml.Name                      `xml:"BI_SupplementaryData1"`
	PlcAndNm string                        `xml:"PlcAndNm"`
	Envlp    *BISupplementaryDataEnvelope1 `xml:"Envlp"`
}

// BISupplementaryDataEnvelope1 is x
type BISupplementaryDataEnvelope1 struct {
	XMLName xml.Name         `xml:"BI_SupplementaryDataEnvelope1"`
	Cstmr   *BIAddtlCstmrInf `xml:"Cstmr"`
}

// BIAddtlCstmrInf is [PS: Customer's Hometown]
type BIAddtlCstmrInf struct {
	XMLName  xml.Name `xml:"BI_AddtlCstmrInf"`
	Tp       string   `xml:"Tp"`
	Id       string   `xml:"Id"`
	RsdntSts string   `xml:"RsdntSts"`
	TwnNm    string   `xml:"TwnNm"`
}
