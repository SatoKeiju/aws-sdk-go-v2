// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the partition indexes associated with a table.
func (c *Client) GetPartitionIndexes(ctx context.Context, params *GetPartitionIndexesInput, optFns ...func(*Options)) (*GetPartitionIndexesOutput, error) {
	if params == nil {
		params = &GetPartitionIndexesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPartitionIndexes", params, optFns, c.addOperationGetPartitionIndexesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPartitionIndexesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPartitionIndexesInput struct {

	// Specifies the name of a database from which you want to retrieve partition
	// indexes.
	//
	// This member is required.
	DatabaseName *string

	// Specifies the name of a table for which you want to retrieve the partition
	// indexes.
	//
	// This member is required.
	TableName *string

	// The catalog ID where the table resides.
	CatalogId *string

	// A continuation token, included if this is a continuation call.
	NextToken *string

	noSmithyDocumentSerde
}

type GetPartitionIndexesOutput struct {

	// A continuation token, present if the current list segment is not the last.
	NextToken *string

	// A list of index descriptors.
	PartitionIndexDescriptorList []types.PartitionIndexDescriptor

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPartitionIndexesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPartitionIndexes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPartitionIndexes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPartitionIndexes"); err != nil {
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
	if err = addOpGetPartitionIndexesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPartitionIndexes(options.Region), middleware.Before); err != nil {
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

// GetPartitionIndexesAPIClient is a client that implements the
// GetPartitionIndexes operation.
type GetPartitionIndexesAPIClient interface {
	GetPartitionIndexes(context.Context, *GetPartitionIndexesInput, ...func(*Options)) (*GetPartitionIndexesOutput, error)
}

var _ GetPartitionIndexesAPIClient = (*Client)(nil)

// GetPartitionIndexesPaginatorOptions is the paginator options for
// GetPartitionIndexes
type GetPartitionIndexesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetPartitionIndexesPaginator is a paginator for GetPartitionIndexes
type GetPartitionIndexesPaginator struct {
	options   GetPartitionIndexesPaginatorOptions
	client    GetPartitionIndexesAPIClient
	params    *GetPartitionIndexesInput
	nextToken *string
	firstPage bool
}

// NewGetPartitionIndexesPaginator returns a new GetPartitionIndexesPaginator
func NewGetPartitionIndexesPaginator(client GetPartitionIndexesAPIClient, params *GetPartitionIndexesInput, optFns ...func(*GetPartitionIndexesPaginatorOptions)) *GetPartitionIndexesPaginator {
	if params == nil {
		params = &GetPartitionIndexesInput{}
	}

	options := GetPartitionIndexesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetPartitionIndexesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetPartitionIndexesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetPartitionIndexes page.
func (p *GetPartitionIndexesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetPartitionIndexesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetPartitionIndexes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetPartitionIndexes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPartitionIndexes",
	}
}
