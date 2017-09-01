package header

import (
	"errors"
	"github.com/tutuvss/sip/core"
	"github.com/tutuvss/sip/address"
	"strconv"
	"strings"
)

/**
 * The generic AuthenticationHeader
 */

type Authentication struct {
	Parameters

	scheme string
}

func NewAuthentication(name string) *Authentication {
	this := &Authentication{}
	this.Parameters.super(name)
	this.parameters.SetSeparator(core.SIPSeparatorNames_COMMA)
	this.scheme = ParameterNames_DIGEST

	return this
}

func (this *Authentication) super(name string) {
	this.Parameters.super(name)
	this.parameters.SetSeparator(core.SIPSeparatorNames_COMMA)
	this.scheme = ParameterNames_DIGEST
}

/** Set the specified parameter.
 * Bug reported by Dominic Sparks.
 *
 * @param name  -- name of the  parameter
 * @param value  -- value of the parameter.
 */
func (this *Authentication) SetParameter(name, value string) error {
	if name == "" {
		return errors.New("NullPointerException: null name")
	}
	nv := this.parameters.GetNameValue(strings.ToLower(name))
	if nv == nil {
		nv = core.NewNameValue(name, value)
		if strings.ToLower(name) == (ParameterNames_QOP) ||
			strings.ToLower(name) == (ParameterNames_REALM) ||
			strings.ToLower(name) == (ParameterNames_CNONCE) ||
			strings.ToLower(name) == (ParameterNames_NONCE) ||
			strings.ToLower(name) == (ParameterNames_USERNAME) ||
			strings.ToLower(name) == (ParameterNames_DOMAIN) ||
			strings.ToLower(name) == (ParameterNames_OPAQUE) ||
			strings.ToLower(name) == (ParameterNames_NEXT_NONCE) ||
			strings.ToLower(name) == (ParameterNames_URI) ||
			strings.ToLower(name) == (ParameterNames_ALGORITHM) ||
			strings.ToLower(name) == (ParameterNames_RESPONSE) {
			//if value == "" {//TODO by LY
			//	return errors.New("NullPointerException: null value")
			//}
			if strings.HasPrefix(value, core.SIPSeparatorNames_DOUBLE_QUOTE) {
				return errors.New("ParseException: " + value + " : Unexpected DOUBLE_QUOTE")
			}
			nv.SetQuotedValue()
		}
		this.parameters.SetNameValue(nv)
	} else {
		nv.SetValue(value)
	}

	return nil
}

/** This is only used for the parser interface.
 *@param challenge -- the challenge from which the parameters are
 * extracted.
 */
func (this *Authentication) SetChallenge(challenge Challenge) {
	this.scheme = challenge.scheme
	this.parameters = challenge.authParams
}

func (this *Authentication) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode in canonical form.
 * @return canonical string.
 */
func (this *Authentication) EncodeBody() string {
	return this.scheme + core.SIPSeparatorNames_SP + this.parameters.String()
}

/**
 * Sets the scheme of the challenge information for this
 * AuthenticationHeaderHeader.  For example, Digest.
 *
 * @param scheme - the new string value that identifies the challenge
 * information scheme.
 *
 */
func (this *Authentication) SetScheme(scheme string) {
	this.scheme = scheme
}

/**
 * Returns the scheme of the challenge information for this
 * AuthenticationHeaderHeader.
 *
 * @return the string value of the challenge information.
 *
 */
func (this *Authentication) GetScheme() string {
	return this.scheme
}

/**
 * Sets the Realm of the WWWAuthenicateHeader to the <var>realm</var>
 * parameter value. Realm strings MUST be globally unique.  It is
 * RECOMMENDED that a realm string contain a hostname or domain name.
 * Realm strings SHOULD present a human-readable identifier that can be
 * rendered to a user.
 *
 * @param realm the new Realm String of this WWWAuthenicateHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the realm.
 *
 */
func (this *Authentication) SetRealm(realm string) (ParseException error) {
	if realm == "" {
		return errors.New("NullPointerException: The realm parameter is null")
	}
	this.SetParameter(ParameterNames_REALM, realm)
	return nil
}

