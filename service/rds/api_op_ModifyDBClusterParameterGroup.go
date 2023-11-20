// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the parameters of a DB cluster parameter group. To modify more than
// one parameter, submit a list of the following: ParameterName , ParameterValue ,
// and ApplyMethod . A maximum of 20 parameters can be modified in a single
// request. After you create a DB cluster parameter group, you should wait at least
// 5 minutes before creating your first DB cluster that uses that DB cluster
// parameter group as the default parameter group. This allows Amazon RDS to fully
// complete the create action before the parameter group is used as the default for
// a new DB cluster. This is especially important for parameters that are critical
// when creating the default database for a DB cluster, such as the character set
// for the default database defined by the character_set_database parameter. You
// can use the Parameter Groups option of the Amazon RDS console (https://console.aws.amazon.com/rds/)
// or the DescribeDBClusterParameters operation to verify that your DB cluster
// parameter group has been created or modified. If the modified DB cluster
// parameter group is used by an Aurora Serverless v1 cluster, Aurora applies the
// update immediately. The cluster restart might interrupt your workload. In that
// case, your application must reopen any connections and retry any transactions
// that were active when the parameter changes took effect. For more information on
// Amazon Aurora DB clusters, see What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) ModifyDBClusterParameterGroup(ctx context.Context, params *ModifyDBClusterParameterGroupInput, optFns ...func(*Options)) (*ModifyDBClusterParameterGroupOutput, error) {
	if params == nil {
		params = &ModifyDBClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterParameterGroup", params, optFns, c.addOperationModifyDBClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterParameterGroupInput struct {

	// The name of the DB cluster parameter group to modify.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// A list of parameters in the DB cluster parameter group to modify. Valid Values
	// (for the application method): immediate | pending-reboot You can use the
	// immediate value with dynamic parameters only. You can use the pending-reboot
	// value for both dynamic and static parameters. When the application method is
	// immediate , changes to dynamic parameters are applied immediately to the DB
	// clusters associated with the parameter group. When the application method is
	// pending-reboot , changes to dynamic and static parameters are applied after a
	// reboot without failover to the DB clusters associated with the parameter group.
	//
	// This member is required.
	Parameters []types.Parameter

	noSmithyDocumentSerde
}

type ModifyDBClusterParameterGroupOutput struct {

	// The name of the DB cluster parameter group. Constraints:
	//   - Must be 1 to 255 letters or numbers.
	//   - First character must be a letter
	//   - Can't end with a hyphen or contain two consecutive hyphens
	// This value is stored as a lowercase string.
	DBClusterParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDBClusterParameterGroup"); err != nil {
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
	if err = addOpModifyDBClusterParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDBClusterParameterGroup",
	}
}
