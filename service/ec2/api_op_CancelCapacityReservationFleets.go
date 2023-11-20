// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels one or more Capacity Reservation Fleets. When you cancel a Capacity
// Reservation Fleet, the following happens:
//   - The Capacity Reservation Fleet's status changes to cancelled .
//   - The individual Capacity Reservations in the Fleet are cancelled. Instances
//     running in the Capacity Reservations at the time of cancelling the Fleet
//     continue to run in shared capacity.
//   - The Fleet stops creating new Capacity Reservations.
func (c *Client) CancelCapacityReservationFleets(ctx context.Context, params *CancelCapacityReservationFleetsInput, optFns ...func(*Options)) (*CancelCapacityReservationFleetsOutput, error) {
	if params == nil {
		params = &CancelCapacityReservationFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelCapacityReservationFleets", params, optFns, c.addOperationCancelCapacityReservationFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelCapacityReservationFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelCapacityReservationFleetsInput struct {

	// The IDs of the Capacity Reservation Fleets to cancel.
	//
	// This member is required.
	CapacityReservationFleetIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type CancelCapacityReservationFleetsOutput struct {

	// Information about the Capacity Reservation Fleets that could not be cancelled.
	FailedFleetCancellations []types.FailedCapacityReservationFleetCancellationResult

	// Information about the Capacity Reservation Fleets that were successfully
	// cancelled.
	SuccessfulFleetCancellations []types.CapacityReservationFleetCancellationState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelCapacityReservationFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCancelCapacityReservationFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCancelCapacityReservationFleets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelCapacityReservationFleets"); err != nil {
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
	if err = addOpCancelCapacityReservationFleetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelCapacityReservationFleets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelCapacityReservationFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelCapacityReservationFleets",
	}
}
