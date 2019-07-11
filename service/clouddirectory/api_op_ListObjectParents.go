// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectParentsRequest
type ListObjectParentsInput struct {
	_ struct{} `type:"structure"`

	// Represents the manner and timing in which the successful write or update
	// of an object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// the object resides. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// When set to True, returns all ListObjectParentsResponse$ParentLinks. There
	// could be multiple links between a parent-child pair.
	IncludeAllLinksToEachParent *bool `type:"boolean"`

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The reference that identifies the object for which parent objects are being
	// listed.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListObjectParentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListObjectParentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListObjectParentsInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectParentsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IncludeAllLinksToEachParent != nil {
		v := *s.IncludeAllLinksToEachParent

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IncludeAllLinksToEachParent", protocol.BoolValue(v), metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-consistency-level", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectParentsResponse
type ListObjectParentsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token.
	NextToken *string `type:"string"`

	// Returns a list of parent reference and LinkName Tuples.
	ParentLinks []ObjectIdentifierAndLinkNameTuple `type:"list"`

	// The parent structure, which is a map with key as the ObjectIdentifier and
	// LinkName as the value.
	Parents map[string]string `type:"map"`
}

// String returns the string representation
func (s ListObjectParentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectParentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParentLinks != nil {
		v := s.ParentLinks

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ParentLinks", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Parents != nil {
		v := s.Parents

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Parents", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opListObjectParents = "ListObjectParents"

// ListObjectParentsRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists parent objects that are associated with a given object in pagination
// fashion.
//
//    // Example sending a request using ListObjectParentsRequest.
//    req := client.ListObjectParentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectParents
func (c *Client) ListObjectParentsRequest(input *ListObjectParentsInput) ListObjectParentsRequest {
	op := &aws.Operation{
		Name:       opListObjectParents,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/parent",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListObjectParentsInput{}
	}

	req := c.newRequest(op, input, &ListObjectParentsOutput{})
	return ListObjectParentsRequest{Request: req, Input: input, Copy: c.ListObjectParentsRequest}
}

// ListObjectParentsRequest is the request type for the
// ListObjectParents API operation.
type ListObjectParentsRequest struct {
	*aws.Request
	Input *ListObjectParentsInput
	Copy  func(*ListObjectParentsInput) ListObjectParentsRequest
}

// Send marshals and sends the ListObjectParents API request.
func (r ListObjectParentsRequest) Send(ctx context.Context) (*ListObjectParentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectParentsResponse{
		ListObjectParentsOutput: r.Request.Data.(*ListObjectParentsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectParentsRequestPaginator returns a paginator for ListObjectParents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectParentsRequest(input)
//   p := clouddirectory.NewListObjectParentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectParentsPaginator(req ListObjectParentsRequest) ListObjectParentsPaginator {
	return ListObjectParentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListObjectParentsInput
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

// ListObjectParentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectParentsPaginator struct {
	aws.Pager
}

func (p *ListObjectParentsPaginator) CurrentPage() *ListObjectParentsOutput {
	return p.Pager.CurrentPage().(*ListObjectParentsOutput)
}

// ListObjectParentsResponse is the response type for the
// ListObjectParents API operation.
type ListObjectParentsResponse struct {
	*ListObjectParentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectParents request.
func (r *ListObjectParentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
