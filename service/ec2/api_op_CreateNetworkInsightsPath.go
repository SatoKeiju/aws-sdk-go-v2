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

// Creates a path to analyze for reachability. Reachability Analyzer enables you
// to analyze and debug network reachability between two resources in your virtual
// private cloud (VPC). For more information, see the Reachability Analyzer Guide (https://docs.aws.amazon.com/vpc/latest/reachability/)
// .
func (c *Client) CreateNetworkInsightsPath(ctx context.Context, params *CreateNetworkInsightsPathInput, optFns ...func(*Options)) (*CreateNetworkInsightsPathOutput, error) {
	if params == nil {
		params = &CreateNetworkInsightsPathInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNetworkInsightsPath", params, optFns, c.addOperationCreateNetworkInsightsPathMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNetworkInsightsPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNetworkInsightsPathInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	//
	// This member is required.
	ClientToken *string

	// The protocol.
	//
	// This member is required.
	Protocol types.Protocol

	// The ID or ARN of the source. If the resource is in another account, you must
	// specify an ARN.
	//
	// This member is required.
	Source *string

	// The ID or ARN of the destination. If the resource is in another account, you
	// must specify an ARN.
	Destination *string

	// The IP address of the destination.
	DestinationIp *string

	// The destination port.
	DestinationPort *int32

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Scopes the analysis to network paths that match specific filters at the
	// destination. If you specify this parameter, you can't specify the parameter for
	// the destination IP address.
	FilterAtDestination *types.PathRequestFilter

	// Scopes the analysis to network paths that match specific filters at the source.
	// If you specify this parameter, you can't specify the parameters for the source
	// IP address or the destination port.
	FilterAtSource *types.PathRequestFilter

	// The IP address of the source.
	SourceIp *string

	// The tags to add to the path.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateNetworkInsightsPathOutput struct {

	// Information about the path.
	NetworkInsightsPath *types.NetworkInsightsPath

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNetworkInsightsPathMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateNetworkInsightsPath{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateNetworkInsightsPath{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNetworkInsightsPath"); err != nil {
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
	if err = addIdempotencyToken_opCreateNetworkInsightsPathMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNetworkInsightsPathValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkInsightsPath(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateNetworkInsightsPath struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNetworkInsightsPath) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNetworkInsightsPath) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNetworkInsightsPathInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNetworkInsightsPathInput ")
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
func addIdempotencyToken_opCreateNetworkInsightsPathMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateNetworkInsightsPath{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNetworkInsightsPath(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNetworkInsightsPath",
	}
}
