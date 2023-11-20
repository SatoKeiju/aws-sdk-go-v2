// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a channel message. Only admins can perform this action. Deletion makes
// messages inaccessible immediately. A background process deletes any revisions
// created by UpdateChannelMessage . The x-amz-chime-bearer request header is
// mandatory. Use the ARN of the AppInstanceUser or AppInstanceBot that makes the
// API call as the value in the header.
func (c *Client) DeleteChannelMessage(ctx context.Context, params *DeleteChannelMessageInput, optFns ...func(*Options)) (*DeleteChannelMessageOutput, error) {
	if params == nil {
		params = &DeleteChannelMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteChannelMessage", params, optFns, c.addOperationDeleteChannelMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteChannelMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteChannelMessageInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The ID of the message being deleted.
	//
	// This member is required.
	MessageId *string

	// The ID of the SubChannel in the request. Only required when deleting messages
	// in a SubChannel that the user belongs to.
	SubChannelId *string

	noSmithyDocumentSerde
}

type DeleteChannelMessageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteChannelMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteChannelMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteChannelMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteChannelMessage"); err != nil {
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
	if err = addOpDeleteChannelMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChannelMessage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteChannelMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteChannelMessage",
	}
}
