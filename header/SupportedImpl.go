package header

import (
	"errors"
	"github.com/tutuvss/sip/core"
)

/**
*Supported SIP Header.
 */
type Supported struct {
	SIPHeader

	/* the Option field
	 */
	optionTag string
}

/** default constructor
 */
func NewSupported() *Supported {
	this := &Supported{}
	this.SIPHeader.super(core.SIPHeaderNames_SUPPORTED)
	return this
}

/** Constructor
 * @param option_tag String to set
 */
func NewSupportedFromString(option_tag string) *Supported {
	this := &Supported{}
	this.SIPHeader.super(core.SIPHeaderNames_SUPPORTED)
	this.optionTag = option_tag
	return this
}

/** Return canonical form of the header.
 * @return encoded header.
 */
func (this *Supported) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Just the encoded body of the header.
*@return the string encoded header body.
 */
func (this *Supported) EncodeBody() string {
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
func (this *Supported) SetOptionTag(optionTag string) (ParseException error) {
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
func (this *Supported) GetOptionTag() string {
	return this.optionTag
}
