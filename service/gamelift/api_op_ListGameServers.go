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

// This operation is used with the Amazon GameLift FleetIQ solution and game
// server groups. Retrieves information on all game servers that are currently
// active in a specified game server group. You can opt to sort the list by game
// server age. Use the pagination parameters to retrieve results in a set of
// sequential segments. Learn more Amazon GameLift FleetIQ Guide (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
func (c *Client) ListGameServers(ctx context.Context, params *ListGameServersInput, optFns ...func(*Options)) (*ListGameServersOutput, error) {
	if params == nil {
		params = &ListGameServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGameServers", params, optFns, c.addOperationListGameServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGameServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGameServersInput struct {

	// An identifier for the game server group to retrieve a list of game servers
	// from. Use either the name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	// Indicates how to sort the returned data based on game server registration
	// timestamp. Use ASCENDING to retrieve oldest game servers first, or use
	// DESCENDING to retrieve newest game servers first. If this parameter is left
	// empty, game servers are returned in no particular order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListGameServersOutput struct {

	// A collection of game server objects that match the request.
	GameServers []types.GameServer

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGameServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGameServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGameServers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGameServers"); err != nil {
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
	if err = addOpListGameServersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGameServers(options.Region), middleware.Before); err != nil {
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

// ListGameServersAPIClient is a client that implements the ListGameServers
// operation.
type ListGameServersAPIClient interface {
	ListGameServers(context.Context, *ListGameServersInput, ...func(*Options)) (*ListGameServersOutput, error)
}

var _ ListGameServersAPIClient = (*Client)(nil)

// ListGameServersPaginatorOptions is the paginator options for ListGameServers
type ListGameServersPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGameServersPaginator is a paginator for ListGameServers
type ListGameServersPaginator struct {
	options   ListGameServersPaginatorOptions
	client    ListGameServersAPIClient
	params    *ListGameServersInput
	nextToken *string
	firstPage bool
}

// NewListGameServersPaginator returns a new ListGameServersPaginator
func NewListGameServersPaginator(client ListGameServersAPIClient, params *ListGameServersInput, optFns ...func(*ListGameServersPaginatorOptions)) *ListGameServersPaginator {
	if params == nil {
		params = &ListGameServersInput{}
	}

	options := ListGameServersPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGameServersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGameServersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGameServers page.
func (p *ListGameServersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGameServersOutput, error) {
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

	result, err := p.client.ListGameServers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGameServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGameServers",
	}
}
