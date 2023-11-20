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

// Accepts a proposal request to attach a virtual private gateway or transit
// gateway to a Direct Connect gateway.
func (c *Client) AcceptDirectConnectGatewayAssociationProposal(ctx context.Context, params *AcceptDirectConnectGatewayAssociationProposalInput, optFns ...func(*Options)) (*AcceptDirectConnectGatewayAssociationProposalOutput, error) {
	if params == nil {
		params = &AcceptDirectConnectGatewayAssociationProposalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptDirectConnectGatewayAssociationProposal", params, optFns, c.addOperationAcceptDirectConnectGatewayAssociationProposalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptDirectConnectGatewayAssociationProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptDirectConnectGatewayAssociationProposalInput struct {

	// The ID of the Amazon Web Services account that owns the virtual private gateway
	// or transit gateway.
	//
	// This member is required.
	AssociatedGatewayOwnerAccount *string

	// The ID of the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayId *string

	// The ID of the request proposal.
	//
	// This member is required.
	ProposalId *string

	// Overrides the Amazon VPC prefixes advertised to the Direct Connect gateway. For
	// information about how to set the prefixes, see Allowed Prefixes (https://docs.aws.amazon.com/directconnect/latest/UserGuide/multi-account-associate-vgw.html#allowed-prefixes)
	// in the Direct Connect User Guide.
	OverrideAllowedPrefixesToDirectConnectGateway []types.RouteFilterPrefix

	noSmithyDocumentSerde
}

type AcceptDirectConnectGatewayAssociationProposalOutput struct {

	// Information about an association between a Direct Connect gateway and a virtual
	// private gateway or transit gateway.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptDirectConnectGatewayAssociationProposalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptDirectConnectGatewayAssociationProposal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptDirectConnectGatewayAssociationProposal{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AcceptDirectConnectGatewayAssociationProposal"); err != nil {
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
	if err = addOpAcceptDirectConnectGatewayAssociationProposalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptDirectConnectGatewayAssociationProposal(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptDirectConnectGatewayAssociationProposal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AcceptDirectConnectGatewayAssociationProposal",
	}
}
