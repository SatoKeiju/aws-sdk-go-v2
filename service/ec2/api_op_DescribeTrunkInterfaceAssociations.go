// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more network interface trunk associations.
func (c *Client) DescribeTrunkInterfaceAssociations(ctx context.Context, params *DescribeTrunkInterfaceAssociationsInput, optFns ...func(*Options)) (*DescribeTrunkInterfaceAssociationsOutput, error) {
	if params == nil {
		params = &DescribeTrunkInterfaceAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrunkInterfaceAssociations", params, optFns, c.addOperationDescribeTrunkInterfaceAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrunkInterfaceAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrunkInterfaceAssociationsInput struct {

	// The IDs of the associations.
	AssociationIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//   - gre-key - The ID of a trunk interface association.
	//   - interface-protocol - The interface protocol. Valid values are VLAN and GRE .
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeTrunkInterfaceAssociationsOutput struct {

	// Information about the trunk associations.
	InterfaceAssociations []types.TrunkInterfaceAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrunkInterfaceAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTrunkInterfaceAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTrunkInterfaceAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTrunkInterfaceAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrunkInterfaceAssociations(options.Region), middleware.Before); err != nil {
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

// DescribeTrunkInterfaceAssociationsAPIClient is a client that implements the
// DescribeTrunkInterfaceAssociations operation.
type DescribeTrunkInterfaceAssociationsAPIClient interface {
	DescribeTrunkInterfaceAssociations(context.Context, *DescribeTrunkInterfaceAssociationsInput, ...func(*Options)) (*DescribeTrunkInterfaceAssociationsOutput, error)
}

var _ DescribeTrunkInterfaceAssociationsAPIClient = (*Client)(nil)

// DescribeTrunkInterfaceAssociationsPaginatorOptions is the paginator options for
// DescribeTrunkInterfaceAssociations
type DescribeTrunkInterfaceAssociationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTrunkInterfaceAssociationsPaginator is a paginator for
// DescribeTrunkInterfaceAssociations
type DescribeTrunkInterfaceAssociationsPaginator struct {
	options   DescribeTrunkInterfaceAssociationsPaginatorOptions
	client    DescribeTrunkInterfaceAssociationsAPIClient
	params    *DescribeTrunkInterfaceAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTrunkInterfaceAssociationsPaginator returns a new
// DescribeTrunkInterfaceAssociationsPaginator
func NewDescribeTrunkInterfaceAssociationsPaginator(client DescribeTrunkInterfaceAssociationsAPIClient, params *DescribeTrunkInterfaceAssociationsInput, optFns ...func(*DescribeTrunkInterfaceAssociationsPaginatorOptions)) *DescribeTrunkInterfaceAssociationsPaginator {
	if params == nil {
		params = &DescribeTrunkInterfaceAssociationsInput{}
	}

	options := DescribeTrunkInterfaceAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTrunkInterfaceAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTrunkInterfaceAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTrunkInterfaceAssociations page.
func (p *DescribeTrunkInterfaceAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTrunkInterfaceAssociationsOutput, error) {
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

	result, err := p.client.DescribeTrunkInterfaceAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTrunkInterfaceAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTrunkInterfaceAssociations",
	}
}
