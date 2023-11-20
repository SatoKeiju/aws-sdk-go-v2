// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of runs for a machine learning transform. Machine learning task
// runs are asynchronous tasks that Glue runs on your behalf as part of various
// machine learning workflows. You can get a sortable, filterable list of machine
// learning task runs by calling GetMLTaskRuns with their parent transform's
// TransformID and other optional parameters as documented in this section. This
// operation returns a list of historic runs and must be paginated.
func (c *Client) GetMLTaskRuns(ctx context.Context, params *GetMLTaskRunsInput, optFns ...func(*Options)) (*GetMLTaskRunsOutput, error) {
	if params == nil {
		params = &GetMLTaskRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMLTaskRuns", params, optFns, c.addOperationGetMLTaskRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMLTaskRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMLTaskRunsInput struct {

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string

	// The filter criteria, in the TaskRunFilterCriteria structure, for the task run.
	Filter *types.TaskRunFilterCriteria

	// The maximum number of results to return.
	MaxResults *int32

	// A token for pagination of the results. The default is empty.
	NextToken *string

	// The sorting criteria, in the TaskRunSortCriteria structure, for the task run.
	Sort *types.TaskRunSortCriteria

	noSmithyDocumentSerde
}

type GetMLTaskRunsOutput struct {

	// A pagination token, if more results are available.
	NextToken *string

	// A list of task runs that are associated with the transform.
	TaskRuns []types.TaskRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMLTaskRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMLTaskRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMLTaskRuns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMLTaskRuns"); err != nil {
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
	if err = addOpGetMLTaskRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMLTaskRuns(options.Region), middleware.Before); err != nil {
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

// GetMLTaskRunsAPIClient is a client that implements the GetMLTaskRuns operation.
type GetMLTaskRunsAPIClient interface {
	GetMLTaskRuns(context.Context, *GetMLTaskRunsInput, ...func(*Options)) (*GetMLTaskRunsOutput, error)
}

var _ GetMLTaskRunsAPIClient = (*Client)(nil)

// GetMLTaskRunsPaginatorOptions is the paginator options for GetMLTaskRuns
type GetMLTaskRunsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetMLTaskRunsPaginator is a paginator for GetMLTaskRuns
type GetMLTaskRunsPaginator struct {
	options   GetMLTaskRunsPaginatorOptions
	client    GetMLTaskRunsAPIClient
	params    *GetMLTaskRunsInput
	nextToken *string
	firstPage bool
}

// NewGetMLTaskRunsPaginator returns a new GetMLTaskRunsPaginator
func NewGetMLTaskRunsPaginator(client GetMLTaskRunsAPIClient, params *GetMLTaskRunsInput, optFns ...func(*GetMLTaskRunsPaginatorOptions)) *GetMLTaskRunsPaginator {
	if params == nil {
		params = &GetMLTaskRunsInput{}
	}

	options := GetMLTaskRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetMLTaskRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetMLTaskRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetMLTaskRuns page.
func (p *GetMLTaskRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetMLTaskRunsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetMLTaskRuns(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetMLTaskRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMLTaskRuns",
	}
}
