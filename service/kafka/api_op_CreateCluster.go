// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new MSK cluster.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, c.addOperationCreateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// Information about the broker nodes in the cluster.
	//
	// This member is required.
	BrokerNodeGroupInfo *types.BrokerNodeGroupInfo

	// The name of the cluster.
	//
	// This member is required.
	ClusterName *string

	// The version of Apache Kafka.
	//
	// This member is required.
	KafkaVersion *string

	// The number of broker nodes in the cluster.
	//
	// This member is required.
	NumberOfBrokerNodes *int32

	// Includes all client authentication related information.
	ClientAuthentication *types.ClientAuthentication

	// Represents the configuration that you want MSK to use for the brokers in a
	// cluster.
	ConfigurationInfo *types.ConfigurationInfo

	// Includes all encryption-related information.
	EncryptionInfo *types.EncryptionInfo

	// Specifies the level of monitoring for the MSK cluster. The possible values are
	// DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER, and PER_TOPIC_PER_PARTITION.
	EnhancedMonitoring types.EnhancedMonitoring

	LoggingInfo *types.LoggingInfo

	// The settings for open monitoring.
	OpenMonitoring *types.OpenMonitoringInfo

	// This controls storage mode for supported storage tiers.
	StorageMode types.StorageMode

	// Create tags when creating the cluster.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateClusterOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The name of the MSK cluster.
	ClusterName *string

	// The state of the cluster. The possible states are ACTIVE, CREATING, DELETING,
	// FAILED, HEALING, MAINTENANCE, REBOOTING_BROKER, and UPDATING.
	State types.ClusterState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCluster"); err != nil {
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
	if err = addOpCreateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCluster",
	}
}
