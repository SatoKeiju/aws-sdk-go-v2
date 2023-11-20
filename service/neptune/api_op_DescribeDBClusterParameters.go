// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the detailed parameter list for a particular DB cluster parameter group.
func (c *Client) DescribeDBClusterParameters(ctx context.Context, params *DescribeDBClusterParametersInput, optFns ...func(*Options)) (*DescribeDBClusterParametersOutput, error) {
	if params == nil {
		params = &DescribeDBClusterParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterParameters", params, optFns, c.addOperationDescribeDBClusterParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClusterParametersInput struct {

	// The name of a specific DB cluster parameter group to return parameter details
	// for. Constraints:
	//   - If supplied, must match the name of an existing DBClusterParameterGroup.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// This parameter is not currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// A value that indicates to return only parameters for a specific source.
	// Parameter sources can be engine , service , or customer .
	Source *string

	noSmithyDocumentSerde
}

type DescribeDBClusterParametersOutput struct {

	// An optional pagination token provided by a previous DescribeDBClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// Provides a list of parameters for the DB cluster parameter group.
	Parameters []types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClusterParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBClusterParameters"); err != nil {
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
	if err = addOpDescribeDBClusterParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterParameters(options.Region), middleware.Before); err != nil {
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

// DescribeDBClusterParametersAPIClient is a client that implements the
// DescribeDBClusterParameters operation.
type DescribeDBClusterParametersAPIClient interface {
	DescribeDBClusterParameters(context.Context, *DescribeDBClusterParametersInput, ...func(*Options)) (*DescribeDBClusterParametersOutput, error)
}

var _ DescribeDBClusterParametersAPIClient = (*Client)(nil)

// DescribeDBClusterParametersPaginatorOptions is the paginator options for
// DescribeDBClusterParameters
type DescribeDBClusterParametersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClusterParametersPaginator is a paginator for
// DescribeDBClusterParameters
type DescribeDBClusterParametersPaginator struct {
	options   DescribeDBClusterParametersPaginatorOptions
	client    DescribeDBClusterParametersAPIClient
	params    *DescribeDBClusterParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClusterParametersPaginator returns a new
// DescribeDBClusterParametersPaginator
func NewDescribeDBClusterParametersPaginator(client DescribeDBClusterParametersAPIClient, params *DescribeDBClusterParametersInput, optFns ...func(*DescribeDBClusterParametersPaginatorOptions)) *DescribeDBClusterParametersPaginator {
	if params == nil {
		params = &DescribeDBClusterParametersInput{}
	}

	options := DescribeDBClusterParametersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBClusterParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClusterParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBClusterParameters page.
func (p *DescribeDBClusterParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClusterParametersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBClusterParameters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBClusterParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBClusterParameters",
	}
}
