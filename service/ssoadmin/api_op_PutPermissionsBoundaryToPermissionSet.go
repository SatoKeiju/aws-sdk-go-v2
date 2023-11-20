// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches an Amazon Web Services managed or customer managed policy to the
// specified PermissionSet as a permissions boundary.
func (c *Client) PutPermissionsBoundaryToPermissionSet(ctx context.Context, params *PutPermissionsBoundaryToPermissionSetInput, optFns ...func(*Options)) (*PutPermissionsBoundaryToPermissionSetOutput, error) {
	if params == nil {
		params = &PutPermissionsBoundaryToPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPermissionsBoundaryToPermissionSet", params, optFns, c.addOperationPutPermissionsBoundaryToPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPermissionsBoundaryToPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPermissionsBoundaryToPermissionSetInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the PermissionSet .
	//
	// This member is required.
	PermissionSetArn *string

	// The permissions boundary that you want to attach to a PermissionSet .
	//
	// This member is required.
	PermissionsBoundary *types.PermissionsBoundary

	noSmithyDocumentSerde
}

type PutPermissionsBoundaryToPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPermissionsBoundaryToPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPermissionsBoundaryToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPermissionsBoundaryToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutPermissionsBoundaryToPermissionSet"); err != nil {
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
	if err = addOpPutPermissionsBoundaryToPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPermissionsBoundaryToPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPermissionsBoundaryToPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutPermissionsBoundaryToPermissionSet",
	}
}
