// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"time"
)

// Document ...
type Document *Document

// ISODateTime ...
type ISODateTime time.Time

// Max20000Text ...
type Max20000Text string

// Max350Text ...
type Max350Text string

// Max35Text ...
type Max35Text string

// MessageReference is x
type MessageReference struct {
	Ref string `xml:"Ref"`
}

// RejectionReason2 is x
type RejectionReason2 struct {
	RjctgPtyRsn string    `xml:"RjctgPtyRsn"`
	RjctnDtTm   time.Time `xml:"RjctnDtTm"`
	ErrLctn     string    `xml:"ErrLctn"`
	RsnDesc     string    `xml:"RsnDesc"`
	AddtlData   string    `xml:"AddtlData"`
}

// MessageRejectV01 is x
type MessageRejectV01 struct {
	RltdRef *MessageReference `xml:"RltdRef"`
	Rsn     *RejectionReason2 `xml:"Rsn"`
}
