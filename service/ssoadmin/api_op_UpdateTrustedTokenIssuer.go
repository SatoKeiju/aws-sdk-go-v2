// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the name of the trusted token issuer, or the path of a source attribute
// or destination attribute for a trusted token issuer configuration. Updating this
// trusted token issuer configuration might cause users to lose access to any
// applications that are configured to use the trusted token issuer.
func (c *Client) UpdateTrustedTokenIssuer(ctx context.Context, params *UpdateTrustedTokenIssuerInput, optFns ...func(*Options)) (*UpdateTrustedTokenIssuerOutput, error) {
	if params == nil {
		params = &UpdateTrustedTokenIssuerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTrustedTokenIssuer", params, optFns, c.addOperationUpdateTrustedTokenIssuerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrustedTokenIssuerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTrustedTokenIssuerInput struct {

	// Specifies the ARN of the trusted token issuer configuration that you want to
	// update.
	//
	// This member is required.
	TrustedTokenIssuerArn *string

	// Specifies the updated name to be applied to the trusted token issuer
	// configuration.
	Name *string

	// Specifies a structure with settings to apply to the specified trusted token
	// issuer. The settings that you can provide are determined by the type of the
	// trusted token issuer that you are updating.
	TrustedTokenIssuerConfiguration types.TrustedTokenIssuerUpdateConfiguration

	noSmithyDocumentSerde
}

type UpdateTrustedTokenIssuerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTrustedTokenIssuerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTrustedTokenIssuer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTrustedTokenIssuer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTrustedTokenIssuer"); err != nil {
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
	if err = addOpUpdateTrustedTokenIssuerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTrustedTokenIssuer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTrustedTokenIssuer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTrustedTokenIssuer",
	}
}
