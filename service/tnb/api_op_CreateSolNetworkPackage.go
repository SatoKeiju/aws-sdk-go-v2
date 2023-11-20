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

// Creates a network package. A network package is a .zip file in CSAR (Cloud
// Service Archive) format defines the function packages you want to deploy and the
// Amazon Web Services infrastructure you want to deploy them on. For more
// information, see Network instances (https://docs.aws.amazon.com/tnb/latest/ug/network-instances.html)
// in the Amazon Web Services Telco Network Builder User Guide. A network package
// consists of a network service descriptor (NSD) file (required) and any
// additional files (optional), such as scripts specific to your needs. For
// example, if you have multiple function packages in your network package, you can
// use the NSD to define which network functions should run in certain VPCs,
// subnets, or EKS clusters. This request creates an empty network package
// container with an ID. Once you create a network package, you can upload the
// network package content using PutSolNetworkPackageContent (https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolNetworkPackageContent.html)
// .
func (c *Client) CreateSolNetworkPackage(ctx context.Context, params *CreateSolNetworkPackageInput, optFns ...func(*Options)) (*CreateSolNetworkPackageOutput, error) {
	if params == nil {
		params = &CreateSolNetworkPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSolNetworkPackage", params, optFns, c.addOperationCreateSolNetworkPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSolNetworkPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSolNetworkPackageInput struct {

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSolNetworkPackageOutput struct {

	// Network package ARN.
	//
	// This member is required.
	Arn *string

	// ID of the network package.
	//
	// This member is required.
	Id *string

	// Onboarding state of the network service descriptor in the network package.
	//
	// This member is required.
	NsdOnboardingState types.NsdOnboardingState

	// Operational state of the network service descriptor in the network package.
	//
	// This member is required.
	NsdOperationalState types.NsdOperationalState

	// Usage state of the network service descriptor in the network package.
	//
	// This member is required.
	NsdUsageState types.NsdUsageState

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSolNetworkPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSolNetworkPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSolNetworkPackage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSolNetworkPackage"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSolNetworkPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSolNetworkPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSolNetworkPackage",
	}
}
