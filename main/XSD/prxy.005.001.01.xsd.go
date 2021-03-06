// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"time"
)

// AccountIdentification4Choice is Account Only Acct ID (Othr) enquiry info group
type AccountIdentification4Choice struct {
	Othr *GenericAccountIdentification1 `xml:"Othr"`
}

// AddressType2Code ...
type AddressType2Code string

// BICFIIdentifier ...
type BICFIIdentifier string

// BranchAndFinancialInstitutionIdentification5 is Account Only Agent FI ID enquiry info group
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

// CashAccount40 is Account Only Acct Type enquiry info group
type CashAccount40 struct {
	Id *AccountIdentification4Choice `xml:"Id"`
	Tp *CashAccountType2ChoiceProxy  `xml:"Tp"`
}

// CashAccountType2ChoiceProxy is Account Type for Account-Only enquiry.
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
	PrxyNqryReq *ProxyEnquiryV01 `xml:"PrxyNqryReq"`
}

// ExternalFinancialInstitutionIdentification1Code ...
type ExternalFinancialInstitutionIdentification1Code string

// FinancialIdentificationSchemeName1Choice ...
type FinancialIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// FinancialInstitutionIdentification8 is Account Only Agent FI ID Othr enquiry info group
type FinancialInstitutionIdentification8 struct {
	BICFI       string                               `xml:"BICFI"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	Nm          string                               `xml:"Nm"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr"`
}

// GenericAccountIdentification1 is Account ID for Account-Only enquiry
type GenericAccountIdentification1 struct {
	Id string `xml:"Id"`
}

// GenericFinancialIdentification1 is Bank ID for Account Only enquiry
type GenericFinancialIdentification1 struct {
	Id      string                                    `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                    `xml:"Issr"`
}

// GroupHeader59 is x
type GroupHeader59 struct {
	MsgId   string         `xml:"MsgId"`
	CreDtTm time.Time      `xml:"CreDtTm"`
	MsgSndr *Party12Choice `xml:"MsgSndr"`
}

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

// Max35Text ...
type Max35Text string

// Max3NumericText ...
type Max3NumericText string

// Max70Text ...
type Max70Text string

// Min11Max11Text ...
type Min11Max11Text string

// Party12Choice is x
type Party12Choice struct {
	Agt *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
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

// ProxyAccountType ...
type ProxyAccountType string

// ProxyDefinition1 ...
type ProxyDefinition1 struct {
	Tp  string `xml:"Tp"`
	Val string `xml:"Val"`
}

// ProxyEnquiry11 ...
type ProxyEnquiry11 struct {
	PrxyRtrvl *ProxyDefinition1 `xml:"PrxyRtrvl"`
	RegnSts   string            `xml:"RegnSts"`
	NbOfItm   string            `xml:"NbOfItm"`
}

// ProxyEnquiry12 is Maximum number of items to return in records in the prxy.006.001.01 Proxy Enquiry response.
type ProxyEnquiry12 struct {
	Agt     *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
	Acct    *CashAccount40                                `xml:"Acct"`
	RegnSts string                                        `xml:"RegnSts"`
	NbOfItm string                                        `xml:"NbOfItm"`
}

// ProxyEnquiry13 ...
type ProxyEnquiry13 struct {
	PrxyNqry *ProxyDefinition1                             `xml:"PrxyNqry"`
	Agt      *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
	Acct     *CashAccount40                                `xml:"Acct"`
}

// ProxyEnquiry14 ...
type ProxyEnquiry14 struct {
	PrxyNqry *ProxyDefinition1 `xml:"PrxyNqry"`
}

// ProxyEnquiryChoice1 is Account Only enquiry info group - Choice.
type ProxyEnquiryChoice1 struct {
	RegnId  string             `xml:"RegnId"`
	ScndId  *ScndIdDefinition1 `xml:"ScndId"`
	AccOnly *ProxyEnquiry12    `xml:"AccOnly"`
}

// ProxyEnquiryV01 is x
type ProxyEnquiryV01 struct {
	GrpHdr *GroupHeader59       `xml:"GrpHdr"`
	Nqry   *ProxyEnquiryChoice1 `xml:"Nqry"`
}

// ProxyRegistrationStatusCode ...
type ProxyRegistrationStatusCode string

// ScndIdDefinition1 is x
type ScndIdDefinition1 struct {
	Tp  string `xml:"Tp"`
	Val string `xml:"Val"`
}
