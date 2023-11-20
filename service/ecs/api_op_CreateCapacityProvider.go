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

// Creates a new capacity provider. Capacity providers are associated with an
// Amazon ECS cluster and are used in capacity provider strategies to facilitate
// cluster auto scaling. Only capacity providers that use an Auto Scaling group can
// be created. Amazon ECS tasks on Fargate use the FARGATE and FARGATE_SPOT
// capacity providers. These providers are available to all accounts in the Amazon
// Web Services Regions that Fargate supports.
func (c *Client) CreateCapacityProvider(ctx context.Context, params *CreateCapacityProviderInput, optFns ...func(*Options)) (*CreateCapacityProviderOutput, error) {
	if params == nil {
		params = &CreateCapacityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCapacityProvider", params, optFns, c.addOperationCreateCapacityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCapacityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCapacityProviderInput struct {

	// The details of the Auto Scaling group for the capacity provider.
	//
	// This member is required.
	AutoScalingGroupProvider *types.AutoScalingGroupProvider

	// The name of the capacity provider. Up to 255 characters are allowed. They
	// include letters (both upper and lowercase letters), numbers, underscores (_),
	// and hyphens (-). The name can't be prefixed with " aws ", " ecs ", or " fargate
	// ".
	//
	// This member is required.
	Name *string

	// The metadata that you apply to the capacity provider to categorize and organize
	// them more conveniently. Each tag consists of a key and an optional value. You
	// define both of them. The following basic restrictions apply to tags:
	//   - Maximum number of tags per resource - 50
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//   - Maximum key length - 128 Unicode characters in UTF-8
	//   - Maximum value length - 256 Unicode characters in UTF-8
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//   - Tag keys and values are case-sensitive.
	//   - Do not use aws: , AWS: , or any upper or lowercase combination of such as a
	//   prefix for either keys or values as it is reserved for Amazon Web Services use.
	//   You cannot edit or delete tag keys or values with this prefix. Tags with this
	//   prefix do not count against your tags per resource limit.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCapacityProviderOutput struct {

	// The full description of the new capacity provider.
	CapacityProvider *types.CapacityProvider

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCapacityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCapacityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCapacityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCapacityProvider"); err != nil {
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
	if err = addOpCreateCapacityProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCapacityProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCapacityProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCapacityProvider",
	}
}
