package parser

import (
	"github.com/tutuvss/sip/core"
	"github.com/tutuvss/sip/header"
)

/** SIPParser for Require header.
 */
type RequireParser struct {
	HeaderParser
}

/** Creates a new instance of RequireParser
 * @param require the header to parse
 */
func NewRequireParser(require string) *RequireParser {
	this := &RequireParser{}
	this.HeaderParser.super(require)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewRequireParserFromLexer(lexer core.Lexer) *RequireParser {
	this := &RequireParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (RequireList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *RequireParser) Parse() (sh header.Header, ParseException error) {
	requireList := header.NewRequireList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_REQUIRE)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		r := header.NewRequire()
		r.SetHeaderName(core.SIPHeaderNames_REQUIRE)

		// Parsing the option tag
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		r.SetOptionTag(token.GetTokenValue())
		lexer.SPorHT()

		requireList.PushBack(r)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			lexer.Match(',')
			lexer.SPorHT()

			r = header.NewRequire()

			// Parsing the option tag
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			r.SetOptionTag(token.GetTokenValue())
			lexer.SPorHT()

			requireList.PushBack(r)
		}

	}

	return requireList, nil
}
