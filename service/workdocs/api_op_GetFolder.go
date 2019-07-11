// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/GetFolderRequest
type GetFolderInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// The ID of the folder.
	//
	// FolderId is a required field
	FolderId *string `location:"uri" locationName:"FolderId" min:"1" type:"string" required:"true"`

	// Set to TRUE to include custom metadata in the response.
	IncludeCustomMetadata *bool `location:"querystring" locationName:"includeCustomMetadata" type:"boolean"`
}

// String returns the string representation
func (s GetFolderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFolderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFolderInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.FolderId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FolderId"))
	}
	if s.FolderId != nil && len(*s.FolderId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FolderId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFolderInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FolderId != nil {
		v := *s.FolderId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FolderId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IncludeCustomMetadata != nil {
		v := *s.IncludeCustomMetadata

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeCustomMetadata", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/GetFolderResponse
type GetFolderOutput struct {
	_ struct{} `type:"structure"`

	// The custom metadata on the folder.
	CustomMetadata map[string]string `min:"1" type:"map"`

	// The metadata of the folder.
	Metadata *FolderMetadata `type:"structure"`
}

// String returns the string representation
func (s GetFolderOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFolderOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CustomMetadata != nil {
		v := s.CustomMetadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "CustomMetadata", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Metadata", v, metadata)
	}
	return nil
}

const opGetFolder = "GetFolder"

// GetFolderRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Retrieves the metadata of the specified folder.
//
//    // Example sending a request using GetFolderRequest.
//    req := client.GetFolderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/GetFolder
func (c *Client) GetFolderRequest(input *GetFolderInput) GetFolderRequest {
	op := &aws.Operation{
		Name:       opGetFolder,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/folders/{FolderId}",
	}

	if input == nil {
		input = &GetFolderInput{}
	}

	req := c.newRequest(op, input, &GetFolderOutput{})
	return GetFolderRequest{Request: req, Input: input, Copy: c.GetFolderRequest}
}

// GetFolderRequest is the request type for the
// GetFolder API operation.
type GetFolderRequest struct {
	*aws.Request
	Input *GetFolderInput
	Copy  func(*GetFolderInput) GetFolderRequest
}

// Send marshals and sends the GetFolder API request.
func (r GetFolderRequest) Send(ctx context.Context) (*GetFolderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFolderResponse{
		GetFolderOutput: r.Request.Data.(*GetFolderOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFolderResponse is the response type for the
// GetFolder API operation.
type GetFolderResponse struct {
	*GetFolderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFolder request.
func (r *GetFolderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
