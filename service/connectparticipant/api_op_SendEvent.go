// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends an event. ConnectionToken is used for invoking this API instead of
// ParticipantToken . The Amazon Connect Participant Service APIs do not use
// Signature Version 4 authentication (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
// .
func (c *Client) SendEvent(ctx context.Context, params *SendEventInput, optFns ...func(*Options)) (*SendEventOutput, error) {
	if params == nil {
		params = &SendEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendEvent", params, optFns, c.addOperationSendEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendEventInput struct {

	// The authentication token associated with the participant's connection.
	//
	// This member is required.
	ConnectionToken *string

	// The content type of the request. Supported types are:
	//   - application/vnd.amazonaws.connect.event.typing
	//   - application/vnd.amazonaws.connect.event.connection.acknowledged
	//   - application/vnd.amazonaws.connect.event.message.delivered
	//   - application/vnd.amazonaws.connect.event.message.read
	//
	// This member is required.
	ContentType *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// The content of the event to be sent (for example, message text). For content
	// related to message receipts, this is supported in the form of a JSON string.
	// Sample Content: "{\"messageId\":\"11111111-aaaa-bbbb-cccc-EXAMPLE01234\"}"
	Content *string

	noSmithyDocumentSerde
}

type SendEventOutput struct {

	// The time when the event was sent. It's specified in ISO 8601 format:
	// yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	AbsoluteTime *string

	// The ID of the response.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendEvent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendEvent"); err != nil {
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
	if err = addIdempotencyToken_opSendEventMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendEvent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSendEvent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendEvent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendEventInput ")
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
func addIdempotencyToken_opSendEventMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendEvent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendEvent",
	}
}
