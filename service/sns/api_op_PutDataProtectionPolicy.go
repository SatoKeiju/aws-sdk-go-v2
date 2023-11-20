// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates an inline policy document that is stored in the specified
// Amazon SNS topic.
func (c *Client) PutDataProtectionPolicy(ctx context.Context, params *PutDataProtectionPolicyInput, optFns ...func(*Options)) (*PutDataProtectionPolicyOutput, error) {
	if params == nil {
		params = &PutDataProtectionPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDataProtectionPolicy", params, optFns, c.addOperationPutDataProtectionPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDataProtectionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDataProtectionPolicyInput struct {

	// The JSON serialization of the topic's DataProtectionPolicy . The
	// DataProtectionPolicy must be in JSON string format. Length Constraints: Maximum
	// length of 30,720.
	//
	// This member is required.
	DataProtectionPolicy *string

	// The ARN of the topic whose DataProtectionPolicy you want to add or update. For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type PutDataProtectionPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDataProtectionPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutDataProtectionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutDataProtectionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutDataProtectionPolicy"); err != nil {
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
	if err = addOpPutDataProtectionPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDataProtectionPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDataProtectionPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutDataProtectionPolicy",
	}
}
