// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an app block. App blocks are an Amazon AppStream 2.0 resource that
// stores the details about the virtual hard disk in an S3 bucket. It also stores
// the setup script with details about how to mount the virtual hard disk. The
// virtual hard disk includes the application binaries and other files necessary to
// launch your applications. Multiple applications can be assigned to a single app
// block. This is only supported for Elastic fleets.
func (c *Client) CreateAppBlock(ctx context.Context, params *CreateAppBlockInput, optFns ...func(*Options)) (*CreateAppBlockOutput, error) {
	if params == nil {
		params = &CreateAppBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppBlock", params, optFns, c.addOperationCreateAppBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppBlockInput struct {

	// The name of the app block.
	//
	// This member is required.
	Name *string

	// The source S3 location of the app block.
	//
	// This member is required.
	SourceS3Location *types.S3Location

	// The description of the app block.
	Description *string

	// The display name of the app block. This is not displayed to the user.
	DisplayName *string

	// The packaging type of the app block.
	PackagingType types.PackagingType

	// The post setup script details of the app block. This can only be provided for
	// the APPSTREAM2 PackagingType.
	PostSetupScriptDetails *types.ScriptDetails

	// The setup script details of the app block. This must be provided for the CUSTOM
	// PackagingType.
	SetupScriptDetails *types.ScriptDetails

	// The tags assigned to the app block.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAppBlockOutput struct {

	// The app block.
	AppBlock *types.AppBlock

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAppBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAppBlock{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAppBlock"); err != nil {
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
	if err = addOpCreateAppBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppBlock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAppBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAppBlock",
	}
}
