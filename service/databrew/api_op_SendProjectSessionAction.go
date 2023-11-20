// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Performs a recipe step within an interactive DataBrew session that's currently
// open.
func (c *Client) SendProjectSessionAction(ctx context.Context, params *SendProjectSessionActionInput, optFns ...func(*Options)) (*SendProjectSessionActionOutput, error) {
	if params == nil {
		params = &SendProjectSessionActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendProjectSessionAction", params, optFns, c.addOperationSendProjectSessionActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendProjectSessionActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendProjectSessionActionInput struct {

	// The name of the project to apply the action to.
	//
	// This member is required.
	Name *string

	// A unique identifier for an interactive session that's currently open and ready
	// for work. The action will be performed on this session.
	ClientSessionId *string

	// If true, the result of the recipe step will be returned, but not applied.
	Preview bool

	// Represents a single step from a DataBrew recipe to be performed.
	RecipeStep *types.RecipeStep

	// The index from which to preview a step. This index is used to preview the
	// result of steps that have already been applied, so that the resulting view frame
	// is from earlier in the view frame stack.
	StepIndex *int32

	// Represents the data being transformed during an action.
	ViewFrame *types.ViewFrame

	noSmithyDocumentSerde
}

type SendProjectSessionActionOutput struct {

	// The name of the project that was affected by the action.
	//
	// This member is required.
	Name *string

	// A unique identifier for the action that was performed.
	ActionId *int32

	// A message indicating the result of performing the action.
	Result *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendProjectSessionActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendProjectSessionAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendProjectSessionAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendProjectSessionAction"); err != nil {
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
	if err = addOpSendProjectSessionActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendProjectSessionAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendProjectSessionAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendProjectSessionAction",
	}
}
