package header

import (
	"container/list"
	"github.com/tutuvss/sip/core"
)

type ParametersHeader interface {
	Header
	GetParameter(name string) string
	GetParameterValue(name string) string //interface{} TODO
	GetParameterNames() *list.List
	HasParameters() bool
	RemoveParameter(name string)
	SetParameter(name, value string) (ParseException error)
	SetQuotedParameter(name, value string)
	HasParameter(parameterName string) bool
	RemoveParameters()
	GetParameters() *core.NameValueList
	SetParameters(parameters *core.NameValueList)
	GetNameValue(parameterName string) *core.NameValue
}
