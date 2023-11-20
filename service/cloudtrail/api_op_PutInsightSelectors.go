// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lets you enable Insights event logging by specifying the Insights selectors
// that you want to enable on an existing trail or event data store. You also use
// PutInsightSelectors to turn off Insights event logging, by passing an empty list
// of Insights types. The valid Insights event types are ApiErrorRateInsight and
// ApiCallRateInsight . To enable Insights on an event data store, you must specify
// the ARNs (or ID suffix of the ARNs) for the source event data store (
// EventDataStore ) and the destination event data store ( InsightsDestination ).
// The source event data store logs management events and enables Insights. The
// destination event data store logs Insights events based upon the management
// event activity of the source event data store. The source and destination event
// data stores must belong to the same Amazon Web Services account. To log Insights
// events for a trail, you must specify the name ( TrailName ) of the CloudTrail
// trail for which you want to change or add Insights selectors. To log CloudTrail
// Insights events on API call volume, the trail or event data store must log write
// management events. To log CloudTrail Insights events on API error rate, the
// trail or event data store must log read or write management events. You can
// call GetEventSelectors on a trail to check whether the trail logs management
// events. You can call GetEventDataStore on an event data store to check whether
// the event data store logs management events. For more information, see Logging
// CloudTrail Insights events (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html)
// in the CloudTrail User Guide.
func (c *Client) PutInsightSelectors(ctx context.Context, params *PutInsightSelectorsInput, optFns ...func(*Options)) (*PutInsightSelectorsOutput, error) {
	if params == nil {
		params = &PutInsightSelectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutInsightSelectors", params, optFns, c.addOperationPutInsightSelectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutInsightSelectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutInsightSelectorsInput struct {

	// A JSON string that contains the Insights types you want to log on a trail or
	// event data store. ApiCallRateInsight and ApiErrorRateInsight are valid Insight
	// types. The ApiCallRateInsight Insights type analyzes write-only management API
	// calls that are aggregated per minute against a baseline API call volume. The
	// ApiErrorRateInsight Insights type analyzes management API calls that result in
	// error codes. The error is shown if the API call is unsuccessful.
	//
	// This member is required.
	InsightSelectors []types.InsightSelector

	// The ARN (or ID suffix of the ARN) of the source event data store for which you
	// want to change or add Insights selectors. To enable Insights on an event data
	// store, you must provide both the EventDataStore and InsightsDestination
	// parameters. You cannot use this parameter with the TrailName parameter.
	EventDataStore *string

	// The ARN (or ID suffix of the ARN) of the destination event data store that logs
	// Insights events. To enable Insights on an event data store, you must provide
	// both the EventDataStore and InsightsDestination parameters. You cannot use this
	// parameter with the TrailName parameter.
	InsightsDestination *string

	// The name of the CloudTrail trail for which you want to change or add Insights
	// selectors. You cannot use this parameter with the EventDataStore and
	// InsightsDestination parameters.
	TrailName *string

	noSmithyDocumentSerde
}

type PutInsightSelectorsOutput struct {

	// The Amazon Resource Name (ARN) of the source event data store for which you
	// want to change or add Insights selectors.
	EventDataStoreArn *string

	// A JSON string that contains the Insights event types that you want to log on a
	// trail or event data store. The valid Insights types are ApiErrorRateInsight and
	// ApiCallRateInsight .
	InsightSelectors []types.InsightSelector

	// The ARN of the destination event data store that logs Insights events.
	InsightsDestination *string

	// The Amazon Resource Name (ARN) of a trail for which you want to change or add
	// Insights selectors.
	TrailARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutInsightSelectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutInsightSelectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutInsightSelectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutInsightSelectors"); err != nil {
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
	if err = addOpPutInsightSelectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutInsightSelectors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutInsightSelectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutInsightSelectors",
	}
}
