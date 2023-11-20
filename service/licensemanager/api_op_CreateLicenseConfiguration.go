// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a license configuration. A license configuration is an abstraction of a
// customer license agreement that can be consumed and enforced by License Manager.
// Components include specifications for the license type (licensing by instance,
// socket, CPU, or vCPU), allowed tenancy (shared tenancy, Dedicated Instance,
// Dedicated Host, or all of these), license affinity to host (how long a license
// must be associated with a host), and the number of licenses purchased and used.
func (c *Client) CreateLicenseConfiguration(ctx context.Context, params *CreateLicenseConfigurationInput, optFns ...func(*Options)) (*CreateLicenseConfigurationOutput, error) {
	if params == nil {
		params = &CreateLicenseConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLicenseConfiguration", params, optFns, c.addOperationCreateLicenseConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLicenseConfigurationInput struct {

	// Dimension used to track the license inventory.
	//
	// This member is required.
	LicenseCountingType types.LicenseCountingType

	// Name of the license configuration.
	//
	// This member is required.
	Name *string

	// Description of the license configuration.
	Description *string

	// When true, disassociates a resource when software is uninstalled.
	DisassociateWhenNotFound *bool

	// Number of licenses managed by the license configuration.
	LicenseCount *int64

	// Indicates whether hard or soft license enforcement is used. Exceeding a hard
	// limit blocks the launch of new instances.
	LicenseCountHardLimit *bool

	// License rules. The syntax is #name=value (for example,
	// #allowedTenancy=EC2-DedicatedHost). The available rules vary by dimension, as
	// follows.
	//   - Cores dimension: allowedTenancy | licenseAffinityToHost | maximumCores |
	//   minimumCores
	//   - Instances dimension: allowedTenancy | maximumCores | minimumCores |
	//   maximumSockets | minimumSockets | maximumVcpus | minimumVcpus
	//   - Sockets dimension: allowedTenancy | licenseAffinityToHost | maximumSockets |
	//   minimumSockets
	//   - vCPUs dimension: allowedTenancy | honorVcpuOptimization | maximumVcpus |
	//   minimumVcpus
	// The unit for licenseAffinityToHost is days and the range is 1 to 180. The
	// possible values for allowedTenancy are EC2-Default , EC2-DedicatedHost , and
	// EC2-DedicatedInstance . The possible values for honorVcpuOptimization are True
	// and False .
	LicenseRules []string

	// Product information.
	ProductInformationList []types.ProductInformation

	// Tags to add to the license configuration.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateLicenseConfigurationOutput struct {

	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLicenseConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLicenseConfiguration"); err != nil {
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
	if err = addOpCreateLicenseConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLicenseConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLicenseConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLicenseConfiguration",
	}
}
