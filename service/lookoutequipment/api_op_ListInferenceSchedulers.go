// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of all inference schedulers currently available for your
// account.
func (c *Client) ListInferenceSchedulers(ctx context.Context, params *ListInferenceSchedulersInput, optFns ...func(*Options)) (*ListInferenceSchedulersOutput, error) {
	if params == nil {
		params = &ListInferenceSchedulersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInferenceSchedulers", params, optFns, c.addOperationListInferenceSchedulersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInferenceSchedulersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInferenceSchedulersInput struct {

	// The beginning of the name of the inference schedulers to be listed.
	InferenceSchedulerNameBeginsWith *string

	// Specifies the maximum number of inference schedulers to list.
	MaxResults *int32

	// The name of the machine learning model used by the inference scheduler to be
	// listed.
	ModelName *string

	// An opaque pagination token indicating where to continue the listing of
	// inference schedulers.
	NextToken *string

	// Specifies the current status of the inference schedulers.
	Status types.InferenceSchedulerStatus

	noSmithyDocumentSerde
}

type ListInferenceSchedulersOutput struct {

	// Provides information about the specified inference scheduler, including data
	// upload frequency, model name and ARN, and status.
	InferenceSchedulerSummaries []types.InferenceSchedulerSummary

	// An opaque pagination token indicating where to continue the listing of
	// inference schedulers.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInferenceSchedulersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListInferenceSchedulers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListInferenceSchedulers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInferenceSchedulers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInferenceSchedulers(options.Region), middleware.Before); err != nil {
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

// ListInferenceSchedulersAPIClient is a client that implements the
// ListInferenceSchedulers operation.
type ListInferenceSchedulersAPIClient interface {
	ListInferenceSchedulers(context.Context, *ListInferenceSchedulersInput, ...func(*Options)) (*ListInferenceSchedulersOutput, error)
}

var _ ListInferenceSchedulersAPIClient = (*Client)(nil)

// ListInferenceSchedulersPaginatorOptions is the paginator options for
// ListInferenceSchedulers
type ListInferenceSchedulersPaginatorOptions struct {
	// Specifies the maximum number of inference schedulers to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInferenceSchedulersPaginator is a paginator for ListInferenceSchedulers
type ListInferenceSchedulersPaginator struct {
	options   ListInferenceSchedulersPaginatorOptions
	client    ListInferenceSchedulersAPIClient
	params    *ListInferenceSchedulersInput
	nextToken *string
	firstPage bool
}

// NewListInferenceSchedulersPaginator returns a new
// ListInferenceSchedulersPaginator
func NewListInferenceSchedulersPaginator(client ListInferenceSchedulersAPIClient, params *ListInferenceSchedulersInput, optFns ...func(*ListInferenceSchedulersPaginatorOptions)) *ListInferenceSchedulersPaginator {
	if params == nil {
		params = &ListInferenceSchedulersInput{}
	}

	options := ListInferenceSchedulersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInferenceSchedulersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInferenceSchedulersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInferenceSchedulers page.
func (p *ListInferenceSchedulersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInferenceSchedulersOutput, error) {
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

	result, err := p.client.ListInferenceSchedulers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInferenceSchedulers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInferenceSchedulers",
	}
}
