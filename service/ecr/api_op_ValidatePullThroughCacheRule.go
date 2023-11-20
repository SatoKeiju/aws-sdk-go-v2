// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates an existing pull through cache rule for an upstream registry that
// requires authentication. This will retrieve the contents of the Amazon Web
// Services Secrets Manager secret, verify the syntax, and then validate that
// authentication to the upstream registry is successful.
func (c *Client) ValidatePullThroughCacheRule(ctx context.Context, params *ValidatePullThroughCacheRuleInput, optFns ...func(*Options)) (*ValidatePullThroughCacheRuleOutput, error) {
	if params == nil {
		params = &ValidatePullThroughCacheRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidatePullThroughCacheRule", params, optFns, c.addOperationValidatePullThroughCacheRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidatePullThroughCacheRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidatePullThroughCacheRuleInput struct {

	// The repository name prefix associated with the pull through cache rule.
	//
	// This member is required.
	EcrRepositoryPrefix *string

	// The registry ID associated with the pull through cache rule. If you do not
	// specify a registry, the default registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type ValidatePullThroughCacheRuleOutput struct {

	// The Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager
	// secret associated with the pull through cache rule.
	CredentialArn *string

	// The Amazon ECR repository prefix associated with the pull through cache rule.
	EcrRepositoryPrefix *string

	// The reason the validation failed. For more details about possible causes and
	// how to address them, see Using pull through cache rules (https://docs.aws.amazon.com/AmazonECR/latest/userguide/pull-through-cache.html)
	// in the Amazon Elastic Container Registry User Guide.
	Failure *string

	// Whether or not the pull through cache rule was validated. If true , Amazon ECR
	// was able to reach the upstream registry and authentication was successful. If
	// false , there was an issue and validation failed. The failure reason indicates
	// the cause.
	IsValid bool

	// The registry ID associated with the request.
	RegistryId *string

	// The upstream registry URL associated with the pull through cache rule.
	UpstreamRegistryUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidatePullThroughCacheRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpValidatePullThroughCacheRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpValidatePullThroughCacheRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ValidatePullThroughCacheRule"); err != nil {
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
	if err = addOpValidatePullThroughCacheRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidatePullThroughCacheRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidatePullThroughCacheRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ValidatePullThroughCacheRule",
	}
}
