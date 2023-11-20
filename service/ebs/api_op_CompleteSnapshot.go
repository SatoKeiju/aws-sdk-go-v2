// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Seals and completes the snapshot after all of the required blocks of data have
// been written to it. Completing the snapshot changes the status to completed .
// You cannot write new blocks to a snapshot after it has been completed. You
// should always retry requests that receive server ( 5xx ) error responses, and
// ThrottlingException and RequestThrottledException client error responses. For
// more information see Error retries (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CompleteSnapshot(ctx context.Context, params *CompleteSnapshotInput, optFns ...func(*Options)) (*CompleteSnapshotOutput, error) {
	if params == nil {
		params = &CompleteSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteSnapshot", params, optFns, c.addOperationCompleteSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteSnapshotInput struct {

	// The number of blocks that were written to the snapshot.
	//
	// This member is required.
	ChangedBlocksCount *int32

	// The ID of the snapshot.
	//
	// This member is required.
	SnapshotId *string

	// An aggregated Base-64 SHA256 checksum based on the checksums of each written
	// block. To generate the aggregated checksum using the linear aggregation method,
	// arrange the checksums for each written block in ascending order of their block
	// index, concatenate them to form a single string, and then generate the checksum
	// on the entire string using the SHA256 algorithm.
	Checksum *string

	// The aggregation method used to generate the checksum. Currently, the only
	// supported aggregation method is LINEAR .
	ChecksumAggregationMethod types.ChecksumAggregationMethod

	// The algorithm used to generate the checksum. Currently, the only supported
	// algorithm is SHA256 .
	ChecksumAlgorithm types.ChecksumAlgorithm

	noSmithyDocumentSerde
}

type CompleteSnapshotOutput struct {

	// The status of the snapshot.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCompleteSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCompleteSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCompleteSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CompleteSnapshot"); err != nil {
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
	if err = addOpCompleteSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCompleteSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CompleteSnapshot",
	}
}
