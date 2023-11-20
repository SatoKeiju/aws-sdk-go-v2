// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an association. For more information about extensions and associations,
// see Working with AppConfig extensions (https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html)
// in the AppConfig User Guide.
func (c *Client) UpdateExtensionAssociation(ctx context.Context, params *UpdateExtensionAssociationInput, optFns ...func(*Options)) (*UpdateExtensionAssociationOutput, error) {
	if params == nil {
		params = &UpdateExtensionAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateExtensionAssociation", params, optFns, c.addOperationUpdateExtensionAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateExtensionAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateExtensionAssociationInput struct {

	// The system-generated ID for the association.
	//
	// This member is required.
	ExtensionAssociationId *string

	// The parameter names and values defined in the extension.
	Parameters map[string]string

	noSmithyDocumentSerde
}

type UpdateExtensionAssociationOutput struct {

	// The system-generated Amazon Resource Name (ARN) for the extension.
	Arn *string

	// The ARN of the extension defined in the association.
	ExtensionArn *string

	// The version number for the extension defined in the association.
	ExtensionVersionNumber int32

	// The system-generated ID for the association.
	Id *string

	// The parameter names and values defined in the association.
	Parameters map[string]string

	// The ARNs of applications, configuration profiles, or environments defined in
	// the association.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateExtensionAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateExtensionAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateExtensionAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateExtensionAssociation"); err != nil {
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
	if err = addOpUpdateExtensionAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateExtensionAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateExtensionAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateExtensionAssociation",
	}
}
