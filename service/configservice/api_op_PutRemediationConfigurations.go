// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates the remediation configuration with a specific Config rule with
// the selected target or action. The API creates the RemediationConfiguration
// object for the Config rule. The Config rule must already exist for you to add a
// remediation configuration. The target (SSM document) must exist and have
// permissions to use the target. If you make backward incompatible changes to the
// SSM document, you must call this again to ensure the remediations can run. This
// API does not support adding remediation configurations for service-linked Config
// Rules such as Organization Config rules, the rules deployed by conformance
// packs, and rules deployed by Amazon Web Services Security Hub. For manual
// remediation configuration, you need to provide a value for automationAssumeRole
// or use a value in the assumeRole field to remediate your resources. The SSM
// automation document can use either as long as it maps to a valid parameter.
// However, for automatic remediation configuration, the only valid assumeRole
// field value is AutomationAssumeRole and you need to provide a value for
// AutomationAssumeRole to remediate your resources.
func (c *Client) PutRemediationConfigurations(ctx context.Context, params *PutRemediationConfigurationsInput, optFns ...func(*Options)) (*PutRemediationConfigurationsOutput, error) {
	if params == nil {
		params = &PutRemediationConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRemediationConfigurations", params, optFns, c.addOperationPutRemediationConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRemediationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRemediationConfigurationsInput struct {

	// A list of remediation configuration objects.
	//
	// This member is required.
	RemediationConfigurations []types.RemediationConfiguration

	noSmithyDocumentSerde
}

type PutRemediationConfigurationsOutput struct {

	// Returns a list of failed remediation batch objects.
	FailedBatches []types.FailedRemediationBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRemediationConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRemediationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRemediationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRemediationConfigurations"); err != nil {
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
	if err = addOpPutRemediationConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRemediationConfigurations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRemediationConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRemediationConfigurations",
	}
}