/**
 * Returns the Realm value of this WWWAuthenicateHeader. This convenience
 * method returns only the realm of the complete Challenge.
 *
 * @return the String representing the Realm information, null if value is
 * not Set.
 *
 */
func (this *Authentication) GetRealm() string {
	return this.GetParameter(ParameterNames_REALM)
}

/**
 * Sets the Nonce of the WWWAuthenicateHeader to the <var>nonce</var>
 * parameter value.
 *
 * @param nonce - the new nonce String of this WWWAuthenicateHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the nonce value.
 *
 */
func (this *Authentication) SetNonce(nonce string) (ParseException error) {
	if nonce == "" {
		return errors.New("NullPointerException: The nonce parameter is null")
	}
	this.SetParameter(ParameterNames_NONCE, nonce)
	return nil
}

/**
 * Returns the Nonce value of this WWWAuthenicateHeader.
 *
 * @return the String representing the nonce information, null if value is
 * not Set.
 *
 */
func (this *Authentication) GetNonce() string {
	return this.GetParameter(ParameterNames_NONCE)
}

/**
 * Sets the URI of the WWWAuthenicateHeader to the <var>uri</var>
 * parameter value.
 *
 * @param uri - the new URI of this WWWAuthenicateHeader.
 *
 */
func (this *Authentication) SetURI(uri address.URI) error {
	if uri != nil {
		nv := core.NewNameValue(ParameterNames_URI, uri.String())
		nv.SetQuotedValue()
		this.parameters.SetNameValue(nv)
		return nil
	} else {
		return errors.New("NullPointerException: Null URI")
	}
}

/**
 * Returns the URI value of this WWWAuthenicateHeader,
 * for example DigestURI.
 *
 * @return the URI representing the URI information, null if value is
 * not Set.
 *
 */
func (this *Authentication) GetURI() address.URI {
	url := this.GetParameter(ParameterNames_URI)
	return address.NewURIImpl(url)
}

/**
 * Sets the Algorithm of the WWWAuthenicateHeader to the new
 * <var>algorithm</var> parameter value.
 *
 * @param algorithm - the new algorithm String of this WWWAuthenicateHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the algorithm value.
 *
 */
func (this *Authentication) SetAlgorithm(algorithm string) (ParseException error) {
	if algorithm == "" {
		return errors.New("NullPointerException: null arg")
	}
	this.SetParameter(ParameterNames_ALGORITHM, algorithm)
	return nil
}

/**
 * Returns the Algorithm value of this WWWAuthenicateHeader.
 *
 * @return the String representing the Algorithm information, null if the
 * value is not Set.
 *
 */
func (this *Authentication) GetAlgorithm() string {
	return this.GetParameter(ParameterNames_ALGORITHM)
}

/**
 * Sets the Qop value of the WWWAuthenicateHeader to the new
 * <var>qop</var> parameter value.
 *
 * @param qop - the new Qop string of this WWWAuthenicateHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the Qop value.
 *
 */
func (this *Authentication) SetQop(qop string) (ParseException error) {
	if qop == "" {
		return errors.New("NullPointerException: null arg")
	}
	this.SetParameter(ParameterNames_QOP, qop)
	return nil
}

/**
 * Returns the Qop value of this WWWAuthenicateHeader.
 *
 * @return the string representing the Qop information, null if the
 * value is not Set.
 *
 */
func (this *Authentication) GetQop() string {
	return this.GetParameter(ParameterNames_QOP)
}

/**
 * Sets the Opaque value of the WWWAuthenicateHeader to the new
 * <var>opaque</var> parameter value.
 *
 * @param opaque - the new Opaque string of this WWWAuthenicateHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the opaque value.
 *
 */
func (this *Authentication) SetOpaque(opaque string) (ParseException error) {
	if opaque == "" {
		return errors.New("NullPointerException: null arg")
	}
	this.SetParameter(ParameterNames_OPAQUE, opaque)
	return nil
}

/**
 * Returns the Opaque value of this WWWAuthenicateHeader.
 *
 * @return the String representing the Opaque information, null if the
 * value is not Set.
 *
 */
