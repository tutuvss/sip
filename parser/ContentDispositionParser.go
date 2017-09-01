package parser

import (
	"github.com/tutuvss/sip/core"
	"github.com/tutuvss/sip/header"
)

/** SIPParser for ContentLanguage header.
 */
type ContentDispositionParser struct {
	ParametersParser
}

/**
 * Creates a new instance of ContentDispositionParser
 * @param contentDisposition the header to parse
 */
func NewContentDispositionParser(contentDisposition string) *ContentDispositionParser {
	this := &ContentDispositionParser{}
	this.ParametersParser.super(contentDisposition)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewContentDispositionParserFromLexer(lexer core.Lexer) *ContentDispositionParser {
	this := &ContentDispositionParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

/** parse the ContentDispositionHeader String header
 * @return Header (ContentDispositionList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ContentDispositionParser) Parse() (sh header.Header, ParseException error) {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_DISPOSITION)

	cd := header.NewContentDisposition()
	cd.SetHeaderName(core.SIPHeaderNames_CONTENT_DISPOSITION)

	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)

	token := lexer.GetNextToken()
	cd.SetDispositionType(token.GetTokenValue())
	lexer.SPorHT()
	if ParseException = this.ParametersParser.Parse(cd); ParseException != nil {
		return nil, ParseException
	}

	lexer.SPorHT()
	lexer.Match('\n')

	return cd, nil
}
