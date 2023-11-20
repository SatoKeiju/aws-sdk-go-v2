// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the specified voice profile domains in the administrator's AWS account.
func (c *Client) ListVoiceProfileDomains(ctx context.Context, params *ListVoiceProfileDomainsInput, optFns ...func(*Options)) (*ListVoiceProfileDomainsOutput, error) {
	if params == nil {
		params = &ListVoiceProfileDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVoiceProfileDomains", params, optFns, c.addOperationListVoiceProfileDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVoiceProfileDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVoiceProfileDomainsInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token used to return the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVoiceProfileDomainsOutput struct {

	// The token used to return the next page of results.
	NextToken *string

	// The list of voice profile domains.
	VoiceProfileDomains []types.VoiceProfileDomainSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVoiceProfileDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVoiceProfileDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVoiceProfileDomains{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVoiceProfileDomains"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVoiceProfileDomains(options.Region), middleware.Before); err != nil {
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

// ListVoiceProfileDomainsAPIClient is a client that implements the
// ListVoiceProfileDomains operation.
type ListVoiceProfileDomainsAPIClient interface {
	ListVoiceProfileDomains(context.Context, *ListVoiceProfileDomainsInput, ...func(*Options)) (*ListVoiceProfileDomainsOutput, error)
}

var _ ListVoiceProfileDomainsAPIClient = (*Client)(nil)

// ListVoiceProfileDomainsPaginatorOptions is the paginator options for
// ListVoiceProfileDomains
type ListVoiceProfileDomainsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVoiceProfileDomainsPaginator is a paginator for ListVoiceProfileDomains
type ListVoiceProfileDomainsPaginator struct {
	options   ListVoiceProfileDomainsPaginatorOptions
	client    ListVoiceProfileDomainsAPIClient
	params    *ListVoiceProfileDomainsInput
	nextToken *string
	firstPage bool
}

// NewListVoiceProfileDomainsPaginator returns a new
// ListVoiceProfileDomainsPaginator
func NewListVoiceProfileDomainsPaginator(client ListVoiceProfileDomainsAPIClient, params *ListVoiceProfileDomainsInput, optFns ...func(*ListVoiceProfileDomainsPaginatorOptions)) *ListVoiceProfileDomainsPaginator {
	if params == nil {
		params = &ListVoiceProfileDomainsInput{}
	}

	options := ListVoiceProfileDomainsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVoiceProfileDomainsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVoiceProfileDomainsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVoiceProfileDomains page.
func (p *ListVoiceProfileDomainsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVoiceProfileDomainsOutput, error) {
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

	result, err := p.client.ListVoiceProfileDomains(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVoiceProfileDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVoiceProfileDomains",
	}
}
