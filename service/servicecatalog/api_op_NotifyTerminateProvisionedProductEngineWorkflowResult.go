// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Notifies the result of the terminate engine execution.
func (c *Client) NotifyTerminateProvisionedProductEngineWorkflowResult(ctx context.Context, params *NotifyTerminateProvisionedProductEngineWorkflowResultInput, optFns ...func(*Options)) (*NotifyTerminateProvisionedProductEngineWorkflowResultOutput, error) {
	if params == nil {
		params = &NotifyTerminateProvisionedProductEngineWorkflowResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NotifyTerminateProvisionedProductEngineWorkflowResult", params, optFns, c.addOperationNotifyTerminateProvisionedProductEngineWorkflowResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NotifyTerminateProvisionedProductEngineWorkflowResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyTerminateProvisionedProductEngineWorkflowResultInput struct {

	// The idempotency token that identifies the terminate engine execution.
	//
	// This member is required.
	IdempotencyToken *string

	// The identifier of the record.
	//
	// This member is required.
	RecordId *string

	// The status of the terminate engine execution.
	//
	// This member is required.
	Status types.EngineWorkflowStatus

	// The encrypted contents of the terminate engine execution payload that Service
	// Catalog sends after the Terraform product terminate workflow starts.
	//
	// This member is required.
	WorkflowToken *string

	// The reason why the terminate engine execution failed.
	FailureReason *string

	noSmithyDocumentSerde
}

type NotifyTerminateProvisionedProductEngineWorkflowResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNotifyTerminateProvisionedProductEngineWorkflowResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpNotifyTerminateProvisionedProductEngineWorkflowResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpNotifyTerminateProvisionedProductEngineWorkflowResult{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "NotifyTerminateProvisionedProductEngineWorkflowResult"); err != nil {
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
	if err = addIdempotencyToken_opNotifyTerminateProvisionedProductEngineWorkflowResultMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpNotifyTerminateProvisionedProductEngineWorkflowResultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyTerminateProvisionedProductEngineWorkflowResult(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpNotifyTerminateProvisionedProductEngineWorkflowResult struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpNotifyTerminateProvisionedProductEngineWorkflowResult) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpNotifyTerminateProvisionedProductEngineWorkflowResult) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*NotifyTerminateProvisionedProductEngineWorkflowResultInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *NotifyTerminateProvisionedProductEngineWorkflowResultInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opNotifyTerminateProvisionedProductEngineWorkflowResultMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpNotifyTerminateProvisionedProductEngineWorkflowResult{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opNotifyTerminateProvisionedProductEngineWorkflowResult(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NotifyTerminateProvisionedProductEngineWorkflowResult",
	}
}
