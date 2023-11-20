// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of browser settings.
func (c *Client) ListBrowserSettings(ctx context.Context, params *ListBrowserSettingsInput, optFns ...func(*Options)) (*ListBrowserSettingsOutput, error) {
	if params == nil {
		params = &ListBrowserSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBrowserSettings", params, optFns, c.addOperationListBrowserSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBrowserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBrowserSettingsInput struct {

	// The maximum number of results to be included in the next page.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBrowserSettingsOutput struct {

	// The browser settings.
	BrowserSettings []types.BrowserSettingsSummary

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBrowserSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBrowserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBrowserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBrowserSettings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBrowserSettings(options.Region), middleware.Before); err != nil {
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

// ListBrowserSettingsAPIClient is a client that implements the
// ListBrowserSettings operation.
type ListBrowserSettingsAPIClient interface {
	ListBrowserSettings(context.Context, *ListBrowserSettingsInput, ...func(*Options)) (*ListBrowserSettingsOutput, error)
}

var _ ListBrowserSettingsAPIClient = (*Client)(nil)

// ListBrowserSettingsPaginatorOptions is the paginator options for
// ListBrowserSettings
type ListBrowserSettingsPaginatorOptions struct {
	// The maximum number of results to be included in the next page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBrowserSettingsPaginator is a paginator for ListBrowserSettings
type ListBrowserSettingsPaginator struct {
	options   ListBrowserSettingsPaginatorOptions
	client    ListBrowserSettingsAPIClient
	params    *ListBrowserSettingsInput
	nextToken *string
	firstPage bool
}

// NewListBrowserSettingsPaginator returns a new ListBrowserSettingsPaginator
func NewListBrowserSettingsPaginator(client ListBrowserSettingsAPIClient, params *ListBrowserSettingsInput, optFns ...func(*ListBrowserSettingsPaginatorOptions)) *ListBrowserSettingsPaginator {
	if params == nil {
		params = &ListBrowserSettingsInput{}
	}

	options := ListBrowserSettingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBrowserSettingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBrowserSettingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBrowserSettings page.
func (p *ListBrowserSettingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBrowserSettingsOutput, error) {
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

	result, err := p.client.ListBrowserSettings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBrowserSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBrowserSettings",
	}
}
