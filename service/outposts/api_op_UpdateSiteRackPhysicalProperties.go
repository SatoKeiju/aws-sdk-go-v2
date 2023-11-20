// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the physical and logistical details for a rack at a site. For more
// information about hardware requirements for racks, see Network readiness
// checklist (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#checklist)
// in the Amazon Web Services Outposts User Guide. To update a rack at a site with
// an order of IN_PROGRESS , you must wait for the order to complete or cancel the
// order.
func (c *Client) UpdateSiteRackPhysicalProperties(ctx context.Context, params *UpdateSiteRackPhysicalPropertiesInput, optFns ...func(*Options)) (*UpdateSiteRackPhysicalPropertiesOutput, error) {
	if params == nil {
		params = &UpdateSiteRackPhysicalPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSiteRackPhysicalProperties", params, optFns, c.addOperationUpdateSiteRackPhysicalPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSiteRackPhysicalPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSiteRackPhysicalPropertiesInput struct {

	// The ID or the Amazon Resource Name (ARN) of the site.
	//
	// This member is required.
	SiteId *string

	// The type of fiber that you will use to attach the Outpost to your network.
	FiberOpticCableType types.FiberOpticCableType

	// The maximum rack weight that this site can support. NO_LIMIT is over 2000lbs.
	MaximumSupportedWeightLbs types.MaximumSupportedWeightLbs

	// The type of optical standard that you will use to attach the Outpost to your
	// network. This field is dependent on uplink speed, fiber type, and distance to
	// the upstream device. For more information about networking requirements for
	// racks, see Network (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#facility-networking)
	// in the Amazon Web Services Outposts User Guide.
	//   - OPTIC_10GBASE_SR : 10GBASE-SR
	//   - OPTIC_10GBASE_IR : 10GBASE-IR
	//   - OPTIC_10GBASE_LR : 10GBASE-LR
	//   - OPTIC_40GBASE_SR : 40GBASE-SR
	//   - OPTIC_40GBASE_ESR : 40GBASE-ESR
	//   - OPTIC_40GBASE_IR4_LR4L : 40GBASE-IR (LR4L)
	//   - OPTIC_40GBASE_LR4 : 40GBASE-LR4
	//   - OPTIC_100GBASE_SR4 : 100GBASE-SR4
	//   - OPTIC_100GBASE_CWDM4 : 100GBASE-CWDM4
	//   - OPTIC_100GBASE_LR4 : 100GBASE-LR4
	//   - OPTIC_100G_PSM4_MSA : 100G PSM4 MSA
	//   - OPTIC_1000BASE_LX : 1000Base-LX
	//   - OPTIC_1000BASE_SX : 1000Base-SX
	OpticalStandard types.OpticalStandard

	// The power connector that Amazon Web Services should plan to provide for
	// connections to the hardware. Note the correlation between PowerPhase and
	// PowerConnector .
	//   - Single-phase AC feed
	//   - L6-30P – (common in US); 30A; single phase
	//   - IEC309 (blue) – P+N+E, 6hr; 32 A; single phase
	//   - Three-phase AC feed
	//   - AH530P7W (red) – 3P+N+E, 7hr; 30A; three phase
	//   - AH532P6W (red) – 3P+N+E, 6hr; 32A; three phase
	PowerConnector types.PowerConnector

	// The power draw, in kVA, available at the hardware placement position for the
	// rack.
	PowerDrawKva types.PowerDrawKva

	// Indicates whether the power feed comes above or below the rack.
	PowerFeedDrop types.PowerFeedDrop

	// The power option that you can provide for hardware.
	//   - Single-phase AC feed: 200 V to 277 V, 50 Hz or 60 Hz
	//   - Three-phase AC feed: 346 V to 480 V, 50 Hz or 60 Hz
	PowerPhase types.PowerPhase

	// Racks come with two Outpost network devices. Depending on the supported uplink
	// speed at the site, the Outpost network devices provide a variable number of
	// uplinks. Specify the number of uplinks for each Outpost network device that you
	// intend to use to connect the rack to your network. Note the correlation between
	// UplinkGbps and UplinkCount .
	//   - 1Gbps - Uplinks available: 1, 2, 4, 6, 8
	//   - 10Gbps - Uplinks available: 1, 2, 4, 8, 12, 16
	//   - 40 and 100 Gbps- Uplinks available: 1, 2, 4
	UplinkCount types.UplinkCount

	// The uplink speed the rack should support for the connection to the Region.
	UplinkGbps types.UplinkGbps

	noSmithyDocumentSerde
}

type UpdateSiteRackPhysicalPropertiesOutput struct {

	// Information about a site.
	Site *types.Site

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSiteRackPhysicalPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSiteRackPhysicalProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSiteRackPhysicalProperties{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSiteRackPhysicalProperties"); err != nil {
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
	if err = addOpUpdateSiteRackPhysicalPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSiteRackPhysicalProperties(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSiteRackPhysicalProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSiteRackPhysicalProperties",
	}
}
