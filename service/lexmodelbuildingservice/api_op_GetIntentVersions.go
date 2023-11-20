// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about all of the versions of an intent. The GetIntentVersions
// operation returns an IntentMetadata object for each version of an intent. For
// example, if an intent has three numbered versions, the GetIntentVersions
// operation returns four IntentMetadata objects in the response, one for each
// numbered version and one for the $LATEST version. The GetIntentVersions
// operation always returns at least one version, the $LATEST version. This
// operation requires permissions for the lex:GetIntentVersions action.
func (c *Client) GetIntentVersions(ctx context.Context, params *GetIntentVersionsInput, optFns ...func(*Options)) (*GetIntentVersionsOutput, error) {
	if params == nil {
		params = &GetIntentVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIntentVersions", params, optFns, c.addOperationGetIntentVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIntentVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntentVersionsInput struct {

	// The name of the intent for which versions should be returned.
	//
	// This member is required.
	Name *string

	// The maximum number of intent versions to return in the response. The default is
	// 10.
	MaxResults *int32

	// A pagination token for fetching the next page of intent versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetIntentVersionsOutput struct {

	// An array of IntentMetadata objects, one for each numbered version of the intent
	// plus one for the $LATEST version.
	Intents []types.IntentMetadata

	// A pagination token for fetching the next page of intent versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIntentVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIntentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIntentVersions"); err != nil {
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
	if err = addOpGetIntentVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntentVersions(options.Region), middleware.Before); err != nil {
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

// GetIntentVersionsAPIClient is a client that implements the GetIntentVersions
// operation.
type GetIntentVersionsAPIClient interface {
	GetIntentVersions(context.Context, *GetIntentVersionsInput, ...func(*Options)) (*GetIntentVersionsOutput, error)
}

var _ GetIntentVersionsAPIClient = (*Client)(nil)

// GetIntentVersionsPaginatorOptions is the paginator options for GetIntentVersions
type GetIntentVersionsPaginatorOptions struct {
	// The maximum number of intent versions to return in the response. The default is
	// 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetIntentVersionsPaginator is a paginator for GetIntentVersions
type GetIntentVersionsPaginator struct {
	options   GetIntentVersionsPaginatorOptions
	client    GetIntentVersionsAPIClient
	params    *GetIntentVersionsInput
	nextToken *string
	firstPage bool
}

// NewGetIntentVersionsPaginator returns a new GetIntentVersionsPaginator
func NewGetIntentVersionsPaginator(client GetIntentVersionsAPIClient, params *GetIntentVersionsInput, optFns ...func(*GetIntentVersionsPaginatorOptions)) *GetIntentVersionsPaginator {
	if params == nil {
		params = &GetIntentVersionsInput{}
	}

	options := GetIntentVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetIntentVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetIntentVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetIntentVersions page.
func (p *GetIntentVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetIntentVersionsOutput, error) {
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

	result, err := p.client.GetIntentVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetIntentVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIntentVersions",
	}
}
