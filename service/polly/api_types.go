// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Provides lexicon name and lexicon content in string format. For more information,
// see Pronunciation Lexicon Specification (PLS) Version 1.0 (https://www.w3.org/TR/pronunciation-lexicon/).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/Lexicon
type Lexicon struct {
	_ struct{} `type:"structure"`

	// Lexicon content in string format. The content of a lexicon must be in PLS
	// format.
	Content *string `type:"string"`

	// Name of the lexicon.
	Name *string `type:"string"`
}

// String returns the string representation
func (s Lexicon) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Lexicon) MarshalFields(e protocol.FieldEncoder) error {
	if s.Content != nil {
		v := *s.Content

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Content", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains metadata describing the lexicon such as the number of lexemes, language
// code, and so on. For more information, see Managing Lexicons (http://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/LexiconAttributes
type LexiconAttributes struct {
	_ struct{} `type:"structure"`

	// Phonetic alphabet used in the lexicon. Valid values are ipa and x-sampa.
	Alphabet *string `type:"string"`

	// Language code that the lexicon applies to. A lexicon with a language code
	// such as "en" would be applied to all English languages (en-GB, en-US, en-AUS,
	// en-WLS, and so on.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// Date lexicon was last modified (a timestamp value).
	LastModified *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Number of lexemes in the lexicon.
	LexemesCount *int64 `type:"integer"`

	// Amazon Resource Name (ARN) of the lexicon.
	LexiconArn *string `type:"string"`

	// Total size of the lexicon, in characters.
	Size *int64 `type:"integer"`
}

// String returns the string representation
func (s LexiconAttributes) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LexiconAttributes) MarshalFields(e protocol.FieldEncoder) error {
	if s.Alphabet != nil {
		v := *s.Alphabet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Alphabet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LanguageCode) > 0 {
		v := s.LanguageCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LanguageCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.LexemesCount != nil {
		v := *s.LexemesCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LexemesCount", protocol.Int64Value(v), metadata)
	}
	if s.LexiconArn != nil {
		v := *s.LexiconArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LexiconArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Size != nil {
		v := *s.Size

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Size", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Describes the content of the lexicon.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/LexiconDescription
type LexiconDescription struct {
	_ struct{} `type:"structure"`

	// Provides lexicon metadata.
	Attributes *LexiconAttributes `type:"structure"`

	// Name of the lexicon.
	Name *string `type:"string"`
}

// String returns the string representation
func (s LexiconDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LexiconDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Attributes", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// SynthesisTask object that provides information about a speech synthesis task.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/SynthesisTask
type SynthesisTask struct {
	_ struct{} `type:"structure"`

	// Timestamp for the time the synthesis task was started.
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Optional language code for a synthesis task. This is only necessary if using
	// a bilingual voice, such as Aditi, which can be used for either Indian English
	// (en-IN) or Hindi (hi-IN).
	//
	// If a bilingual voice is used and no language code is specified, Amazon Polly
	// will use the default language of the bilingual voice. The default language
	// for any voice is the one returned by the DescribeVoices (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html)
	// operation for the LanguageCode parameter. For example, if no language code
	// is specified, Aditi will use Indian English rather than Hindi.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon
	// is the same as the language of the voice.
	LexiconNames []string `type:"list"`

	// The format in which the returned output will be encoded. For audio stream,
	// this will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json.
	OutputFormat OutputFormat `type:"string" enum:"true"`

	// Pathway for the output speech file.
	OutputUri *string `type:"string"`

	// Number of billable characters synthesized.
	RequestCharacters *int64 `type:"integer"`

	// The audio frequency specified in Hz.
	//
	// The valid values for mp3 and ogg_vorbis are "8000", "16000", and "22050".
	// The default value is "22050".
	//
	// Valid values for pcm are "8000" and "16000" The default value is "16000".
	SampleRate *string `type:"string"`

	// ARN for the SNS topic optionally used for providing status notification for
	// a speech synthesis task.
	SnsTopicArn *string `type:"string"`

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []SpeechMarkType `type:"list"`

	// The Amazon Polly generated identifier for a speech synthesis task.
	TaskId *string `type:"string"`

	// Current status of the individual speech synthesis task.
	TaskStatus TaskStatus `type:"string" enum:"true"`

	// Reason for the current status of a specific speech synthesis task, including
	// errors if the task has failed.
	TaskStatusReason *string `type:"string"`

	// Specifies whether the input text is plain text or SSML. The default value
	// is plain text.
	TextType TextType `type:"string" enum:"true"`

	// Voice ID to use for the synthesis.
	VoiceId VoiceId `type:"string" enum:"true"`
}

// String returns the string representation
func (s SynthesisTask) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SynthesisTask) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.LanguageCode) > 0 {
		v := s.LanguageCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LanguageCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LexiconNames != nil {
		v := s.LexiconNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "LexiconNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.OutputFormat) > 0 {
		v := s.OutputFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputFormat", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.OutputUri != nil {
		v := *s.OutputUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestCharacters != nil {
		v := *s.RequestCharacters

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestCharacters", protocol.Int64Value(v), metadata)
	}
	if s.SampleRate != nil {
		v := *s.SampleRate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SampleRate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SnsTopicArn != nil {
		v := *s.SnsTopicArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SnsTopicArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SpeechMarkTypes != nil {
		v := s.SpeechMarkTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SpeechMarkTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.TaskId != nil {
		v := *s.TaskId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TaskId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.TaskStatus) > 0 {
		v := s.TaskStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TaskStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TaskStatusReason != nil {
		v := *s.TaskStatusReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TaskStatusReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.TextType) > 0 {
		v := s.TextType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TextType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.VoiceId) > 0 {
		v := s.VoiceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VoiceId", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Description of the voice.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/Voice
type Voice struct {
	_ struct{} `type:"structure"`

	// Additional codes for languages available for the specified voice in addition
	// to its default language.
	//
	// For example, the default language for Aditi is Indian English (en-IN) because
	// it was first used for that language. Since Aditi is bilingual and fluent
	// in both Indian English and Hindi, this parameter would show the code hi-IN.
	AdditionalLanguageCodes []LanguageCode `type:"list"`

	// Gender of the voice.
	Gender Gender `type:"string" enum:"true"`

	// Amazon Polly assigned voice ID. This is the ID that you specify when calling
	// the SynthesizeSpeech operation.
	Id VoiceId `type:"string" enum:"true"`

	// Language code of the voice.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// Human readable name of the language in English.
	LanguageName *string `type:"string"`

	// Name of the voice (for example, Salli, Kendra, etc.). This provides a human
	// readable voice name that you might display in your application.
	Name *string `type:"string"`
}

// String returns the string representation
func (s Voice) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Voice) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdditionalLanguageCodes != nil {
		v := s.AdditionalLanguageCodes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AdditionalLanguageCodes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.Gender) > 0 {
		v := s.Gender

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Gender", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Id) > 0 {
		v := s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.LanguageCode) > 0 {
		v := s.LanguageCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LanguageCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LanguageName != nil {
		v := *s.LanguageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LanguageName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
