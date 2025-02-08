// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package faq

import (
	json "encoding/json"
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

func easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersFaq(in *jlexer.Lexer, out *postQuestionRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
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
		case "title":
			out.Title = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "category_id":
			out.CategoryID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersFaq(out *jwriter.Writer, in postQuestionRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"category_id\":"
		out.RawString(prefix)
		out.String(string(in.CategoryID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v postQuestionRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersFaq(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v postQuestionRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersFaq(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *postQuestionRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersFaq(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *postQuestionRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersFaq(l, v)
}
