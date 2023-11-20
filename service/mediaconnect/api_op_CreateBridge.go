// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new bridge. The request must include one source.
func (c *Client) CreateBridge(ctx context.Context, params *CreateBridgeInput, optFns ...func(*Options)) (*CreateBridgeOutput, error) {
	if params == nil {
		params = &CreateBridgeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBridge", params, optFns, c.addOperationCreateBridgeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBridgeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new bridge. The request must include one source.
type CreateBridgeInput struct {

	// The name of the bridge. This name can not be modified after the bridge is
	// created.
	//
	// This member is required.
	Name *string

	// The bridge placement Amazon Resource Number (ARN).
	//
	// This member is required.
	PlacementArn *string

	// The sources that you want to add to this bridge.
	//
	// This member is required.
	Sources []types.AddBridgeSourceRequest

	// Create a bridge with the egress bridge type. An egress bridge is a
	// cloud-to-ground bridge. The content comes from an existing MediaConnect flow and
	// is delivered to your premises.
	EgressGatewayBridge *types.AddEgressGatewayBridgeRequest

	// Create a bridge with the ingress bridge type. An ingress bridge is a
	// ground-to-cloud bridge. The content originates at your premises and is delivered
	// to the cloud.
	IngressGatewayBridge *types.AddIngressGatewayBridgeRequest

	// The outputs that you want to add to this bridge.
	Outputs []types.AddBridgeOutputRequest

	// The settings for source failover.
	SourceFailoverConfig *types.FailoverConfig

	noSmithyDocumentSerde
}

type CreateBridgeOutput struct {

	// A Bridge is the connection between your datacenter's Instances and the AWS
	// cloud. A bridge can be used to send video from the AWS cloud to your datacenter
	// or from your datacenter to the AWS cloud.
	Bridge *types.Bridge

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBridgeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBridge{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBridge{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBridge"); err != nil {
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
	if err = addOpCreateBridgeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBridge(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBridge(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBridge",
	}
}
