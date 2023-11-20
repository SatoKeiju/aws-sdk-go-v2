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

// Accepts a request to associate subnets with a transit gateway multicast domain.
func (c *Client) AcceptTransitGatewayMulticastDomainAssociations(ctx context.Context, params *AcceptTransitGatewayMulticastDomainAssociationsInput, optFns ...func(*Options)) (*AcceptTransitGatewayMulticastDomainAssociationsOutput, error) {
	if params == nil {
		params = &AcceptTransitGatewayMulticastDomainAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptTransitGatewayMulticastDomainAssociations", params, optFns, c.addOperationAcceptTransitGatewayMulticastDomainAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptTransitGatewayMulticastDomainAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptTransitGatewayMulticastDomainAssociationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The IDs of the subnets to associate with the transit gateway multicast domain.
	SubnetIds []string

	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId *string

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string

	noSmithyDocumentSerde
}

type AcceptTransitGatewayMulticastDomainAssociationsOutput struct {

	// Information about the multicast domain associations.
	Associations *types.TransitGatewayMulticastDomainAssociations

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptTransitGatewayMulticastDomainAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpAcceptTransitGatewayMulticastDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAcceptTransitGatewayMulticastDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AcceptTransitGatewayMulticastDomainAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptTransitGatewayMulticastDomainAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptTransitGatewayMulticastDomainAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AcceptTransitGatewayMulticastDomainAssociations",
	}
}
