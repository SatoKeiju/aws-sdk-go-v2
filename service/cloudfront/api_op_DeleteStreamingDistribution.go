// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a streaming distribution. To delete an RTMP distribution using the
// CloudFront API, perform the following steps. To delete an RTMP distribution
// using the CloudFront API:
//   - Disable the RTMP distribution.
//   - Submit a GET Streaming Distribution Config request to get the current
//     configuration and the Etag header for the distribution.
//   - Update the XML document that was returned in the response to your GET
//     Streaming Distribution Config request to change the value of Enabled to false
//     .
//   - Submit a PUT Streaming Distribution Config request to update the
//     configuration for your distribution. In the request body, include the XML
//     document that you updated in Step 3. Then set the value of the HTTP If-Match
//     header to the value of the ETag header that CloudFront returned when you
//     submitted the GET Streaming Distribution Config request in Step 2.
//   - Review the response to the PUT Streaming Distribution Config request to
//     confirm that the distribution was successfully disabled.
//   - Submit a GET Streaming Distribution Config request to confirm that your
//     changes have propagated. When propagation is complete, the value of Status is
//     Deployed .
//   - Submit a DELETE Streaming Distribution request. Set the value of the HTTP
//     If-Match header to the value of the ETag header that CloudFront returned when
//     you submitted the GET Streaming Distribution Config request in Step 2.
//   - Review the response to your DELETE Streaming Distribution request to confirm
//     that the distribution was successfully deleted.
//
// For information about deleting a distribution using the CloudFront console, see
// Deleting a Distribution (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/HowToDeleteDistribution.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) DeleteStreamingDistribution(ctx context.Context, params *DeleteStreamingDistributionInput, optFns ...func(*Options)) (*DeleteStreamingDistributionOutput, error) {
	if params == nil {
		params = &DeleteStreamingDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStreamingDistribution", params, optFns, c.addOperationDeleteStreamingDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStreamingDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to delete a streaming distribution.
type DeleteStreamingDistributionInput struct {

	// The distribution ID.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when you disabled the streaming
	// distribution. For example: E2QWRUHAPOMQZL .
	IfMatch *string

	noSmithyDocumentSerde
}

type DeleteStreamingDistributionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteStreamingDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteStreamingDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteStreamingDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteStreamingDistribution"); err != nil {
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
	if err = addOpDeleteStreamingDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStreamingDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteStreamingDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteStreamingDistribution",
	}
}
