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

// Creates an Amazon Aurora DB cluster from MySQL data stored in an Amazon S3
// bucket. Amazon RDS must be authorized to access the Amazon S3 bucket and the
// data must be created using the Percona XtraBackup utility as described in
// Migrating Data from MySQL by Using an Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Migrating.ExtMySQL.html#AuroraMySQL.Migrating.ExtMySQL.S3)
// in the Amazon Aurora User Guide. This action only restores the DB cluster, not
// the DB instances for that DB cluster. You must invoke the CreateDBInstance
// action to create DB instances for the restored DB cluster, specifying the
// identifier of the restored DB cluster in DBClusterIdentifier . You can create DB
// instances only after the RestoreDBClusterFromS3 action has completed and the DB
// cluster is available. For more information on Amazon Aurora, see What is Amazon
// Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
// The source DB engine must be MySQL.
func (c *Client) RestoreDBClusterFromS3(ctx context.Context, params *RestoreDBClusterFromS3Input, optFns ...func(*Options)) (*RestoreDBClusterFromS3Output, error) {
	if params == nil {
		params = &RestoreDBClusterFromS3Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBClusterFromS3", params, optFns, c.addOperationRestoreDBClusterFromS3Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBClusterFromS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBClusterFromS3Input struct {

	// The name of the DB cluster to create from the source data in the Amazon S3
	// bucket. This parameter isn't case-sensitive. Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - First character must be a letter.
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	// Example: my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// The name of the database engine to be used for this DB cluster. Valid Values:
	// aurora-mysql (for Aurora MySQL)
	//
	// This member is required.
	Engine *string

	// The name of the master user for the restored DB cluster. Constraints:
	//   - Must be 1 to 16 letters or numbers.
	//   - First character must be a letter.
	//   - Can't be a reserved word for the chosen database engine.
	//
	// This member is required.
	MasterUsername *string

	// The name of the Amazon S3 bucket that contains the data used to create the
	// Amazon Aurora DB cluster.
	//
	// This member is required.
	S3BucketName *string

	// The Amazon Resource Name (ARN) of the Amazon Web Services Identity and Access
	// Management (IAM) role that authorizes Amazon RDS to access the Amazon S3 bucket
	// on your behalf.
	//
	// This member is required.
	S3IngestionRoleArn *string

	// The identifier for the database engine that was backed up to create the files
	// stored in the Amazon S3 bucket. Valid Values: mysql
	//
	// This member is required.
	SourceEngine *string

	// The version of the database that the backup files were created from. MySQL
	// versions 5.7 and 8.0 are supported. Example: 5.7.40 , 8.0.28
	//
	// This member is required.
	SourceEngineVersion *string

	// A list of Availability Zones (AZs) where instances in the restored DB cluster
	// can be created.
	AvailabilityZones []string

	// The target backtrack window, in seconds. To disable backtracking, set this
	// value to 0. Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	// Default: 0 Constraints:
	//   - If specified, this value must be set to a number from 0 to 259,200 (72
	//   hours).
	BacktrackWindow *int64

	// The number of days for which automated backups of the restored DB cluster are
	// retained. You must specify a minimum value of 1. Default: 1 Constraints:
	//   - Must be a value from 1 to 35
	BackupRetentionPeriod *int32

	// A value that indicates that the restored DB cluster should be associated with
	// the specified CharacterSet.
	CharacterSetName *string

	// Specifies whether to copy all tags from the restored DB cluster to snapshots of
	// the restored DB cluster. The default is not to copy them.
	CopyTagsToSnapshot *bool

	// The name of the DB cluster parameter group to associate with the restored DB
	// cluster. If this argument is omitted, the default parameter group for the engine
	// version is used. Constraints:
	//   - If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string

	// A DB subnet group to associate with the restored DB cluster. Constraints: If
	// supplied, must match the name of an existing DBSubnetGroup. Example:
	// mydbsubnetgroup
	DBSubnetGroupName *string

	// The database name for the restored DB cluster.
	DatabaseName *string

	// Specifies whether to enable deletion protection for the DB cluster. The
	// database can't be deleted when deletion protection is enabled. By default,
	// deletion protection isn't enabled.
	DeletionProtection *bool

	// Specify the Active Directory directory ID to restore the DB cluster in. The
	// domain must be created prior to this operation. For Amazon Aurora DB clusters,
	// Amazon RDS can use Kerberos Authentication to authenticate users that connect to
	// the DB cluster. For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html)
	// in the Amazon Aurora User Guide.
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string

	// The list of logs that the restored DB cluster is to export to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. Aurora MySQL Possible
	// values are audit , error , general , and slowquery . For more information about
	// exporting CloudWatch Logs for Amazon Aurora, see Publishing Database Logs to
	// Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	EnableCloudwatchLogsExports []string

	// Specifies whether to enable mapping of Amazon Web Services Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping isn't
	// enabled. For more information, see IAM Database Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide.
	EnableIAMDatabaseAuthentication *bool

	// The version number of the database engine to use. To list all of the available
	// engine versions for aurora-mysql (Aurora MySQL), use the following command: aws
	// rds describe-db-engine-versions --engine aurora-mysql --query
	// "DBEngineVersions[].EngineVersion" Aurora MySQL Examples:
	// 5.7.mysql_aurora.2.12.0 , 8.0.mysql_aurora.3.04.0
	EngineVersion *string

	// The Amazon Web Services KMS key identifier for an encrypted DB cluster. The
	// Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or
	// alias name for the KMS key. To use a KMS key in a different Amazon Web Services
	// account, specify the key ARN or alias ARN. If the StorageEncrypted parameter is
	// enabled, and you do not specify a value for the KmsKeyId parameter, then Amazon
	// RDS will use your default KMS key. There is a default KMS key for your Amazon
	// Web Services account. Your Amazon Web Services account has a different default
	// KMS key for each Amazon Web Services Region.
	KmsKeyId *string

	// Specifies whether to manage the master user password with Amazon Web Services
	// Secrets Manager. For more information, see Password management with Amazon Web
	// Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html)
	// in the Amazon RDS User Guide and Password management with Amazon Web Services
	// Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html)
	// in the Amazon Aurora User Guide. Constraints:
	//   - Can't manage the master user password with Amazon Web Services Secrets
	//   Manager if MasterUserPassword is specified.
	ManageMasterUserPassword *bool

	// The password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@". Constraints:
	//   - Must contain from 8 to 41 characters.
	//   - Can't be specified if ManageMasterUserPassword is turned on.
	MasterUserPassword *string

	// The Amazon Web Services KMS key identifier to encrypt a secret that is
	// automatically generated and managed in Amazon Web Services Secrets Manager. This
	// setting is valid only if the master user password is managed by RDS in Amazon
	// Web Services Secrets Manager for the DB cluster. The Amazon Web Services KMS key
	// identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To
	// use a KMS key in a different Amazon Web Services account, specify the key ARN or
	// alias ARN. If you don't specify MasterUserSecretKmsKeyId , then the
	// aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a
	// different Amazon Web Services account, then you can't use the aws/secretsmanager
	// KMS key to encrypt the secret, and you must use a customer managed KMS key.
	// There is a default KMS key for your Amazon Web Services account. Your Amazon Web
	// Services account has a different default KMS key for each Amazon Web Services
	// Region.
	MasterUserSecretKmsKeyId *string

	// The network type of the DB cluster. Valid Values:
	//   - IPV4
	//   - DUAL
	// The network type is determined by the DBSubnetGroup specified for the DB
	// cluster. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the
	// IPv6 protocols ( DUAL ). For more information, see  Working with a DB instance
	// in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)
	// in the Amazon Aurora User Guide.
	NetworkType *string

	// A value that indicates that the restored DB cluster should be associated with
	// the specified option group. Permanent options can't be removed from an option
	// group. An option group can't be removed from a DB cluster once it is associated
	// with a DB cluster.
	OptionGroupName *string

	// The port number on which the instances in the restored DB cluster accept
	// connections. Default: 3306
	Port *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each Amazon
	// Web Services Region. To view the time blocks available, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow)
	// in the Amazon Aurora User Guide. Constraints:
	//   - Must be in the format hh24:mi-hh24:mi .
	//   - Must be in Universal Coordinated Time (UTC).
	//   - Must not conflict with the preferred maintenance window.
	//   - Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each Amazon
	// Web Services Region, occurring on a random day of the week. To see the time
	// blocks available, see Adjusting the Preferred Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// The prefix for all of the file names that contain the data used to create the
	// Amazon Aurora DB cluster. If you do not specify a SourceS3Prefix value, then the
	// Amazon Aurora DB cluster is created by using all of the files in the Amazon S3
	// bucket.
	S3Prefix *string

	// Contains the scaling configuration of an Aurora Serverless v2 DB cluster. For
	// more information, see Using Amazon Aurora Serverless v2 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html)
	// in the Amazon Aurora User Guide.
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// Specifies whether the restored DB cluster is encrypted.
	StorageEncrypted *bool

	// Specifies the storage type to be associated with the DB cluster. Valid Values:
	// aurora , aurora-iopt1 Default: aurora Valid for: Aurora DB clusters only
	StorageType *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []types.Tag

	// A list of EC2 VPC security groups to associate with the restored DB cluster.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBClusterFromS3Output struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster. For
	// an Amazon Aurora DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , PromoteReadReplicaDBCluster ,
	// RestoreDBClusterFromS3 , RestoreDBClusterFromSnapshot ,
	// RestoreDBClusterToPointInTime , StartDBCluster , and StopDBCluster . For a
	// Multi-AZ DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , RebootDBCluster ,
	// RestoreDBClusterFromSnapshot , and RestoreDBClusterToPointInTime . For more
	// information on Amazon Aurora DB clusters, see What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
	// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
	// see Multi-AZ deployments with two readable standby DB instances (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
	// in the Amazon RDS User Guide.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBClusterFromS3Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterFromS3{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterFromS3{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreDBClusterFromS3"); err != nil {
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
	if err = addOpRestoreDBClusterFromS3ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterFromS3(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBClusterFromS3(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreDBClusterFromS3",
	}
}
