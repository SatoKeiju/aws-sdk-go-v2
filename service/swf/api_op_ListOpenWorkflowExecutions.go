// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of open workflow executions in the specified domain that meet
// the filtering criteria. The results may be split into multiple pages. To
// retrieve subsequent pages, make the call again using the nextPageToken returned
// by the initial call. This operation is eventually consistent. The results are
// best effort and may not exactly reflect recent updates and changes. Access
// Control You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//   - Use an Action element to allow or deny permission to call this action.
//   - Constrain the following parameters by using a Condition element with the
//     appropriate keys.
//   - tagFilter.tag : String constraint. The key is swf:tagFilter.tag .
//   - typeFilter.name : String constraint. The key is swf:typeFilter.name .
//   - typeFilter.version : String constraint. The key is swf:typeFilter.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) ListOpenWorkflowExecutions(ctx context.Context, params *ListOpenWorkflowExecutionsInput, optFns ...func(*Options)) (*ListOpenWorkflowExecutionsOutput, error) {
	if params == nil {
		params = &ListOpenWorkflowExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOpenWorkflowExecutions", params, optFns, c.addOperationListOpenWorkflowExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOpenWorkflowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOpenWorkflowExecutionsInput struct {

	// The name of the domain that contains the workflow executions to list.
	//
	// This member is required.
	Domain *string

	// Workflow executions are included in the returned results based on whether their
	// start times are within the range specified by this filter.
	//
	// This member is required.
	StartTimeFilter *types.ExecutionTimeFilter

	// If specified, only workflow executions matching the workflow ID specified in
	// the filter are returned. executionFilter , typeFilter and tagFilter are
	// mutually exclusive. You can specify at most one of these in a request.
	ExecutionFilter *types.WorkflowExecutionFilter

	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	MaximumPageSize int32

	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return a 400 error: " Specified token has exceeded its
	// maximum lifetime ". The configured maximumPageSize determines how many results
	// can be returned in a single call.
	NextPageToken *string

	// When set to true , returns the results in reverse order. By default the results
	// are returned in descending order of the start time of the executions.
	ReverseOrder bool

	// If specified, only executions that have the matching tag are listed.
	// executionFilter , typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	TagFilter *types.TagFilter

	// If specified, only executions of the type specified in the filter are returned.
	// executionFilter , typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	TypeFilter *types.WorkflowTypeFilter

	noSmithyDocumentSerde
}

// Contains a paginated list of information about workflow executions.
type ListOpenWorkflowExecutionsOutput struct {

	// The list of workflow information structures.
	//
	// This member is required.
	ExecutionInfos []types.WorkflowExecutionInfo

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken . Keep all other arguments unchanged. The
	// configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOpenWorkflowExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListOpenWorkflowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListOpenWorkflowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOpenWorkflowExecutions"); err != nil {
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
	if err = addOpListOpenWorkflowExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOpenWorkflowExecutions(options.Region), middleware.Before); err != nil {
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

// ListOpenWorkflowExecutionsAPIClient is a client that implements the
// ListOpenWorkflowExecutions operation.
type ListOpenWorkflowExecutionsAPIClient interface {
	ListOpenWorkflowExecutions(context.Context, *ListOpenWorkflowExecutionsInput, ...func(*Options)) (*ListOpenWorkflowExecutionsOutput, error)
}

var _ ListOpenWorkflowExecutionsAPIClient = (*Client)(nil)

// ListOpenWorkflowExecutionsPaginatorOptions is the paginator options for
// ListOpenWorkflowExecutions
type ListOpenWorkflowExecutionsPaginatorOptions struct {
	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOpenWorkflowExecutionsPaginator is a paginator for
// ListOpenWorkflowExecutions
type ListOpenWorkflowExecutionsPaginator struct {
	options   ListOpenWorkflowExecutionsPaginatorOptions
	client    ListOpenWorkflowExecutionsAPIClient
	params    *ListOpenWorkflowExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListOpenWorkflowExecutionsPaginator returns a new
// ListOpenWorkflowExecutionsPaginator
func NewListOpenWorkflowExecutionsPaginator(client ListOpenWorkflowExecutionsAPIClient, params *ListOpenWorkflowExecutionsInput, optFns ...func(*ListOpenWorkflowExecutionsPaginatorOptions)) *ListOpenWorkflowExecutionsPaginator {
	if params == nil {
		params = &ListOpenWorkflowExecutionsInput{}
	}

	options := ListOpenWorkflowExecutionsPaginatorOptions{}
	if params.MaximumPageSize != 0 {
		options.Limit = params.MaximumPageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOpenWorkflowExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextPageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOpenWorkflowExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOpenWorkflowExecutions page.
func (p *ListOpenWorkflowExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOpenWorkflowExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextPageToken = p.nextToken

	params.MaximumPageSize = p.options.Limit

	result, err := p.client.ListOpenWorkflowExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListOpenWorkflowExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOpenWorkflowExecutions",
	}
}
