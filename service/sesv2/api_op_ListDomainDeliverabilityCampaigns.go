// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a domain
// only if you enabled the Deliverability dashboard for the domain.
func (c *Client) ListDomainDeliverabilityCampaigns(ctx context.Context, params *ListDomainDeliverabilityCampaignsInput, optFns ...func(*Options)) (*ListDomainDeliverabilityCampaignsOutput, error) {
	if params == nil {
		params = &ListDomainDeliverabilityCampaignsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomainDeliverabilityCampaigns", params, optFns, c.addOperationListDomainDeliverabilityCampaignsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainDeliverabilityCampaignsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a domain
// only if you enabled the Deliverability dashboard.
type ListDomainDeliverabilityCampaignsInput struct {

	// The last day that you want to obtain deliverability data for. This value has to
	// be less than or equal to 30 days after the value of the StartDate parameter.
	//
	// This member is required.
	EndDate *time.Time

	// The first day that you want to obtain deliverability data for.
	//
	// This member is required.
	StartDate *time.Time

	// The domain to obtain deliverability data for.
	//
	// This member is required.
	SubscribedDomain *string

	// A token that’s returned from a previous call to the
	// ListDomainDeliverabilityCampaigns operation. This token indicates the position
	// of a campaign in the list of campaigns.
	NextToken *string

	// The maximum number of results to include in response to a single call to the
	// ListDomainDeliverabilityCampaigns operation. If the number of results is larger
	// than the number that you specify in this parameter, the response includes a
	// NextToken element, which you can use to obtain additional results.
	PageSize *int32

	noSmithyDocumentSerde
}

// An array of objects that provide deliverability data for all the campaigns that
// used a specific domain to send email during a specified time range. This data is
// available for a domain only if you enabled the Deliverability dashboard for the
// domain.
type ListDomainDeliverabilityCampaignsOutput struct {

	// An array of responses, one for each campaign that used the domain to send email
	// during the specified time range.
	//
	// This member is required.
	DomainDeliverabilityCampaigns []types.DomainDeliverabilityCampaign

	// A token that’s returned from a previous call to the
	// ListDomainDeliverabilityCampaigns operation. This token indicates the position
	// of the campaign in the list of campaigns.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDomainDeliverabilityCampaignsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomainDeliverabilityCampaigns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainDeliverabilityCampaigns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDomainDeliverabilityCampaigns"); err != nil {
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
	if err = addOpListDomainDeliverabilityCampaignsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainDeliverabilityCampaigns(options.Region), middleware.Before); err != nil {
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

// ListDomainDeliverabilityCampaignsAPIClient is a client that implements the
// ListDomainDeliverabilityCampaigns operation.
type ListDomainDeliverabilityCampaignsAPIClient interface {
	ListDomainDeliverabilityCampaigns(context.Context, *ListDomainDeliverabilityCampaignsInput, ...func(*Options)) (*ListDomainDeliverabilityCampaignsOutput, error)
}

var _ ListDomainDeliverabilityCampaignsAPIClient = (*Client)(nil)

// ListDomainDeliverabilityCampaignsPaginatorOptions is the paginator options for
// ListDomainDeliverabilityCampaigns
type ListDomainDeliverabilityCampaignsPaginatorOptions struct {
	// The maximum number of results to include in response to a single call to the
	// ListDomainDeliverabilityCampaigns operation. If the number of results is larger
	// than the number that you specify in this parameter, the response includes a
	// NextToken element, which you can use to obtain additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDomainDeliverabilityCampaignsPaginator is a paginator for
// ListDomainDeliverabilityCampaigns
type ListDomainDeliverabilityCampaignsPaginator struct {
	options   ListDomainDeliverabilityCampaignsPaginatorOptions
	client    ListDomainDeliverabilityCampaignsAPIClient
	params    *ListDomainDeliverabilityCampaignsInput
	nextToken *string
	firstPage bool
}

// NewListDomainDeliverabilityCampaignsPaginator returns a new
// ListDomainDeliverabilityCampaignsPaginator
func NewListDomainDeliverabilityCampaignsPaginator(client ListDomainDeliverabilityCampaignsAPIClient, params *ListDomainDeliverabilityCampaignsInput, optFns ...func(*ListDomainDeliverabilityCampaignsPaginatorOptions)) *ListDomainDeliverabilityCampaignsPaginator {
	if params == nil {
		params = &ListDomainDeliverabilityCampaignsInput{}
	}

	options := ListDomainDeliverabilityCampaignsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDomainDeliverabilityCampaignsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDomainDeliverabilityCampaignsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDomainDeliverabilityCampaigns page.
func (p *ListDomainDeliverabilityCampaignsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDomainDeliverabilityCampaignsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListDomainDeliverabilityCampaigns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDomainDeliverabilityCampaigns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDomainDeliverabilityCampaigns",
	}
}
