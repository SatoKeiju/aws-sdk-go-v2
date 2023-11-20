// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of schemas detected by Fleet Advisor Collectors in your account.
func (c *Client) DescribeFleetAdvisorSchemas(ctx context.Context, params *DescribeFleetAdvisorSchemasInput, optFns ...func(*Options)) (*DescribeFleetAdvisorSchemasOutput, error) {
	if params == nil {
		params = &DescribeFleetAdvisorSchemasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetAdvisorSchemas", params, optFns, c.addOperationDescribeFleetAdvisorSchemasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetAdvisorSchemasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetAdvisorSchemasInput struct {

	// If you specify any of the following filters, the output includes information
	// for only those schemas that meet the filter criteria:
	//   - complexity – The schema's complexity, for example Simple .
	//   - database-id – The ID of the schema's database.
	//   - database-ip-address – The IP address of the schema's database.
	//   - database-name – The name of the schema's database.
	//   - database-engine – The name of the schema database's engine.
	//   - original-schema-name – The name of the schema's database's main schema.
	//   - schema-id – The ID of the schema, for example 15 .
	//   - schema-name – The name of the schema.
	//   - server-ip-address – The IP address of the schema database's server.
	// An example is: describe-fleet-advisor-schemas --filter
	// Name="schema-id",Values="50"
	Filters []types.Filter

	// Sets the maximum number of records returned in the response.
	MaxRecords *int32

	// If NextToken is returned by a previous response, there are more results
	// available. The value of NextToken is a unique pagination token for each page.
	// Make the call again using the returned token to retrieve the next page. Keep all
	// other arguments unchanged.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetAdvisorSchemasOutput struct {

	// A collection of SchemaResponse objects.
	FleetAdvisorSchemas []types.SchemaResponse

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetAdvisorSchemasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetAdvisorSchemas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetAdvisorSchemas{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleetAdvisorSchemas"); err != nil {
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
	if err = addOpDescribeFleetAdvisorSchemasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetAdvisorSchemas(options.Region), middleware.Before); err != nil {
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

// DescribeFleetAdvisorSchemasAPIClient is a client that implements the
// DescribeFleetAdvisorSchemas operation.
type DescribeFleetAdvisorSchemasAPIClient interface {
	DescribeFleetAdvisorSchemas(context.Context, *DescribeFleetAdvisorSchemasInput, ...func(*Options)) (*DescribeFleetAdvisorSchemasOutput, error)
}

var _ DescribeFleetAdvisorSchemasAPIClient = (*Client)(nil)

// DescribeFleetAdvisorSchemasPaginatorOptions is the paginator options for
// DescribeFleetAdvisorSchemas
type DescribeFleetAdvisorSchemasPaginatorOptions struct {
	// Sets the maximum number of records returned in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetAdvisorSchemasPaginator is a paginator for
// DescribeFleetAdvisorSchemas
type DescribeFleetAdvisorSchemasPaginator struct {
	options   DescribeFleetAdvisorSchemasPaginatorOptions
	client    DescribeFleetAdvisorSchemasAPIClient
	params    *DescribeFleetAdvisorSchemasInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetAdvisorSchemasPaginator returns a new
// DescribeFleetAdvisorSchemasPaginator
func NewDescribeFleetAdvisorSchemasPaginator(client DescribeFleetAdvisorSchemasAPIClient, params *DescribeFleetAdvisorSchemasInput, optFns ...func(*DescribeFleetAdvisorSchemasPaginatorOptions)) *DescribeFleetAdvisorSchemasPaginator {
	if params == nil {
		params = &DescribeFleetAdvisorSchemasInput{}
	}

	options := DescribeFleetAdvisorSchemasPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetAdvisorSchemasPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetAdvisorSchemasPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetAdvisorSchemas page.
func (p *DescribeFleetAdvisorSchemasPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetAdvisorSchemasOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeFleetAdvisorSchemas(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFleetAdvisorSchemas(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleetAdvisorSchemas",
	}
}
