package header

import "github.com/tutuvss/sip/core"

/**
* Authorization SIP header.
*
* @see ProxyAuthorization
 */
type Authorization struct {
	Authentication
}

/** Default constructor.
 */
func NewAuthorization() *Authorization {
	this := &Authorization{}
	this.Authentication.super(core.SIPHeaderNames_AUTHORIZATION)
	return this
}
