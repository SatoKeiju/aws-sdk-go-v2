// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists blueprint configurations for a Amazon DataZone environment.
func (c *Client) ListEnvironmentBlueprintConfigurations(ctx context.Context, params *ListEnvironmentBlueprintConfigurationsInput, optFns ...func(*Options)) (*ListEnvironmentBlueprintConfigurationsOutput, error) {
	if params == nil {
		params = &ListEnvironmentBlueprintConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironmentBlueprintConfigurations", params, optFns, c.addOperationListEnvironmentBlueprintConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentBlueprintConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentBlueprintConfigurationsInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// The maximum number of blueprint configurations to return in a single call to
	// ListEnvironmentBlueprintConfigurations . When the number of configurations to be
	// listed is greater than the value of MaxResults , the response contains a
	// NextToken value that you can use in a subsequent call to
	// ListEnvironmentBlueprintConfigurations to list the next set of configurations.
	MaxResults *int32

	// When the number of blueprint configurations is greater than the default value
	// for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of configurations, the response includes
	// a pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListEnvironmentBlueprintConfigurations to list the next set
	// of configurations.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnvironmentBlueprintConfigurationsOutput struct {

	// The results of the ListEnvironmentBlueprintConfigurations action.
	Items []types.EnvironmentBlueprintConfigurationItem

	// When the number of blueprint configurations is greater than the default value
	// for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of configurations, the response includes
	// a pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListEnvironmentBlueprintConfigurations to list the next set
	// of configurations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentBlueprintConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnvironmentBlueprintConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnvironmentBlueprintConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEnvironmentBlueprintConfigurations"); err != nil {
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
	if err = addOpListEnvironmentBlueprintConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironmentBlueprintConfigurations(options.Region), middleware.Before); err != nil {
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

// ListEnvironmentBlueprintConfigurationsAPIClient is a client that implements the
// ListEnvironmentBlueprintConfigurations operation.
type ListEnvironmentBlueprintConfigurationsAPIClient interface {
	ListEnvironmentBlueprintConfigurations(context.Context, *ListEnvironmentBlueprintConfigurationsInput, ...func(*Options)) (*ListEnvironmentBlueprintConfigurationsOutput, error)
}

var _ ListEnvironmentBlueprintConfigurationsAPIClient = (*Client)(nil)

// ListEnvironmentBlueprintConfigurationsPaginatorOptions is the paginator options
// for ListEnvironmentBlueprintConfigurations
type ListEnvironmentBlueprintConfigurationsPaginatorOptions struct {
	// The maximum number of blueprint configurations to return in a single call to
	// ListEnvironmentBlueprintConfigurations . When the number of configurations to be
	// listed is greater than the value of MaxResults , the response contains a
	// NextToken value that you can use in a subsequent call to
	// ListEnvironmentBlueprintConfigurations to list the next set of configurations.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentBlueprintConfigurationsPaginator is a paginator for
// ListEnvironmentBlueprintConfigurations
type ListEnvironmentBlueprintConfigurationsPaginator struct {
	options   ListEnvironmentBlueprintConfigurationsPaginatorOptions
	client    ListEnvironmentBlueprintConfigurationsAPIClient
	params    *ListEnvironmentBlueprintConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentBlueprintConfigurationsPaginator returns a new
// ListEnvironmentBlueprintConfigurationsPaginator
func NewListEnvironmentBlueprintConfigurationsPaginator(client ListEnvironmentBlueprintConfigurationsAPIClient, params *ListEnvironmentBlueprintConfigurationsInput, optFns ...func(*ListEnvironmentBlueprintConfigurationsPaginatorOptions)) *ListEnvironmentBlueprintConfigurationsPaginator {
	if params == nil {
		params = &ListEnvironmentBlueprintConfigurationsInput{}
	}

	options := ListEnvironmentBlueprintConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentBlueprintConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentBlueprintConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironmentBlueprintConfigurations page.
func (p *ListEnvironmentBlueprintConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentBlueprintConfigurationsOutput, error) {
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

	result, err := p.client.ListEnvironmentBlueprintConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnvironmentBlueprintConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEnvironmentBlueprintConfigurations",
	}
}
