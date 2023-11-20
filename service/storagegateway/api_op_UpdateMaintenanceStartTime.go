// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a gateway's weekly maintenance start time information, including day
// and time of the week. The maintenance time is the time in your gateway's time
// zone.
func (c *Client) UpdateMaintenanceStartTime(ctx context.Context, params *UpdateMaintenanceStartTimeInput, optFns ...func(*Options)) (*UpdateMaintenanceStartTimeOutput, error) {
	if params == nil {
		params = &UpdateMaintenanceStartTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMaintenanceStartTime", params, optFns, c.addOperationUpdateMaintenanceStartTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMaintenanceStartTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the following fields:
//   - UpdateMaintenanceStartTimeInput$DayOfMonth
//   - UpdateMaintenanceStartTimeInput$DayOfWeek
//   - UpdateMaintenanceStartTimeInput$HourOfDay
//   - UpdateMaintenanceStartTimeInput$MinuteOfHour
type UpdateMaintenanceStartTimeInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// The hour component of the maintenance start time represented as hh, where hh is
	// the hour (00 to 23). The hour of the day is in the time zone of the gateway.
	//
	// This member is required.
	HourOfDay *int32

	// The minute component of the maintenance start time represented as mm, where mm
	// is the minute (00 to 59). The minute of the hour is in the time zone of the
	// gateway.
	//
	// This member is required.
	MinuteOfHour *int32

	// The day of the month component of the maintenance start time represented as an
	// ordinal number from 1 to 28, where 1 represents the first day of the month and
	// 28 represents the last day of the month.
	DayOfMonth *int32

	// The day of the week component of the maintenance start time week represented as
	// an ordinal number from 0 to 6, where 0 represents Sunday and 6 Saturday.
	DayOfWeek *int32

	noSmithyDocumentSerde
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway whose
// maintenance start time is updated.
type UpdateMaintenanceStartTimeOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMaintenanceStartTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMaintenanceStartTime"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addV4DetectSkewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateMaintenanceStartTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMaintenanceStartTime(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateMaintenanceStartTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMaintenanceStartTime",
	}
}
