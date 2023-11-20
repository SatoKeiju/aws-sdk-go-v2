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

// Modify the configurations of an IPAM pool. For more information, see Modify a
// pool (https://docs.aws.amazon.com/vpc/latest/ipam/mod-pool-ipam.html) in the
// Amazon VPC IPAM User Guide.
func (c *Client) ModifyIpamPool(ctx context.Context, params *ModifyIpamPoolInput, optFns ...func(*Options)) (*ModifyIpamPoolOutput, error) {
	if params == nil {
		params = &ModifyIpamPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyIpamPool", params, optFns, c.addOperationModifyIpamPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyIpamPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyIpamPoolInput struct {

	// The ID of the IPAM pool you want to modify.
	//
	// This member is required.
	IpamPoolId *string

	// Add tag allocation rules to a pool. For more information about allocation
	// rules, see Create a top-level pool (https://docs.aws.amazon.com/vpc/latest/ipam/create-top-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	AddAllocationResourceTags []types.RequestIpamResourceTag

	// The default netmask length for allocations added to this pool. If, for example,
	// the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new
	// allocations will default to 10.0.0.0/16.
	AllocationDefaultNetmaskLength *int32

	// The maximum netmask length possible for CIDR allocations in this IPAM pool to
	// be compliant. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible
	// netmask lengths for IPv6 addresses are 0 - 128.The maximum netmask length must
	// be greater than the minimum netmask length.
	AllocationMaxNetmaskLength *int32

	// The minimum netmask length required for CIDR allocations in this IPAM pool to
	// be compliant. Possible netmask lengths for IPv4 addresses are 0 - 32. Possible
	// netmask lengths for IPv6 addresses are 0 - 128. The minimum netmask length must
	// be less than the maximum netmask length.
	AllocationMinNetmaskLength *int32

	// If true, IPAM will continuously look for resources within the CIDR range of
	// this pool and automatically import them as allocations into your IPAM. The CIDRs
	// that will be allocated for these resources must not already be allocated to
	// other resources in order for the import to succeed. IPAM will import a CIDR
	// regardless of its compliance with the pool's allocation rules, so a resource
	// might be imported and subsequently marked as noncompliant. If IPAM discovers
	// multiple CIDRs that overlap, IPAM will import the largest CIDR only. If IPAM
	// discovers multiple CIDRs with matching CIDRs, IPAM will randomly import one of
	// them only. A locale must be set on the pool for this feature to work.
	AutoImport *bool

	// Clear the default netmask length allocation rule for this pool.
	ClearAllocationDefaultNetmaskLength *bool

	// The description of the IPAM pool you want to modify.
	Description *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Remove tag allocation rules from a pool.
	RemoveAllocationResourceTags []types.RequestIpamResourceTag

	noSmithyDocumentSerde
}

type ModifyIpamPoolOutput struct {

	// The results of the modification.
	IpamPool *types.IpamPool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyIpamPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyIpamPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyIpamPool{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyIpamPool"); err != nil {
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
	if err = addOpModifyIpamPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyIpamPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyIpamPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyIpamPool",
	}
}
