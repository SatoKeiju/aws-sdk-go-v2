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

// Disassociates a CIDR block from a VPC. To disassociate the CIDR block, you must
// specify its association ID. You can get the association ID by using DescribeVpcs
// . You must detach or delete all gateways and resources that are associated with
// the CIDR block before you can disassociate it. You cannot disassociate the CIDR
// block with which you originally created the VPC (the primary CIDR block).
func (c *Client) DisassociateVpcCidrBlock(ctx context.Context, params *DisassociateVpcCidrBlockInput, optFns ...func(*Options)) (*DisassociateVpcCidrBlockOutput, error) {
	if params == nil {
		params = &DisassociateVpcCidrBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateVpcCidrBlock", params, optFns, c.addOperationDisassociateVpcCidrBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateVpcCidrBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateVpcCidrBlockInput struct {

	// The association ID for the CIDR block.
	//
	// This member is required.
	AssociationId *string

	noSmithyDocumentSerde
}

type DisassociateVpcCidrBlockOutput struct {

	// Information about the IPv4 CIDR block association.
	CidrBlockAssociation *types.VpcCidrBlockAssociation

	// Information about the IPv6 CIDR block association.
	Ipv6CidrBlockAssociation *types.VpcIpv6CidrBlockAssociation

	// The ID of the VPC.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateVpcCidrBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisassociateVpcCidrBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateVpcCidrBlock{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateVpcCidrBlock"); err != nil {
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
	if err = addOpDisassociateVpcCidrBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateVpcCidrBlock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateVpcCidrBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateVpcCidrBlock",
	}
}
