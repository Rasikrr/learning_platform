// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entity

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

func easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity(in *jlexer.Lexer, out *TopicContent) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "topic_id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.TopicID).UnmarshalText(data))
			}
		case "content":
			out.Content = string(in.String())
		case "additional_resources":
			if in.IsNull() {
				in.Skip()
				out.AdditionalResources = nil
			} else {
				in.Delim('[')
				if out.AdditionalResources == nil {
					if !in.IsDelim(']') {
						out.AdditionalResources = make([]string, 0, 4)
					} else {
						out.AdditionalResources = []string{}
					}
				} else {
					out.AdditionalResources = (out.AdditionalResources)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.AdditionalResources = append(out.AdditionalResources, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "video_urls":
			if in.IsNull() {
				in.Skip()
				out.VideoURLs = nil
			} else {
				in.Delim('[')
				if out.VideoURLs == nil {
					if !in.IsDelim(']') {
						out.VideoURLs = make([]string, 0, 4)
					} else {
						out.VideoURLs = []string{}
					}
				} else {
					out.VideoURLs = (out.VideoURLs)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.VideoURLs = append(out.VideoURLs, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "image_urls":
			if in.IsNull() {
				in.Skip()
				out.ImageURLs = nil
			} else {
				in.Delim('[')
				if out.ImageURLs == nil {
					if !in.IsDelim(']') {
						out.ImageURLs = make([]string, 0, 4)
					} else {
						out.ImageURLs = []string{}
					}
				} else {
					out.ImageURLs = (out.ImageURLs)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.ImageURLs = append(out.ImageURLs, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity(out *jwriter.Writer, in TopicContent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"topic_id\":"
		out.RawString(prefix)
		out.RawText((in.TopicID).MarshalText())
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"additional_resources\":"
		out.RawString(prefix)
		if in.AdditionalResources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.AdditionalResources {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"video_urls\":"
		out.RawString(prefix)
		if in.VideoURLs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.VideoURLs {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"image_urls\":"
		out.RawString(prefix)
		if in.ImageURLs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.ImageURLs {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TopicContent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TopicContent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TopicContent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TopicContent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity(l, v)
}
func easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity1(in *jlexer.Lexer, out *Topic) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "course_id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.CourseID).UnmarshalText(data))
			}
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "content":
			if in.IsNull() {
				in.Skip()
				out.Content = nil
			} else {
				if out.Content == nil {
					out.Content = new(TopicContent)
				}
				(*out.Content).UnmarshalEasyJSON(in)
			}
		case "quizzes":
			if in.IsNull() {
				in.Skip()
				out.Quizzes = nil
			} else {
				in.Delim('[')
				if out.Quizzes == nil {
					if !in.IsDelim(']') {
						out.Quizzes = make([]*Quiz, 0, 8)
					} else {
						out.Quizzes = []*Quiz{}
					}
				} else {
					out.Quizzes = (out.Quizzes)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *Quiz
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(Quiz)
						}
						(*v10).UnmarshalEasyJSON(in)
					}
					out.Quizzes = append(out.Quizzes, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "practical_tasks":
			if in.IsNull() {
				in.Skip()
				out.PracticalTasks = nil
			} else {
				in.Delim('[')
				if out.PracticalTasks == nil {
					if !in.IsDelim(']') {
						out.PracticalTasks = make([]*PracticalTask, 0, 8)
					} else {
						out.PracticalTasks = []*PracticalTask{}
					}
				} else {
					out.PracticalTasks = (out.PracticalTasks)[:0]
				}
				for !in.IsDelim(']') {
					var v11 *PracticalTask
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(PracticalTask)
						}
						(*v11).UnmarshalEasyJSON(in)
					}
					out.PracticalTasks = append(out.PracticalTasks, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "order_number":
			out.OrderNumber = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity1(out *jwriter.Writer, in Topic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"course_id\":"
		out.RawString(prefix)
		out.RawText((in.CourseID).MarshalText())
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	if in.Content != nil {
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		(*in.Content).MarshalEasyJSON(out)
	}
	if len(in.Quizzes) != 0 {
		const prefix string = ",\"quizzes\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v12, v13 := range in.Quizzes {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.PracticalTasks) != 0 {
		const prefix string = ",\"practical_tasks\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v14, v15 := range in.PracticalTasks {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"order_number\":"
		out.RawString(prefix)
		out.Int(int(in.OrderNumber))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Topic) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Topic) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Topic) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Topic) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity1(l, v)
}
func easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity2(in *jlexer.Lexer, out *Quiz) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "topic_id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.TopicID).UnmarshalText(data))
			}
		case "question":
			out.Question = string(in.String())
		case "options":
			if in.IsNull() {
				in.Skip()
				out.Options = nil
			} else {
				in.Delim('[')
				if out.Options == nil {
					if !in.IsDelim(']') {
						out.Options = make([]string, 0, 4)
					} else {
						out.Options = []string{}
					}
				} else {
					out.Options = (out.Options)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.Options = append(out.Options, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "correct_answers":
			if in.IsNull() {
				in.Skip()
				out.CorrectAnswers = nil
			} else {
				in.Delim('[')
				if out.CorrectAnswers == nil {
					if !in.IsDelim(']') {
						out.CorrectAnswers = make([]bool, 0, 64)
					} else {
						out.CorrectAnswers = []bool{}
					}
				} else {
					out.CorrectAnswers = (out.CorrectAnswers)[:0]
				}
				for !in.IsDelim(']') {
					var v17 bool
					v17 = bool(in.Bool())
					out.CorrectAnswers = append(out.CorrectAnswers, v17)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "multiple_choice":
			out.MultipleChoice = bool(in.Bool())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity2(out *jwriter.Writer, in Quiz) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"topic_id\":"
		out.RawString(prefix)
		out.RawText((in.TopicID).MarshalText())
	}
	{
		const prefix string = ",\"question\":"
		out.RawString(prefix)
		out.String(string(in.Question))
	}
	{
		const prefix string = ",\"options\":"
		out.RawString(prefix)
		if in.Options == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.Options {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.String(string(v19))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"correct_answers\":"
		out.RawString(prefix)
		if in.CorrectAnswers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.CorrectAnswers {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.Bool(bool(v21))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"multiple_choice\":"
		out.RawString(prefix)
		out.Bool(bool(in.MultipleChoice))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Quiz) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Quiz) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Quiz) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Quiz) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity2(l, v)
}
func easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity3(in *jlexer.Lexer, out *PracticalTask) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "topic_id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.TopicID).UnmarshalText(data))
			}
		case "description":
			out.Description = string(in.String())
		case "difficulty_level":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DifficultyLevel).UnmarshalJSON(data))
			}
		case "starter_code":
			out.StarterCode = string(in.String())
		case "expected_output":
			out.ExpectedOutput = string(in.String())
		case "order_number":
			out.OrderNumber = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity3(out *jwriter.Writer, in PracticalTask) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"topic_id\":"
		out.RawString(prefix)
		out.RawText((in.TopicID).MarshalText())
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"difficulty_level\":"
		out.RawString(prefix)
		out.Raw((in.DifficultyLevel).MarshalJSON())
	}
	{
		const prefix string = ",\"starter_code\":"
		out.RawString(prefix)
		out.String(string(in.StarterCode))
	}
	{
		const prefix string = ",\"expected_output\":"
		out.RawString(prefix)
		out.String(string(in.ExpectedOutput))
	}
	{
		const prefix string = ",\"order_number\":"
		out.RawString(prefix)
		out.Int(int(in.OrderNumber))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PracticalTask) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PracticalTask) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PracticalTask) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PracticalTask) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity3(l, v)
}
func easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity4(in *jlexer.Lexer, out *Course) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "title":
			out.Title = string(in.String())
		case "image_url":
			if in.IsNull() {
				in.Skip()
				out.ImageURL = nil
			} else {
				if out.ImageURL == nil {
					out.ImageURL = new(string)
				}
				*out.ImageURL = string(in.String())
			}
		case "category":
			(out.Category).UnmarshalEasyJSON(in)
		case "description":
			out.Description = string(in.String())
		case "topics":
			if in.IsNull() {
				in.Skip()
				out.Topics = nil
			} else {
				in.Delim('[')
				if out.Topics == nil {
					if !in.IsDelim(']') {
						out.Topics = make([]*Topic, 0, 8)
					} else {
						out.Topics = []*Topic{}
					}
				} else {
					out.Topics = (out.Topics)[:0]
				}
				for !in.IsDelim(']') {
					var v22 *Topic
					if in.IsNull() {
						in.Skip()
						v22 = nil
					} else {
						if v22 == nil {
							v22 = new(Topic)
						}
						(*v22).UnmarshalEasyJSON(in)
					}
					out.Topics = append(out.Topics, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "created_by":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.CreatedBy).UnmarshalText(data))
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity4(out *jwriter.Writer, in Course) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"image_url\":"
		out.RawString(prefix)
		if in.ImageURL == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ImageURL))
		}
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		(in.Category).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	if len(in.Topics) != 0 {
		const prefix string = ",\"topics\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v23, v24 := range in.Topics {
				if v23 > 0 {
					out.RawByte(',')
				}
				if v24 == nil {
					out.RawString("null")
				} else {
					(*v24).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"created_by\":"
		out.RawString(prefix)
		out.RawText((in.CreatedBy).MarshalText())
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Course) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Course) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Course) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Course) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity4(l, v)
}
func easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity5(in *jlexer.Lexer, out *Category) {
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
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "name":
			out.Name = string(in.String())
		case "created_by":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.CreatedBy).UnmarshalText(data))
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity5(out *jwriter.Writer, in Category) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"created_by\":"
		out.RawString(prefix)
		out.RawText((in.CreatedBy).MarshalText())
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Category) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Category) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCc42051bEncodeGithubComRasikrrLearningPlatformInternalDomainEntity5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Category) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Category) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCc42051bDecodeGithubComRasikrrLearningPlatformInternalDomainEntity5(l, v)
}
