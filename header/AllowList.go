package header

import (
	"container/list"
	"errors"
	"github.com/tutuvss/sip/core"
)

/**
* List of ALLOW headers. The sip message can have multiple Allow headers
*
 */
type AllowList struct {
	SIPHeaderList
}

/** default constructor
 */
func NewAllowList() *AllowList {
	this := &AllowList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ALLOW)
	return this
}

/**
 * Gets an Iterator of all the methods of the AllowHeader. Returns an empty
 *
 * Iterator if no methods are defined in this Allow Header.
 *
 *
 *
 * @return Iterator of String objects each identifing the methods of
 *
 * AllowHeader.
 *
 */

func (this *AllowList) GetMethods() *list.List {
	ll := list.New()
	for e := this.Front(); e != nil; e = e.Next() {
		ll.PushBack(e.Value.(*Allow).GetMethod())
	}
	return ll
}

/**
 * Sets the methods supported defined by this AllowHeader.
 *
 *
 *
 * @param methods - the Iterator of Strings defining the methods supported
 *
 * in this AllowHeader
 *
 * @throws ParseException which signals that an error has been reached
 *
 * unexpectedly while parsing the Strings defining the methods supported.
 *
 */

func (this *AllowList) SetMethods(methods *list.List) (ParseException error) {

	for e := methods.Front(); e != nil; e = e.Next() {
		allow := NewAllow()
		if str, ok := e.Value.(string); ok {
			if ParseException = allow.SetMethod(str); ParseException != nil {
				return ParseException
			}
		} else {
			return errors.New("ParseException: the method parameter is not string")
		}

		this.PushBack(allow)
	}
	return nil
}
