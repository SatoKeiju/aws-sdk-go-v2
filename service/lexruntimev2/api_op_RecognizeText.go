// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends user input to Amazon Lex V2. Client applications use this API to send
// requests to Amazon Lex V2 at runtime. Amazon Lex V2 then interprets the user
// input using the machine learning model that it build for the bot. In response,
// Amazon Lex V2 returns the next message to convey to the user and an optional
// response card to display. If the optional post-fulfillment response is
// specified, the messages are returned as follows. For more information, see
// PostFulfillmentStatusSpecification (https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html)
// .
//   - Success message - Returned if the Lambda function completes successfully
//     and the intent state is fulfilled or ready fulfillment if the message is
//     present.
//   - Failed message - The failed message is returned if the Lambda function
//     throws an exception or if the Lambda function returns a failed intent state
//     without a message.
//   - Timeout message - If you don't configure a timeout message and a timeout,
//     and the Lambda function doesn't return within 30 seconds, the timeout message is
//     returned. If you configure a timeout, the timeout message is returned when the
//     period times out.
//
// For more information, see Completion message (https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html)
// .
func (c *Client) RecognizeText(ctx context.Context, params *RecognizeTextInput, optFns ...func(*Options)) (*RecognizeTextOutput, error) {
	if params == nil {
		params = &RecognizeTextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RecognizeText", params, optFns, c.addOperationRecognizeTextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RecognizeTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecognizeTextInput struct {

	// The alias identifier in use for the bot that processes the request.
	//
	// This member is required.
	BotAliasId *string

	// The identifier of the bot that processes the request.
	//
	// This member is required.
	BotId *string

	// The locale where the session is in use.
	//
	// This member is required.
	LocaleId *string

	// The identifier of the user session that is having the conversation.
	//
	// This member is required.
	SessionId *string

	// The text that the user entered. Amazon Lex V2 interprets this text.
	//
	// This member is required.
	Text *string

	// Request-specific information passed between the client application and Amazon
	// Lex V2 The namespace x-amz-lex: is reserved for special attributes. Don't
	// create any request attributes with the prefix x-amz-lex: .
	RequestAttributes map[string]string

	// The current state of the dialog between the user and the bot.
	SessionState *types.SessionState

	noSmithyDocumentSerde
}

type RecognizeTextOutput struct {

	// A list of intents that Amazon Lex V2 determined might satisfy the user's
	// utterance. Each interpretation includes the intent, a score that indicates now
	// confident Amazon Lex V2 is that the interpretation is the correct one, and an
	// optional sentiment response that indicates the sentiment expressed in the
	// utterance.
	Interpretations []types.Interpretation

	// A list of messages last sent to the user. The messages are ordered based on the
	// order that you returned the messages from your Lambda function or the order that
	// the messages are defined in the bot.
	Messages []types.Message

	// The bot member that recognized the text.
	RecognizedBotMember *types.RecognizedBotMember

	// The attributes sent in the request.
	RequestAttributes map[string]string

	// The identifier of the session in use.
	SessionId *string

	// Represents the current state of the dialog between the user and the bot. Use
	// this to determine the progress of the conversation and what the next action may
	// be.
	SessionState *types.SessionState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRecognizeTextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRecognizeText{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRecognizeText{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RecognizeText"); err != nil {
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
	if err = addOpRecognizeTextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRecognizeText(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRecognizeText(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RecognizeText",
	}
}
