// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a cache policy configuration. To get a cache policy configuration, you
// must provide the policy's identifier. If the cache policy is attached to a
// distribution's cache behavior, you can get the policy's identifier using
// ListDistributions or GetDistribution . If the cache policy is not attached to a
// cache behavior, you can get the identifier using ListCachePolicies .
func (c *Client) GetCachePolicyConfig(ctx context.Context, params *GetCachePolicyConfigInput, optFns ...func(*Options)) (*GetCachePolicyConfigOutput, error) {
	if params == nil {
		params = &GetCachePolicyConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCachePolicyConfig", params, optFns, c.addOperationGetCachePolicyConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCachePolicyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCachePolicyConfigInput struct {

	// The unique identifier for the cache policy. If the cache policy is attached to
	// a distribution's cache behavior, you can get the policy's identifier using
	// ListDistributions or GetDistribution . If the cache policy is not attached to a
	// cache behavior, you can get the identifier using ListCachePolicies .
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetCachePolicyConfigOutput struct {

	// The cache policy configuration.
	CachePolicyConfig *types.CachePolicyConfig

	// The current version of the cache policy.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCachePolicyConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetCachePolicyConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetCachePolicyConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCachePolicyConfig"); err != nil {
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
	if err = addOpGetCachePolicyConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCachePolicyConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCachePolicyConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCachePolicyConfig",
	}
}
