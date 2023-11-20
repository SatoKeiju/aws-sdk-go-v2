// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a delivery. A delivery is a connection between a logical delivery
// source and a logical delivery destination that you have already created. Only
// some Amazon Web Services services support being configured as a delivery source
// using this operation. These services are listed as Supported [V2 Permissions] in
// the table at Enabling logging from Amazon Web Services services. (https://docs.aws.amazon.com/     AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html#AWS-vended-logs-permissions)
// A delivery destination can represent a log group in CloudWatch Logs, an Amazon
// S3 bucket, or a delivery stream in Kinesis Data Firehose. To configure logs
// delivery between a supported Amazon Web Services service and a destination, you
// must do the following:
//   - Create a delivery source, which is a logical object that represents the
//     resource that is actually sending the logs. For more information, see
//     PutDeliverySource (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html)
//     .
//   - Create a delivery destination, which is a logical object that represents
//     the actual delivery destination. For more information, see
//     PutDeliveryDestination (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html)
//     .
//   - If you are delivering logs cross-account, you must use
//     PutDeliveryDestinationPolicy (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationolicy.html)
//     in the destination account to assign an IAM policy to the destination. This
//     policy allows delivery to that destination.
//   - Use CreateDelivery to create a delivery by pairing exactly one delivery
//     source and one delivery destination.
//
// You can configure a single delivery source to send logs to multiple
// destinations by creating multiple deliveries. You can also create multiple
// deliveries to configure multiple delivery sources to send logs to the same
// delivery destination. You can't update an existing delivery. You can only create
// and delete deliveries.
func (c *Client) CreateDelivery(ctx context.Context, params *CreateDeliveryInput, optFns ...func(*Options)) (*CreateDeliveryOutput, error) {
	if params == nil {
		params = &CreateDeliveryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDelivery", params, optFns, c.addOperationCreateDeliveryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeliveryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeliveryInput struct {

	// The ARN of the delivery destination to use for this delivery.
	//
	// This member is required.
	DeliveryDestinationArn *string

	// The name of the delivery source to use for this delivery.
	//
	// This member is required.
	DeliverySourceName *string

	// An optional list of key-value pairs to associate with the resource. For more
	// information about tagging, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDeliveryOutput struct {

	// A structure that contains information about the delivery that you just created.
	Delivery *types.Delivery

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeliveryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDelivery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDelivery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDelivery"); err != nil {
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
	if err = addOpCreateDeliveryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDelivery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDelivery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDelivery",
	}
}
