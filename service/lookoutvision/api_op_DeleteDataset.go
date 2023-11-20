// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing Amazon Lookout for Vision dataset . If your the project has
// a single dataset, you must create a new dataset before you can create a model.
// If you project has a training dataset and a test dataset consider the following.
//
//   - If you delete the test dataset, your project reverts to a single dataset
//     project. If you then train the model, Amazon Lookout for Vision internally
//     splits the remaining dataset into a training and test dataset.
//   - If you delete the training dataset, you must create a training dataset
//     before you can create a model.
//
// This operation requires permissions to perform the lookoutvision:DeleteDataset
// operation.
func (c *Client) DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) {
	if params == nil {
		params = &DeleteDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataset", params, optFns, c.addOperationDeleteDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDatasetInput struct {

	// The type of the dataset to delete. Specify train to delete the training
	// dataset. Specify test to delete the test dataset. To delete the dataset in a
	// single dataset project, specify train .
	//
	// This member is required.
	DatasetType *string

	// The name of the project that contains the dataset that you want to delete.
	//
	// This member is required.
	ProjectName *string

	// ClientToken is an idempotency token that ensures a call to DeleteDataset
	// completes only once. You choose the value to pass. For example, An issue might
	// prevent you from getting a response from DeleteDataset . In this case, safely
	// retry your call to DeleteDataset by using the same ClientToken parameter value.
	// If you don't supply a value for ClientToken , the AWS SDK you are using inserts
	// a value for you. This prevents retries after a network error from making
	// multiple deletetion requests. You'll need to provide your own value for other
	// use cases. An error occurs if the other input parameters are not the same as in
	// the first request. Using a different value for ClientToken is considered a new
	// call to DeleteDataset . An idempotency token is active for 8 hours.
	ClientToken *string

	noSmithyDocumentSerde
}

type DeleteDatasetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteDataset"); err != nil {
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
	if err = addIdempotencyToken_opDeleteDatasetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataset(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteDataset struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteDataset) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteDataset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteDatasetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteDatasetInput ")
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
func addIdempotencyToken_opDeleteDatasetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteDataset{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteDataset",
	}
}
