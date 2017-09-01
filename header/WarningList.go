package header

import "github.com/tutuvss/sip/core"

/**
* A Warning SIPObject. (A list of Warning headers).
 */
type WarningList struct {
	SIPHeaderList
}

/** Constructor.
 *
 */
func NewWarningList() *WarningList {
	this := &WarningList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_WARNING)
	return this
}
