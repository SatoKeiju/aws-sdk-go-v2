// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an external destination to your SQL-based Kinesis Data Analytics
// application. If you want Kinesis Data Analytics to deliver data from an
// in-application stream within your application to an external destination (such
// as an Kinesis data stream, a Kinesis Data Firehose delivery stream, or an Amazon
// Lambda function), you add the relevant configuration to your application using
// this operation. You can configure one or more outputs for your application. Each
// output configuration maps an in-application stream and an external destination.
// You can use one of the output configurations to deliver data from your
// in-application error stream to an external destination so that you can analyze
// the errors. Any configuration update, including adding a streaming source using
// this operation, results in a new version of the application. You can use the
// DescribeApplication operation to find the current application version.
func (c *Client) AddApplicationOutput(ctx context.Context, params *AddApplicationOutputInput, optFns ...func(*Options)) (*AddApplicationOutputOutput, error) {
	if params == nil {
		params = &AddApplicationOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationOutput", params, optFns, c.addOperationAddApplicationOutputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationOutputInput struct {

	// The name of the application to which you want to add the output configuration.
	//
	// This member is required.
	ApplicationName *string

	// The version of the application to which you want to add the output
	// configuration. You can use the DescribeApplication operation to get the current
	// application version. If the version specified is not the current version, the
	// ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// An array of objects, each describing one output configuration. In the output
	// configuration, you specify the name of an in-application stream, a destination
	// (that is, a Kinesis data stream, a Kinesis Data Firehose delivery stream, or an
	// Amazon Lambda function), and record the formation to use when writing to the
	// destination.
	//
	// This member is required.
	Output *types.Output

	noSmithyDocumentSerde
}

type AddApplicationOutputOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string

	// The updated application version ID. Kinesis Data Analytics increments this ID
	// when the application is updated.
	ApplicationVersionId *int64

	// Describes the application output configuration. For more information, see
	// Configuring Application Output (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html)
	// .
	OutputDescriptions []types.OutputDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddApplicationOutputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationOutput{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddApplicationOutput"); err != nil {
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
	if err = addOpAddApplicationOutputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationOutput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationOutput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddApplicationOutput",
	}
}
