// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the cluster configuration of the specified Elasticsearch domain,
// setting as setting the instance type and the number of instances.
func (c *Client) UpdateElasticsearchDomainConfig(ctx context.Context, params *UpdateElasticsearchDomainConfigInput, optFns ...func(*Options)) (*UpdateElasticsearchDomainConfigOutput, error) {
	if params == nil {
		params = &UpdateElasticsearchDomainConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateElasticsearchDomainConfig", params, optFns, c.addOperationUpdateElasticsearchDomainConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateElasticsearchDomainConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the UpdateElasticsearchDomain operation.
// Specifies the type and number of instances in the domain cluster.
type UpdateElasticsearchDomainConfigInput struct {

	// The name of the Elasticsearch domain that you are updating.
	//
	// This member is required.
	DomainName *string

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string

	// Modifies the advanced option to allow references to indices in an HTTP request
	// body. Must be false when configuring access to individual sub-resources. By
	// default, the value is true . See Configuration Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options)
	// for more information.
	AdvancedOptions map[string]string

	// Specifies advanced security options.
	AdvancedSecurityOptions *types.AdvancedSecurityOptionsInput

	// Specifies Auto-Tune options.
	AutoTuneOptions *types.AutoTuneOptions

	// Options to specify the Cognito user and identity pools for Kibana
	// authentication. For more information, see Amazon Cognito Authentication for
	// Kibana (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html)
	// .
	CognitoOptions *types.CognitoOptions

	// Options to specify configuration that will be applied to the domain endpoint.
	DomainEndpointOptions *types.DomainEndpointOptions

	// This flag, when set to True, specifies whether the UpdateElasticsearchDomain
	// request should return the results of validation checks without actually applying
	// the change. This flag, when set to True, specifies the deployment mechanism
	// through which the update shall be applied on the domain. This will not actually
	// perform the Update.
	DryRun *bool

	// Specify the type and size of the EBS volume that you want to use.
	EBSOptions *types.EBSOptions

	// The type and number of instances to instantiate for the domain cluster.
	ElasticsearchClusterConfig *types.ElasticsearchClusterConfig

	// Specifies the Encryption At Rest Options.
	EncryptionAtRestOptions *types.EncryptionAtRestOptions

	// Map of LogType and LogPublishingOption , each containing options to publish a
	// given type of Elasticsearch log.
	LogPublishingOptions map[string]types.LogPublishingOption

	// Specifies the NodeToNodeEncryptionOptions.
	NodeToNodeEncryptionOptions *types.NodeToNodeEncryptionOptions

	// Option to set the time, in UTC format, for the daily automated snapshot.
	// Default value is 0 hours.
	SnapshotOptions *types.SnapshotOptions

	// Options to specify the subnets and security groups for VPC endpoint. For more
	// information, see Creating a VPC (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-creating-vpc)
	// in VPC Endpoints for Amazon Elasticsearch Service Domains
	VPCOptions *types.VPCOptions

	noSmithyDocumentSerde
}

// The result of an UpdateElasticsearchDomain request. Contains the status of the
// Elasticsearch domain being updated.
type UpdateElasticsearchDomainConfigOutput struct {

	// The status of the updated Elasticsearch domain.
	//
	// This member is required.
	DomainConfig *types.ElasticsearchDomainConfig

	// Contains result of DryRun.
	DryRunResults *types.DryRunResults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateElasticsearchDomainConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateElasticsearchDomainConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateElasticsearchDomainConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateElasticsearchDomainConfig"); err != nil {
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
	if err = addOpUpdateElasticsearchDomainConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateElasticsearchDomainConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateElasticsearchDomainConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateElasticsearchDomainConfig",
	}
}
