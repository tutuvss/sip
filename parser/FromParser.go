package parser

import (
	"github.com/tutuvss/sip/core"
	"github.com/tutuvss/sip/address"
	"github.com/tutuvss/sip/header"
)

/** From header parser.
 */
type FromParser struct {
	AddressParametersParser
}

func NewFromParser(from string) *FromParser {
	this := &FromParser{}
	this.AddressParametersParser.super(from)
	return this
}

func (this *FromParser) super(from string) {
	this.AddressParametersParser.super(from)
}

func NewFromParserFromLexer(lexer core.Lexer) *FromParser {
	this := &FromParser{}
	this.AddressParametersParser.superFromLexer(lexer)
	return this
}

func (this *FromParser) Parse() (sh header.Header, ParseException error) {
	from := header.NewFrom()
	this.HeaderName(TokenTypes_FROM)

	if ParseException = this.AddressParametersParser.Parse(from); ParseException != nil {
		return nil, ParseException
	}

	this.GetLexer().Match('\n')
	addr := from.GetAddress().(*address.AddressImpl)
	if addr.GetAddressType() == address.ADDRESS_SPEC {
		// the parameters are header parameters.
		if from.GetAddress().GetURI().IsSipURI() {
			sipUri, _ := from.GetAddress().GetURI().(*address.SipURIImpl)
			parms := sipUri.GetUriParms()
			if parms != nil && parms.Len() > 0 {
				from.SetParameters(parms)
				sipUri.RemoveUriParms()
			}
		}
	}

	return from, nil
}
