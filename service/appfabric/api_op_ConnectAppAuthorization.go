// Code generated by smithy-go-codegen DO NOT EDIT.

package appfabric

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appfabric/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Establishes a connection between Amazon Web Services AppFabric and an
// application, which allows AppFabric to call the APIs of the application.
func (c *Client) ConnectAppAuthorization(ctx context.Context, params *ConnectAppAuthorizationInput, optFns ...func(*Options)) (*ConnectAppAuthorizationOutput, error) {
	if params == nil {
		params = &ConnectAppAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConnectAppAuthorization", params, optFns, c.addOperationConnectAppAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConnectAppAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConnectAppAuthorizationInput struct {

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app
	// authorization to use for the request.
	//
	// This member is required.
	AppAuthorizationIdentifier *string

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app
	// bundle that contains the app authorization to use for the request.
	//
	// This member is required.
	AppBundleIdentifier *string

	// Contains OAuth2 authorization information. This is required if the app
	// authorization for the request is configured with an OAuth2 ( oauth2 )
	// authorization type.
	AuthRequest *types.AuthRequest

	noSmithyDocumentSerde
}

type ConnectAppAuthorizationOutput struct {

	// Contains a summary of the app authorization.
	//
	// This member is required.
	AppAuthorizationSummary *types.AppAuthorizationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConnectAppAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConnectAppAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConnectAppAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ConnectAppAuthorization"); err != nil {
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
	if err = addOpConnectAppAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConnectAppAuthorization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConnectAppAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConnectAppAuthorization",
	}
}
