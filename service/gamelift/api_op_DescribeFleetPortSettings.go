// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a fleet's inbound connection permissions. Connection permissions
// specify the range of IP addresses and port settings that incoming traffic can
// use to access server processes in the fleet. Game sessions that are running on
// instances in the fleet must use connections that fall in this range. This
// operation can be used in the following ways:
//   - To retrieve the inbound connection permissions for a fleet, identify the
//     fleet's unique identifier.
//   - To check the status of recent updates to a fleet remote location, specify
//     the fleet ID and a location. Port setting updates can take time to propagate
//     across all locations.
//
// If successful, a set of IpPermission objects is returned for the requested
// fleet ID. When a location is specified, a pending status is included. If the
// requested fleet has been deleted, the result set is empty. Learn more Setting
// up Amazon GameLift fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
func (c *Client) DescribeFleetPortSettings(ctx context.Context, params *DescribeFleetPortSettingsInput, optFns ...func(*Options)) (*DescribeFleetPortSettingsOutput, error) {
	if params == nil {
		params = &DescribeFleetPortSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetPortSettings", params, optFns, c.addOperationDescribeFleetPortSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetPortSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetPortSettingsInput struct {

	// A unique identifier for the fleet to retrieve port settings for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// A remote location to check for status of port setting updates. Use the Amazon
	// Web Services Region code format, such as us-west-2 .
	Location *string

	noSmithyDocumentSerde
}

type DescribeFleetPortSettingsOutput struct {

	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) that is assigned to a Amazon GameLift fleet resource and uniquely identifies
	// it. ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	FleetArn *string

	// A unique identifier for the fleet that was requested.
	FleetId *string

	// The port settings for the requested fleet ID.
	InboundPermissions []types.IpPermission

	// The requested fleet location, expressed as an Amazon Web Services Region code,
	// such as us-west-2 .
	Location *string

	// The current status of updates to the fleet's port settings in the requested
	// fleet location. A status of PENDING_UPDATE indicates that an update was
	// requested for the fleet but has not yet been completed for the location.
	UpdateStatus types.LocationUpdateStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetPortSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetPortSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetPortSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleetPortSettings"); err != nil {
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
	if err = addOpDescribeFleetPortSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetPortSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFleetPortSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleetPortSettings",
	}
}
