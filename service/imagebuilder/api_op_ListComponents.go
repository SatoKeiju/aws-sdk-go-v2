// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of components that can be filtered by name, or by using the
// listed filters to streamline results. Newly created components can take up to
// two minutes to appear in the ListComponents API Results. The semantic version
// has four nodes: ../. You can assign values for the first three, and can filter
// on all of them. Filtering: With semantic versioning, you have the flexibility to
// use wildcards (x) to specify the most recent versions or nodes when selecting
// the base image or components for your recipe. When you use a wildcard in any
// node, all nodes to the right of the first wildcard must also be wildcards.
func (c *Client) ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) {
	if params == nil {
		params = &ListComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponents", params, optFns, c.addOperationListComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentsInput struct {

	// Returns the list of components for the specified name.
	ByName bool

	// Use the following filters to streamline results:
	//   - description
	//   - name
	//   - platform
	//   - supportedOsVersion
	//   - type
	//   - version
	Filters []types.Filter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	// Filters results based on the type of owner for the component. By default, this
	// request returns a list of components that your account owns. To see results for
	// other types of owners, you can specify components that Amazon manages, third
	// party components, or components that other accounts have shared with you.
	Owner types.Ownership

	noSmithyDocumentSerde
}

type ListComponentsOutput struct {

	// The list of component semantic versions. The semantic version has four nodes:
	// ../. You can assign values for the first three, and can filter on all of them.
	ComponentVersionList []types.ComponentVersion

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListComponents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListComponents"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponents(options.Region), middleware.Before); err != nil {
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

// ListComponentsAPIClient is a client that implements the ListComponents
// operation.
type ListComponentsAPIClient interface {
	ListComponents(context.Context, *ListComponentsInput, ...func(*Options)) (*ListComponentsOutput, error)
}

var _ ListComponentsAPIClient = (*Client)(nil)

// ListComponentsPaginatorOptions is the paginator options for ListComponents
type ListComponentsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComponentsPaginator is a paginator for ListComponents
type ListComponentsPaginator struct {
	options   ListComponentsPaginatorOptions
	client    ListComponentsAPIClient
	params    *ListComponentsInput
	nextToken *string
	firstPage bool
}

// NewListComponentsPaginator returns a new ListComponentsPaginator
func NewListComponentsPaginator(client ListComponentsAPIClient, params *ListComponentsInput, optFns ...func(*ListComponentsPaginatorOptions)) *ListComponentsPaginator {
	if params == nil {
		params = &ListComponentsInput{}
	}

	options := ListComponentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComponentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComponents page.
func (p *ListComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComponentsOutput, error) {
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

	result, err := p.client.ListComponents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListComponents",
	}
}
