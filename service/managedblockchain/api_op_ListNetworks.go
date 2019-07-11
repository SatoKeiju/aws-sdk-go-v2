// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListNetworksInput
type ListNetworksInput struct {
	_ struct{} `type:"structure"`

	// An optional framework specifier. If provided, only networks of this framework
	// type are listed.
	Framework Framework `location:"querystring" locationName:"framework" type:"string" enum:"true"`

	// The maximum number of networks to list.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The name of the network.
	Name *string `location:"querystring" locationName:"name" type:"string"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// An optional status specifier. If provided, only networks currently in this
	// status are listed.
	Status NetworkStatus `location:"querystring" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListNetworksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNetworksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNetworksInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNetworksInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.Framework) > 0 {
		v := s.Framework

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "framework", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListNetworksOutput
type ListNetworksOutput struct {
	_ struct{} `type:"structure"`

	// An array of NetworkSummary objects that contain configuration properties
	// for each network.
	Networks []NetworkSummary `type:"list"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListNetworksOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNetworksOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Networks != nil {
		v := s.Networks

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Networks", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListNetworks = "ListNetworks"

// ListNetworksRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns information about the networks in which the current AWS account has
// members.
//
//    // Example sending a request using ListNetworksRequest.
//    req := client.ListNetworksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListNetworks
func (c *Client) ListNetworksRequest(input *ListNetworksInput) ListNetworksRequest {
	op := &aws.Operation{
		Name:       opListNetworks,
		HTTPMethod: "GET",
		HTTPPath:   "/networks",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListNetworksInput{}
	}

	req := c.newRequest(op, input, &ListNetworksOutput{})
	return ListNetworksRequest{Request: req, Input: input, Copy: c.ListNetworksRequest}
}

// ListNetworksRequest is the request type for the
// ListNetworks API operation.
type ListNetworksRequest struct {
	*aws.Request
	Input *ListNetworksInput
	Copy  func(*ListNetworksInput) ListNetworksRequest
}

// Send marshals and sends the ListNetworks API request.
func (r ListNetworksRequest) Send(ctx context.Context) (*ListNetworksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNetworksResponse{
		ListNetworksOutput: r.Request.Data.(*ListNetworksOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNetworksRequestPaginator returns a paginator for ListNetworks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNetworksRequest(input)
//   p := managedblockchain.NewListNetworksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNetworksPaginator(req ListNetworksRequest) ListNetworksPaginator {
	return ListNetworksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListNetworksInput
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

// ListNetworksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNetworksPaginator struct {
	aws.Pager
}

func (p *ListNetworksPaginator) CurrentPage() *ListNetworksOutput {
	return p.Pager.CurrentPage().(*ListNetworksOutput)
}

// ListNetworksResponse is the response type for the
// ListNetworks API operation.
type ListNetworksResponse struct {
	*ListNetworksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNetworks request.
func (r *ListNetworksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
