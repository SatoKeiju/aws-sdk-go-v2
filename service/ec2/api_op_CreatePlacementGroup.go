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

// Creates a placement group in which to launch instances. The strategy of the
// placement group determines how the instances are organized within the group. A
// cluster placement group is a logical grouping of instances within a single
// Availability Zone that benefit from low network latency, high network
// throughput. A spread placement group places instances on distinct hardware. A
// partition placement group places groups of instances in different partitions,
// where instances in one partition do not share the same hardware with instances
// in another partition. For more information, see Placement groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html)
// in the Amazon EC2 User Guide.
func (c *Client) CreatePlacementGroup(ctx context.Context, params *CreatePlacementGroupInput, optFns ...func(*Options)) (*CreatePlacementGroupOutput, error) {
	if params == nil {
		params = &CreatePlacementGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlacementGroup", params, optFns, c.addOperationCreatePlacementGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePlacementGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePlacementGroupInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// A name for the placement group. Must be unique within the scope of your account
	// for the Region. Constraints: Up to 255 ASCII characters
	GroupName *string

	// The number of partitions. Valid only when Strategy is set to partition .
	PartitionCount *int32

	// Determines how placement groups spread instances.
	//   - Host – You can use host only with Outpost placement groups.
	//   - Rack – No usage restrictions.
	SpreadLevel types.SpreadLevel

	// The placement strategy.
	Strategy types.PlacementStrategy

	// The tags to apply to the new placement group.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreatePlacementGroupOutput struct {

	// Information about the placement group.
	PlacementGroup *types.PlacementGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePlacementGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreatePlacementGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreatePlacementGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePlacementGroup"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlacementGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePlacementGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePlacementGroup",
	}
}
