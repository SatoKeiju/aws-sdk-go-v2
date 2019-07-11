// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetIntentsRequest
type GetIntentsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of intents to return in the response. The default is 10.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Substring to match in intent names. An intent will be returned if any part
	// of its name matches the substring. For example, "xyz" matches both "xyzabc"
	// and "abcxyz."
	NameContains *string `location:"querystring" locationName:"nameContains" min:"1" type:"string"`

	// A pagination token that fetches the next page of intents. If the response
	// to this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of intents, specify the pagination token
	// in the next request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetIntentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIntentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIntentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIntentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NameContains != nil {
		v := *s.NameContains

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nameContains", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetIntentsResponse
type GetIntentsOutput struct {
	_ struct{} `type:"structure"`

	// An array of Intent objects. For more information, see PutBot.
	Intents []IntentMetadata `locationName:"intents" type:"list"`

	// If the response is truncated, the response includes a pagination token that
	// you can specify in your next request to fetch the next page of intents.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetIntentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIntentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Intents != nil {
		v := s.Intents

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "intents", metadata)
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

const opGetIntents = "GetIntents"

// GetIntentsRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns intent information as follows:
//
//    * If you specify the nameContains field, returns the $LATEST version of
//    all intents that contain the specified string.
//
//    * If you don't specify the nameContains field, returns information about
//    the $LATEST version of all intents.
//
// The operation requires permission for the lex:GetIntents action.
//
//    // Example sending a request using GetIntentsRequest.
//    req := client.GetIntentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetIntents
func (c *Client) GetIntentsRequest(input *GetIntentsInput) GetIntentsRequest {
	op := &aws.Operation{
		Name:       opGetIntents,
		HTTPMethod: "GET",
		HTTPPath:   "/intents/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetIntentsInput{}
	}

	req := c.newRequest(op, input, &GetIntentsOutput{})
	return GetIntentsRequest{Request: req, Input: input, Copy: c.GetIntentsRequest}
}

// GetIntentsRequest is the request type for the
// GetIntents API operation.
type GetIntentsRequest struct {
	*aws.Request
	Input *GetIntentsInput
	Copy  func(*GetIntentsInput) GetIntentsRequest
}

// Send marshals and sends the GetIntents API request.
func (r GetIntentsRequest) Send(ctx context.Context) (*GetIntentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIntentsResponse{
		GetIntentsOutput: r.Request.Data.(*GetIntentsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetIntentsRequestPaginator returns a paginator for GetIntents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetIntentsRequest(input)
//   p := lexmodelbuildingservice.NewGetIntentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetIntentsPaginator(req GetIntentsRequest) GetIntentsPaginator {
	return GetIntentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetIntentsInput
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

// GetIntentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetIntentsPaginator struct {
	aws.Pager
}

func (p *GetIntentsPaginator) CurrentPage() *GetIntentsOutput {
	return p.Pager.CurrentPage().(*GetIntentsOutput)
}

// GetIntentsResponse is the response type for the
// GetIntents API operation.
type GetIntentsResponse struct {
	*GetIntentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIntents request.
func (r *GetIntentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
