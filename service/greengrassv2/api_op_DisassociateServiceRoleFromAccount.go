// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates the Greengrass service role from IoT Greengrass for your Amazon
// Web Services account in this Amazon Web Services Region. Without a service role,
// IoT Greengrass can't verify the identity of client devices or manage core device
// connectivity information. For more information, see Greengrass service role (https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-service-role.html)
// in the IoT Greengrass Version 2 Developer Guide.
func (c *Client) DisassociateServiceRoleFromAccount(ctx context.Context, params *DisassociateServiceRoleFromAccountInput, optFns ...func(*Options)) (*DisassociateServiceRoleFromAccountOutput, error) {
	if params == nil {
		params = &DisassociateServiceRoleFromAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateServiceRoleFromAccount", params, optFns, c.addOperationDisassociateServiceRoleFromAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateServiceRoleFromAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateServiceRoleFromAccountInput struct {
	noSmithyDocumentSerde
}

type DisassociateServiceRoleFromAccountOutput struct {

	// The time when the service role was disassociated from IoT Greengrass for your
	// Amazon Web Services account in this Amazon Web Services Region.
	DisassociatedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateServiceRoleFromAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateServiceRoleFromAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateServiceRoleFromAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateServiceRoleFromAccount"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateServiceRoleFromAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateServiceRoleFromAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateServiceRoleFromAccount",
	}
}
