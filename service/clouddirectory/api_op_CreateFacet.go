// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateFacetRequest
type CreateFacetInput struct {
	_ struct{} `type:"structure"`

	// The attributes that are associated with the Facet.
	Attributes []FacetAttribute `type:"list"`

	// There are two different styles that you can define on any given facet, Static
	// and Dynamic. For static facets, all attributes must be defined in the schema.
	// For dynamic facets, attributes can be defined during data plane operations.
	FacetStyle FacetStyle `type:"string" enum:"true"`

	// The name of the Facet, which is unique for a given schema.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Specifies whether a given object created from this facet is of type node,
	// leaf node, policy or index.
	//
	//    * Node: Can have multiple children but one parent.
	//
	//    * Leaf node: Cannot have children but can have multiple parents.
	//
	//    * Policy: Allows you to store a policy document and policy type. For more
	//    information, see Policies (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies).
	//
	//    * Index: Can be created with the Index API.
	ObjectType ObjectType `type:"string" enum:"true"`

	// The schema ARN in which the new Facet will be created. For more information,
	// see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateFacetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFacetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFacetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFacetInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attributes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.FacetStyle) > 0 {
		v := s.FacetStyle

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FacetStyle", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ObjectType) > 0 {
		v := s.ObjectType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ObjectType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateFacetResponse
type CreateFacetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateFacetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFacetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateFacet = "CreateFacet"

// CreateFacetRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Creates a new Facet in a schema. Facet creation is allowed only in development
// or applied schemas.
//
//    // Example sending a request using CreateFacetRequest.
//    req := client.CreateFacetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateFacet
func (c *Client) CreateFacetRequest(input *CreateFacetInput) CreateFacetRequest {
	op := &aws.Operation{
		Name:       opCreateFacet,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/facet/create",
	}

	if input == nil {
		input = &CreateFacetInput{}
	}

	req := c.newRequest(op, input, &CreateFacetOutput{})
	return CreateFacetRequest{Request: req, Input: input, Copy: c.CreateFacetRequest}
}

// CreateFacetRequest is the request type for the
// CreateFacet API operation.
type CreateFacetRequest struct {
	*aws.Request
	Input *CreateFacetInput
	Copy  func(*CreateFacetInput) CreateFacetRequest
}

// Send marshals and sends the CreateFacet API request.
func (r CreateFacetRequest) Send(ctx context.Context) (*CreateFacetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFacetResponse{
		CreateFacetOutput: r.Request.Data.(*CreateFacetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFacetResponse is the response type for the
// CreateFacet API operation.
type CreateFacetResponse struct {
	*CreateFacetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFacet request.
func (r *CreateFacetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
