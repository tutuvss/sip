package header

import (
	"github.com/tutuvss/sip/core"
)

/**
* AcceptEncodingList of AccepEncoding headers.
 */
type AcceptEncodingList struct {
	SIPHeaderList
}

/** default constructor
 */
func NewAcceptEncodingList() *AcceptEncodingList {
	this := &AcceptEncodingList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ACCEPT_ENCODING)
	return this
}
