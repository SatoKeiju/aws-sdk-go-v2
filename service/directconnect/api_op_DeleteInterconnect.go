// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified interconnect. Intended for use by Direct Connect Partners
// only.
func (c *Client) DeleteInterconnect(ctx context.Context, params *DeleteInterconnectInput, optFns ...func(*Options)) (*DeleteInterconnectOutput, error) {
	if params == nil {
		params = &DeleteInterconnectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteInterconnect", params, optFns, c.addOperationDeleteInterconnectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteInterconnectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInterconnectInput struct {

	// The ID of the interconnect.
	//
	// This member is required.
	InterconnectId *string

	noSmithyDocumentSerde
}

type DeleteInterconnectOutput struct {

	// The state of the interconnect. The following are the possible values:
	//   - requested : The initial state of an interconnect. The interconnect stays in
	//   the requested state until the Letter of Authorization (LOA) is sent to the
	//   customer.
	//   - pending : The interconnect is approved, and is being initialized.
	//   - available : The network link is up, and the interconnect is ready for use.
	//   - down : The network link is down.
	//   - deleting : The interconnect is being deleted.
	//   - deleted : The interconnect is deleted.
	//   - unknown : The state of the interconnect is not available.
	InterconnectState types.InterconnectState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteInterconnectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteInterconnect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteInterconnect{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteInterconnect"); err != nil {
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
	if err = addOpDeleteInterconnectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInterconnect(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteInterconnect(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteInterconnect",
	}
}
