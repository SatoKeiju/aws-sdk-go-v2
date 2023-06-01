// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches address books and lists the ones that meet a set of filter and sort
// criteria.
//
// Deprecated: Alexa For Business is no longer supported
func (c *Client) SearchAddressBooks(ctx context.Context, params *SearchAddressBooksInput, optFns ...func(*Options)) (*SearchAddressBooksOutput, error) {
	if params == nil {
		params = &SearchAddressBooksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchAddressBooks", params, optFns, c.addOperationSearchAddressBooksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchAddressBooksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchAddressBooksInput struct {

	// The filters to use to list a specified set of address books. The supported
	// filter key is AddressBookName.
	Filters []types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response only
	// includes results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The sort order to use in listing the specified set of address books. The
	// supported sort key is AddressBookName.
	SortCriteria []types.Sort

	noSmithyDocumentSerde
}

type SearchAddressBooksOutput struct {

	// The address books that meet the specified set of filter criteria, in sort order.
	AddressBooks []types.AddressBookData

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The total number of address books returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchAddressBooksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchAddressBooks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchAddressBooks{}, middleware.After)
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
	if err = addOpSearchAddressBooksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchAddressBooks(options.Region), middleware.Before); err != nil {
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

// SearchAddressBooksAPIClient is a client that implements the SearchAddressBooks
// operation.
type SearchAddressBooksAPIClient interface {
	SearchAddressBooks(context.Context, *SearchAddressBooksInput, ...func(*Options)) (*SearchAddressBooksOutput, error)
}

var _ SearchAddressBooksAPIClient = (*Client)(nil)

// SearchAddressBooksPaginatorOptions is the paginator options for
// SearchAddressBooks
type SearchAddressBooksPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchAddressBooksPaginator is a paginator for SearchAddressBooks
type SearchAddressBooksPaginator struct {
	options   SearchAddressBooksPaginatorOptions
	client    SearchAddressBooksAPIClient
	params    *SearchAddressBooksInput
	nextToken *string
	firstPage bool
}

// NewSearchAddressBooksPaginator returns a new SearchAddressBooksPaginator
func NewSearchAddressBooksPaginator(client SearchAddressBooksAPIClient, params *SearchAddressBooksInput, optFns ...func(*SearchAddressBooksPaginatorOptions)) *SearchAddressBooksPaginator {
	if params == nil {
		params = &SearchAddressBooksInput{}
	}

	options := SearchAddressBooksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchAddressBooksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchAddressBooksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchAddressBooks page.
func (p *SearchAddressBooksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchAddressBooksOutput, error) {
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

	result, err := p.client.SearchAddressBooks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchAddressBooks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchAddressBooks",
	}
}
