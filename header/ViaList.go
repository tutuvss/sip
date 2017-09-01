package header

import (
	"github.com/tutuvss/sip/core"
)

/**
*  Keeps a list and a hashtable of via header functions.
 */
type ViaList struct {
	SIPHeaderList

	stringRep string
}

/**
 * Default Constructor.
 */
func NewViaList() *ViaList {
	this := &ViaList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_VIA)
	return this
}

func (this *ViaList) super(name string) {
	this.SIPHeaderList.super(name)
}

/**
         * make a clone of this header list. This supercedes the parent
         * function of the same signature - here for speed. Cloning based
	 * on introspection is slower.
	  *
         * @return clone of this Header list.
*/
func (this *ViaList) Clone() interface{} {
	vlist := NewViaList()
	//ListIterator it = this.listIterator();
	for e := this.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Via).Clone().(*Via)
		vlist.PushBack(v)
	}
	return vlist
}
