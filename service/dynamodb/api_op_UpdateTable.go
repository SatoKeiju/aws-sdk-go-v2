// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the provisioned throughput settings, global secondary indexes, or
// DynamoDB Streams settings for a given table. This operation only applies to
// Version 2019.11.21 (Current) (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
// of global tables. You can only perform one of the following operations at once:
//   - Modify the provisioned throughput settings of the table.
//   - Remove a global secondary index from the table.
//   - Create a new global secondary index on the table. After the index begins
//     backfilling, you can use UpdateTable to perform other operations.
//
// UpdateTable is an asynchronous operation; while it is executing, the table
// status changes from ACTIVE to UPDATING . While it is UPDATING , you cannot issue
// another UpdateTable request. When the table returns to the ACTIVE state, the
// UpdateTable operation is complete.
func (c *Client) UpdateTable(ctx context.Context, params *UpdateTableInput, optFns ...func(*Options)) (*UpdateTableOutput, error) {
	if params == nil {
		params = &UpdateTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTable", params, optFns, c.addOperationUpdateTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an UpdateTable operation.
type UpdateTableInput struct {

	// The name of the table to be updated.
	//
	// This member is required.
	TableName *string

	// An array of attributes that describe the key schema for the table and indexes.
	// If you are adding a new global secondary index to the table,
	// AttributeDefinitions must include the key element(s) of the new index.
	AttributeDefinitions []types.AttributeDefinition

	// Controls how you are charged for read and write throughput and how you manage
	// capacity. When switching from pay-per-request to provisioned capacity, initial
	// provisioned capacity values must be set. The initial provisioned capacity values
	// are estimated based on the consumed read and write capacity of your table and
	// global secondary indexes over the past 30 minutes.
	//   - PROVISIONED - We recommend using PROVISIONED for predictable workloads.
	//   PROVISIONED sets the billing mode to Provisioned Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual)
	//   .
	//   - PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable
	//   workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand)
	//   .
	BillingMode types.BillingMode

	// Indicates whether deletion protection is to be enabled (true) or disabled
	// (false) on the table.
	DeletionProtectionEnabled *bool

	// An array of one or more global secondary indexes for the table. For each index
	// in the array, you can request one action:
	//   - Create - add a new global secondary index to the table.
	//   - Update - modify the provisioned throughput settings of an existing global
	//   secondary index.
	//   - Delete - remove a global secondary index from the table.
	// You can create or delete only one global secondary index per UpdateTable
	// operation. For more information, see Managing Global Secondary Indexes (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.OnlineOps.html)
	// in the Amazon DynamoDB Developer Guide.
	GlobalSecondaryIndexUpdates []types.GlobalSecondaryIndexUpdate

	// The new provisioned throughput settings for the specified table or index.
	ProvisionedThroughput *types.ProvisionedThroughput

	// A list of replica update actions (create, delete, or update) for the table.
	// This property only applies to Version 2019.11.21 (Current) (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
	// of global tables.
	ReplicaUpdates []types.ReplicationGroupUpdate

	// The new server-side encryption settings for the specified table.
	SSESpecification *types.SSESpecification

	// Represents the DynamoDB Streams configuration for the table. You receive a
	// ResourceInUseException if you try to enable a stream on a table that already has
	// a stream, or if you try to disable a stream on a table that doesn't have a
	// stream.
	StreamSpecification *types.StreamSpecification

	// The table class of the table to be updated. Valid values are STANDARD and
	// STANDARD_INFREQUENT_ACCESS .
	TableClass types.TableClass

	noSmithyDocumentSerde
}

// Represents the output of an UpdateTable operation.
type UpdateTableOutput struct {

	// Represents the properties of the table.
	TableDescription *types.TableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTable{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTable"); err != nil {
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
	if err = addOpUpdateTableDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addV4DetectSkewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTable(options.Region), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
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

func addOpUpdateTableDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpUpdateTableDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpUpdateTableDiscoverEndpoint(ctx context.Context, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*UpdateTableInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opUpdateTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTable",
	}
}
