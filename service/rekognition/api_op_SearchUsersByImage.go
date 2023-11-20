// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for UserIDs using a supplied image. It first detects the largest face
// in the image, and then searches a specified collection for matching UserIDs. The
// operation returns an array of UserIDs that match the face in the supplied image,
// ordered by similarity score with the highest similarity first. It also returns a
// bounding box for the face found in the input image. Information about faces
// detected in the supplied image, but not used for the search, is returned in an
// array of UnsearchedFace objects. If no valid face is detected in the image, the
// response will contain an empty UserMatches list and no SearchedFace object.
func (c *Client) SearchUsersByImage(ctx context.Context, params *SearchUsersByImageInput, optFns ...func(*Options)) (*SearchUsersByImageOutput, error) {
	if params == nil {
		params = &SearchUsersByImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchUsersByImage", params, optFns, c.addOperationSearchUsersByImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchUsersByImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchUsersByImageInput struct {

	// The ID of an existing collection containing the UserID.
	//
	// This member is required.
	CollectionId *string

	// Provides the input image either as bytes or an S3 object. You pass image bytes
	// to an Amazon Rekognition API operation by using the Bytes property. For
	// example, you would use the Bytes property to pass an image loaded from a local
	// file system. Image bytes passed by using the Bytes property must be
	// base64-encoded. Your code may not need to encode image bytes if you are using an
	// AWS SDK to call Amazon Rekognition API operations. For more information, see
	// Analyzing an Image Loaded from a Local File System in the Amazon Rekognition
	// Developer Guide. You pass images stored in an S3 bucket to an Amazon Rekognition
	// API operation by using the S3Object property. Images stored in an S3 bucket do
	// not need to be base64-encoded. The region for the S3 bucket containing the S3
	// object must match the region you use for Amazon Rekognition operations. If you
	// use the AWS CLI to call Amazon Rekognition operations, passing image bytes using
	// the Bytes property is not supported. You must first upload the image to an
	// Amazon S3 bucket and then call the operation using the S3Object property. For
	// Amazon Rekognition to process an S3 object, the user must have permission to
	// access the S3 object. For more information, see How Amazon Rekognition works
	// with IAM in the Amazon Rekognition Developer Guide.
	//
	// This member is required.
	Image *types.Image

	// Maximum number of UserIDs to return.
	MaxUsers *int32

	// A filter that specifies a quality bar for how much filtering is done to
	// identify faces. Filtered faces aren't searched for in the collection. The
	// default value is NONE.
	QualityFilter types.QualityFilter

	// Specifies the minimum confidence in the UserID match to return. Default value
	// is 80.
	UserMatchThreshold *float32

	noSmithyDocumentSerde
}

type SearchUsersByImageOutput struct {

	// Version number of the face detection model associated with the input collection
	// CollectionId.
	FaceModelVersion *string

	// A list of FaceDetail objects containing the BoundingBox for the largest face in
	// image, as well as the confidence in the bounding box, that was searched for
	// matches. If no valid face is detected in the image the response will contain no
	// SearchedFace object.
	SearchedFace *types.SearchedFaceDetails

	// List of UnsearchedFace objects. Contains the face details infered from the
	// specified image but not used for search. Contains reasons that describe why a
	// face wasn't used for Search.
	UnsearchedFaces []types.UnsearchedFace

	// An array of UserID objects that matched the input face, along with the
	// confidence in the match. The returned structure will be empty if there are no
	// matches. Returned if the SearchUsersByImageResponse action is successful.
	UserMatches []types.UserMatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchUsersByImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchUsersByImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchUsersByImage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchUsersByImage"); err != nil {
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
	if err = addOpSearchUsersByImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchUsersByImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchUsersByImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchUsersByImage",
	}
}
