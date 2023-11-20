// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays details about an event bus in your account. This can include the
// external Amazon Web Services accounts that are permitted to write events to your
// default event bus, and the associated policy. For custom event buses and partner
// event buses, it displays the name, ARN, policy, state, and creation time. To
// enable your account to receive events from other accounts on its default event
// bus, use PutPermission (https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html)
// . For more information about partner event buses, see CreateEventBus (https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateEventBus.html)
// .
func (c *Client) DescribeEventBus(ctx context.Context, params *DescribeEventBusInput, optFns ...func(*Options)) (*DescribeEventBusOutput, error) {
	if params == nil {
		params = &DescribeEventBusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventBus", params, optFns, c.addOperationDescribeEventBusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventBusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventBusInput struct {

	// The name or ARN of the event bus to show details for. If you omit this, the
	// default event bus is displayed.
	Name *string

	noSmithyDocumentSerde
}

type DescribeEventBusOutput struct {

	// The Amazon Resource Name (ARN) of the account permitted to write events to the
	// current account.
	Arn *string

	// The name of the event bus. Currently, this is always default .
	Name *string

	// The policy that enables the external account to send events to your account.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventBusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventBus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventBus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEventBus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventBus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEventBus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEventBus",
	}
}
