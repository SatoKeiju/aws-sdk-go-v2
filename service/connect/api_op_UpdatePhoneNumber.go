// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates your claimed phone number from its current Amazon Connect instance or
// traffic distribution group to another Amazon Connect instance or traffic
// distribution group in the same Amazon Web Services Region. After using this API,
// you must verify that the phone number is attached to the correct flow in the
// target instance or traffic distribution group. You need to do this because the
// API switches only the phone number to a new instance or traffic distribution
// group. It doesn't migrate the flow configuration of the phone number, too. You
// can call DescribePhoneNumber (https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html)
// API to verify the status of a previous UpdatePhoneNumber (https://docs.aws.amazon.com/connect/latest/APIReference/API_UpdatePhoneNumber.html)
// operation.
func (c *Client) UpdatePhoneNumber(ctx context.Context, params *UpdatePhoneNumberInput, optFns ...func(*Options)) (*UpdatePhoneNumberOutput, error) {
	if params == nil {
		params = &UpdatePhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePhoneNumber", params, optFns, c.addOperationUpdatePhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePhoneNumberInput struct {

	// A unique identifier for the phone number.
	//
	// This member is required.
	PhoneNumberId *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// The identifier of the Amazon Connect instance that phone numbers are claimed
	// to. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance. You must enter InstanceId or
	// TargetArn .
	InstanceId *string

	// The Amazon Resource Name (ARN) for Amazon Connect instances or traffic
	// distribution groups that phone number inbound traffic is routed through. You
	// must enter InstanceId or TargetArn .
	TargetArn *string

	noSmithyDocumentSerde
}

type UpdatePhoneNumberOutput struct {

	// The Amazon Resource Name (ARN) of the phone number.
	PhoneNumberArn *string

	// A unique identifier for the phone number.
	PhoneNumberId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePhoneNumber"); err != nil {
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
	if err = addIdempotencyToken_opUpdatePhoneNumberMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdatePhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePhoneNumber(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdatePhoneNumber struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdatePhoneNumber) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdatePhoneNumber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdatePhoneNumberInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdatePhoneNumberInput ")
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
func addIdempotencyToken_opUpdatePhoneNumberMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdatePhoneNumber{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdatePhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePhoneNumber",
	}
}
