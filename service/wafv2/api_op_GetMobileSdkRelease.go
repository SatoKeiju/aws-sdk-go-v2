// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information for the specified mobile SDK release, including release
// notes and tags. The mobile SDK is not generally available. Customers who have
// access to the mobile SDK can use it to establish and manage WAF tokens for use
// in HTTP(S) requests from a mobile device to WAF. For more information, see WAF
// client application integration (https://docs.aws.amazon.com/waf/latest/developerguide/waf-application-integration.html)
// in the WAF Developer Guide.
func (c *Client) GetMobileSdkRelease(ctx context.Context, params *GetMobileSdkReleaseInput, optFns ...func(*Options)) (*GetMobileSdkReleaseOutput, error) {
	if params == nil {
		params = &GetMobileSdkReleaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMobileSdkRelease", params, optFns, c.addOperationGetMobileSdkReleaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMobileSdkReleaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMobileSdkReleaseInput struct {

	// The device platform.
	//
	// This member is required.
	Platform types.Platform

	// The release version. For the latest available version, specify LATEST .
	//
	// This member is required.
	ReleaseVersion *string

	noSmithyDocumentSerde
}

type GetMobileSdkReleaseOutput struct {

	// Information for a specified SDK release, including release notes and tags.
	MobileSdkRelease *types.MobileSdkRelease

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMobileSdkReleaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMobileSdkRelease{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMobileSdkRelease{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMobileSdkRelease"); err != nil {
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
	if err = addOpGetMobileSdkReleaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMobileSdkRelease(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMobileSdkRelease(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMobileSdkRelease",
	}
}
