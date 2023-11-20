// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified Amazon Web Services account alias. For information about
// using an Amazon Web Services account alias, see Creating, deleting, and listing
// an Amazon Web Services account alias (https://docs.aws.amazon.com/signin/latest/userguide/CreateAccountAlias.html)
// in the Amazon Web Services Sign-In User Guide.
func (c *Client) DeleteAccountAlias(ctx context.Context, params *DeleteAccountAliasInput, optFns ...func(*Options)) (*DeleteAccountAliasOutput, error) {
	if params == nil {
		params = &DeleteAccountAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccountAlias", params, optFns, c.addOperationDeleteAccountAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccountAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccountAliasInput struct {

	// The name of the account alias to delete. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex) ) a string of characters
	// consisting of lowercase letters, digits, and dashes. You cannot start or finish
	// with a dash, nor can you have two dashes in a row.
	//
	// This member is required.
	AccountAlias *string

	noSmithyDocumentSerde
}

type DeleteAccountAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAccountAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAccountAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAccountAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAccountAlias"); err != nil {
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
	if err = addOpDeleteAccountAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccountAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAccountAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAccountAlias",
	}
}
