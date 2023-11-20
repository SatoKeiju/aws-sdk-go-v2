// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a transit gateway. You can use a transit gateway to interconnect your
// virtual private clouds (VPC) and on-premises networks. After the transit gateway
// enters the available state, you can attach your VPCs and VPN connections to the
// transit gateway. To attach your VPCs, use CreateTransitGatewayVpcAttachment . To
// attach a VPN connection, use CreateCustomerGateway to create a customer gateway
// and specify the ID of the customer gateway and the ID of the transit gateway in
// a call to CreateVpnConnection . When you create a transit gateway, we create a
// default transit gateway route table and use it as the default association route
// table and the default propagation route table. You can use
// CreateTransitGatewayRouteTable to create additional transit gateway route
// tables. If you disable automatic route propagation, we do not create a default
// transit gateway route table. You can use
// EnableTransitGatewayRouteTablePropagation to propagate routes from a resource
// attachment to a transit gateway route table. If you disable automatic
// associations, you can use AssociateTransitGatewayRouteTable to associate a
// resource attachment with a transit gateway route table.
func (c *Client) CreateTransitGateway(ctx context.Context, params *CreateTransitGatewayInput, optFns ...func(*Options)) (*CreateTransitGatewayOutput, error) {
	if params == nil {
		params = &CreateTransitGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransitGateway", params, optFns, c.addOperationCreateTransitGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransitGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitGatewayInput struct {

	// A description of the transit gateway.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The transit gateway options.
	Options *types.TransitGatewayRequestOptions

	// The tags to apply to the transit gateway.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateTransitGatewayOutput struct {

	// Information about the transit gateway.
	TransitGateway *types.TransitGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTransitGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTransitGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTransitGateway{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTransitGateway"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTransitGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTransitGateway",
	}
}
