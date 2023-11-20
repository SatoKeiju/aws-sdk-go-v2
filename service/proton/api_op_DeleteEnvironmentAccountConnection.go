// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// In an environment account, delete an environment account connection. After you
// delete an environment account connection that’s in use by an Proton environment,
// Proton can’t manage the environment infrastructure resources until a new
// environment account connection is accepted for the environment account and
// associated environment. You're responsible for cleaning up provisioned resources
// that remain without an environment connection. For more information, see
// Environment account connections (https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html)
// in the Proton User guide.
func (c *Client) DeleteEnvironmentAccountConnection(ctx context.Context, params *DeleteEnvironmentAccountConnectionInput, optFns ...func(*Options)) (*DeleteEnvironmentAccountConnectionOutput, error) {
	if params == nil {
		params = &DeleteEnvironmentAccountConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEnvironmentAccountConnection", params, optFns, c.addOperationDeleteEnvironmentAccountConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEnvironmentAccountConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEnvironmentAccountConnectionInput struct {

	// The ID of the environment account connection to delete.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DeleteEnvironmentAccountConnectionOutput struct {

	// The detailed data of the environment account connection being deleted.
	EnvironmentAccountConnection *types.EnvironmentAccountConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEnvironmentAccountConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteEnvironmentAccountConnection"); err != nil {
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
	if err = addOpDeleteEnvironmentAccountConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEnvironmentAccountConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEnvironmentAccountConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteEnvironmentAccountConnection",
	}
}
