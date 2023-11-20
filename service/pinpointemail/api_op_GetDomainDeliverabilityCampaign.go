// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain that
// the Deliverability dashboard is enabled for ( PutDeliverabilityDashboardOption
// operation).
func (c *Client) GetDomainDeliverabilityCampaign(ctx context.Context, params *GetDomainDeliverabilityCampaignInput, optFns ...func(*Options)) (*GetDomainDeliverabilityCampaignOutput, error) {
	if params == nil {
		params = &GetDomainDeliverabilityCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainDeliverabilityCampaign", params, optFns, c.addOperationGetDomainDeliverabilityCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainDeliverabilityCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain that
// the Deliverability dashboard is enabled for ( PutDeliverabilityDashboardOption
// operation).
type GetDomainDeliverabilityCampaignInput struct {

	// The unique identifier for the campaign. Amazon Pinpoint automatically generates
	// and assigns this identifier to a campaign. This value is not the same as the
	// campaign identifier that Amazon Pinpoint assigns to campaigns that you create
	// and manage by using the Amazon Pinpoint API or the Amazon Pinpoint console.
	//
	// This member is required.
	CampaignId *string

	noSmithyDocumentSerde
}

// An object that contains all the deliverability data for a specific campaign.
// This data is available for a campaign only if the campaign sent email by using a
// domain that the Deliverability dashboard is enabled for (
// PutDeliverabilityDashboardOption operation).
type GetDomainDeliverabilityCampaignOutput struct {

	// An object that contains the deliverability data for the campaign.
	//
	// This member is required.
	DomainDeliverabilityCampaign *types.DomainDeliverabilityCampaign

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDomainDeliverabilityCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainDeliverabilityCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainDeliverabilityCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDomainDeliverabilityCampaign"); err != nil {
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
	if err = addOpGetDomainDeliverabilityCampaignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainDeliverabilityCampaign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomainDeliverabilityCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDomainDeliverabilityCampaign",
	}
}
