// Code generated by smithy-go-codegen DO NOT EDIT.

package workmailmessageflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the raw content of an in-transit email message, in MIME format. This
// example describes how to update in-transit email message. For more information
// and examples for using this API, see Updating message content with AWS Lambda (https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html)
// . Updates to an in-transit message only appear when you call
// PutRawMessageContent from an AWS Lambda function configured with a synchronous
// Run Lambda (https://docs.aws.amazon.com/workmail/latest/adminguide/lambda.html#synchronous-rules)
// rule. If you call PutRawMessageContent on a delivered or sent message, the
// message remains unchanged, even though GetRawMessageContent (https://docs.aws.amazon.com/workmail/latest/APIReference/API_messageflow_GetRawMessageContent.html)
// returns an updated message.
func (c *Client) PutRawMessageContent(ctx context.Context, params *PutRawMessageContentInput, optFns ...func(*Options)) (*PutRawMessageContentOutput, error) {
	if params == nil {
		params = &PutRawMessageContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRawMessageContent", params, optFns, c.addOperationPutRawMessageContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRawMessageContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRawMessageContentInput struct {

	// Describes the raw message content of the updated email message.
	//
	// This member is required.
	Content *types.RawMessageContent

	// The identifier of the email message being updated.
	//
	// This member is required.
	MessageId *string

	noSmithyDocumentSerde
}

type PutRawMessageContentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRawMessageContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutRawMessageContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRawMessageContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRawMessageContent"); err != nil {
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
	if err = addOpPutRawMessageContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRawMessageContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRawMessageContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRawMessageContent",
	}
}
