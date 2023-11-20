// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all of the configuration sets associated with your account in the current
// region. Configuration sets are groups of rules that you can apply to the emails
// you send. You apply a configuration set to an email by including a reference to
// the configuration set in the headers of the email. When you apply a
// configuration set to an email, all of the rules in that configuration set are
// applied to the email.
func (c *Client) ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) {
	if params == nil {
		params = &ListConfigurationSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurationSets", params, optFns, c.addOperationListConfigurationSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain a list of configuration sets for your Amazon SES account in
// the current Amazon Web Services Region.
type ListConfigurationSetsInput struct {

	// A token returned from a previous call to ListConfigurationSets to indicate the
	// position in the list of configuration sets.
	NextToken *string

	// The number of results to show in a single call to ListConfigurationSets . If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32

	noSmithyDocumentSerde
}

// A list of configuration sets in your Amazon SES account in the current Amazon
// Web Services Region.
type ListConfigurationSetsOutput struct {

	// An array that contains all of the configuration sets in your Amazon SES account
	// in the current Amazon Web Services Region.
	ConfigurationSets []string

	// A token that indicates that there are additional configuration sets to list. To
	// view additional configuration sets, issue another request to
	// ListConfigurationSets , and pass this token in the NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConfigurationSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurationSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurationSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConfigurationSets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationSets(options.Region), middleware.Before); err != nil {
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

// ListConfigurationSetsAPIClient is a client that implements the
// ListConfigurationSets operation.
type ListConfigurationSetsAPIClient interface {
	ListConfigurationSets(context.Context, *ListConfigurationSetsInput, ...func(*Options)) (*ListConfigurationSetsOutput, error)
}

var _ ListConfigurationSetsAPIClient = (*Client)(nil)

// ListConfigurationSetsPaginatorOptions is the paginator options for
// ListConfigurationSets
type ListConfigurationSetsPaginatorOptions struct {
	// The number of results to show in a single call to ListConfigurationSets . If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConfigurationSetsPaginator is a paginator for ListConfigurationSets
type ListConfigurationSetsPaginator struct {
	options   ListConfigurationSetsPaginatorOptions
	client    ListConfigurationSetsAPIClient
	params    *ListConfigurationSetsInput
	nextToken *string
	firstPage bool
}

// NewListConfigurationSetsPaginator returns a new ListConfigurationSetsPaginator
func NewListConfigurationSetsPaginator(client ListConfigurationSetsAPIClient, params *ListConfigurationSetsInput, optFns ...func(*ListConfigurationSetsPaginatorOptions)) *ListConfigurationSetsPaginator {
	if params == nil {
		params = &ListConfigurationSetsInput{}
	}

	options := ListConfigurationSetsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConfigurationSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConfigurationSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConfigurationSets page.
func (p *ListConfigurationSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) {
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

	result, err := p.client.ListConfigurationSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConfigurationSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConfigurationSets",
	}
}
