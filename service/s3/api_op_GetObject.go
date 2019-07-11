// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectRequest
type GetObjectInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Return the object only if its entity tag (ETag) is the same as the one specified,
	// otherwise return a 412 (precondition failed).
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`

	// Return the object only if it has been modified since the specified time,
	// otherwise return a 304 (not modified).
	IfModifiedSince *time.Time `location:"header" locationName:"If-Modified-Since" type:"timestamp" timestampFormat:"rfc822"`

	// Return the object only if its entity tag (ETag) is different from the one
	// specified, otherwise return a 304 (not modified).
	IfNoneMatch *string `location:"header" locationName:"If-None-Match" type:"string"`

	// Return the object only if it has not been modified since the specified time,
	// otherwise return a 412 (precondition failed).
	IfUnmodifiedSince *time.Time `location:"header" locationName:"If-Unmodified-Since" type:"timestamp" timestampFormat:"rfc822"`

	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Part number of the object being read. This is a positive integer between
	// 1 and 10,000. Effectively performs a 'ranged' GET request for the part specified.
	// Useful for downloading just a part of an object.
	PartNumber *int64 `location:"querystring" locationName:"partNumber" type:"integer"`

	// Downloads the specified range bytes of an object. For more information about
	// the HTTP Range header, go to http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35.
	Range *string `location:"header" locationName:"Range" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Sets the Cache-Control header of the response.
	ResponseCacheControl *string `location:"querystring" locationName:"response-cache-control" type:"string"`

	// Sets the Content-Disposition header of the response
	ResponseContentDisposition *string `location:"querystring" locationName:"response-content-disposition" type:"string"`

	// Sets the Content-Encoding header of the response.
	ResponseContentEncoding *string `location:"querystring" locationName:"response-content-encoding" type:"string"`

	// Sets the Content-Language header of the response.
	ResponseContentLanguage *string `location:"querystring" locationName:"response-content-language" type:"string"`

	// Sets the Content-Type header of the response.
	ResponseContentType *string `location:"querystring" locationName:"response-content-type" type:"string"`

	// Sets the Expires header of the response.
	ResponseExpires *time.Time `location:"querystring" locationName:"response-expires" type:"timestamp" timestampFormat:"iso8601"`

	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// does not store the encryption key. The key must be appropriate for use with
	// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// VersionId used to reference a specific version of the object.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *GetObjectInput) getSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IfMatch != nil {
		v := *s.IfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Match", protocol.StringValue(v), metadata)
	}
	if s.IfModifiedSince != nil {
		v := *s.IfModifiedSince

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Modified-Since", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.IfNoneMatch != nil {
		v := *s.IfNoneMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-None-Match", protocol.StringValue(v), metadata)
	}
	if s.IfUnmodifiedSince != nil {
		v := *s.IfUnmodifiedSince

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Unmodified-Since", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.Range != nil {
		v := *s.Range

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Range", protocol.StringValue(v), metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKey != nil {
		v := *s.SSECustomerKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.PartNumber != nil {
		v := *s.PartNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "partNumber", protocol.Int64Value(v), metadata)
	}
	if s.ResponseCacheControl != nil {
		v := *s.ResponseCacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-cache-control", protocol.StringValue(v), metadata)
	}
	if s.ResponseContentDisposition != nil {
		v := *s.ResponseContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-content-disposition", protocol.StringValue(v), metadata)
	}
	if s.ResponseContentEncoding != nil {
		v := *s.ResponseContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-content-encoding", protocol.StringValue(v), metadata)
	}
	if s.ResponseContentLanguage != nil {
		v := *s.ResponseContentLanguage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-content-language", protocol.StringValue(v), metadata)
	}
	if s.ResponseContentType != nil {
		v := *s.ResponseContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-content-type", protocol.StringValue(v), metadata)
	}
	if s.ResponseExpires != nil {
		v := *s.ResponseExpires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-expires", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectOutput
type GetObjectOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	AcceptRanges *string `location:"header" locationName:"accept-ranges" type:"string"`

	// Object data.
	Body io.ReadCloser `type:"blob"`

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// Size of the body in bytes.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The portion of the object returned in the response.
	ContentRange *string `location:"header" locationName:"Content-Range" type:"string"`

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool `location:"header" locationName:"x-amz-delete-marker" type:"boolean"`

	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *string `location:"header" locationName:"Expires" type:"string"`

	// Last modified date of the object
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp" timestampFormat:"rfc822"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP,
	// you can create metadata whose values are not legal HTTP headers.
	MissingMeta *int64 `location:"header" locationName:"x-amz-missing-meta" type:"integer"`

	// Indicates whether this object has an active legal hold. This field is only
	// returned if you have permission to view an object's legal hold status.
	ObjectLockLegalHoldStatus ObjectLockLegalHoldStatus `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"true"`

	// The Object Lock mode currently in place for this object.
	ObjectLockMode ObjectLockMode `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"true"`

	// The date and time when this object's Object Lock will expire.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"rfc822"`

	// The count of parts this object has.
	PartsCount *int64 `location:"header" locationName:"x-amz-mp-parts-count" type:"integer"`

	ReplicationStatus ReplicationStatus `location:"header" locationName:"x-amz-replication-status" type:"string" enum:"true"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// Provides information about object restoration operation and expiration time
	// of the restored object copy.
	Restore *string `location:"header" locationName:"x-amz-restore" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	StorageClass StorageClass `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"true"`

	// The number of tags, if any, on the object.
	TagCount *int64 `location:"header" locationName:"x-amz-tagging-count" type:"integer"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s GetObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AcceptRanges != nil {
		v := *s.AcceptRanges

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "accept-ranges", protocol.StringValue(v), metadata)
	}
	if s.CacheControl != nil {
		v := *s.CacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Cache-Control", protocol.StringValue(v), metadata)
	}
	if s.ContentDisposition != nil {
		v := *s.ContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Disposition", protocol.StringValue(v), metadata)
	}
	if s.ContentEncoding != nil {
		v := *s.ContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Encoding", protocol.StringValue(v), metadata)
	}
	if s.ContentLanguage != nil {
		v := *s.ContentLanguage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Language", protocol.StringValue(v), metadata)
	}
	if s.ContentLength != nil {
		v := *s.ContentLength

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Length", protocol.Int64Value(v), metadata)
	}
	if s.ContentRange != nil {
		v := *s.ContentRange

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Range", protocol.StringValue(v), metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue(v), metadata)
	}
	if s.DeleteMarker != nil {
		v := *s.DeleteMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-delete-marker", protocol.BoolValue(v), metadata)
	}
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Expiration != nil {
		v := *s.Expiration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-expiration", protocol.StringValue(v), metadata)
	}
	if s.Expires != nil {
		v := *s.Expires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Expires", protocol.StringValue(v), metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Last-Modified", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.MissingMeta != nil {
		v := *s.MissingMeta

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-missing-meta", protocol.Int64Value(v), metadata)
	}
	if len(s.ObjectLockLegalHoldStatus) > 0 {
		v := s.ObjectLockLegalHoldStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-legal-hold", v, metadata)
	}
	if len(s.ObjectLockMode) > 0 {
		v := s.ObjectLockMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-mode", v, metadata)
	}
	if s.ObjectLockRetainUntilDate != nil {
		v := *s.ObjectLockRetainUntilDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-retain-until-date", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.PartsCount != nil {
		v := *s.PartsCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-mp-parts-count", protocol.Int64Value(v), metadata)
	}
	if len(s.ReplicationStatus) > 0 {
		v := s.ReplicationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-replication-status", v, metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.Restore != nil {
		v := *s.Restore

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-restore", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if len(s.StorageClass) > 0 {
		v := s.StorageClass

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-storage-class", v, metadata)
	}
	if s.TagCount != nil {
		v := *s.TagCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-tagging-count", protocol.Int64Value(v), metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	if s.WebsiteRedirectLocation != nil {
		v := *s.WebsiteRedirectLocation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-website-redirect-location", protocol.StringValue(v), metadata)
	}
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.HeadersTarget, "x-amz-meta-", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.StringValue(v1))
		}
		ms0.End()

	}
	// Skipping Body Output type's body not valid.
	return nil
}

const opGetObject = "GetObject"

// GetObjectRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Retrieves objects from Amazon S3.
//
//    // Example sending a request using GetObjectRequest.
//    req := client.GetObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObject
func (c *Client) GetObjectRequest(input *GetObjectInput) GetObjectRequest {
	op := &aws.Operation{
		Name:       opGetObject,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &GetObjectInput{}
	}

	req := c.newRequest(op, input, &GetObjectOutput{})
	return GetObjectRequest{Request: req, Input: input, Copy: c.GetObjectRequest}
}

// GetObjectRequest is the request type for the
// GetObject API operation.
type GetObjectRequest struct {
	*aws.Request
	Input *GetObjectInput
	Copy  func(*GetObjectInput) GetObjectRequest
}

// Send marshals and sends the GetObject API request.
func (r GetObjectRequest) Send(ctx context.Context) (*GetObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectResponse{
		GetObjectOutput: r.Request.Data.(*GetObjectOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectResponse is the response type for the
// GetObject API operation.
type GetObjectResponse struct {
	*GetObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObject request.
func (r *GetObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
