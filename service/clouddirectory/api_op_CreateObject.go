// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateObjectRequest
type CreateObjectInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory in which
	// the object will be created. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The name of link that is used to attach this object to a parent.
	LinkName *string `min:"1" type:"string"`

	// The attribute map whose attribute ARN contains the key and attribute value
	// as the map value.
	ObjectAttributeList []AttributeKeyAndValue `type:"list"`

	// If specified, the parent reference to which this object will be attached.
	ParentReference *ObjectReference `type:"structure"`

	// A list of schema facets to be associated with the object. Do not provide
	// minor version components. See SchemaFacet for details.
	//
	// SchemaFacets is a required field
	SchemaFacets []SchemaFacet `type:"list" required:"true"`
}

// String returns the string representation
func (s CreateObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateObjectInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.LinkName != nil && len(*s.LinkName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LinkName", 1))
	}

	if s.SchemaFacets == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaFacets"))
	}
	if s.ObjectAttributeList != nil {
		for i, v := range s.ObjectAttributeList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ObjectAttributeList", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SchemaFacets != nil {
		for i, v := range s.SchemaFacets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SchemaFacets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.LinkName != nil {
		v := *s.LinkName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LinkName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ObjectAttributeList != nil {
		v := s.ObjectAttributeList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ObjectAttributeList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ParentReference != nil {
		v := s.ParentReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ParentReference", v, metadata)
	}
	if s.SchemaFacets != nil {
		v := s.SchemaFacets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SchemaFacets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateObjectResponse
type CreateObjectOutput struct {
	_ struct{} `type:"structure"`

	// The identifier that is associated with the object.
	ObjectIdentifier *string `type:"string"`
}

// String returns the string representation
func (s CreateObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ObjectIdentifier != nil {
		v := *s.ObjectIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ObjectIdentifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateObject = "CreateObject"

// CreateObjectRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Creates an object in a Directory. Additionally attaches the object to a parent,
// if a parent reference and LinkName is specified. An object is simply a collection
// of Facet attributes. You can also use this API call to create a policy object,
// if the facet from which you create the object is a policy facet.
//
//    // Example sending a request using CreateObjectRequest.
//    req := client.CreateObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateObject
func (c *Client) CreateObjectRequest(input *CreateObjectInput) CreateObjectRequest {
	op := &aws.Operation{
		Name:       opCreateObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object",
	}

	if input == nil {
		input = &CreateObjectInput{}
	}

	req := c.newRequest(op, input, &CreateObjectOutput{})
	return CreateObjectRequest{Request: req, Input: input, Copy: c.CreateObjectRequest}
}

// CreateObjectRequest is the request type for the
// CreateObject API operation.
type CreateObjectRequest struct {
	*aws.Request
	Input *CreateObjectInput
	Copy  func(*CreateObjectInput) CreateObjectRequest
}

// Send marshals and sends the CreateObject API request.
func (r CreateObjectRequest) Send(ctx context.Context) (*CreateObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateObjectResponse{
		CreateObjectOutput: r.Request.Data.(*CreateObjectOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateObjectResponse is the response type for the
// CreateObject API operation.
type CreateObjectResponse struct {
	*CreateObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateObject request.
func (r *CreateObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
