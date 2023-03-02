package tests

import (
	"github.com/betprophet1/easyjson"
	"github.com/betprophet1/easyjson/jwriter"
)

//easyjson:json
type NestedInterfaces struct {
	Value interface{}
	Slice []interface{}
	Map   map[string]interface{}
}

type NestedEasyMarshaler struct {
	EasilyMarshaled bool
}

var _ easyjson.Marshaler = &NestedEasyMarshaler{}

func (i *NestedEasyMarshaler) MarshalEasyJSON(w *jwriter.Writer) {
	// We use this method only to indicate that easyjson.Marshaler
	// interface was really used while encoding.
	i.EasilyMarshaled = true
}