func (this *Authentication) GetOpaque() string {
	return this.GetParameter(ParameterNames_OPAQUE)
}

/**
 * Sets the Domain of the WWWAuthenicateHeader to the <var>domain</var>
 * parameter value.
 *
 * @param domain - the new Domain string of this WWWAuthenicateHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the domain.
 *
 */
func (this *Authentication) SetDomain(domain string) (ParseException error) {
	if domain == "" {
		return errors.New("NullPointerException: null arg")
	}
	this.SetParameter(ParameterNames_DOMAIN, domain)
	return nil
}

/**
 * Returns the Domain value of this WWWAuthenicateHeader.
 *
 * @return the String representing the Domain information, null if value is
 * not Set.
 *
 */
func (this *Authentication) GetDomain() string {
	return this.GetParameter(ParameterNames_DOMAIN)
}

/**
 * Sets the value of the stale parameter of the WWWAuthenicateHeader to the
 * <var>stale</var> parameter value.
 *
 * @param stale - the new boolean value of the stale parameter.
 *
 */
func (this *Authentication) SetStale(stale bool) {
	if stale == true {
		this.SetParameter(ParameterNames_STALE, "true")
	} else {
		this.SetParameter(ParameterNames_STALE, "false")
	}

}

/**
 * Returns the boolean value of the state paramater of this
 * WWWAuthenicateHeader.
 *
 * @return the boolean representing if the challenge is stale.
 *
 */
func (this *Authentication) IsStale() bool {
	stale := this.GetParameter(ParameterNames_STALE)
	if stale == "true" {
		return true
	} else {
		return false
	}
}

/** Set the CNonce.
 *
 * @param cnonce -- a nonce string.
 */
func (this *Authentication) SetCNonce(cnonce string) (ParseException error) {
	this.SetParameter(ParameterNames_CNONCE, cnonce)
	return nil
}

/** Get the CNonce.
 *
 *@return the cnonce value.
 */
func (this *Authentication) GetCNonce() string {
	return this.GetParameter(ParameterNames_CNONCE)
}

func (this *Authentication) GetNonceCount() int {
	//return this.GetParameterAsHexInt(ParameterNames_NC);
	s := this.GetParameter(ParameterNames_NONCE_COUNT)
	nCount, _ := strconv.ParseInt(s, 10, 32)
	return int(nCount)
}

/** Set the nonce count pakrameter.
 * Bug fix sent in by Andreas Byström
 */

func (this *Authentication) SetNonceCount(nonceCount int) (ParseException error) {
	if nonceCount < 0 {
		return errors.New("ParseException: bad value")
	}

	nc := strconv.FormatUint(uint64(nonceCount), 16)

	base := "00000000"
	nc = base[0:8-len(nc)] + nc
	this.SetParameter(ParameterNames_NC, nc)
	return nil
}

/**
 * Get the RESPONSE value (or null if it does not exist).
 *
 * @return String response parameter value.
 */
func (this *Authentication) GetResponse() string {
	return this.GetParameterValue(ParameterNames_RESPONSE)
}

/** Set the Response.
 *
 *@param response to Set.
 */
func (this *Authentication) SetResponse(response string) (ParseException error) {
	if response == "" {
		return errors.New("NullPointerException: Null parameter")
	}

	this.SetParameter(ParameterNames_RESPONSE, response)
	return nil
}

/**
 * Returns the Username value of this AuthenticationHeader.
 * This convenience method returns only the username of the
 * complete Response.
 *
 * @return the String representing the Username information,
 * null if value is not Set.
 *
 * @since JAIN SIP v1.1
 *
 */
func (this *Authentication) GetUsername() string {
	return this.GetParameter(ParameterNames_USERNAME)
}

/**
 * Sets the Username of the AuthenticationHeader to
 * the <var>username</var> parameter value.
 *
 * @param username the new Username String of this AuthenticationHeader.
 *
 * @throws ParseException which signals that an error has been reached
 *
 * unexpectedly while parsing the username.
 *
 * @since JAIN SIP v1.1
 *
 */
func (this *Authentication) SetUsername(username string) (ParseException error) {
	return this.SetParameter(ParameterNames_USERNAME, username)
}
