// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This target has been deprecated. Given a data set type and a from date,
// asynchronously publishes the requested customer support data to the specified S3
// bucket and notifies the specified SNS topic once the data is available. Returns
// a unique request identifier that can be used to correlate requests with
// notifications from the SNS topic. Data sets will be published in comma-separated
// values (CSV) format with the file name
// {data_set_type}_YYYY-MM-DD'T'HH-mm-ss'Z'.csv. If a file with the same name
// already exists (e.g. if the same data set is requested twice), the original file
// will be overwritten by the new file. Requires a Role with an attached
// permissions policy providing Allow permissions for the following actions:
// s3:PutObject, s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish,
// iam:GetRolePolicy.
//
// Deprecated: This target has been deprecated. As of December 2022 Product
// Support Connection is no longer supported.
func (c *Client) StartSupportDataExport(ctx context.Context, params *StartSupportDataExportInput, optFns ...func(*Options)) (*StartSupportDataExportOutput, error) {
	if params == nil {
		params = &StartSupportDataExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSupportDataExport", params, optFns, c.addOperationStartSupportDataExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSupportDataExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This target has been deprecated. Container for the parameters to the
// StartSupportDataExport operation.
type StartSupportDataExportInput struct {

	// This target has been deprecated. Specifies the data set type to be written to
	// the output csv file. The data set types customer_support_contacts_data and
	// test_customer_support_contacts_data both result in a csv file containing the
	// following fields: Product Id, Product Code, Customer Guid, Subscription Guid,
	// Subscription Start Date, Organization, AWS Account Id, Given Name, Surname,
	// Telephone Number, Email, Title, Country Code, ZIP Code, Operation Type, and
	// Operation Time.
	//   - customer_support_contacts_data Customer support contact data. The data set
	//   will contain all changes (Creates, Updates, and Deletes) to customer support
	//   contact data from the date specified in the from_date parameter.
	//   - test_customer_support_contacts_data An example data set containing static
	//   test data in the same format as customer_support_contacts_data
	//
	// This member is required.
	DataSetType types.SupportDataSetType

	// This target has been deprecated. The name (friendly name, not ARN) of the
	// destination S3 bucket.
	//
	// This member is required.
	DestinationS3BucketName *string

	// This target has been deprecated. The start date from which to retrieve the data
	// set in UTC. This parameter only affects the customer_support_contacts_data data
	// set type.
	//
	// This member is required.
	FromDate *time.Time

	// This target has been deprecated. The Amazon Resource Name (ARN) of the Role
	// with an attached permissions policy to interact with the provided AWS services.
	//
	// This member is required.
	RoleNameArn *string

	// This target has been deprecated. Amazon Resource Name (ARN) for the SNS Topic
	// that will be notified when the data set has been published or if an error has
	// occurred.
	//
	// This member is required.
	SnsTopicArn *string

	// This target has been deprecated. (Optional) Key-value pairs which will be
	// returned, unmodified, in the Amazon SNS notification message and the data set
	// metadata file.
	CustomerDefinedValues map[string]string

	// This target has been deprecated. (Optional) The desired S3 prefix for the
	// published data set, similar to a directory path in standard file systems. For
	// example, if given the bucket name "mybucket" and the prefix
	// "myprefix/mydatasets", the output file "outputfile" would be published to
	// "s3://mybucket/myprefix/mydatasets/outputfile". If the prefix directory
	// structure does not exist, it will be created. If no prefix is provided, the data
	// set will be published to the S3 bucket root.
	DestinationS3Prefix *string

	noSmithyDocumentSerde
}

// This target has been deprecated. Container for the result of the
// StartSupportDataExport operation.
type StartSupportDataExportOutput struct {

	// This target has been deprecated. A unique identifier representing a specific
	// request to the StartSupportDataExport operation. This identifier can be used to
	// correlate a request with notifications from the SNS topic.
	DataSetRequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSupportDataExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartSupportDataExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSupportDataExport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartSupportDataExport"); err != nil {
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
	if err = addOpStartSupportDataExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSupportDataExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartSupportDataExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartSupportDataExport",
	}
}
