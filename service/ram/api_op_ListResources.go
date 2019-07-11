// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListResourcesRequest
type ListResourcesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The principal.
	Principal *string `locationName:"principal" type:"string"`

	// The Amazon Resource Names (ARN) of the resources.
	ResourceArns []string `locationName:"resourceArns" type:"list"`

	// The type of owner.
	//
	// ResourceOwner is a required field
	ResourceOwner ResourceOwner `locationName:"resourceOwner" type:"string" required:"true" enum:"true"`

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []string `locationName:"resourceShareArns" type:"list"`

	// The resource type.
	ResourceType *string `locationName:"resourceType" type:"string"`
}

// String returns the string representation
func (s ListResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResourcesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.ResourceOwner) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceOwner"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListResourcesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Principal != nil {
		v := *s.Principal

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "principal", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArns != nil {
		v := s.ResourceArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.ResourceOwner) > 0 {
		v := s.ResourceOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceOwner", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ResourceShareArns != nil {
		v := s.ResourceShareArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceShareArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListResourcesResponse
type ListResourcesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the resources.
	Resources []Resource `locationName:"resources" type:"list"`
}

// String returns the string representation
func (s ListResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListResourcesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Resources != nil {
		v := s.Resources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListResources = "ListResources"

// ListResourcesRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Lists the resources that the specified principal can access.
//
//    // Example sending a request using ListResourcesRequest.
//    req := client.ListResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListResources
func (c *Client) ListResourcesRequest(input *ListResourcesInput) ListResourcesRequest {
	op := &aws.Operation{
		Name:       opListResources,
		HTTPMethod: "POST",
		HTTPPath:   "/listresources",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListResourcesInput{}
	}

	req := c.newRequest(op, input, &ListResourcesOutput{})
	return ListResourcesRequest{Request: req, Input: input, Copy: c.ListResourcesRequest}
}

// ListResourcesRequest is the request type for the
// ListResources API operation.
type ListResourcesRequest struct {
	*aws.Request
	Input *ListResourcesInput
	Copy  func(*ListResourcesInput) ListResourcesRequest
}

// Send marshals and sends the ListResources API request.
func (r ListResourcesRequest) Send(ctx context.Context) (*ListResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourcesResponse{
		ListResourcesOutput: r.Request.Data.(*ListResourcesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResourcesRequestPaginator returns a paginator for ListResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResourcesRequest(input)
//   p := ram.NewListResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResourcesPaginator(req ListResourcesRequest) ListResourcesPaginator {
	return ListResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListResourcesInput
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

// ListResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResourcesPaginator struct {
	aws.Pager
}

func (p *ListResourcesPaginator) CurrentPage() *ListResourcesOutput {
	return p.Pager.CurrentPage().(*ListResourcesOutput)
}

// ListResourcesResponse is the response type for the
// ListResources API operation.
type ListResourcesResponse struct {
	*ListResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResources request.
func (r *ListResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
