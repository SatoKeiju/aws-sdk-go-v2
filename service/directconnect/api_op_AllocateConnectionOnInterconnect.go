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
	"time"
)

// Deprecated. Use AllocateHostedConnection instead. Creates a hosted connection
// on an interconnect. Allocates a VLAN number and a specified amount of bandwidth
// for use by a hosted connection on the specified interconnect. Intended for use
// by Direct Connect Partners only.
//
// Deprecated: This operation has been deprecated.
func (c *Client) AllocateConnectionOnInterconnect(ctx context.Context, params *AllocateConnectionOnInterconnectInput, optFns ...func(*Options)) (*AllocateConnectionOnInterconnectOutput, error) {
	if params == nil {
		params = &AllocateConnectionOnInterconnectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AllocateConnectionOnInterconnect", params, optFns, c.addOperationAllocateConnectionOnInterconnectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AllocateConnectionOnInterconnectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllocateConnectionOnInterconnectInput struct {

	// The bandwidth of the connection. The possible values are 50Mbps, 100Mbps,
	// 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, and 10Gbps. Note that
	// only those Direct Connect Partners who have met specific requirements are
	// allowed to create a 1Gbps, 2Gbps, 5Gbps or 10Gbps hosted connection.
	//
	// This member is required.
	Bandwidth *string

	// The name of the provisioned connection.
	//
	// This member is required.
	ConnectionName *string

	// The ID of the interconnect on which the connection will be provisioned.
	//
	// This member is required.
	InterconnectId *string

	// The ID of the Amazon Web Services account of the customer for whom the
	// connection will be provisioned.
	//
	// This member is required.
	OwnerAccount *string

	// The dedicated VLAN provisioned to the connection.
	//
	// This member is required.
	Vlan int32

	noSmithyDocumentSerde
}

// Information about an Direct Connect connection.
type AllocateConnectionOnInterconnectOutput struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The ID of the connection.
	ConnectionId *string

	// The name of the connection.
	ConnectionName *string

	// The state of the connection. The following are the possible values:
	//   - ordering : The initial state of a hosted connection provisioned on an
	//   interconnect. The connection stays in the ordering state until the owner of the
	//   hosted connection confirms or declines the connection order.
	//   - requested : The initial state of a standard connection. The connection stays
	//   in the requested state until the Letter of Authorization (LOA) is sent to the
	//   customer.
	//   - pending : The connection has been approved and is being initialized.
	//   - available : The network link is up and the connection is ready for use.
	//   - down : The network link is down.
	//   - deleting : The connection is being deleted.
	//   - deleted : The connection has been deleted.
	//   - rejected : A hosted connection in the ordering state enters the rejected
	//   state if it is deleted by the customer.
	//   - unknown : The state of the connection is not available.
	ConnectionState types.ConnectionState

	// The MAC Security (MACsec) connection encryption mode. The valid values are
	// no_encrypt , should_encrypt , and must_encrypt .
	EncryptionMode *string

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// Indicates whether the connection supports MAC Security (MACsec).
	MacSecCapable *bool

	// The MAC Security (MACsec) security keys associated with the connection.
	MacSecKeys []types.MacSecKey

	// The ID of the Amazon Web Services account that owns the connection.
	OwnerAccount *string

	// The name of the Direct Connect service provider associated with the connection.
	PartnerName *string

	// The MAC Security (MACsec) port link status of the connection. The valid values
	// are Encryption Up , which means that there is an active Connection Key Name, or
	// Encryption Down .
	PortEncryptionStatus *string

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The Amazon Web Services Region where the connection is located.
	Region *string

	// The tags associated with the connection.
	Tags []types.Tag

	// The ID of the VLAN.
	Vlan int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAllocateConnectionOnInterconnectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAllocateConnectionOnInterconnect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAllocateConnectionOnInterconnect{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AllocateConnectionOnInterconnect"); err != nil {
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
	if err = addOpAllocateConnectionOnInterconnectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAllocateConnectionOnInterconnect(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAllocateConnectionOnInterconnect(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AllocateConnectionOnInterconnect",
	}
}
