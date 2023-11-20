// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The SendTestEventNotification operation causes Amazon Mechanical Turk to send a
// notification message as if a HIT event occurred, according to the provided
// notification specification. This allows you to test notifications without
// setting up notifications for a real HIT type and trying to trigger them using
// the website. When you call this operation, the service attempts to send the test
// notification immediately.
func (c *Client) SendTestEventNotification(ctx context.Context, params *SendTestEventNotificationInput, optFns ...func(*Options)) (*SendTestEventNotificationOutput, error) {
	if params == nil {
		params = &SendTestEventNotificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendTestEventNotification", params, optFns, c.addOperationSendTestEventNotificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendTestEventNotificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendTestEventNotificationInput struct {

	// The notification specification to test. This value is identical to the value
	// you would provide to the UpdateNotificationSettings operation when you establish
	// the notification specification for a HIT type.
	//
	// This member is required.
	Notification *types.NotificationSpecification

	// The event to simulate to test the notification specification. This event is
	// included in the test message even if the notification specification does not
	// include the event type. The notification specification does not filter out the
	// test event.
	//
	// This member is required.
	TestEventType types.EventType

	noSmithyDocumentSerde
}

type SendTestEventNotificationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendTestEventNotificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendTestEventNotification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendTestEventNotification{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendTestEventNotification"); err != nil {
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
	if err = addOpSendTestEventNotificationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendTestEventNotification(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendTestEventNotification(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendTestEventNotification",
	}
}
