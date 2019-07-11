// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The PATCH request to grant a temporary extension to the remaining quota of
// a usage plan associated with a specified API key.
type UpdateUsageInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the API key associated with the usage plan in
	// which a temporary extension is granted to the remaining quota.
	//
	// KeyId is a required field
	KeyId *string `location:"uri" locationName:"keyId" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`

	// [Required] The Id of the usage plan associated with the usage data.
	//
	// UsagePlanId is a required field
	UsagePlanId *string `location:"uri" locationName:"usageplanId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUsageInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}

	if s.UsagePlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsagePlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUsageInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.PatchOperations != nil {
		v := s.PatchOperations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "patchOperations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.KeyId != nil {
		v := *s.KeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "keyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UsagePlanId != nil {
		v := *s.UsagePlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "usageplanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents the usage data of a usage plan.
//
// Create and Use Usage Plans (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html),
// Manage Usage in a Usage Plan (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-usage-plans-with-console.html#api-gateway-usage-plan-manage-usage)
type UpdateUsageOutput struct {
	_ struct{} `type:"structure"`

	// The ending date of the usage data.
	EndDate *string `locationName:"endDate" type:"string"`

	// The usage data, as daily logs of used and remaining quotas, over the specified
	// time interval indexed over the API keys in a usage plan. For example, {...,
	// "values" : { "{api_key}" : [ [0, 100], [10, 90], [100, 10]]}, where {api_key}
	// stands for an API key value and the daily log entry is of the format [used
	// quota, remaining quota].
	Items map[string][][]int64 `locationName:"values" type:"map"`

	Position *string `locationName:"position" type:"string"`

	// The starting date of the usage data.
	StartDate *string `locationName:"startDate" type:"string"`

	// The plan Id associated with this usage data.
	UsagePlanId *string `locationName:"usagePlanId" type:"string"`
}

// String returns the string representation
func (s UpdateUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUsageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EndDate != nil {
		v := *s.EndDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endDate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "values", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ls1 := ms0.List(k1)
			ls1.Start()
			for _, v2 := range v1 {
				ls2 := ls1.List()
				ls2.Start()
				for _, v3 := range v2 {
					ls2.ListAddValue(protocol.Int64Value(v3))
				}
				ls2.End()
			}
			ls1.End()
		}
		ms0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartDate != nil {
		v := *s.StartDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startDate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UsagePlanId != nil {
		v := *s.UsagePlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "usagePlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateUsage = "UpdateUsage"

// UpdateUsageRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Grants a temporary extension to the remaining quota of a usage plan associated
// with a specified API key.
//
//    // Example sending a request using UpdateUsageRequest.
//    req := client.UpdateUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateUsageRequest(input *UpdateUsageInput) UpdateUsageRequest {
	op := &aws.Operation{
		Name:       opUpdateUsage,
		HTTPMethod: "PATCH",
		HTTPPath:   "/usageplans/{usageplanId}/keys/{keyId}/usage",
	}

	if input == nil {
		input = &UpdateUsageInput{}
	}

	req := c.newRequest(op, input, &UpdateUsageOutput{})
	return UpdateUsageRequest{Request: req, Input: input, Copy: c.UpdateUsageRequest}
}

// UpdateUsageRequest is the request type for the
// UpdateUsage API operation.
type UpdateUsageRequest struct {
	*aws.Request
	Input *UpdateUsageInput
	Copy  func(*UpdateUsageInput) UpdateUsageRequest
}

// Send marshals and sends the UpdateUsage API request.
func (r UpdateUsageRequest) Send(ctx context.Context) (*UpdateUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUsageResponse{
		UpdateUsageOutput: r.Request.Data.(*UpdateUsageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUsageResponse is the response type for the
// UpdateUsage API operation.
type UpdateUsageResponse struct {
	*UpdateUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUsage request.
func (r *UpdateUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
