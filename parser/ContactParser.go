package parser

import (
	"github.com/tutuvss/sip/core"
	"github.com/tutuvss/sip/header"
	"strconv"
)

/**
* A parser for The SIP contact header.
 */

type ContactParser struct {
	AddressParametersParser
}

func NewContactParser(contact string) *ContactParser {
	this := &ContactParser{}
	this.AddressParametersParser.super(contact)
	return this
}

func NewContactParserFromLexer(lexer core.Lexer) *ContactParser {
	this := &ContactParser{}
	this.AddressParametersParser.superFromLexer(lexer)
	return this
}

func (this *ContactParser) Parse() (sh header.Header, ParseException error) {
	retval := header.NewContactList()
	// past the header name and the colon.
	lexer := this.GetLexer()
	var la byte
	this.HeaderName(TokenTypes_CONTACT)
	for {
		contact := header.NewContact()
		if la, _ = lexer.LookAheadK(0); la == '*' {
			lexer.Match('*')
			contact.SetWildCardFlag(true)
		} else {
			if ParseException = this.AddressParametersParser.Parse(contact); ParseException != nil {
				return nil, ParseException
			}

			if contact.HasParameter(header.ParameterNames_EXPIRES) {
				if _, ParseException = strconv.Atoi(contact.GetParameter(header.ParameterNames_EXPIRES)); ParseException != nil {
					return nil, ParseException
				}
			}
		}
		retval.AddContact(contact)
		lexer.SPorHT()
		if la, _ = lexer.LookAheadK(0); la == ',' {
			lexer.Match(',')
			lexer.SPorHT()
		} else if la, _ = lexer.LookAheadK(0); la == '\n' {
			break
		} else {
			return nil, this.CreateParseException("unexpected char")
		}
	}
	return retval, nil
}
