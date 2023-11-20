// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Edit or change an OpsItem. You must have permission in Identity and Access
// Management (IAM) to update an OpsItem. For more information, see Set up
// OpsCenter (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setup.html)
// in the Amazon Web Services Systems Manager User Guide. Operations engineers and
// IT professionals use Amazon Web Services Systems Manager OpsCenter to view,
// investigate, and remediate operational issues impacting the performance and
// health of their Amazon Web Services resources. For more information, see
// OpsCenter (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html)
// in the Amazon Web Services Systems Manager User Guide.
func (c *Client) UpdateOpsItem(ctx context.Context, params *UpdateOpsItemInput, optFns ...func(*Options)) (*UpdateOpsItemOutput, error) {
	if params == nil {
		params = &UpdateOpsItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateOpsItem", params, optFns, c.addOperationUpdateOpsItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateOpsItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateOpsItemInput struct {

	// The ID of the OpsItem.
	//
	// This member is required.
	OpsItemId *string

	// The time a runbook workflow ended. Currently reported only for the OpsItem type
	// /aws/changerequest .
	ActualEndTime *time.Time

	// The time a runbook workflow started. Currently reported only for the OpsItem
	// type /aws/changerequest .
	ActualStartTime *time.Time

	// Specify a new category for an OpsItem.
	Category *string

	// User-defined text that contains information about the OpsItem, in Markdown
	// format.
	Description *string

	// The Amazon Resource Name (ARN) of an SNS topic where notifications are sent
	// when this OpsItem is edited or changed.
	Notifications []types.OpsItemNotification

	// Add new keys or edit existing key-value pairs of the OperationalData map in the
	// OpsItem object. Operational data is custom data that provides useful reference
	// details about the OpsItem. For example, you can specify log files, error
	// strings, license keys, troubleshooting tips, or other relevant data. You enter
	// operational data as key-value pairs. The key has a maximum length of 128
	// characters. The value has a maximum size of 20 KB. Operational data keys can't
	// begin with the following: amazon , aws , amzn , ssm , /amazon , /aws , /amzn ,
	// /ssm . You can choose to make the data searchable by other users in the account
	// or you can restrict search access. Searchable data means that all users with
	// access to the OpsItem Overview page (as provided by the DescribeOpsItems API
	// operation) can view and search on the specified data. Operational data that
	// isn't searchable is only viewable by users who have access to the OpsItem (as
	// provided by the GetOpsItem API operation). Use the /aws/resources key in
	// OperationalData to specify a related resource in the request. Use the
	// /aws/automations key in OperationalData to associate an Automation runbook with
	// the OpsItem. To view Amazon Web Services CLI example commands that use these
	// keys, see Creating OpsItems manually (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-manually-create-OpsItems.html)
	// in the Amazon Web Services Systems Manager User Guide.
	OperationalData map[string]types.OpsItemDataValue

	// Keys that you want to remove from the OperationalData map.
	OperationalDataToDelete []string

	// The OpsItem Amazon Resource Name (ARN).
	OpsItemArn *string

	// The time specified in a change request for a runbook workflow to end. Currently
	// supported only for the OpsItem type /aws/changerequest .
	PlannedEndTime *time.Time

	// The time specified in a change request for a runbook workflow to start.
	// Currently supported only for the OpsItem type /aws/changerequest .
	PlannedStartTime *time.Time

	// The importance of this OpsItem in relation to other OpsItems in the system.
	Priority *int32

	// One or more OpsItems that share something in common with the current OpsItems.
	// For example, related OpsItems can include OpsItems with similar error messages,
	// impacted resources, or statuses for the impacted resource.
	RelatedOpsItems []types.RelatedOpsItem

	// Specify a new severity for an OpsItem.
	Severity *string

	// The OpsItem status. Status can be Open , In Progress , or Resolved . For more
	// information, see Editing OpsItem details (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-working-with-OpsItems-editing-details.html)
	// in the Amazon Web Services Systems Manager User Guide.
	Status types.OpsItemStatus

	// A short heading that describes the nature of the OpsItem and the impacted
	// resource.
	Title *string

	noSmithyDocumentSerde
}

type UpdateOpsItemOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateOpsItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateOpsItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateOpsItem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateOpsItem"); err != nil {
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
	if err = addOpUpdateOpsItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateOpsItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateOpsItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateOpsItem",
	}
}
