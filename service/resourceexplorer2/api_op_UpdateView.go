// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies some of the details of a view. You can change the filter string and
// the list of included properties. You can't change the name of the view.
func (c *Client) UpdateView(ctx context.Context, params *UpdateViewInput, optFns ...func(*Options)) (*UpdateViewOutput, error) {
	if params == nil {
		params = &UpdateViewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateView", params, optFns, c.addOperationUpdateViewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateViewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateViewInput struct {

	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the view that you want to modify.
	//
	// This member is required.
	ViewArn *string

	// An array of strings that specify which resources are included in the results of
	// queries made using this view. When you use this view in a Search operation, the
	// filter string is combined with the search's QueryString parameter using a
	// logical AND operator. For information about the supported syntax, see Search
	// query reference for Resource Explorer (https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html)
	// in the Amazon Web Services Resource Explorer User Guide. This query string in
	// the context of this operation supports only filter prefixes (https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-filters)
	// with optional operators (https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-operators)
	// . It doesn't support free-form text. For example, the string region:us*
	// service:ec2 -tag:stage=prod includes all Amazon EC2 resources in any Amazon Web
	// Services Region that begins with the letters us and is not tagged with a key
	// Stage that has the value prod .
	Filters *types.SearchFilter

	// Specifies optional fields that you want included in search results from this
	// view. It is a list of objects that each describe a field to include. The default
	// is an empty list, with no optional fields included in the results.
	IncludedProperties []types.IncludedProperty

	noSmithyDocumentSerde
}

type UpdateViewOutput struct {

	// Details about the view that you changed with this operation.
	View *types.View

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateViewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateView{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateView{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateView"); err != nil {
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
	if err = addOpUpdateViewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateView(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateView(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateView",
	}
}
