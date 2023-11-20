// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The action to start test set execution.
func (c *Client) StartTestExecution(ctx context.Context, params *StartTestExecutionInput, optFns ...func(*Options)) (*StartTestExecutionOutput, error) {
	if params == nil {
		params = &StartTestExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTestExecution", params, optFns, c.addOperationStartTestExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTestExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTestExecutionInput struct {

	// Indicates whether we use streaming or non-streaming APIs for the test set
	// execution. For streaming, StartConversation Runtime API is used. Whereas, for
	// non-streaming, RecognizeUtterance and RecognizeText Amazon Lex Runtime API are
	// used.
	//
	// This member is required.
	ApiMode types.TestExecutionApiMode

	// The target bot for the test set execution.
	//
	// This member is required.
	Target *types.TestExecutionTarget

	// The test set Id for the test set execution.
	//
	// This member is required.
	TestSetId *string

	// Indicates whether audio or text is used.
	TestExecutionModality types.TestExecutionModality

	noSmithyDocumentSerde
}

type StartTestExecutionOutput struct {

	// Indicates whether we use streaming or non-streaming APIs for the test set
	// execution. For streaming, StartConversation Amazon Lex Runtime API is used.
	// Whereas for non-streaming, RecognizeUtterance and RecognizeText Amazon Lex
	// Runtime API are used.
	ApiMode types.TestExecutionApiMode

	// The creation date and time for the test set execution.
	CreationDateTime *time.Time

	// The target bot for the test set execution.
	Target *types.TestExecutionTarget

	// The unique identifier of the test set execution.
	TestExecutionId *string

	// Indicates whether audio or text is used.
	TestExecutionModality types.TestExecutionModality

	// The test set Id for the test set execution.
	TestSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTestExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartTestExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartTestExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartTestExecution"); err != nil {
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
	if err = addOpStartTestExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTestExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartTestExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartTestExecution",
	}
}
