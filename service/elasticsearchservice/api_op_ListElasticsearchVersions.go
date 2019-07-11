// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Container for the parameters to the ListElasticsearchVersions operation.
// Use MaxResults to control the maximum number of results to retrieve in a
// single call.
//
// Use NextToken in response to retrieve more results. If the received response
// does not contain a NextToken, then there are no more results to retrieve.
type ListElasticsearchVersionsInput struct {
	_ struct{} `type:"structure"`

	// Set this value to limit the number of results returned. Value provided must
	// be greater than 10 else it wont be honored.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// Paginated APIs accepts NextToken input to returns next page results and provides
	// a NextToken output in the response which can be used by the client to retrieve
	// more results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListElasticsearchVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListElasticsearchVersionsInput) MarshalFields(e protocol.FieldEncoder) error {

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

// Container for the parameters for response received from ListElasticsearchVersions
// operation.
type ListElasticsearchVersionsOutput struct {
	_ struct{} `type:"structure"`

	// List of supported elastic search versions.
	ElasticsearchVersions []string `type:"list"`

	// Paginated APIs accepts NextToken input to returns next page results and provides
	// a NextToken output in the response which can be used by the client to retrieve
	// more results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListElasticsearchVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListElasticsearchVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ElasticsearchVersions != nil {
		v := s.ElasticsearchVersions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ElasticsearchVersions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
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

const opListElasticsearchVersions = "ListElasticsearchVersions"

// ListElasticsearchVersionsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// List all supported Elasticsearch versions
//
//    // Example sending a request using ListElasticsearchVersionsRequest.
//    req := client.ListElasticsearchVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListElasticsearchVersionsRequest(input *ListElasticsearchVersionsInput) ListElasticsearchVersionsRequest {
	op := &aws.Operation{
		Name:       opListElasticsearchVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/versions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListElasticsearchVersionsInput{}
	}

	req := c.newRequest(op, input, &ListElasticsearchVersionsOutput{})
	return ListElasticsearchVersionsRequest{Request: req, Input: input, Copy: c.ListElasticsearchVersionsRequest}
}

// ListElasticsearchVersionsRequest is the request type for the
// ListElasticsearchVersions API operation.
type ListElasticsearchVersionsRequest struct {
	*aws.Request
	Input *ListElasticsearchVersionsInput
	Copy  func(*ListElasticsearchVersionsInput) ListElasticsearchVersionsRequest
}

// Send marshals and sends the ListElasticsearchVersions API request.
func (r ListElasticsearchVersionsRequest) Send(ctx context.Context) (*ListElasticsearchVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListElasticsearchVersionsResponse{
		ListElasticsearchVersionsOutput: r.Request.Data.(*ListElasticsearchVersionsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListElasticsearchVersionsRequestPaginator returns a paginator for ListElasticsearchVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListElasticsearchVersionsRequest(input)
//   p := elasticsearchservice.NewListElasticsearchVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListElasticsearchVersionsPaginator(req ListElasticsearchVersionsRequest) ListElasticsearchVersionsPaginator {
	return ListElasticsearchVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListElasticsearchVersionsInput
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

// ListElasticsearchVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListElasticsearchVersionsPaginator struct {
	aws.Pager
}

func (p *ListElasticsearchVersionsPaginator) CurrentPage() *ListElasticsearchVersionsOutput {
	return p.Pager.CurrentPage().(*ListElasticsearchVersionsOutput)
}

// ListElasticsearchVersionsResponse is the response type for the
// ListElasticsearchVersions API operation.
type ListElasticsearchVersionsResponse struct {
	*ListElasticsearchVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListElasticsearchVersions request.
func (r *ListElasticsearchVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
