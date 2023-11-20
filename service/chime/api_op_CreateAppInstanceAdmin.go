// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Promotes an AppInstanceUser to an AppInstanceAdmin . The promoted user can
// perform the following actions. This API is is no longer supported and will not
// be updated. We recommend using the latest version, CreateAppInstanceAdmin (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_identity-chime_CreateAppInstanceAdmin.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//   - ChannelModerator actions across all channels in the AppInstance .
//   - DeleteChannelMessage actions.
//
// Only an AppInstanceUser can be promoted to an AppInstanceAdmin role.
//
// Deprecated: Replaced by CreateAppInstanceAdmin in the Amazon Chime SDK Identity
// Namespace
func (c *Client) CreateAppInstanceAdmin(ctx context.Context, params *CreateAppInstanceAdminInput, optFns ...func(*Options)) (*CreateAppInstanceAdminOutput, error) {
	if params == nil {
		params = &CreateAppInstanceAdminInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppInstanceAdmin", params, optFns, c.addOperationCreateAppInstanceAdminMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppInstanceAdminOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppInstanceAdminInput struct {

	// The ARN of the administrator of the current AppInstance .
	//
	// This member is required.
	AppInstanceAdminArn *string

	// The ARN of the AppInstance .
	//
	// This member is required.
	AppInstanceArn *string

	noSmithyDocumentSerde
}

type CreateAppInstanceAdminOutput struct {

	// The name and ARN of the admin for the AppInstance .
	AppInstanceAdmin *types.Identity

	// The ARN of the of the admin for the AppInstance .
	AppInstanceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppInstanceAdminMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAppInstanceAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAppInstanceAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAppInstanceAdmin"); err != nil {
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
	if err = addEndpointPrefix_opCreateAppInstanceAdminMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAppInstanceAdminValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppInstanceAdmin(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateAppInstanceAdminMiddleware struct {
}

func (*endpointPrefix_opCreateAppInstanceAdminMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateAppInstanceAdminMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "identity-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateAppInstanceAdminMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateAppInstanceAdminMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateAppInstanceAdmin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAppInstanceAdmin",
	}
}
