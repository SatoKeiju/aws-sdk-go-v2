// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the configuration of the specified Amazon Web Services Verified Access
// instance.
func (c *Client) ModifyVerifiedAccessInstance(ctx context.Context, params *ModifyVerifiedAccessInstanceInput, optFns ...func(*Options)) (*ModifyVerifiedAccessInstanceOutput, error) {
	if params == nil {
		params = &ModifyVerifiedAccessInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVerifiedAccessInstance", params, optFns, c.addOperationModifyVerifiedAccessInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVerifiedAccessInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVerifiedAccessInstanceInput struct {

	// The ID of the Verified Access instance.
	//
	// This member is required.
	VerifiedAccessInstanceId *string

	// A unique, case-sensitive token that you provide to ensure idempotency of your
	// modification request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// A description for the Verified Access instance.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyVerifiedAccessInstanceOutput struct {

	// Details about the Verified Access instance.
	VerifiedAccessInstance *types.VerifiedAccessInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVerifiedAccessInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVerifiedAccessInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVerifiedAccessInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyVerifiedAccessInstance"); err != nil {
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
	if err = addIdempotencyToken_opModifyVerifiedAccessInstanceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpModifyVerifiedAccessInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVerifiedAccessInstance(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpModifyVerifiedAccessInstance struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpModifyVerifiedAccessInstance) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpModifyVerifiedAccessInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ModifyVerifiedAccessInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ModifyVerifiedAccessInstanceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opModifyVerifiedAccessInstanceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpModifyVerifiedAccessInstance{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opModifyVerifiedAccessInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyVerifiedAccessInstance",
	}
}
