// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the webhook associated with an CodeBuild build project. If you use
// Bitbucket for your repository, rotateSecret is ignored.
func (c *Client) UpdateWebhook(ctx context.Context, params *UpdateWebhookInput, optFns ...func(*Options)) (*UpdateWebhookOutput, error) {
	if params == nil {
		params = &UpdateWebhookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWebhook", params, optFns, c.addOperationUpdateWebhookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWebhookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWebhookInput struct {

	// The name of the CodeBuild project.
	//
	// This member is required.
	ProjectName *string

	// A regular expression used to determine which repository branches are built when
	// a webhook is triggered. If the name of a branch matches the regular expression,
	// then it is built. If branchFilter is empty, then all branches are built. It is
	// recommended that you use filterGroups instead of branchFilter .
	BranchFilter *string

	// Specifies the type of build this webhook will trigger.
	BuildType types.WebhookBuildType

	// An array of arrays of WebhookFilter objects used to determine if a webhook
	// event can trigger a build. A filter group must contain at least one EVENT
	// WebhookFilter .
	FilterGroups [][]types.WebhookFilter

	// A boolean value that specifies whether the associated GitHub repository's
	// secret token should be updated. If you use Bitbucket for your repository,
	// rotateSecret is ignored.
	RotateSecret bool

	noSmithyDocumentSerde
}

type UpdateWebhookOutput struct {

	// Information about a repository's webhook that is associated with a project in
	// CodeBuild.
	Webhook *types.Webhook

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWebhookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWebhook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWebhook{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateWebhook"); err != nil {
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
	if err = addOpUpdateWebhookValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWebhook(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWebhook(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateWebhook",
	}
}
