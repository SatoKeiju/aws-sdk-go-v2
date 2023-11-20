// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the settings for an existing DB proxy endpoint.
func (c *Client) ModifyDBProxyEndpoint(ctx context.Context, params *ModifyDBProxyEndpointInput, optFns ...func(*Options)) (*ModifyDBProxyEndpointOutput, error) {
	if params == nil {
		params = &ModifyDBProxyEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBProxyEndpoint", params, optFns, c.addOperationModifyDBProxyEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBProxyEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBProxyEndpointInput struct {

	// The name of the DB proxy sociated with the DB proxy endpoint that you want to
	// modify.
	//
	// This member is required.
	DBProxyEndpointName *string

	// The new identifier for the DBProxyEndpoint . An identifier must begin with a
	// letter and must contain only ASCII letters, digits, and hyphens; it can't end
	// with a hyphen or contain two consecutive hyphens.
	NewDBProxyEndpointName *string

	// The VPC security group IDs for the DB proxy endpoint. When the DB proxy
	// endpoint uses a different VPC than the original proxy, you also specify a
	// different set of security group IDs than for the original proxy.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type ModifyDBProxyEndpointOutput struct {

	// The DBProxyEndpoint object representing the new settings for the DB proxy
	// endpoint.
	DBProxyEndpoint *types.DBProxyEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBProxyEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBProxyEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBProxyEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDBProxyEndpoint"); err != nil {
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
	if err = addOpModifyDBProxyEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBProxyEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBProxyEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDBProxyEndpoint",
	}
}
