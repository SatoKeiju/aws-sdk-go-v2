// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the associations between your Direct Connect gateways and virtual private
// gateways and transit gateways. You must specify one of the following:
//   - A Direct Connect gateway The response contains all virtual private gateways
//     and transit gateways associated with the Direct Connect gateway.
//   - A virtual private gateway The response contains the Direct Connect gateway.
//   - A transit gateway The response contains the Direct Connect gateway.
//   - A Direct Connect gateway and a virtual private gateway The response
//     contains the association between the Direct Connect gateway and virtual private
//     gateway.
//   - A Direct Connect gateway and a transit gateway The response contains the
//     association between the Direct Connect gateway and transit gateway.
func (c *Client) DescribeDirectConnectGatewayAssociations(ctx context.Context, params *DescribeDirectConnectGatewayAssociationsInput, optFns ...func(*Options)) (*DescribeDirectConnectGatewayAssociationsOutput, error) {
	if params == nil {
		params = &DescribeDirectConnectGatewayAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDirectConnectGatewayAssociations", params, optFns, c.addOperationDescribeDirectConnectGatewayAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDirectConnectGatewayAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDirectConnectGatewayAssociationsInput struct {

	// The ID of the associated gateway.
	AssociatedGatewayId *string

	// The ID of the Direct Connect gateway association.
	AssociationId *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. If
	// MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int32

	// The token provided in the previous call to retrieve the next page.
	NextToken *string

	// The ID of the virtual private gateway or transit gateway.
	VirtualGatewayId *string

	noSmithyDocumentSerde
}

type DescribeDirectConnectGatewayAssociationsOutput struct {

	// Information about the associations.
	DirectConnectGatewayAssociations []types.DirectConnectGatewayAssociation

	// The token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDirectConnectGatewayAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDirectConnectGatewayAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDirectConnectGatewayAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDirectConnectGatewayAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDirectConnectGatewayAssociations",
	}
}
