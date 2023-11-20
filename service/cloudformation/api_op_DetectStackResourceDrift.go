// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about whether a resource's actual configuration differs, or
// has drifted, from its expected configuration, as defined in the stack template
// and any values specified as template parameters. This information includes
// actual and expected property values for resources in which CloudFormation
// detects drift. Only resource properties explicitly defined in the stack template
// are checked for drift. For more information about stack and resource drift, see
// Detecting Unregulated Configuration Changes to Stacks and Resources (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html)
// . Use DetectStackResourceDrift to detect drift on individual resources, or
// DetectStackDrift to detect drift on all resources in a given stack that support
// drift detection. Resources that don't currently support drift detection can't be
// checked. For a list of resources that support drift detection, see Resources
// that Support Drift Detection (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html)
// .
func (c *Client) DetectStackResourceDrift(ctx context.Context, params *DetectStackResourceDriftInput, optFns ...func(*Options)) (*DetectStackResourceDriftOutput, error) {
	if params == nil {
		params = &DetectStackResourceDriftInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectStackResourceDrift", params, optFns, c.addOperationDetectStackResourceDriftMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectStackResourceDriftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectStackResourceDriftInput struct {

	// The logical name of the resource for which to return drift information.
	//
	// This member is required.
	LogicalResourceId *string

	// The name of the stack to which the resource belongs.
	//
	// This member is required.
	StackName *string

	noSmithyDocumentSerde
}

type DetectStackResourceDriftOutput struct {

	// Information about whether the resource's actual configuration has drifted from
	// its expected template configuration, including actual and expected property
	// values and any differences detected.
	//
	// This member is required.
	StackResourceDrift *types.StackResourceDrift

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectStackResourceDriftMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDetectStackResourceDrift{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDetectStackResourceDrift{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DetectStackResourceDrift"); err != nil {
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
	if err = addOpDetectStackResourceDriftValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectStackResourceDrift(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectStackResourceDrift(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DetectStackResourceDrift",
	}
}
