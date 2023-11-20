// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports the root certificate of a certificate authority (CA) used to obtain TLS
// certificates used by associated websites within the company network.
//
// Deprecated: Amazon WorkLink is no longer supported. This will be removed in a
// future version of the SDK.
func (c *Client) AssociateWebsiteCertificateAuthority(ctx context.Context, params *AssociateWebsiteCertificateAuthorityInput, optFns ...func(*Options)) (*AssociateWebsiteCertificateAuthorityOutput, error) {
	if params == nil {
		params = &AssociateWebsiteCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateWebsiteCertificateAuthority", params, optFns, c.addOperationAssociateWebsiteCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateWebsiteCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWebsiteCertificateAuthorityInput struct {

	// The root certificate of the CA.
	//
	// This member is required.
	Certificate *string

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The certificate name to display.
	DisplayName *string

	noSmithyDocumentSerde
}

type AssociateWebsiteCertificateAuthorityOutput struct {

	// A unique identifier for the CA.
	WebsiteCaId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateWebsiteCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateWebsiteCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateWebsiteCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateWebsiteCertificateAuthority"); err != nil {
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
	if err = addOpAssociateWebsiteCertificateAuthorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWebsiteCertificateAuthority(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateWebsiteCertificateAuthority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateWebsiteCertificateAuthority",
	}
}
