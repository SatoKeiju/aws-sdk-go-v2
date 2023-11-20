// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables an account setting for a specified user, role, or the root user for an
// account.
func (c *Client) DeleteAccountSetting(ctx context.Context, params *DeleteAccountSettingInput, optFns ...func(*Options)) (*DeleteAccountSettingOutput, error) {
	if params == nil {
		params = &DeleteAccountSettingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccountSetting", params, optFns, c.addOperationDeleteAccountSettingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccountSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccountSettingInput struct {

	// The resource name to disable the account setting for. If serviceLongArnFormat
	// is specified, the ARN for your Amazon ECS services is affected. If
	// taskLongArnFormat is specified, the ARN and resource ID for your Amazon ECS
	// tasks is affected. If containerInstanceLongArnFormat is specified, the ARN and
	// resource ID for your Amazon ECS container instances is affected. If
	// awsvpcTrunking is specified, the ENI limit for your Amazon ECS container
	// instances is affected.
	//
	// This member is required.
	Name types.SettingName

	// The Amazon Resource Name (ARN) of the principal. It can be an user, role, or
	// the root user. If you specify the root user, it disables the account setting for
	// all users, roles, and the root user of the account unless a user or role
	// explicitly overrides these settings. If this field is omitted, the setting is
	// changed only for the authenticated user.
	PrincipalArn *string

	noSmithyDocumentSerde
}

type DeleteAccountSettingOutput struct {

	// The account setting for the specified principal ARN.
	Setting *types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAccountSettingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAccountSetting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAccountSetting{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAccountSetting"); err != nil {
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
	if err = addOpDeleteAccountSettingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccountSetting(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAccountSetting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAccountSetting",
	}
}
