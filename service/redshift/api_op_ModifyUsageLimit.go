// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies a usage limit in a cluster. You can't modify the feature type or
// period of a usage limit.
func (c *Client) ModifyUsageLimit(ctx context.Context, params *ModifyUsageLimitInput, optFns ...func(*Options)) (*ModifyUsageLimitOutput, error) {
	if params == nil {
		params = &ModifyUsageLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyUsageLimit", params, optFns, c.addOperationModifyUsageLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyUsageLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyUsageLimitInput struct {

	// The identifier of the usage limit to modify.
	//
	// This member is required.
	UsageLimitId *string

	// The new limit amount. For more information about this parameter, see UsageLimit .
	Amount *int64

	// The new action that Amazon Redshift takes when the limit is reached. For more
	// information about this parameter, see UsageLimit .
	BreachAction types.UsageLimitBreachAction

	noSmithyDocumentSerde
}

// Describes a usage limit object for a cluster.
type ModifyUsageLimitOutput struct {

	// The limit amount. If time-based, this amount is in minutes. If data-based, this
	// amount is in terabytes (TB).
	Amount *int64

	// The action that Amazon Redshift takes when the limit is reached. Possible
	// values are:
	//   - log - To log an event in a system table. The default is log.
	//   - emit-metric - To emit CloudWatch metrics.
	//   - disable - To disable the feature until the next usage period begins.
	BreachAction types.UsageLimitBreachAction

	// The identifier of the cluster with a usage limit.
	ClusterIdentifier *string

	// The Amazon Redshift feature to which the limit applies.
	FeatureType types.UsageLimitFeatureType

	// The type of limit. Depending on the feature type, this can be based on a time
	// duration or data size.
	LimitType types.UsageLimitLimitType

	// The time period that the amount applies to. A weekly period begins on Sunday.
	// The default is monthly .
	Period types.UsageLimitPeriod

	// A list of tag instances.
	Tags []types.Tag

	// The identifier of the usage limit.
	UsageLimitId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyUsageLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyUsageLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyUsageLimit{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyUsageLimit"); err != nil {
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
	if err = addOpModifyUsageLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyUsageLimit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyUsageLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyUsageLimit",
	}
}
