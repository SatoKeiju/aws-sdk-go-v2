// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists specific permissions of users and groups with access to your Amazon
// Kendra experience.
func (c *Client) ListEntityPersonas(ctx context.Context, params *ListEntityPersonasInput, optFns ...func(*Options)) (*ListEntityPersonasOutput, error) {
	if params == nil {
		params = &ListEntityPersonasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntityPersonas", params, optFns, c.addOperationListEntityPersonasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntityPersonasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntityPersonasInput struct {

	// The identifier of your Amazon Kendra experience.
	//
	// This member is required.
	Id *string

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	// The maximum number of returned users or groups.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Kendra returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of users or groups.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEntityPersonasOutput struct {

	// If the response is truncated, Amazon Kendra returns this token, which you can
	// use in a later request to retrieve the next set of users or groups.
	NextToken *string

	// An array of summary information for one or more users or groups.
	SummaryItems []types.PersonasSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEntityPersonasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEntityPersonas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntityPersonas{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEntityPersonas"); err != nil {
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
	if err = addOpListEntityPersonasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntityPersonas(options.Region), middleware.Before); err != nil {
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

// ListEntityPersonasAPIClient is a client that implements the ListEntityPersonas
// operation.
type ListEntityPersonasAPIClient interface {
	ListEntityPersonas(context.Context, *ListEntityPersonasInput, ...func(*Options)) (*ListEntityPersonasOutput, error)
}

var _ ListEntityPersonasAPIClient = (*Client)(nil)

// ListEntityPersonasPaginatorOptions is the paginator options for
// ListEntityPersonas
type ListEntityPersonasPaginatorOptions struct {
	// The maximum number of returned users or groups.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEntityPersonasPaginator is a paginator for ListEntityPersonas
type ListEntityPersonasPaginator struct {
	options   ListEntityPersonasPaginatorOptions
	client    ListEntityPersonasAPIClient
	params    *ListEntityPersonasInput
	nextToken *string
	firstPage bool
}

// NewListEntityPersonasPaginator returns a new ListEntityPersonasPaginator
func NewListEntityPersonasPaginator(client ListEntityPersonasAPIClient, params *ListEntityPersonasInput, optFns ...func(*ListEntityPersonasPaginatorOptions)) *ListEntityPersonasPaginator {
	if params == nil {
		params = &ListEntityPersonasInput{}
	}

	options := ListEntityPersonasPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEntityPersonasPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEntityPersonasPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEntityPersonas page.
func (p *ListEntityPersonasPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEntityPersonasOutput, error) {
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

	result, err := p.client.ListEntityPersonas(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEntityPersonas(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEntityPersonas",
	}
}
