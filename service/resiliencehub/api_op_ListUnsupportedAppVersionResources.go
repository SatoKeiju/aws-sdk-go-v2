// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resources that are not currently supported in Resilience Hub. An
// unsupported resource is a resource that exists in the object that was used to
// create an app, but is not supported by Resilience Hub.
func (c *Client) ListUnsupportedAppVersionResources(ctx context.Context, params *ListUnsupportedAppVersionResourcesInput, optFns ...func(*Options)) (*ListUnsupportedAppVersionResourcesOutput, error) {
	if params == nil {
		params = &ListUnsupportedAppVersionResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUnsupportedAppVersionResources", params, optFns, c.addOperationListUnsupportedAppVersionResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUnsupportedAppVersionResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUnsupportedAppVersionResourcesInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The version of the application.
	//
	// This member is required.
	AppVersion *string

	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	// The identifier for a specific resolution.
	ResolutionId *string

	noSmithyDocumentSerde
}

type ListUnsupportedAppVersionResourcesOutput struct {

	// The identifier for a specific resolution.
	//
	// This member is required.
	ResolutionId *string

	// The unsupported resources for the application.
	//
	// This member is required.
	UnsupportedResources []types.UnsupportedResource

	// Token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUnsupportedAppVersionResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUnsupportedAppVersionResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUnsupportedAppVersionResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListUnsupportedAppVersionResources"); err != nil {
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
	if err = addOpListUnsupportedAppVersionResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUnsupportedAppVersionResources(options.Region), middleware.Before); err != nil {
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

// ListUnsupportedAppVersionResourcesAPIClient is a client that implements the
// ListUnsupportedAppVersionResources operation.
type ListUnsupportedAppVersionResourcesAPIClient interface {
	ListUnsupportedAppVersionResources(context.Context, *ListUnsupportedAppVersionResourcesInput, ...func(*Options)) (*ListUnsupportedAppVersionResourcesOutput, error)
}

var _ ListUnsupportedAppVersionResourcesAPIClient = (*Client)(nil)

// ListUnsupportedAppVersionResourcesPaginatorOptions is the paginator options for
// ListUnsupportedAppVersionResources
type ListUnsupportedAppVersionResourcesPaginatorOptions struct {
	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUnsupportedAppVersionResourcesPaginator is a paginator for
// ListUnsupportedAppVersionResources
type ListUnsupportedAppVersionResourcesPaginator struct {
	options   ListUnsupportedAppVersionResourcesPaginatorOptions
	client    ListUnsupportedAppVersionResourcesAPIClient
	params    *ListUnsupportedAppVersionResourcesInput
	nextToken *string
	firstPage bool
}

// NewListUnsupportedAppVersionResourcesPaginator returns a new
// ListUnsupportedAppVersionResourcesPaginator
func NewListUnsupportedAppVersionResourcesPaginator(client ListUnsupportedAppVersionResourcesAPIClient, params *ListUnsupportedAppVersionResourcesInput, optFns ...func(*ListUnsupportedAppVersionResourcesPaginatorOptions)) *ListUnsupportedAppVersionResourcesPaginator {
	if params == nil {
		params = &ListUnsupportedAppVersionResourcesInput{}
	}

	options := ListUnsupportedAppVersionResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUnsupportedAppVersionResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUnsupportedAppVersionResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUnsupportedAppVersionResources page.
func (p *ListUnsupportedAppVersionResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUnsupportedAppVersionResourcesOutput, error) {
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

	result, err := p.client.ListUnsupportedAppVersionResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUnsupportedAppVersionResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListUnsupportedAppVersionResources",
	}
}
