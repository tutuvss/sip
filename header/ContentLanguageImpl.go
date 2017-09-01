package header

import (
	"github.com/tutuvss/sip/core"
)

/**
* ContentLanguage header
* <pre>
*Fielding, et al.            Standards Track                   [Page 118]
*RFC 2616                        HTTP/1.1                       June 1999
*
*  14.12 Content-Language
*
*   The Content-Language entity-header field describes the natural
*   language(s) of the intended audience for the enclosed entity. Note
*   that this might not be equivalent to all the languages used within
*   the entity-body.
*
*       Content-Language  = "Content-Language" ":" 1#language-tag
*
*   Language tags are defined in section 3.10. The primary purpose of
*   Content-Language is to allow a user to identify and differentiate
*   entities according to the user's own preferred language. Thus, if the
*   body content is intended only for a Danish-literate audience, the
*   appropriate field is
*
*       Content-Language: da
*
*   If no Content-Language is specified, the default is that the content
*   is intended for all language audiences. This might mean that the
*   sender does not consider it to be specific to any natural language,
*   or that the sender does not know for which language it is intended.
*
*   Multiple languages MAY be listed for content that is intended for
*   multiple audiences. For example, a rendition of the "Treaty of
*   Waitangi," presented simultaneously in the original Maori and English
*   versions, would call for
*
*       Content-Language: mi, en
*
*   However, just because multiple languages are present within an entity
*   does not mean that it is intended for multiple linguistic audiences.
*   An example would be a beginner's language primer, such as "A First
*   Lesson in Latin," which is clearly intended to be used by an
*   English-literate audience. In this case, the Content-Language would
*   properly only include "en".
*
*   Content-Language MAY be applied to any media type -- it is not
*   limited to textual documents.
*</pre>
 */
type ContentLanguage struct {
	SIPHeader

	/** languageTag field.
	 */
	locale string
}

func NewContentLanguage() *ContentLanguage {
	this := &ContentLanguage{}
	this.SIPHeader.super(core.SIPHeaderNames_CONTENT_LANGUAGE)
	return this

}

/** Default constructor.
 * @param lang String to set
 */
func NewContentLanguageFromString(languageTag string) *ContentLanguage {
	this := &ContentLanguage{}
	this.SIPHeader.super(core.SIPHeaderNames_CONTENT_LANGUAGE)
	this.locale = languageTag
	return this
}

func (this *ContentLanguage) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Canonical encoding of the  value of the header.
 * @return encoded body of header.
 */
func (this *ContentLanguage) EncodeBody() string {
	return this.locale
}

/** get the languageTag field.
 * @return String
 */
func (this *ContentLanguage) GetLanguageTag() string {
	return this.locale
}

/** set the languageTag field
 * @param languageTag -- language tag to set.
 */
func (this *ContentLanguage) SetLanguageTag(languageTag string) {
	this.locale = languageTag
}

/**
 * Gets the language value of the ContentLanguageHeader.
 *
 *
 *
 * @return the Locale value of this ContentLanguageHeader
 *
 */
func (this *ContentLanguage) GetContentLanguage() string {
	return this.locale
}

/**
 * Sets the language parameter of this ContentLanguageHeader.
 *
 *
 *
 * @param language - the new Locale value of the language of
 *
 * ContentLanguageHeader
 *
 */
func (this *ContentLanguage) SetContentLanguage(language string) {
	this.locale = language
}
