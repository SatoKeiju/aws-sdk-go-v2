// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Kinesis data stream and all its shards and data. You must shut down
// any applications that are operating on the stream before you delete the stream.
// If an application attempts to operate on a deleted stream, it receives the
// exception ResourceNotFoundException . When invoking this API, you must use
// either the StreamARN or the StreamName parameter, or both. It is recommended
// that you use the StreamARN input parameter when you invoke this API. If the
// stream is in the ACTIVE state, you can delete it. After a DeleteStream request,
// the specified stream is in the DELETING state until Kinesis Data Streams
// completes the deletion. Note: Kinesis Data Streams might continue to accept data
// read and write operations, such as PutRecord , PutRecords , and GetRecords , on
// a stream in the DELETING state until the stream deletion is complete. When you
// delete a stream, any shards in that stream are also deleted, and any tags are
// dissociated from the stream. You can use the DescribeStreamSummary operation to
// check the state of the stream, which is returned in StreamStatus . DeleteStream
// has a limit of five transactions per second per account.
func (c *Client) DeleteStream(ctx context.Context, params *DeleteStreamInput, optFns ...func(*Options)) (*DeleteStreamOutput, error) {
	if params == nil {
		params = &DeleteStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStream", params, optFns, c.addOperationDeleteStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for DeleteStream .
type DeleteStreamInput struct {

	// If this parameter is unset ( null ) or if you set it to false , and the stream
	// has registered consumers, the call to DeleteStream fails with a
	// ResourceInUseException .
	EnforceConsumerDeletion *bool

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream to delete.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *DeleteStreamInput) bindEndpointParams(p *EndpointParameters) {
	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

type DeleteStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&smithyRpcv2cbor_serializeOpDeleteStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&smithyRpcv2cbor_deserializeOpDeleteStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteStream"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStream(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDeleteStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteStream",
	}
}
