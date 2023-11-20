// Code generated by smithy-go-codegen DO NOT EDIT.

package savingsplans

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified Savings Plans offering rates.
func (c *Client) DescribeSavingsPlansOfferingRates(ctx context.Context, params *DescribeSavingsPlansOfferingRatesInput, optFns ...func(*Options)) (*DescribeSavingsPlansOfferingRatesOutput, error) {
	if params == nil {
		params = &DescribeSavingsPlansOfferingRatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSavingsPlansOfferingRates", params, optFns, c.addOperationDescribeSavingsPlansOfferingRatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSavingsPlansOfferingRatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSavingsPlansOfferingRatesInput struct {

	// The filters.
	Filters []types.SavingsPlanOfferingRateFilterElement

	// The maximum number of results to return with a single call. To retrieve
	// additional results, make another call with the returned token value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string

	// The specific AWS operation for the line item in the billing report.
	Operations []string

	// The AWS products.
	Products []types.SavingsPlanProductType

	// The IDs of the offerings.
	SavingsPlanOfferingIds []string

	// The payment options.
	SavingsPlanPaymentOptions []types.SavingsPlanPaymentOption

	// The plan types.
	SavingsPlanTypes []types.SavingsPlanType

	// The services.
	ServiceCodes []types.SavingsPlanRateServiceCode

	// The usage details of the line item in the billing report.
	UsageTypes []string

	noSmithyDocumentSerde
}

type DescribeSavingsPlansOfferingRatesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the Savings Plans offering rates.
	SearchResults []types.SavingsPlanOfferingRate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSavingsPlansOfferingRatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSavingsPlansOfferingRates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSavingsPlansOfferingRates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSavingsPlansOfferingRates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSavingsPlansOfferingRates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSavingsPlansOfferingRates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSavingsPlansOfferingRates",
	}
}
