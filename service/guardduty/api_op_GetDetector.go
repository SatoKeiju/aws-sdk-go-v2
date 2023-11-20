// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves an Amazon GuardDuty detector specified by the detectorId. There might
// be regional differences because some data sources might not be available in all
// the Amazon Web Services Regions where GuardDuty is presently supported. For more
// information, see Regions and endpoints (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html)
// .
func (c *Client) GetDetector(ctx context.Context, params *GetDetectorInput, optFns ...func(*Options)) (*GetDetectorOutput, error) {
	if params == nil {
		params = &GetDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDetector", params, optFns, c.addOperationGetDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDetectorInput struct {

	// The unique ID of the detector that you want to get.
	//
	// This member is required.
	DetectorId *string

	noSmithyDocumentSerde
}

type GetDetectorOutput struct {

	// The GuardDuty service role.
	//
	// This member is required.
	ServiceRole *string

	// The detector status.
	//
	// This member is required.
	Status types.DetectorStatus

	// The timestamp of when the detector was created.
	CreatedAt *string

	// Describes which data sources are enabled for the detector.
	//
	// Deprecated: This parameter is deprecated, use Features instead
	DataSources *types.DataSourceConfigurationsResult

	// Describes the features that have been enabled for the detector.
	Features []types.DetectorFeatureConfigurationResult

	// The publishing frequency of the finding.
	FindingPublishingFrequency types.FindingPublishingFrequency

	// The tags of the detector resource.
	Tags map[string]string

	// The last-updated timestamp for the detector.
	UpdatedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDetector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDetector"); err != nil {
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
	if err = addOpGetDetectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDetector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDetector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDetector",
	}
}
