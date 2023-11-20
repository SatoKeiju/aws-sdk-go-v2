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

// Creates a staging distribution using the configuration of the provided primary
// distribution. A staging distribution is a copy of an existing distribution
// (called the primary distribution) that you can use in a continuous deployment
// workflow. After you create a staging distribution, you can use
// UpdateDistribution to modify the staging distribution's configuration. Then you
// can use CreateContinuousDeploymentPolicy to incrementally move traffic to the
// staging distribution. This API operation requires the following IAM permissions:
//
//   - GetDistribution (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistribution.html)
//   - CreateDistribution (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html)
//   - CopyDistribution (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CopyDistribution.html)
func (c *Client) CopyDistribution(ctx context.Context, params *CopyDistributionInput, optFns ...func(*Options)) (*CopyDistributionOutput, error) {
	if params == nil {
		params = &CopyDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDistribution", params, optFns, c.addOperationCopyDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyDistributionInput struct {

	// A value that uniquely identifies a request to create a resource. This helps to
	// prevent CloudFront from creating a duplicate resource if you accidentally
	// resubmit an identical request.
	//
	// This member is required.
	CallerReference *string

	// The identifier of the primary distribution whose configuration you are copying.
	// To get a distribution ID, use ListDistributions .
	//
	// This member is required.
	PrimaryDistributionId *string

	// A Boolean flag to specify the state of the staging distribution when it's
	// created. When you set this value to True , the staging distribution is enabled.
	// When you set this value to False , the staging distribution is disabled. If you
	// omit this field, the default value is True .
	Enabled *bool

	// The version identifier of the primary distribution whose configuration you are
	// copying. This is the ETag value returned in the response to GetDistribution and
	// GetDistributionConfig .
	IfMatch *string

	// The type of distribution that your primary distribution will be copied to. The
	// only valid value is True , indicating that you are copying to a staging
	// distribution.
	Staging *bool

	noSmithyDocumentSerde
}

type CopyDistributionOutput struct {

	// A distribution tells CloudFront where you want content to be delivered from,
	// and the details about how to track and manage content delivery.
	Distribution *types.Distribution

	// The version identifier for the current version of the staging distribution.
	ETag *string

	// The URL of the staging distribution.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCopyDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCopyDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyDistribution"); err != nil {
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
	if err = addOpCopyDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyDistribution",
	}
}
