// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Amazon Lightsail bucket. Use this action to update the
// configuration of an existing bucket, such as versioning, public accessibility,
// and the Amazon Web Services accounts that can access the bucket.
func (c *Client) UpdateBucket(ctx context.Context, params *UpdateBucketInput, optFns ...func(*Options)) (*UpdateBucketOutput, error) {
	if params == nil {
		params = &UpdateBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBucket", params, optFns, c.addOperationUpdateBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBucketInput struct {

	// The name of the bucket to update.
	//
	// This member is required.
	BucketName *string

	// An object that describes the access log configuration for the bucket.
	AccessLogConfig *types.BucketAccessLogConfig

	// An object that sets the public accessibility of objects in the specified bucket.
	AccessRules *types.AccessRules

	// An array of strings to specify the Amazon Web Services account IDs that can
	// access the bucket. You can give a maximum of 10 Amazon Web Services accounts
	// access to a bucket.
	ReadonlyAccessAccounts []string

	// Specifies whether to enable or suspend versioning of objects in the bucket. The
	// following options can be specified:
	//   - Enabled - Enables versioning of objects in the specified bucket.
	//   - Suspended - Suspends versioning of objects in the specified bucket. Existing
	//   object versions are retained.
	Versioning *string

	noSmithyDocumentSerde
}

type UpdateBucketOutput struct {

	// An object that describes the bucket that is updated.
	Bucket *types.Bucket

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateBucket"); err != nil {
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
	if err = addOpUpdateBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBucket(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateBucket",
	}
}
