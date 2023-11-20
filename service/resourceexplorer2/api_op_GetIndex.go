// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details about the Amazon Web Services Resource Explorer index in the
// Amazon Web Services Region in which you invoked the operation.
func (c *Client) GetIndex(ctx context.Context, params *GetIndexInput, optFns ...func(*Options)) (*GetIndexOutput, error) {
	if params == nil {
		params = &GetIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIndex", params, optFns, c.addOperationGetIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIndexInput struct {
	noSmithyDocumentSerde
}

type GetIndexOutput struct {

	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the index.
	Arn *string

	// The date and time when the index was originally created.
	CreatedAt *time.Time

	// The date and time when the index was last updated.
	LastUpdatedAt *time.Time

	// This response value is present only if this index is Type=AGGREGATOR . A list of
	// the Amazon Web Services Regions that replicate their content to the index in
	// this Region.
	ReplicatingFrom []string

	// This response value is present only if this index is Type=LOCAL . The Amazon Web
	// Services Region that contains the aggregator index, if one exists. If an
	// aggregator index does exist then the Region in which you called this operation
	// replicates its index information to the Region specified in this response value.
	ReplicatingTo []string

	// The current state of the index in this Amazon Web Services Region.
	State types.IndexState

	// Tag key and value pairs that are attached to the index.
	Tags map[string]string

	// The type of the index in this Region. For information about the aggregator
	// index and how it differs from a local index, see Turning on cross-Region search
	// by creating an aggregator index (https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html)
	// .
	Type types.IndexType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIndex{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIndex"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIndex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIndex",
	}
}
