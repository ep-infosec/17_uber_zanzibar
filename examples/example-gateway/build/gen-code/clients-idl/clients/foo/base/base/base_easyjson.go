// Code generated by zanzibar
// @generated
// Checksum : f2fvPlVVW04RAlxcFgKV2Q==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package base

import (
	json "encoding/json"
	fmt "fmt"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson25720c23DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsFooBaseBase(in *jlexer.Lexer, out *Message) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var BodySet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "body":
			out.Body = string(in.String())
			BodySet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !BodySet {
		in.AddError(fmt.Errorf("key 'body' is required"))
	}
}
func easyjson25720c23EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsFooBaseBase(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix[1:])
		out.String(string(in.Body))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson25720c23EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsFooBaseBase(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson25720c23EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsFooBaseBase(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson25720c23DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsFooBaseBase(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson25720c23DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsFooBaseBase(l, v)
}
