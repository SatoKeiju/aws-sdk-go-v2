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

// Detects faces in the input image and adds them to the specified collection.
// Amazon Rekognition doesn't save the actual faces that are detected. Instead, the
// underlying detection algorithm first detects the faces in the input image. For
// each face, the algorithm extracts facial features into a feature vector, and
// stores it in the backend database. Amazon Rekognition uses feature vectors when
// it performs face match and search operations using the SearchFaces and
// SearchFacesByImage operations. For more information, see Adding faces to a
// collection in the Amazon Rekognition Developer Guide. To get the number of faces
// in a collection, call DescribeCollection . If you're using version 1.0 of the
// face detection model, IndexFaces indexes the 15 largest faces in the input
// image. Later versions of the face detection model index the 100 largest faces in
// the input image. If you're using version 4 or later of the face model, image
// orientation information is not returned in the OrientationCorrection field. To
// determine which version of the model you're using, call DescribeCollection and
// supply the collection ID. You can also get the model version from the value of
// FaceModelVersion in the response from IndexFaces For more information, see
// Model Versioning in the Amazon Rekognition Developer Guide. If you provide the
// optional ExternalImageId for the input image you provided, Amazon Rekognition
// associates this ID with all faces that it detects. When you call the ListFaces
// operation, the response returns the external ID. You can use this external image
// ID to create a client-side index to associate the faces with each image. You can
// then use the index to find all faces in an image. You can specify the maximum
// number of faces to index with the MaxFaces input parameter. This is useful when
// you want to index the largest faces in an image and don't want to index smaller
// faces, such as those belonging to people standing in the background. The
// QualityFilter input parameter allows you to filter out detected faces that don’t
// meet a required quality bar. The quality bar is based on a variety of common use
// cases. By default, IndexFaces chooses the quality bar that's used to filter
// faces. You can also explicitly choose the quality bar. Use QualityFilter , to
// set the quality bar by specifying LOW , MEDIUM , or HIGH . If you do not want to
// filter detected faces, specify NONE . To use quality filtering, you need a
// collection associated with version 3 of the face model or higher. To get the
// version of the face model associated with a collection, call DescribeCollection
// . Information about faces detected in an image, but not indexed, is returned in
// an array of UnindexedFace objects, UnindexedFaces . Faces aren't indexed for
// reasons such as:
//   - The number of faces detected exceeds the value of the MaxFaces request
//     parameter.
//   - The face is too small compared to the image dimensions.
//   - The face is too blurry.
//   - The image is too dark.
//   - The face has an extreme pose.
//   - The face doesn’t have enough detail to be suitable for face search.
//
// In response, the IndexFaces operation returns an array of metadata for all
// detected faces, FaceRecords . This includes:
//   - The bounding box, BoundingBox , of the detected face.
//   - A confidence value, Confidence , which indicates the confidence that the
//     bounding box contains a face.
//   - A face ID, FaceId , assigned by the service for each face that's detected
//     and stored.
//   - An image ID, ImageId , assigned by the service for the input image.
//
// If you request ALL or specific facial attributes (e.g., FACE_OCCLUDED ) by using
// the detectionAttributes parameter, Amazon Rekognition returns detailed facial
// attributes, such as facial landmarks (for example, location of eye and mouth),
// facial occlusion, and other facial attributes. If you provide the same image,
// specify the same collection, and use the same external ID in the IndexFaces
// operation, Amazon Rekognition doesn't save duplicate face metadata. The input
// image is passed either as base64-encoded image bytes, or as a reference to an
// image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, passing image bytes isn't supported. The image must be formatted as
// a PNG or JPEG file. This operation requires permissions to perform the
// rekognition:IndexFaces action.
func (c *Client) IndexFaces(ctx context.Context, params *IndexFacesInput, optFns ...func(*Options)) (*IndexFacesOutput, error) {
	if params == nil {
		params = &IndexFacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IndexFaces", params, optFns, c.addOperationIndexFacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IndexFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IndexFacesInput struct {

	// The ID of an existing collection to which you want to add the faces that are
	// detected in the input images.
	//
	// This member is required.
	CollectionId *string

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes isn't
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	Image *types.Image

	// An array of facial attributes you want to be returned. A DEFAULT subset of
	// facial attributes - BoundingBox , Confidence , Pose , Quality , and Landmarks -
	// will always be returned. You can request for specific facial attributes (in
	// addition to the default list) - by using ["DEFAULT", "FACE_OCCLUDED"] or just
	// ["FACE_OCCLUDED"] . You can request for all facial attributes by using ["ALL"] .
	// Requesting more attributes may increase response time. If you provide both,
	// ["ALL", "DEFAULT"] , the service uses a logical AND operator to determine which
	// attributes to return (in this case, all attributes).
	DetectionAttributes []types.Attribute

	// The ID you want to assign to all the faces detected in the image.
	ExternalImageId *string

	// The maximum number of faces to index. The value of MaxFaces must be greater
	// than or equal to 1. IndexFaces returns no more than 100 detected faces in an
	// image, even if you specify a larger value for MaxFaces . If IndexFaces detects
	// more faces than the value of MaxFaces , the faces with the lowest quality are
	// filtered out first. If there are still more faces than the value of MaxFaces ,
	// the faces with the smallest bounding boxes are filtered out (up to the number
	// that's needed to satisfy the value of MaxFaces ). Information about the
	// unindexed faces is available in the UnindexedFaces array. The faces that are
	// returned by IndexFaces are sorted by the largest face bounding box size to the
	// smallest size, in descending order. MaxFaces can be used with a collection
	// associated with any version of the face model.
	MaxFaces *int32

	// A filter that specifies a quality bar for how much filtering is done to
	// identify faces. Filtered faces aren't indexed. If you specify AUTO , Amazon
	// Rekognition chooses the quality bar. If you specify LOW , MEDIUM , or HIGH ,
	// filtering removes all faces that don’t meet the chosen quality bar. The default
	// value is AUTO . The quality bar is based on a variety of common use cases.
	// Low-quality detections can occur for a number of reasons. Some examples are an
	// object that's misidentified as a face, a face that's too blurry, or a face with
	// a pose that's too extreme to use. If you specify NONE , no filtering is
	// performed. To use quality filtering, the collection you are using must be
	// associated with version 3 of the face model or higher.
	QualityFilter types.QualityFilter

	noSmithyDocumentSerde
}

type IndexFacesOutput struct {

	// The version number of the face detection model that's associated with the input
	// collection ( CollectionId ).
	FaceModelVersion *string

	// An array of faces detected and added to the collection. For more information,
	// see Searching Faces in a Collection in the Amazon Rekognition Developer Guide.
	FaceRecords []types.FaceRecord

	// If your collection is associated with a face detection model that's later than
	// version 3.0, the value of OrientationCorrection is always null and no
	// orientation information is returned. If your collection is associated with a
	// face detection model that's version 3.0 or earlier, the following applies:
	//   - If the input image is in .jpeg format, it might contain exchangeable image
	//   file format (Exif) metadata that includes the image's orientation. Amazon
	//   Rekognition uses this orientation information to perform image correction - the
	//   bounding box coordinates are translated to represent object locations after the
	//   orientation information in the Exif metadata is used to correct the image
	//   orientation. Images in .png format don't contain Exif metadata. The value of
	//   OrientationCorrection is null.
	//   - If the image doesn't contain orientation information in its Exif metadata,
	//   Amazon Rekognition returns an estimated orientation (ROTATE_0, ROTATE_90,
	//   ROTATE_180, ROTATE_270). Amazon Rekognition doesn’t perform image correction for
	//   images. The bounding box coordinates aren't translated and represent the object
	//   locations before the image is rotated.
	// Bounding box information is returned in the FaceRecords array. You can get the
	// version of the face detection model by calling DescribeCollection .
	OrientationCorrection types.OrientationCorrection

	// An array of faces that were detected in the image but weren't indexed. They
	// weren't indexed because the quality filter identified them as low quality, or
	// the MaxFaces request parameter filtered them out. To use the quality filter,
	// you specify the QualityFilter request parameter.
	UnindexedFaces []types.UnindexedFace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationIndexFacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpIndexFaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpIndexFaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "IndexFaces"); err != nil {
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
	if err = addOpIndexFacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIndexFaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIndexFaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "IndexFaces",
	}
}
