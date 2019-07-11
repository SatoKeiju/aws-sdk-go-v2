// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure used to request a project be deleted.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DeleteProjectRequest
type DeleteProjectInput struct {
	_ struct{} `type:"structure"`

	// Unique project identifier.
	//
	// ProjectId is a required field
	ProjectId *string `location:"uri" locationName:"projectId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProjectInput"}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteProjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ProjectId != nil {
		v := *s.ProjectId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "projectId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure used in response to request to delete a project.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DeleteProjectResult
type DeleteProjectOutput struct {
	_ struct{} `type:"structure"`

	// Resources which were deleted.
	DeletedResources []Resource `locationName:"deletedResources" type:"list"`

	// Resources which were not deleted, due to a risk of losing potentially important
	// data or files.
	OrphanedResources []Resource `locationName:"orphanedResources" type:"list"`
}

// String returns the string representation
func (s DeleteProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeletedResources != nil {
		v := s.DeletedResources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "deletedResources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.OrphanedResources != nil {
		v := s.OrphanedResources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "orphanedResources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDeleteProject = "DeleteProject"

// DeleteProjectRequest returns a request value for making API operation for
// AWS Mobile.
//
// Delets a project in AWS Mobile Hub.
//
//    // Example sending a request using DeleteProjectRequest.
//    req := client.DeleteProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DeleteProject
func (c *Client) DeleteProjectRequest(input *DeleteProjectInput) DeleteProjectRequest {
	op := &aws.Operation{
		Name:       opDeleteProject,
		HTTPMethod: "DELETE",
		HTTPPath:   "/projects/{projectId}",
	}

	if input == nil {
		input = &DeleteProjectInput{}
	}

	req := c.newRequest(op, input, &DeleteProjectOutput{})
	return DeleteProjectRequest{Request: req, Input: input, Copy: c.DeleteProjectRequest}
}

// DeleteProjectRequest is the request type for the
// DeleteProject API operation.
type DeleteProjectRequest struct {
	*aws.Request
	Input *DeleteProjectInput
	Copy  func(*DeleteProjectInput) DeleteProjectRequest
}

// Send marshals and sends the DeleteProject API request.
func (r DeleteProjectRequest) Send(ctx context.Context) (*DeleteProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProjectResponse{
		DeleteProjectOutput: r.Request.Data.(*DeleteProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProjectResponse is the response type for the
// DeleteProject API operation.
type DeleteProjectResponse struct {
	*DeleteProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProject request.
func (r *DeleteProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
