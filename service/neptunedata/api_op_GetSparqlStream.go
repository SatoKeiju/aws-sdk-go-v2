// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a stream for an RDF graph. With the Neptune Streams feature, you can
// generate a complete sequence of change-log entries that record every change made
// to your graph data as it happens. GetSparqlStream lets you collect these
// change-log entries for an RDF graph. The Neptune streams feature needs to be
// enabled on your Neptune DBcluster. To enable streams, set the neptune_streams (https://docs.aws.amazon.com/neptune/latest/userguide/parameters.html#parameters-db-cluster-parameters-neptune_streams)
// DB cluster parameter to 1 . See Capturing graph changes in real time using
// Neptune streams (https://docs.aws.amazon.com/neptune/latest/userguide/streams.html)
// . When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the neptune-db:GetStreamRecords (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getstreamrecords)
// IAM action in that cluster. Note that the neptune-db:QueryLanguage:Sparql (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys)
// IAM condition key can be used in the policy document to restrict the use of
// SPARQL queries (see Condition keys available in Neptune IAM data-access policy
// statements (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html)
// ).
func (c *Client) GetSparqlStream(ctx context.Context, params *GetSparqlStreamInput, optFns ...func(*Options)) (*GetSparqlStreamOutput, error) {
	if params == nil {
		params = &GetSparqlStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSparqlStream", params, optFns, c.addOperationGetSparqlStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSparqlStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSparqlStreamInput struct {

	// The commit number of the starting record to read from the change-log stream.
	// This parameter is required when iteratorType is AT_SEQUENCE_NUMBER or
	// AFTER_SEQUENCE_NUMBER , and ignored when iteratorType is TRIM_HORIZON or LATEST .
	CommitNum *int64

	// If set to TRUE, Neptune compresses the response using gzip encoding.
	Encoding types.Encoding

	// Can be one of:
	//   - AT_SEQUENCE_NUMBER – Indicates that reading should start from the event
	//   sequence number specified jointly by the commitNum and opNum parameters.
	//   - AFTER_SEQUENCE_NUMBER – Indicates that reading should start right after the
	//   event sequence number specified jointly by the commitNum and opNum parameters.
	//   - TRIM_HORIZON – Indicates that reading should start at the last untrimmed
	//   record in the system, which is the oldest unexpired (not yet deleted) record in
	//   the change-log stream.
	//   - LATEST – Indicates that reading should start at the most recent record in
	//   the system, which is the latest unexpired (not yet deleted) record in the
	//   change-log stream.
	IteratorType types.IteratorType

	// Specifies the maximum number of records to return. There is also a size limit
	// of 10 MB on the response that can't be modified and that takes precedence over
	// the number of records specified in the limit parameter. The response does
	// include a threshold-breaching record if the 10 MB limit was reached. The range
	// for limit is 1 to 100,000, with a default of 10.
	Limit *int64

	// The operation sequence number within the specified commit to start reading from
	// in the change-log stream data. The default is 1 .
	OpNum *int64

	noSmithyDocumentSerde
}

type GetSparqlStreamOutput struct {

	// Serialization format for the change records being returned. Currently, the only
	// supported value is NQUADS .
	//
	// This member is required.
	Format *string

	// Sequence identifier of the last change in the stream response. An event ID is
	// composed of two fields: a commitNum , which identifies a transaction that
	// changed the graph, and an opNum , which identifies a specific operation within
	// that transaction:
	//
	// This member is required.
	LastEventId map[string]string

	// The time at which the commit for the transaction was requested, in milliseconds
	// from the Unix epoch.
	//
	// This member is required.
	LastTrxTimestampInMillis *int64

	// An array of serialized change-log stream records included in the response.
	//
	// This member is required.
	Records []types.SparqlRecord

	// The total number of records in the response.
	//
	// This member is required.
	TotalRecords *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSparqlStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSparqlStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSparqlStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSparqlStream"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSparqlStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSparqlStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSparqlStream",
	}
}
