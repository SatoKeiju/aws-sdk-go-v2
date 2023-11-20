// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the list of all inference experiments.
func (c *Client) ListInferenceExperiments(ctx context.Context, params *ListInferenceExperimentsInput, optFns ...func(*Options)) (*ListInferenceExperimentsOutput, error) {
	if params == nil {
		params = &ListInferenceExperimentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInferenceExperiments", params, optFns, c.addOperationListInferenceExperimentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInferenceExperimentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInferenceExperimentsInput struct {

	// Selects inference experiments which were created after this timestamp.
	CreationTimeAfter *time.Time

	// Selects inference experiments which were created before this timestamp.
	CreationTimeBefore *time.Time

	// Selects inference experiments which were last modified after this timestamp.
	LastModifiedTimeAfter *time.Time

	// Selects inference experiments which were last modified before this timestamp.
	LastModifiedTimeBefore *time.Time

	// The maximum number of results to select.
	MaxResults *int32

	// Selects inference experiments whose names contain this name.
	NameContains *string

	// The response from the last list when returning a list large enough to need
	// tokening.
	NextToken *string

	// The column by which to sort the listed inference experiments.
	SortBy types.SortInferenceExperimentsBy

	// The direction of sorting (ascending or descending).
	SortOrder types.SortOrder

	// Selects inference experiments which are in this status. For the possible
	// statuses, see DescribeInferenceExperiment (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeInferenceExperiment.html)
	// .
	StatusEquals types.InferenceExperimentStatus

	// Selects inference experiments of this type. For the possible types of inference
	// experiments, see CreateInferenceExperiment (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateInferenceExperiment.html)
	// .
	Type types.InferenceExperimentType

	noSmithyDocumentSerde
}

type ListInferenceExperimentsOutput struct {

	// List of inference experiments.
	InferenceExperiments []types.InferenceExperimentSummary

	// The token to use when calling the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInferenceExperimentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInferenceExperiments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInferenceExperiments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInferenceExperiments"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInferenceExperiments(options.Region), middleware.Before); err != nil {
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

// ListInferenceExperimentsAPIClient is a client that implements the
// ListInferenceExperiments operation.
type ListInferenceExperimentsAPIClient interface {
	ListInferenceExperiments(context.Context, *ListInferenceExperimentsInput, ...func(*Options)) (*ListInferenceExperimentsOutput, error)
}

var _ ListInferenceExperimentsAPIClient = (*Client)(nil)

// ListInferenceExperimentsPaginatorOptions is the paginator options for
// ListInferenceExperiments
type ListInferenceExperimentsPaginatorOptions struct {
	// The maximum number of results to select.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInferenceExperimentsPaginator is a paginator for ListInferenceExperiments
type ListInferenceExperimentsPaginator struct {
	options   ListInferenceExperimentsPaginatorOptions
	client    ListInferenceExperimentsAPIClient
	params    *ListInferenceExperimentsInput
	nextToken *string
	firstPage bool
}

// NewListInferenceExperimentsPaginator returns a new
// ListInferenceExperimentsPaginator
func NewListInferenceExperimentsPaginator(client ListInferenceExperimentsAPIClient, params *ListInferenceExperimentsInput, optFns ...func(*ListInferenceExperimentsPaginatorOptions)) *ListInferenceExperimentsPaginator {
	if params == nil {
		params = &ListInferenceExperimentsInput{}
	}

	options := ListInferenceExperimentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInferenceExperimentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInferenceExperimentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInferenceExperiments page.
func (p *ListInferenceExperimentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInferenceExperimentsOutput, error) {
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

	result, err := p.client.ListInferenceExperiments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInferenceExperiments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInferenceExperiments",
	}
}
