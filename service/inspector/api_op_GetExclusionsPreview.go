// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the exclusions preview (a list of ExclusionPreview objects) specified
// by the preview token. You can obtain the preview token by running the
// CreateExclusionsPreview API.
func (c *Client) GetExclusionsPreview(ctx context.Context, params *GetExclusionsPreviewInput, optFns ...func(*Options)) (*GetExclusionsPreviewOutput, error) {
	if params == nil {
		params = &GetExclusionsPreviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExclusionsPreview", params, optFns, c.addOperationGetExclusionsPreviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExclusionsPreviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExclusionsPreviewInput struct {

	// The ARN that specifies the assessment template for which the exclusions preview
	// was requested.
	//
	// This member is required.
	AssessmentTemplateArn *string

	// The unique identifier associated of the exclusions preview.
	//
	// This member is required.
	PreviewToken *string

	// The locale into which you want to translate the exclusion's title, description,
	// and recommendation.
	Locale types.Locale

	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 100. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the GetExclusionsPreviewRequest action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// nextToken from the previous response to continue listing data.
	NextToken *string

	noSmithyDocumentSerde
}

type GetExclusionsPreviewOutput struct {

	// Specifies the status of the request to generate an exclusions preview.
	//
	// This member is required.
	PreviewStatus types.PreviewStatus

	// Information about the exclusions included in the preview.
	ExclusionPreviews []types.ExclusionPreview

	// When a response is generated, if there is more data to be listed, this
	// parameters is present in the response and contains the value to use for the
	// nextToken parameter in a subsequent pagination request. If there is no more data
	// to be listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetExclusionsPreviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetExclusionsPreview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetExclusionsPreview{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetExclusionsPreview"); err != nil {
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
	if err = addOpGetExclusionsPreviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExclusionsPreview(options.Region), middleware.Before); err != nil {
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

// GetExclusionsPreviewAPIClient is a client that implements the
// GetExclusionsPreview operation.
type GetExclusionsPreviewAPIClient interface {
	GetExclusionsPreview(context.Context, *GetExclusionsPreviewInput, ...func(*Options)) (*GetExclusionsPreviewOutput, error)
}

var _ GetExclusionsPreviewAPIClient = (*Client)(nil)

// GetExclusionsPreviewPaginatorOptions is the paginator options for
// GetExclusionsPreview
type GetExclusionsPreviewPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 100. The maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetExclusionsPreviewPaginator is a paginator for GetExclusionsPreview
type GetExclusionsPreviewPaginator struct {
	options   GetExclusionsPreviewPaginatorOptions
	client    GetExclusionsPreviewAPIClient
	params    *GetExclusionsPreviewInput
	nextToken *string
	firstPage bool
}

// NewGetExclusionsPreviewPaginator returns a new GetExclusionsPreviewPaginator
func NewGetExclusionsPreviewPaginator(client GetExclusionsPreviewAPIClient, params *GetExclusionsPreviewInput, optFns ...func(*GetExclusionsPreviewPaginatorOptions)) *GetExclusionsPreviewPaginator {
	if params == nil {
		params = &GetExclusionsPreviewInput{}
	}

	options := GetExclusionsPreviewPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetExclusionsPreviewPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetExclusionsPreviewPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetExclusionsPreview page.
func (p *GetExclusionsPreviewPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetExclusionsPreviewOutput, error) {
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

	result, err := p.client.GetExclusionsPreview(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetExclusionsPreview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetExclusionsPreview",
	}
}
