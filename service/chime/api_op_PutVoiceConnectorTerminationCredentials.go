// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorTerminationCredentialsRequest
type PutVoiceConnectorTerminationCredentialsInput struct {
	_ struct{} `type:"structure"`

	// The termination SIP credentials.
	Credentials []Credential `type:"list"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutVoiceConnectorTerminationCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutVoiceConnectorTerminationCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutVoiceConnectorTerminationCredentialsInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorTerminationCredentialsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Credentials != nil {
		v := s.Credentials

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Credentials", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorTerminationCredentialsOutput
type PutVoiceConnectorTerminationCredentialsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutVoiceConnectorTerminationCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorTerminationCredentialsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutVoiceConnectorTerminationCredentials = "PutVoiceConnectorTerminationCredentials"

// PutVoiceConnectorTerminationCredentialsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Adds termination SIP credentials for the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using PutVoiceConnectorTerminationCredentialsRequest.
//    req := client.PutVoiceConnectorTerminationCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorTerminationCredentials
func (c *Client) PutVoiceConnectorTerminationCredentialsRequest(input *PutVoiceConnectorTerminationCredentialsInput) PutVoiceConnectorTerminationCredentialsRequest {
	op := &aws.Operation{
		Name:       opPutVoiceConnectorTerminationCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/termination/credentials?operation=put",
	}

	if input == nil {
		input = &PutVoiceConnectorTerminationCredentialsInput{}
	}

	req := c.newRequest(op, input, &PutVoiceConnectorTerminationCredentialsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutVoiceConnectorTerminationCredentialsRequest{Request: req, Input: input, Copy: c.PutVoiceConnectorTerminationCredentialsRequest}
}

// PutVoiceConnectorTerminationCredentialsRequest is the request type for the
// PutVoiceConnectorTerminationCredentials API operation.
type PutVoiceConnectorTerminationCredentialsRequest struct {
	*aws.Request
	Input *PutVoiceConnectorTerminationCredentialsInput
	Copy  func(*PutVoiceConnectorTerminationCredentialsInput) PutVoiceConnectorTerminationCredentialsRequest
}

// Send marshals and sends the PutVoiceConnectorTerminationCredentials API request.
func (r PutVoiceConnectorTerminationCredentialsRequest) Send(ctx context.Context) (*PutVoiceConnectorTerminationCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutVoiceConnectorTerminationCredentialsResponse{
		PutVoiceConnectorTerminationCredentialsOutput: r.Request.Data.(*PutVoiceConnectorTerminationCredentialsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutVoiceConnectorTerminationCredentialsResponse is the response type for the
// PutVoiceConnectorTerminationCredentials API operation.
type PutVoiceConnectorTerminationCredentialsResponse struct {
	*PutVoiceConnectorTerminationCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutVoiceConnectorTerminationCredentials request.
func (r *PutVoiceConnectorTerminationCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
