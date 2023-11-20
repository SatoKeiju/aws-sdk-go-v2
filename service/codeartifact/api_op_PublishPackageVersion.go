// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Creates a new package version containing one or more assets (or files). The
// unfinished flag can be used to keep the package version in the Unfinished state
// until all of its assets have been uploaded (see Package version status (https://docs.aws.amazon.com/codeartifact/latest/ug/packages-overview.html#package-version-status.html#package-version-status)
// in the CodeArtifact user guide). To set the package version’s status to
// Published , omit the unfinished flag when uploading the final asset, or set the
// status using UpdatePackageVersionStatus (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdatePackageVersionsStatus.html)
// . Once a package version’s status is set to Published , it cannot change back to
// Unfinished . Only generic packages can be published using this API. For more
// information, see Using generic packages (https://docs.aws.amazon.com/codeartifact/latest/ug/using-generic.html)
// in the CodeArtifact User Guide.
func (c *Client) PublishPackageVersion(ctx context.Context, params *PublishPackageVersionInput, optFns ...func(*Options)) (*PublishPackageVersionOutput, error) {
	if params == nil {
		params = &PublishPackageVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishPackageVersion", params, optFns, c.addOperationPublishPackageVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishPackageVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishPackageVersionInput struct {

	// The content of the asset to publish.
	//
	// This member is required.
	AssetContent io.Reader

	// The name of the asset to publish. Asset names can include Unicode letters and
	// numbers, and the following special characters: ~ ! @ ^ & ( ) - ` _ + [ ] { } ;
	// , . `
	//
	// This member is required.
	AssetName *string

	// The SHA256 hash of the assetContent to publish. This value must be calculated
	// by the caller and provided with the request (see Publishing a generic package (https://docs.aws.amazon.com/codeartifact/latest/ug/using-generic.html#publishing-generic-packages)
	// in the CodeArtifact User Guide). This value is used as an integrity check to
	// verify that the assetContent has not changed after it was originally sent.
	//
	// This member is required.
	AssetSHA256 *string

	// The name of the domain that contains the repository that contains the package
	// version to publish.
	//
	// This member is required.
	Domain *string

	// A format that specifies the type of the package version with the requested
	// asset file. The only supported value is generic .
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package version to publish.
	//
	// This member is required.
	Package *string

	// The package version to publish (for example, 3.5.2 ).
	//
	// This member is required.
	PackageVersion *string

	// The name of the repository that the package version will be published to.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	DomainOwner *string

	// The namespace of the package version to publish.
	Namespace *string

	// Specifies whether the package version should remain in the unfinished state. If
	// omitted, the package version status will be set to Published (see Package
	// version status (https://docs.aws.amazon.com/codeartifact/latest/ug/packages-overview.html#package-version-status)
	// in the CodeArtifact User Guide). Valid values: unfinished
	Unfinished *bool

	noSmithyDocumentSerde
}

type PublishPackageVersionOutput struct {

	// An AssetSummary (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_AssetSummary.html)
	// for the published asset.
	Asset *types.AssetSummary

	// The format of the package version.
	Format types.PackageFormat

	// The namespace of the package version.
	Namespace *string

	// The name of the package.
	Package *string

	// A string that contains the status of the package version. For more information,
	// see Package version status (https://docs.aws.amazon.com/codeartifact/latest/ug/packages-overview.html#package-version-status.html#package-version-status)
	// in the CodeArtifact User Guide.
	Status types.PackageVersionStatus

	// The version of the package.
	Version *string

	// The revision of the package version.
	VersionRevision *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPublishPackageVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPublishPackageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPublishPackageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PublishPackageVersion"); err != nil {
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
	if err = addOpPublishPackageVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublishPackageVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPublishPackageVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PublishPackageVersion",
	}
}
