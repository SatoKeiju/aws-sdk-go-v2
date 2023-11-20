// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Mark a user, group, or resource as no longer used in WorkMail. This action
// disassociates the mailbox and schedules it for clean-up. WorkMail keeps
// mailboxes for 30 days before they are permanently removed. The functionality in
// the console is Disable.
func (c *Client) DeregisterFromWorkMail(ctx context.Context, params *DeregisterFromWorkMailInput, optFns ...func(*Options)) (*DeregisterFromWorkMailOutput, error) {
	if params == nil {
		params = &DeregisterFromWorkMailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterFromWorkMail", params, optFns, c.addOperationDeregisterFromWorkMailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterFromWorkMailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterFromWorkMailInput struct {

	// The identifier for the member to be updated. The identifier can be UserId,
	// ResourceId, or Group Id, Username, Resourcename, or Groupname, or email.
	//   - Entity ID: 12345678-1234-1234-1234-123456789012,
	//   r-0123456789a0123456789b0123456789, or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//   - Email address: entity@domain.tld
	//   - Entity name: entity
	//
	// This member is required.
	EntityId *string

	// The identifier for the organization under which the WorkMail entity exists.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type DeregisterFromWorkMailOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterFromWorkMailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterFromWorkMail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterFromWorkMail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeregisterFromWorkMail"); err != nil {
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
	if err = addOpDeregisterFromWorkMailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterFromWorkMail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterFromWorkMail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeregisterFromWorkMail",
	}
}
