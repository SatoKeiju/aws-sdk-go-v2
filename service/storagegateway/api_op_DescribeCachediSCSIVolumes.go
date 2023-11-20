// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a description of the gateway volumes specified in the request. This
// operation is only supported in the cached volume gateway types. The list of
// gateway volumes in the request must be from one gateway. In the response,
// Storage Gateway returns volume information sorted by volume Amazon Resource Name
// (ARN).
func (c *Client) DescribeCachediSCSIVolumes(ctx context.Context, params *DescribeCachediSCSIVolumesInput, optFns ...func(*Options)) (*DescribeCachediSCSIVolumesOutput, error) {
	if params == nil {
		params = &DescribeCachediSCSIVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCachediSCSIVolumes", params, optFns, c.addOperationDescribeCachediSCSIVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCachediSCSIVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCachediSCSIVolumesInput struct {

	// An array of strings where each string represents the Amazon Resource Name (ARN)
	// of a cached volume. All of the specified cached volumes must be from the same
	// gateway. Use ListVolumes to get volume ARNs for a gateway.
	//
	// This member is required.
	VolumeARNs []string

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
type DescribeCachediSCSIVolumesOutput struct {

	// An array of objects where each object contains metadata about one cached volume.
	CachediSCSIVolumes []types.CachediSCSIVolume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCachediSCSIVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCachediSCSIVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCachediSCSIVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCachediSCSIVolumes"); err != nil {
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
	if err = addOpDescribeCachediSCSIVolumesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCachediSCSIVolumes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCachediSCSIVolumes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCachediSCSIVolumes",
	}
}
