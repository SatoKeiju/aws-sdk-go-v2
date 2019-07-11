// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DescribeReservationRequest
type DescribeReservationInput struct {
	_ struct{} `type:"structure"`

	// ReservationId is a required field
	ReservationId *string `location:"uri" locationName:"reservationId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeReservationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeReservationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeReservationInput"}

	if s.ReservationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeReservationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ReservationId != nil {
		v := *s.ReservationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "reservationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DescribeReservationResponse
type DescribeReservationOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	Count *int64 `locationName:"count" type:"integer"`

	CurrencyCode *string `locationName:"currencyCode" type:"string"`

	Duration *int64 `locationName:"duration" type:"integer"`

	// Units for duration, e.g. 'MONTHS'
	DurationUnits OfferingDurationUnits `locationName:"durationUnits" type:"string" enum:"true"`

	End *string `locationName:"end" type:"string"`

	FixedPrice *float64 `locationName:"fixedPrice" type:"double"`

	Name *string `locationName:"name" type:"string"`

	OfferingDescription *string `locationName:"offeringDescription" type:"string"`

	OfferingId *string `locationName:"offeringId" type:"string"`

	// Offering type, e.g. 'NO_UPFRONT'
	OfferingType OfferingType `locationName:"offeringType" type:"string" enum:"true"`

	Region *string `locationName:"region" type:"string"`

	ReservationId *string `locationName:"reservationId" type:"string"`

	// Resource configuration (codec, resolution, bitrate, ...)
	ResourceSpecification *ReservationResourceSpecification `locationName:"resourceSpecification" type:"structure"`

	Start *string `locationName:"start" type:"string"`

	// Current reservation state
	State ReservationState `locationName:"state" type:"string" enum:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`

	UsagePrice *float64 `locationName:"usagePrice" type:"double"`
}

// String returns the string representation
func (s DescribeReservationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeReservationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Count != nil {
		v := *s.Count

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "count", protocol.Int64Value(v), metadata)
	}
	if s.CurrencyCode != nil {
		v := *s.CurrencyCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currencyCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Duration != nil {
		v := *s.Duration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "duration", protocol.Int64Value(v), metadata)
	}
	if len(s.DurationUnits) > 0 {
		v := s.DurationUnits

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "durationUnits", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.End != nil {
		v := *s.End

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "end", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FixedPrice != nil {
		v := *s.FixedPrice

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fixedPrice", protocol.Float64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OfferingDescription != nil {
		v := *s.OfferingDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "offeringDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OfferingId != nil {
		v := *s.OfferingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "offeringId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.OfferingType) > 0 {
		v := s.OfferingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "offeringType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Region != nil {
		v := *s.Region

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "region", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReservationId != nil {
		v := *s.ReservationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "reservationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceSpecification != nil {
		v := s.ResourceSpecification

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourceSpecification", v, metadata)
	}
	if s.Start != nil {
		v := *s.Start

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "start", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.UsagePrice != nil {
		v := *s.UsagePrice

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "usagePrice", protocol.Float64Value(v), metadata)
	}
	return nil
}

const opDescribeReservation = "DescribeReservation"

// DescribeReservationRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Get details for a reservation.
//
//    // Example sending a request using DescribeReservationRequest.
//    req := client.DescribeReservationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DescribeReservation
func (c *Client) DescribeReservationRequest(input *DescribeReservationInput) DescribeReservationRequest {
	op := &aws.Operation{
		Name:       opDescribeReservation,
		HTTPMethod: "GET",
		HTTPPath:   "/prod/reservations/{reservationId}",
	}

	if input == nil {
		input = &DescribeReservationInput{}
	}

	req := c.newRequest(op, input, &DescribeReservationOutput{})
	return DescribeReservationRequest{Request: req, Input: input, Copy: c.DescribeReservationRequest}
}

// DescribeReservationRequest is the request type for the
// DescribeReservation API operation.
type DescribeReservationRequest struct {
	*aws.Request
	Input *DescribeReservationInput
	Copy  func(*DescribeReservationInput) DescribeReservationRequest
}

// Send marshals and sends the DescribeReservation API request.
func (r DescribeReservationRequest) Send(ctx context.Context) (*DescribeReservationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservationResponse{
		DescribeReservationOutput: r.Request.Data.(*DescribeReservationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeReservationResponse is the response type for the
// DescribeReservation API operation.
type DescribeReservationResponse struct {
	*DescribeReservationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservation request.
func (r *DescribeReservationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
