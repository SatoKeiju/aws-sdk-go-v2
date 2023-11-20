// Code generated by smithy-go-codegen DO NOT EDIT.

package trustedadvisor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/trustedadvisor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List a filterable set of Recommendations within an Organization. This API only
// supports prioritized recommendations.
func (c *Client) ListOrganizationRecommendations(ctx context.Context, params *ListOrganizationRecommendationsInput, optFns ...func(*Options)) (*ListOrganizationRecommendationsOutput, error) {
	if params == nil {
		params = &ListOrganizationRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationRecommendations", params, optFns, c.addOperationListOrganizationRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationRecommendationsInput struct {

	// After the last update of the Recommendation
	AfterLastUpdatedAt *time.Time

	// The aws service associated with the Recommendation
	AwsService *string

	// Before the last update of the Recommendation
	BeforeLastUpdatedAt *time.Time

	// The check identifier of the Recommendation
	CheckIdentifier *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The pillar of the Recommendation
	Pillar types.RecommendationPillar

	// The source of the Recommendation
	Source types.RecommendationSource

	// The status of the Recommendation
	Status types.RecommendationStatus

	// The type of the Recommendation
	Type types.RecommendationType

	noSmithyDocumentSerde
}

type ListOrganizationRecommendationsOutput struct {

	// The list of Recommendations
	//
	// This member is required.
	OrganizationRecommendationSummaries []types.OrganizationRecommendationSummary

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOrganizationRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOrganizationRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrganizationRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOrganizationRecommendations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationRecommendations(options.Region), middleware.Before); err != nil {
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

// ListOrganizationRecommendationsAPIClient is a client that implements the
// ListOrganizationRecommendations operation.
type ListOrganizationRecommendationsAPIClient interface {
	ListOrganizationRecommendations(context.Context, *ListOrganizationRecommendationsInput, ...func(*Options)) (*ListOrganizationRecommendationsOutput, error)
}

var _ ListOrganizationRecommendationsAPIClient = (*Client)(nil)

// ListOrganizationRecommendationsPaginatorOptions is the paginator options for
// ListOrganizationRecommendations
type ListOrganizationRecommendationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOrganizationRecommendationsPaginator is a paginator for
// ListOrganizationRecommendations
type ListOrganizationRecommendationsPaginator struct {
	options   ListOrganizationRecommendationsPaginatorOptions
	client    ListOrganizationRecommendationsAPIClient
	params    *ListOrganizationRecommendationsInput
	nextToken *string
	firstPage bool
}

// NewListOrganizationRecommendationsPaginator returns a new
// ListOrganizationRecommendationsPaginator
func NewListOrganizationRecommendationsPaginator(client ListOrganizationRecommendationsAPIClient, params *ListOrganizationRecommendationsInput, optFns ...func(*ListOrganizationRecommendationsPaginatorOptions)) *ListOrganizationRecommendationsPaginator {
	if params == nil {
		params = &ListOrganizationRecommendationsInput{}
	}

	options := ListOrganizationRecommendationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOrganizationRecommendationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOrganizationRecommendationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOrganizationRecommendations page.
func (p *ListOrganizationRecommendationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOrganizationRecommendationsOutput, error) {
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

	result, err := p.client.ListOrganizationRecommendations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOrganizationRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOrganizationRecommendations",
	}
}
