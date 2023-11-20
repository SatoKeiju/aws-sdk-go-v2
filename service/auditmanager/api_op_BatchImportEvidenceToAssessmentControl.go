// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more pieces of evidence to a control in an Audit Manager
// assessment. You can import manual evidence from any S3 bucket by specifying the
// S3 URI of the object. You can also upload a file from your browser, or enter
// plain text in response to a risk assessment question. The following restrictions
// apply to this action:
//   - manualEvidence can be only one of the following: evidenceFileName ,
//     s3ResourcePath , or textResponse
//   - Maximum size of an individual evidence file: 100 MB
//   - Number of daily manual evidence uploads per control: 100
//   - Supported file formats: See Supported file types for manual evidence (https://docs.aws.amazon.com/audit-manager/latest/userguide/upload-evidence.html#supported-manual-evidence-files)
//     in the Audit Manager User Guide
//
// For more information about Audit Manager service restrictions, see Quotas and
// restrictions for Audit Manager (https://docs.aws.amazon.com/audit-manager/latest/userguide/service-quotas.html)
// .
func (c *Client) BatchImportEvidenceToAssessmentControl(ctx context.Context, params *BatchImportEvidenceToAssessmentControlInput, optFns ...func(*Options)) (*BatchImportEvidenceToAssessmentControlOutput, error) {
	if params == nil {
		params = &BatchImportEvidenceToAssessmentControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchImportEvidenceToAssessmentControl", params, optFns, c.addOperationBatchImportEvidenceToAssessmentControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchImportEvidenceToAssessmentControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchImportEvidenceToAssessmentControlInput struct {

	// The identifier for the assessment.
	//
	// This member is required.
	AssessmentId *string

	// The identifier for the control.
	//
	// This member is required.
	ControlId *string

	// The identifier for the control set.
	//
	// This member is required.
	ControlSetId *string

	// The list of manual evidence objects.
	//
	// This member is required.
	ManualEvidence []types.ManualEvidence

	noSmithyDocumentSerde
}

type BatchImportEvidenceToAssessmentControlOutput struct {

	// A list of errors that the BatchImportEvidenceToAssessmentControl API returned.
	Errors []types.BatchImportEvidenceToAssessmentControlError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchImportEvidenceToAssessmentControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchImportEvidenceToAssessmentControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchImportEvidenceToAssessmentControl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchImportEvidenceToAssessmentControl"); err != nil {
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
	if err = addOpBatchImportEvidenceToAssessmentControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchImportEvidenceToAssessmentControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchImportEvidenceToAssessmentControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchImportEvidenceToAssessmentControl",
	}
}
