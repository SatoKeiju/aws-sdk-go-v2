// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of the components that a Greengrass core device
// runs. By default, this list doesn't include components that are deployed as
// dependencies of other components. To include dependencies in the response, set
// the topologyFilter parameter to ALL . IoT Greengrass relies on individual
// devices to send status updates to the Amazon Web Services Cloud. If the IoT
// Greengrass Core software isn't running on the device, or if device isn't
// connected to the Amazon Web Services Cloud, then the reported status of that
// device might not reflect its current status. The status timestamp indicates when
// the device status was last updated. Core devices send status updates at the
// following times:
//   - When the IoT Greengrass Core software starts
//   - When the core device receives a deployment from the Amazon Web Services
//     Cloud
//   - When the status of any component on the core device becomes BROKEN
//   - At a regular interval that you can configure (https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html#greengrass-nucleus-component-configuration-fss)
//     , which defaults to 24 hours
//   - For IoT Greengrass Core v2.7.0, the core device sends status updates upon
//     local deployment and cloud deployment
func (c *Client) ListInstalledComponents(ctx context.Context, params *ListInstalledComponentsInput, optFns ...func(*Options)) (*ListInstalledComponentsOutput, error) {
	if params == nil {
		params = &ListInstalledComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstalledComponents", params, optFns, c.addOperationListInstalledComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstalledComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstalledComponentsInput struct {

	// The name of the core device. This is also the name of the IoT thing.
	//
	// This member is required.
	CoreDeviceThingName *string

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The filter for the list of components. Choose from the following options:
	//   - ALL – The list includes all components installed on the core device.
	//   - ROOT – The list includes only root components, which are components that you
	//   specify in a deployment. When you choose this option, the list doesn't include
	//   components that the core device installs as dependencies of other components.
	// Default: ROOT
	TopologyFilter types.InstalledComponentTopologyFilter

	noSmithyDocumentSerde
}

type ListInstalledComponentsOutput struct {

	// A list that summarizes each component on the core device. Greengrass nucleus
	// v2.7.0 or later is required to get an accurate lastStatusChangeTimestamp
	// response. This response can be inaccurate in earlier Greengrass nucleus
	// versions. Greengrass nucleus v2.8.0 or later is required to get an accurate
	// lastInstallationSource and lastReportedTimestamp response. This response can be
	// inaccurate or null in earlier Greengrass nucleus versions.
	InstalledComponents []types.InstalledComponent

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInstalledComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInstalledComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInstalledComponents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInstalledComponents"); err != nil {
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
	if err = addOpListInstalledComponentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstalledComponents(options.Region), middleware.Before); err != nil {
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

// ListInstalledComponentsAPIClient is a client that implements the
// ListInstalledComponents operation.
type ListInstalledComponentsAPIClient interface {
	ListInstalledComponents(context.Context, *ListInstalledComponentsInput, ...func(*Options)) (*ListInstalledComponentsOutput, error)
}

var _ ListInstalledComponentsAPIClient = (*Client)(nil)

// ListInstalledComponentsPaginatorOptions is the paginator options for
// ListInstalledComponents
type ListInstalledComponentsPaginatorOptions struct {
	// The maximum number of results to be returned per paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstalledComponentsPaginator is a paginator for ListInstalledComponents
type ListInstalledComponentsPaginator struct {
	options   ListInstalledComponentsPaginatorOptions
	client    ListInstalledComponentsAPIClient
	params    *ListInstalledComponentsInput
	nextToken *string
	firstPage bool
}

// NewListInstalledComponentsPaginator returns a new
// ListInstalledComponentsPaginator
func NewListInstalledComponentsPaginator(client ListInstalledComponentsAPIClient, params *ListInstalledComponentsInput, optFns ...func(*ListInstalledComponentsPaginatorOptions)) *ListInstalledComponentsPaginator {
	if params == nil {
		params = &ListInstalledComponentsInput{}
	}

	options := ListInstalledComponentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInstalledComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstalledComponentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInstalledComponents page.
func (p *ListInstalledComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstalledComponentsOutput, error) {
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

	result, err := p.client.ListInstalledComponents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInstalledComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInstalledComponents",
	}
}
