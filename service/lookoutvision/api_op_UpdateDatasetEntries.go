// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutvision/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates one or more JSON Line entries in a dataset. A JSON Line
// includes information about an image used for training or testing an Amazon
// Lookout for Vision model. To update an existing JSON Line, use the source-ref
// field to identify the JSON Line. The JSON line that you supply replaces the
// existing JSON line. Any existing annotations that are not in the new JSON line
// are removed from the dataset. For more information, see Defining JSON lines for
// anomaly classification in the Amazon Lookout for Vision Developer Guide. The
// images you reference in the source-ref field of a JSON line, must be in the
// same S3 bucket as the existing images in the dataset. Updating a dataset might
// take a while to complete. To check the current status, call DescribeDataset and
// check the Status field in the response. This operation requires permissions to
// perform the lookoutvision:UpdateDatasetEntries operation.
func (c *Client) UpdateDatasetEntries(ctx context.Context, params *UpdateDatasetEntriesInput, optFns ...func(*Options)) (*UpdateDatasetEntriesOutput, error) {
	if params == nil {
		params = &UpdateDatasetEntriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDatasetEntries", params, optFns, c.addOperationUpdateDatasetEntriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDatasetEntriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDatasetEntriesInput struct {

	// The entries to add to the dataset.
	//
	// This member is required.
	Changes []byte

	// The type of the dataset that you want to update. Specify train to update the
	// training dataset. Specify test to update the test dataset. If you have a single
	// dataset project, specify train .
	//
	// This member is required.
	DatasetType *string

	// The name of the project that contains the dataset that you want to update.
	//
	// This member is required.
	ProjectName *string

	// ClientToken is an idempotency token that ensures a call to UpdateDatasetEntries
	// completes only once. You choose the value to pass. For example, An issue might
	// prevent you from getting a response from UpdateDatasetEntries . In this case,
	// safely retry your call to UpdateDatasetEntries by using the same ClientToken
	// parameter value. If you don't supply a value for ClientToken , the AWS SDK you
	// are using inserts a value for you. This prevents retries after a network error
	// from making multiple updates with the same dataset entries. You'll need to
	// provide your own value for other use cases. An error occurs if the other input
	// parameters are not the same as in the first request. Using a different value for
	// ClientToken is considered a new call to UpdateDatasetEntries . An idempotency
	// token is active for 8 hours.
	ClientToken *string

	noSmithyDocumentSerde
}

type UpdateDatasetEntriesOutput struct {

	// The status of the dataset update.
	Status types.DatasetStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDatasetEntriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDatasetEntries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDatasetEntries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDatasetEntries"); err != nil {
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
	if err = addIdempotencyToken_opUpdateDatasetEntriesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateDatasetEntriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDatasetEntries(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateDatasetEntries struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateDatasetEntries) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateDatasetEntries) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateDatasetEntriesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateDatasetEntriesInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateDatasetEntriesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateDatasetEntries{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateDatasetEntries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDatasetEntries",
	}
}
