// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package commands

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

func easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands(in *jlexer.Lexer, out *submitQuizRequest) {
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
		case "answers":
			if in.IsNull() {
				in.Skip()
				out.Answers = nil
			} else {
				in.Delim('[')
				if out.Answers == nil {
					if !in.IsDelim(']') {
						out.Answers = make(answers, 0, 1)
					} else {
						out.Answers = answers{}
					}
				} else {
					out.Answers = (out.Answers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 answerQuiz
					(v1).UnmarshalEasyJSON(in)
					out.Answers = append(out.Answers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands(out *jwriter.Writer, in submitQuizRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"answers\":"
		out.RawString(prefix[1:])
		if in.Answers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Answers {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v submitQuizRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v submitQuizRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *submitQuizRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *submitQuizRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands(l, v)
}
func easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands1(in *jlexer.Lexer, out *courseAndTopic) {
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
		case "CourseID":
			out.CourseID = string(in.String())
		case "TopicID":
			out.TopicID = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands1(out *jwriter.Writer, in courseAndTopic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"CourseID\":"
		out.RawString(prefix[1:])
		out.String(string(in.CourseID))
	}
	{
		const prefix string = ",\"TopicID\":"
		out.RawString(prefix)
		out.String(string(in.TopicID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v courseAndTopic) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v courseAndTopic) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *courseAndTopic) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *courseAndTopic) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands1(l, v)
}
func easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands2(in *jlexer.Lexer, out *answerQuiz) {
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
		case "question_id":
			out.QuestionID = string(in.String())
		case "answer":
			if in.IsNull() {
				in.Skip()
				out.Answer = nil
			} else {
				in.Delim('[')
				if out.Answer == nil {
					if !in.IsDelim(']') {
						out.Answer = make([]bool, 0, 64)
					} else {
						out.Answer = []bool{}
					}
				} else {
					out.Answer = (out.Answer)[:0]
				}
				for !in.IsDelim(']') {
					var v4 bool
					v4 = bool(in.Bool())
					out.Answer = append(out.Answer, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands2(out *jwriter.Writer, in answerQuiz) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"question_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.QuestionID))
	}
	{
		const prefix string = ",\"answer\":"
		out.RawString(prefix)
		if in.Answer == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Answer {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Bool(bool(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v answerQuiz) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v answerQuiz) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *answerQuiz) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *answerQuiz) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComRasikrrLearningPlatformInternalPortsHttpHandlersCoursesCommands2(l, v)
}
