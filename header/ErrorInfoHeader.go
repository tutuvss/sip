package header

import "github.com/tutuvss/sip/address"

/**

 * The Error-Info header field provides a pointer to additional information

 * about the error status response.

 * SIP UACs have user interface capabilities ranging from pop-up windows and

 * audio on PC softclients to audio-only on "black" phones or endpoints

 * connected via gateways.  Rather than forcing a server generating an error

 * to choose between sending an error status code with a detailed reason

 * phrase and playing an audio recording, the Error-Info header field allows

 * both to be sent. The UAC then has the choice of which error indicator to

 * render to the caller.

 * <p>

 * A UAC MAY treat a SIP or SIPS URI in an Error-Info header field as if it

 * were a Contact in a redirect and generate a new INVITE, resulting in a

 * recorded announcement session being established.  A non-SIP URI MAY be

 * rendered to the user.

 * <p>

 * Examples:<br>

 * <code>SIP/2.0 404 The number you have dialed is not in service<br>

 * Error-Info: sip:not-in-service-recording@atlanta.com</code>

 */

type ErrorInfoHeader interface {
	ParametersHeader

	/**

	 * Sets the ErrorInfo of the ErrorInfoHeader to the <var>errorInfo</var>

	 * parameter value.

	 *

	 * @param errorInfo the new ErrorInfo of this ErrorInfoHeader.

	 */

	SetErrorInfo(errorInfo address.URI)

	/**

	 * Returns the ErrorInfo value of this ErrorInfoHeader. This message

	 * may return null if a String message identifies the ErrorInfo.

	 *

	 * @return the URI representing the ErrorInfo.

	 */

	GetErrorInfo() address.URI

	/**

	 * Sets the Error information message to the new <var>message</var> value

	 * supplied to this method.

	 *

	 * @param message - the new string value that represents the error message.

	 * @throws ParseException which signals that an error has been reached

	 * unexpectedly while parsing the error message.

	 */

	SetErrorMessage(message string) (ParseException error)

	/**

	 * Get the Error information message of this ErrorInfoHeader.

	 *

	 * @return the stringified version of the ErrorInfo header.

	 */

	GetErrorMessage() string
}
