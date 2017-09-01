package parser

import (
	"github.com/tutuvss/sip/core"
	"github.com/tutuvss/sip/header"
	"strings"
)

type Parser interface {
	Parse() (sh header.Header, ParseException error)
}

/** Generic header parser class. The parsers for various headers extend this
* class. To create a parser for a new header, extend this class and change
* the createParser class.
 */
type HeaderParser struct {
	SIPParser
}

/** Creates new HeaderParser
 * @param String to parse.
 */
func NewHeaderParser(header string) *HeaderParser {
	this := &HeaderParser{}

	this.SIPParser.super(header)
	this.SIPParser.GetLexer().SetLexerName("command_keywordLexer")

	return this
}

func NewHeaderParserFromLexer(lexer core.Lexer) *HeaderParser {
	this := &HeaderParser{}

	this.SIPParser.SetLexer(lexer)
	this.SIPParser.GetLexer().SetLexerName("command_keywordLexer")

	return this
}

func (this *HeaderParser) super(header string) {
	this.SIPParser.super(header)
	this.SIPParser.GetLexer().SetLexerName("command_keywordLexer")
}

func (this *HeaderParser) superFromLexer(lexer core.Lexer) {
	this.SetLexer(lexer)
	this.SIPParser.GetLexer().SetLexerName("command_keywordLexer")
}

/** Parse the SIP header from the buffer and return a parsed
 * structure.
 *@throws ParseException if there was an error parsing.
 */
func (this *HeaderParser) Parse() (sh header.Header, ParseException error) {
	lexer := this.GetLexer()
	var name string
	if name, ParseException = lexer.GetNextTokenByDelim(':'); ParseException != nil {
		return nil, ParseException
	}
	lexer.ConsumeK(1)
	body := strings.TrimSpace(lexer.GetLine())

	// we dont set any fields because the header is ok
	retval := header.NewExtension(name)
	retval.SetValue(body)
	return retval, nil
}

/** Parse the header name until the colon  and chew WS after that.
 */
func (this *HeaderParser) HeaderName(tok int) {
	this.GetLexer().Match(tok)
	this.GetLexer().SPorHT()
	this.GetLexer().Match(':')
	this.GetLexer().SPorHT()
}
