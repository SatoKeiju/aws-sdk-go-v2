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

// Terminates active Client VPN endpoint connections. This action can be used to
// terminate a specific client connection, or up to five connections established by
// a specific user.
func (c *Client) TerminateClientVpnConnections(ctx context.Context, params *TerminateClientVpnConnectionsInput, optFns ...func(*Options)) (*TerminateClientVpnConnectionsOutput, error) {
	if params == nil {
		params = &TerminateClientVpnConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateClientVpnConnections", params, optFns, c.addOperationTerminateClientVpnConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateClientVpnConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateClientVpnConnectionsInput struct {

	// The ID of the Client VPN endpoint to which the client is connected.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The ID of the client connection to be terminated.
	ConnectionId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The name of the user who initiated the connection. Use this option to terminate
	// all active connections for the specified user. This option can only be used if
	// the user has established up to five connections.
	Username *string

	noSmithyDocumentSerde
}

type TerminateClientVpnConnectionsOutput struct {

	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string

	// The current state of the client connections.
	ConnectionStatuses []types.TerminateConnectionStatus

	// The user who established the terminated client connections.
	Username *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTerminateClientVpnConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpTerminateClientVpnConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpTerminateClientVpnConnections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TerminateClientVpnConnections"); err != nil {
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
	if err = addOpTerminateClientVpnConnectionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateClientVpnConnections(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTerminateClientVpnConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TerminateClientVpnConnections",
	}
}
