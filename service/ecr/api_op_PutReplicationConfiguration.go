// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the replication configuration for a registry. The existing
// replication configuration for a repository can be retrieved with the
// DescribeRegistry API action. The first time the PutReplicationConfiguration API
// is called, a service-linked IAM role is created in your account for the
// replication process. For more information, see Using service-linked roles for
// Amazon ECR (https://docs.aws.amazon.com/AmazonECR/latest/userguide/using-service-linked-roles.html)
// in the Amazon Elastic Container Registry User Guide. When configuring
// cross-account replication, the destination account must grant the source account
// permission to replicate. This permission is controlled using a registry
// permissions policy. For more information, see PutRegistryPolicy .
func (c *Client) PutReplicationConfiguration(ctx context.Context, params *PutReplicationConfigurationInput, optFns ...func(*Options)) (*PutReplicationConfigurationOutput, error) {
	if params == nil {
		params = &PutReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutReplicationConfiguration", params, optFns, c.addOperationPutReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutReplicationConfigurationInput struct {

	// An object representing the replication configuration for a registry.
	//
	// This member is required.
	ReplicationConfiguration *types.ReplicationConfiguration

	noSmithyDocumentSerde
}

type PutReplicationConfigurationOutput struct {

	// The contents of the replication configuration for the registry.
	ReplicationConfiguration *types.ReplicationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutReplicationConfiguration"); err != nil {
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
	if err = addOpPutReplicationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutReplicationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutReplicationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutReplicationConfiguration",
	}
}
