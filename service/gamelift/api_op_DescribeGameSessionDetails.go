// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves additional game session properties, including the game session
// protection policy in force, a set of one or more game sessions in a specific
// fleet location. You can optionally filter the results by current game session
// status. This operation can be used in the following ways:
//   - To retrieve details for all game sessions that are currently running on all
//     locations in a fleet, provide a fleet or alias ID, with an optional status
//     filter. This approach returns details from the fleet's home Region and all
//     remote locations.
//   - To retrieve details for all game sessions that are currently running on a
//     specific fleet location, provide a fleet or alias ID and a location name, with
//     optional status filter. The location can be the fleet's home Region or any
//     remote location.
//   - To retrieve details for a specific game session, provide the game session
//     ID. This approach looks for the game session ID in all fleets that reside in the
//     Amazon Web Services Region defined in the request.
//
// Use the pagination parameters to retrieve results as a set of sequential pages.
// If successful, a GameSessionDetail object is returned for each game session
// that matches the request. Learn more Find a game session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-client-api.html#gamelift-sdk-client-api-find)
// All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DescribeGameSessionDetails(ctx context.Context, params *DescribeGameSessionDetailsInput, optFns ...func(*Options)) (*DescribeGameSessionDetailsOutput, error) {
	if params == nil {
		params = &DescribeGameSessionDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameSessionDetails", params, optFns, c.addOperationDescribeGameSessionDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameSessionDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGameSessionDetailsInput struct {

	// A unique identifier for the alias associated with the fleet to retrieve all
	// game sessions for. You can use either the alias ID or ARN value.
	AliasId *string

	// A unique identifier for the fleet to retrieve all game sessions active on the
	// fleet. You can use either the fleet ID or ARN value.
	FleetId *string

	// A unique identifier for the game session to retrieve.
	GameSessionId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A fleet location to get game session details for. You can specify a fleet's
	// home Region or a remote location. Use the Amazon Web Services Region code
	// format, such as us-west-2 .
	Location *string

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	// Game session status to filter results on. Possible game session statuses
	// include ACTIVE , TERMINATED , ACTIVATING and TERMINATING (the last two are
	// transitory).
	StatusFilter *string

	noSmithyDocumentSerde
}

type DescribeGameSessionDetailsOutput struct {

	// A collection of properties for each game session that matches the request.
	GameSessionDetails []types.GameSessionDetail

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGameSessionDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameSessionDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameSessionDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeGameSessionDetails"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameSessionDetails(options.Region), middleware.Before); err != nil {
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

// DescribeGameSessionDetailsAPIClient is a client that implements the
// DescribeGameSessionDetails operation.
type DescribeGameSessionDetailsAPIClient interface {
	DescribeGameSessionDetails(context.Context, *DescribeGameSessionDetailsInput, ...func(*Options)) (*DescribeGameSessionDetailsOutput, error)
}

var _ DescribeGameSessionDetailsAPIClient = (*Client)(nil)

// DescribeGameSessionDetailsPaginatorOptions is the paginator options for
// DescribeGameSessionDetails
type DescribeGameSessionDetailsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeGameSessionDetailsPaginator is a paginator for
// DescribeGameSessionDetails
type DescribeGameSessionDetailsPaginator struct {
	options   DescribeGameSessionDetailsPaginatorOptions
	client    DescribeGameSessionDetailsAPIClient
	params    *DescribeGameSessionDetailsInput
	nextToken *string
	firstPage bool
}

// NewDescribeGameSessionDetailsPaginator returns a new
// DescribeGameSessionDetailsPaginator
func NewDescribeGameSessionDetailsPaginator(client DescribeGameSessionDetailsAPIClient, params *DescribeGameSessionDetailsInput, optFns ...func(*DescribeGameSessionDetailsPaginatorOptions)) *DescribeGameSessionDetailsPaginator {
	if params == nil {
		params = &DescribeGameSessionDetailsInput{}
	}

	options := DescribeGameSessionDetailsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeGameSessionDetailsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeGameSessionDetailsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeGameSessionDetails page.
func (p *DescribeGameSessionDetailsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeGameSessionDetailsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeGameSessionDetails(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeGameSessionDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeGameSessionDetails",
	}
}
