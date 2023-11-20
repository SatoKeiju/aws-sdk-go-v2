// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a network instance. A network instance is a single network created in
// Amazon Web Services TNB that can be deployed and on which life-cycle operations
// (like terminate, update, and delete) can be performed.
func (c *Client) UpdateSolNetworkInstance(ctx context.Context, params *UpdateSolNetworkInstanceInput, optFns ...func(*Options)) (*UpdateSolNetworkInstanceOutput, error) {
	if params == nil {
		params = &UpdateSolNetworkInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSolNetworkInstance", params, optFns, c.addOperationUpdateSolNetworkInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSolNetworkInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSolNetworkInstanceInput struct {

	// ID of the network instance.
	//
	// This member is required.
	NsInstanceId *string

	// The type of update.
	//
	// This member is required.
	UpdateType types.UpdateSolNetworkType

	// Identifies the network function information parameters and/or the configurable
	// properties of the network function to be modified.
	ModifyVnfInfoData *types.UpdateSolNetworkModify

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. When you use this API, the tags are
	// transferred to the network operation that is created. Use tags to search and
	// filter your resources or track your Amazon Web Services costs.
	Tags map[string]string

	noSmithyDocumentSerde
}

type UpdateSolNetworkInstanceOutput struct {

	// The identifier of the network operation.
	NsLcmOpOccId *string

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. When you use this API, the tags are
	// transferred to the network operation that is created. Use tags to search and
	// filter your resources or track your Amazon Web Services costs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSolNetworkInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSolNetworkInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSolNetworkInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSolNetworkInstance"); err != nil {
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
	if err = addOpUpdateSolNetworkInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSolNetworkInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSolNetworkInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSolNetworkInstance",
	}
}
