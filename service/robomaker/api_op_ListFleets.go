// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListFleetsRequest
type ListFleetsInput struct {
	_ struct{} `type:"structure"`

	// Optional filters to limit results.
	//
	// The filter name name is supported. When filtering, you must use the complete
	// value of the filtered item. You can use up to three filters.
	Filters []Filter `locationName:"filters" min:"1" type:"list"`

	// The maximum number of deployment job results returned by ListFleets in paginated
	// output. When this parameter is used, ListFleets only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListFleets request
	// with the returned nextToken value. This value can be between 1 and 100. If
	// this parameter is not used, then ListFleets returns up to 100 results and
	// a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListFleets request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListFleetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFleetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFleetsInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFleetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
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
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListFleetsResponse
type ListFleetsOutput struct {
	_ struct{} `type:"structure"`

	// A list of fleet details meeting the request criteria.
	FleetDetails []Fleet `locationName:"fleetDetails" type:"list"`

	// The nextToken value to include in a future ListDeploymentJobs request. When
	// the results of a ListFleets request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListFleetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFleetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FleetDetails != nil {
		v := s.FleetDetails

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "fleetDetails", metadata)
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

const opListFleets = "ListFleets"

// ListFleetsRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Returns a list of fleets. You can optionally provide filters to retrieve
// specific fleets.
//
//    // Example sending a request using ListFleetsRequest.
//    req := client.ListFleetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListFleets
func (c *Client) ListFleetsRequest(input *ListFleetsInput) ListFleetsRequest {
	op := &aws.Operation{
		Name:       opListFleets,
		HTTPMethod: "POST",
		HTTPPath:   "/listFleets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFleetsInput{}
	}

	req := c.newRequest(op, input, &ListFleetsOutput{})
	return ListFleetsRequest{Request: req, Input: input, Copy: c.ListFleetsRequest}
}

// ListFleetsRequest is the request type for the
// ListFleets API operation.
type ListFleetsRequest struct {
	*aws.Request
	Input *ListFleetsInput
	Copy  func(*ListFleetsInput) ListFleetsRequest
}

// Send marshals and sends the ListFleets API request.
func (r ListFleetsRequest) Send(ctx context.Context) (*ListFleetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFleetsResponse{
		ListFleetsOutput: r.Request.Data.(*ListFleetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFleetsRequestPaginator returns a paginator for ListFleets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFleetsRequest(input)
//   p := robomaker.NewListFleetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFleetsPaginator(req ListFleetsRequest) ListFleetsPaginator {
	return ListFleetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFleetsInput
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

// ListFleetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFleetsPaginator struct {
	aws.Pager
}

func (p *ListFleetsPaginator) CurrentPage() *ListFleetsOutput {
	return p.Pager.CurrentPage().(*ListFleetsOutput)
}

// ListFleetsResponse is the response type for the
// ListFleets API operation.
type ListFleetsResponse struct {
	*ListFleetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFleets request.
func (r *ListFleetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
