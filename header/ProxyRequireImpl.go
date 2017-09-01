package header

import (
	"errors"
	"github.com/tutuvss/sip/core"
)

/**
* ProxyRequire Header.
 */
type ProxyRequire struct {
	SIPHeader

	/** optiontag field
	 */
	optionTag string
}

/** Default  Constructor
 * @param s String to set
 */
func NewProxyRequire() *ProxyRequire {
	this := &ProxyRequire{}
	this.SIPHeader.super(core.SIPHeaderNames_PROXY_REQUIRE)
	return this
}

/** Constructor
 * @param s String to set
 */
func NewProxyRequireFromString(s string) *ProxyRequire {
	this := &ProxyRequire{}
	this.SIPHeader.super(core.SIPHeaderNames_PROXY_REQUIRE)
	this.optionTag = s
	return this
}

func (this *ProxyRequire) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode in canonical form.
 * @return String
 */
func (this *ProxyRequire) EncodeBody() string {
	return this.optionTag
}

/**
 * Sets the option tag value to the new supplied <var>optionTag</var>
 * parameter.
 *
 * @param optionTag - the new string value of the option tag.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the optionTag value.
 */
func (this *ProxyRequire) SetOptionTag(optionTag string) (ParseException error) {
	if optionTag == "" {
		return errors.New("NullPointerException: the optionTag parameter is null")
	}
	this.optionTag = optionTag
	return nil
}

/**
 * Gets the option tag of this OptionTag class.
 *
 * @return the string that identifies the option tag value.
 */
func (this *ProxyRequire) GetOptionTag() string {
	return this.optionTag
}
