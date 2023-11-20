// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Unassigns one or more IPv6 addresses IPv4 Prefix Delegation prefixes from a
// network interface.
func (c *Client) UnassignIpv6Addresses(ctx context.Context, params *UnassignIpv6AddressesInput, optFns ...func(*Options)) (*UnassignIpv6AddressesOutput, error) {
	if params == nil {
		params = &UnassignIpv6AddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnassignIpv6Addresses", params, optFns, c.addOperationUnassignIpv6AddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnassignIpv6AddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnassignIpv6AddressesInput struct {

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// The IPv6 addresses to unassign from the network interface.
	Ipv6Addresses []string

	// The IPv6 prefixes to unassign from the network interface.
	Ipv6Prefixes []string

	noSmithyDocumentSerde
}

type UnassignIpv6AddressesOutput struct {

	// The ID of the network interface.
	NetworkInterfaceId *string

	// The IPv6 addresses that have been unassigned from the network interface.
	UnassignedIpv6Addresses []string

	// The IPv4 prefixes that have been unassigned from the network interface.
	UnassignedIpv6Prefixes []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUnassignIpv6AddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpUnassignIpv6Addresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpUnassignIpv6Addresses{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UnassignIpv6Addresses"); err != nil {
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
	if err = addOpUnassignIpv6AddressesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnassignIpv6Addresses(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUnassignIpv6Addresses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UnassignIpv6Addresses",
	}
}
