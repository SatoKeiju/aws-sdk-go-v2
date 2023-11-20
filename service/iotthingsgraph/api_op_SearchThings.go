// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for things associated with the specified entity. You can search by
// both device and device model. For example, if two different devices, camera1 and
// camera2, implement the camera device model, the user can associate thing1 to
// camera1 and thing2 to camera2. SearchThings(camera2) will return only thing2,
// but SearchThings(camera) will return both thing1 and thing2. This action
// searches for exact matches and doesn't perform partial text matching.
//
// Deprecated: since: 2022-08-30
func (c *Client) SearchThings(ctx context.Context, params *SearchThingsInput, optFns ...func(*Options)) (*SearchThingsOutput, error) {
	if params == nil {
		params = &SearchThingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchThings", params, optFns, c.addOperationSearchThingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchThingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchThingsInput struct {

	// The ID of the entity to which the things are associated. The IDs should be in
	// the following format. urn:tdm:REGION/ACCOUNT ID/default:device:DEVICENAME
	//
	// This member is required.
	EntityId *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchThingsOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of things in the result set.
	Things []types.Thing

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchThingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchThings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchThings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchThings"); err != nil {
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
	if err = addOpSearchThingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchThings(options.Region), middleware.Before); err != nil {
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

// SearchThingsAPIClient is a client that implements the SearchThings operation.
type SearchThingsAPIClient interface {
	SearchThings(context.Context, *SearchThingsInput, ...func(*Options)) (*SearchThingsOutput, error)
}

var _ SearchThingsAPIClient = (*Client)(nil)

// SearchThingsPaginatorOptions is the paginator options for SearchThings
type SearchThingsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchThingsPaginator is a paginator for SearchThings
type SearchThingsPaginator struct {
	options   SearchThingsPaginatorOptions
	client    SearchThingsAPIClient
	params    *SearchThingsInput
	nextToken *string
	firstPage bool
}

// NewSearchThingsPaginator returns a new SearchThingsPaginator
func NewSearchThingsPaginator(client SearchThingsAPIClient, params *SearchThingsInput, optFns ...func(*SearchThingsPaginatorOptions)) *SearchThingsPaginator {
	if params == nil {
		params = &SearchThingsInput{}
	}

	options := SearchThingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchThingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchThingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchThings page.
func (p *SearchThingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchThingsOutput, error) {
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

	result, err := p.client.SearchThings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchThings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchThings",
	}
}
