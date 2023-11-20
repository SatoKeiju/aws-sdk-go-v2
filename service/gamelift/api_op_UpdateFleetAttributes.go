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

// Updates a fleet's mutable attributes, including game session protection and
// resource creation limits. To update fleet attributes, specify the fleet ID and
// the property values that you want to change. If successful, an updated
// FleetAttributes object is returned. Learn more Setting up Amazon GameLift fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
func (c *Client) UpdateFleetAttributes(ctx context.Context, params *UpdateFleetAttributesInput, optFns ...func(*Options)) (*UpdateFleetAttributesOutput, error) {
	if params == nil {
		params = &UpdateFleetAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFleetAttributes", params, optFns, c.addOperationUpdateFleetAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFleetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFleetAttributesInput struct {

	// A unique identifier for the fleet to update attribute metadata for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// Amazon GameLift Anywhere configuration options.
	AnywhereConfiguration *types.AnywhereConfiguration

	// A human-readable description of a fleet.
	Description *string

	// The name of a metric group to add this fleet to. Use a metric group in Amazon
	// CloudWatch to aggregate the metrics from multiple fleets. Provide an existing
	// metric group name, or create a new metric group by providing a new name. A fleet
	// can only be in one metric group at a time.
	MetricGroups []string

	// A descriptive label that is associated with a fleet. Fleet names do not need to
	// be unique.
	Name *string

	// The game session protection policy to apply to all new game sessions created in
	// this fleet. Game sessions that already exist are not affected. You can set
	// protection for individual game sessions using UpdateGameSession (https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateGameSession.html)
	// .
	//   - NoProtection -- The game session can be terminated during a scale-down
	//   event.
	//   - FullProtection -- If the game session is in an ACTIVE status, it cannot be
	//   terminated during a scale-down event.
	NewGameSessionProtectionPolicy types.ProtectionPolicy

	// Policy settings that limit the number of game sessions an individual player can
	// create over a span of time.
	ResourceCreationLimitPolicy *types.ResourceCreationLimitPolicy

	noSmithyDocumentSerde
}

type UpdateFleetAttributesOutput struct {

	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) that is assigned to a Amazon GameLift fleet resource and uniquely identifies
	// it. ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	FleetArn *string

	// A unique identifier for the fleet that was updated.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFleetAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleetAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleetAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFleetAttributes"); err != nil {
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
	if err = addOpUpdateFleetAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFleetAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFleetAttributes",
	}
}
