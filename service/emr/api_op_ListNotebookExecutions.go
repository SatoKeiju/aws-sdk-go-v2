// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides summaries of all notebook executions. You can filter the list based on
// multiple criteria such as status, time range, and editor id. Returns a maximum
// of 50 notebook executions and a marker to track the paging of a longer notebook
// execution list across multiple ListNotebookExecutions calls.
func (c *Client) ListNotebookExecutions(ctx context.Context, params *ListNotebookExecutionsInput, optFns ...func(*Options)) (*ListNotebookExecutionsOutput, error) {
	if params == nil {
		params = &ListNotebookExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotebookExecutions", params, optFns, c.addOperationListNotebookExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotebookExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotebookExecutionsInput struct {

	// The unique ID of the editor associated with the notebook execution.
	EditorId *string

	// The unique ID of the execution engine.
	ExecutionEngineId *string

	// The beginning of time range filter for listing notebook executions. The default
	// is the timestamp of 30 days ago.
	From *time.Time

	// The pagination token, returned by a previous ListNotebookExecutions call, that
	// indicates the start of the list for this ListNotebookExecutions call.
	Marker *string

	// The status filter for listing notebook executions.
	//   - START_PENDING indicates that the cluster has received the execution request
	//   but execution has not begun.
	//   - STARTING indicates that the execution is starting on the cluster.
	//   - RUNNING indicates that the execution is being processed by the cluster.
	//   - FINISHING indicates that execution processing is in the final stages.
	//   - FINISHED indicates that the execution has completed without error.
	//   - FAILING indicates that the execution is failing and will not finish
	//   successfully.
	//   - FAILED indicates that the execution failed.
	//   - STOP_PENDING indicates that the cluster has received a StopNotebookExecution
	//   request and the stop is pending.
	//   - STOPPING indicates that the cluster is in the process of stopping the
	//   execution as a result of a StopNotebookExecution request.
	//   - STOPPED indicates that the execution stopped because of a
	//   StopNotebookExecution request.
	Status types.NotebookExecutionStatus

	// The end of time range filter for listing notebook executions. The default is
	// the current timestamp.
	To *time.Time

	noSmithyDocumentSerde
}

type ListNotebookExecutionsOutput struct {

	// A pagination token that a subsequent ListNotebookExecutions can use to
	// determine the next set of results to retrieve.
	Marker *string

	// A list of notebook executions.
	NotebookExecutions []types.NotebookExecutionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotebookExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNotebookExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNotebookExecutions{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotebookExecutions(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListNotebookExecutionsAPIClient is a client that implements the
// ListNotebookExecutions operation.
type ListNotebookExecutionsAPIClient interface {
	ListNotebookExecutions(context.Context, *ListNotebookExecutionsInput, ...func(*Options)) (*ListNotebookExecutionsOutput, error)
}

var _ ListNotebookExecutionsAPIClient = (*Client)(nil)

// ListNotebookExecutionsPaginatorOptions is the paginator options for
// ListNotebookExecutions
type ListNotebookExecutionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNotebookExecutionsPaginator is a paginator for ListNotebookExecutions
type ListNotebookExecutionsPaginator struct {
	options   ListNotebookExecutionsPaginatorOptions
	client    ListNotebookExecutionsAPIClient
	params    *ListNotebookExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListNotebookExecutionsPaginator returns a new ListNotebookExecutionsPaginator
func NewListNotebookExecutionsPaginator(client ListNotebookExecutionsAPIClient, params *ListNotebookExecutionsInput, optFns ...func(*ListNotebookExecutionsPaginatorOptions)) *ListNotebookExecutionsPaginator {
	if params == nil {
		params = &ListNotebookExecutionsInput{}
	}

	options := ListNotebookExecutionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNotebookExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNotebookExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNotebookExecutions page.
func (p *ListNotebookExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNotebookExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListNotebookExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListNotebookExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListNotebookExecutions",
	}
}
