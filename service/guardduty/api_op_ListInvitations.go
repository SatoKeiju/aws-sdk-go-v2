// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListInvitationsRequest
type ListInvitationsInput struct {
	_ struct{} `type:"structure"`

	// You can use this parameter to indicate the maximum number of items that you
	// want in the response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListInvitationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListInvitationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListInvitationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListInvitationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// ListInvitations response object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListInvitationsResponse
type ListInvitationsOutput struct {
	_ struct{} `type:"structure"`

	// A list of invitation descriptions.
	Invitations []Invitation `locationName:"invitations" type:"list"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls
	// to the action fill nextToken in the request with the value of NextToken from
	// the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListInvitationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListInvitationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Invitations != nil {
		v := s.Invitations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "invitations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListInvitations = "ListInvitations"

// ListInvitationsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Lists all GuardDuty membership invitations that were sent to the current
// AWS account.
//
//    // Example sending a request using ListInvitationsRequest.
//    req := client.ListInvitationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListInvitations
func (c *Client) ListInvitationsRequest(input *ListInvitationsInput) ListInvitationsRequest {
	op := &aws.Operation{
		Name:       opListInvitations,
		HTTPMethod: "GET",
		HTTPPath:   "/invitation",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListInvitationsInput{}
	}

	req := c.newRequest(op, input, &ListInvitationsOutput{})
	return ListInvitationsRequest{Request: req, Input: input, Copy: c.ListInvitationsRequest}
}

// ListInvitationsRequest is the request type for the
// ListInvitations API operation.
type ListInvitationsRequest struct {
	*aws.Request
	Input *ListInvitationsInput
	Copy  func(*ListInvitationsInput) ListInvitationsRequest
}

// Send marshals and sends the ListInvitations API request.
func (r ListInvitationsRequest) Send(ctx context.Context) (*ListInvitationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListInvitationsResponse{
		ListInvitationsOutput: r.Request.Data.(*ListInvitationsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListInvitationsRequestPaginator returns a paginator for ListInvitations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListInvitationsRequest(input)
//   p := guardduty.NewListInvitationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListInvitationsPaginator(req ListInvitationsRequest) ListInvitationsPaginator {
	return ListInvitationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListInvitationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListInvitationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListInvitationsPaginator struct {
	aws.Pager
}

func (p *ListInvitationsPaginator) CurrentPage() *ListInvitationsOutput {
	return p.Pager.CurrentPage().(*ListInvitationsOutput)
}

// ListInvitationsResponse is the response type for the
// ListInvitations API operation.
type ListInvitationsResponse struct {
	*ListInvitationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListInvitations request.
func (r *ListInvitationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
