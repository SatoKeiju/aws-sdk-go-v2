// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified origination identity from an existing pool. If the
// origination identity isn't associated with the specified pool, an error is
// returned.
func (c *Client) DisassociateOriginationIdentity(ctx context.Context, params *DisassociateOriginationIdentityInput, optFns ...func(*Options)) (*DisassociateOriginationIdentityOutput, error) {
	if params == nil {
		params = &DisassociateOriginationIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateOriginationIdentity", params, optFns, c.addOperationDisassociateOriginationIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateOriginationIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateOriginationIdentityInput struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// The origination identity to use such as a PhoneNumberId, PhoneNumberArn,
	// SenderId or SenderIdArn. You can use DescribePhoneNumbers find the values for
	// PhoneNumberId and PhoneNumberArn, or use DescribeSenderIds to get the values
	// for SenderId and SenderIdArn.
	//
	// This member is required.
	OriginationIdentity *string

	// The unique identifier for the pool to disassociate with the origination
	// identity. This value can be either the PoolId or PoolArn.
	//
	// This member is required.
	PoolId *string

	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. If you don't specify a client token, a randomly generated token is used
	// for the request to ensure idempotency.
	ClientToken *string

	noSmithyDocumentSerde
}

type DisassociateOriginationIdentityOutput struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	IsoCountryCode *string

	// The PhoneNumberId or SenderId of the origination identity.
	OriginationIdentity *string

	// The PhoneNumberArn or SenderIdArn of the origination identity.
	OriginationIdentityArn *string

	// The Amazon Resource Name (ARN) of the pool.
	PoolArn *string

	// The PoolId of the pool no longer associated with the origination identity.
	PoolId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateOriginationIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDisassociateOriginationIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDisassociateOriginationIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateOriginationIdentity"); err != nil {
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
	if err = addIdempotencyToken_opDisassociateOriginationIdentityMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDisassociateOriginationIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateOriginationIdentity(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDisassociateOriginationIdentity struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDisassociateOriginationIdentity) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDisassociateOriginationIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DisassociateOriginationIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DisassociateOriginationIdentityInput ")
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
func addIdempotencyToken_opDisassociateOriginationIdentityMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDisassociateOriginationIdentity{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDisassociateOriginationIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateOriginationIdentity",
	}
}
