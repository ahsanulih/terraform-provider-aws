// Code generated by internal/generate/awsclient/main.go; DO NOT EDIT.
package conns

import (
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/fis"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/ivschat"
	"github.com/aws/aws-sdk-go-v2/service/kendra"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	s3control_sdkv2 "github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	ssm_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
	"github.com/aws/aws-sdk-go/service/account"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplifybackend"
	"github.com/aws/aws-sdk-go/service/amplifyuibuilder"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/appconfig"
	"github.com/aws/aws-sdk-go/service/appconfigdata"
	"github.com/aws/aws-sdk-go/service/appflow"
	"github.com/aws/aws-sdk-go/service/appintegrationsservice"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationcostprofiler"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go/service/applicationinsights"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appregistry"
	"github.com/aws/aws-sdk-go/service/apprunner"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/auditmanager"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/backupgateway"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/aws/aws-sdk-go/service/billingconductor"
	"github.com/aws/aws-sdk-go/service/braket"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/chime"
	"github.com/aws/aws-sdk-go/service/chimesdkidentity"
	"github.com/aws/aws-sdk-go/service/chimesdkmeetings"
	"github.com/aws/aws-sdk-go/service/chimesdkmessaging"
	"github.com/aws/aws-sdk-go/service/cloud9"
	"github.com/aws/aws-sdk-go/service/cloudcontrolapi"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchevidently"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchrum"
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/codestar"
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"github.com/aws/aws-sdk-go/service/codestarnotifications"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connectcontactlens"
	"github.com/aws/aws-sdk-go/service/connectparticipant"
	"github.com/aws/aws-sdk-go/service/connectwisdomservice"
	"github.com/aws/aws-sdk-go/service/controltower"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/customerprofiles"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/detective"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devopsguru"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/drs"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/ebs"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecrpublic"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/emrcontainers"
	"github.com/aws/aws-sdk-go/service/emrserverless"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/finspace"
	"github.com/aws/aws-sdk-go/service/finspacedata"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/fms"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/gamelift"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/gluedatabrew"
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/greengrassv2"
	"github.com/aws/aws-sdk-go/service/groundstation"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/healthlake"
	"github.com/aws/aws-sdk-go/service/honeycode"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/iotdeviceadvisor"
	"github.com/aws/aws-sdk-go/service/iotevents"
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
	"github.com/aws/aws-sdk-go/service/iotfleethub"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"github.com/aws/aws-sdk-go/service/iotsitewise"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/iottwinmaker"
	"github.com/aws/aws-sdk-go/service/iotwireless"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/kafkaconnect"
	"github.com/aws/aws-sdk-go/service/keyspaces"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go/service/lexmodelsv2"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go/service/lexruntimev2"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/locationservice"
	"github.com/aws/aws-sdk-go/service/lookoutequipment"
	"github.com/aws/aws-sdk-go/service/lookoutforvision"
	"github.com/aws/aws-sdk-go/service/lookoutmetrics"
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"github.com/aws/aws-sdk-go/service/macie"
	"github.com/aws/aws-sdk-go/service/macie2"
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"github.com/aws/aws-sdk-go/service/managedgrafana"
	"github.com/aws/aws-sdk-go/service/marketplacecatalog"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/memorydb"
	"github.com/aws/aws-sdk-go/service/mgn"
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go/service/migrationhubrefactorspaces"
	"github.com/aws/aws-sdk-go/service/migrationhubstrategyrecommendations"
	"github.com/aws/aws-sdk-go/service/mobile"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/mwaa"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/networkfirewall"
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"github.com/aws/aws-sdk-go/service/nimblestudio"
	"github.com/aws/aws-sdk-go/service/opensearchservice"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/outposts"
	"github.com/aws/aws-sdk-go/service/panorama"
	"github.com/aws/aws-sdk-go/service/personalize"
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"github.com/aws/aws-sdk-go/service/pi"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/prometheusservice"
	"github.com/aws/aws-sdk-go/service/proton"
	"github.com/aws/aws-sdk-go/service/qldb"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/aws/aws-sdk-go/service/recyclebin"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
	"github.com/aws/aws-sdk-go/service/redshiftserverless"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/resiliencehub"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/robomaker"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53recoverycluster"
	"github.com/aws/aws-sdk-go/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/aws/aws-sdk-go/service/s3outposts"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/sagemakeredgemanager"
	"github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go/service/savingsplans"
	"github.com/aws/aws-sdk-go/service/schemas"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/aws/aws-sdk-go/service/signer"
	"github.com/aws/aws-sdk-go/service/simpledb"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/snowball"
	"github.com/aws/aws-sdk-go/service/snowdevicemanagement"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssmcontacts"
	"github.com/aws/aws-sdk-go/service/ssmincidents"
	"github.com/aws/aws-sdk-go/service/sso"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/synthetics"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/aws/aws-sdk-go/service/timestreamquery"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/aws/aws-sdk-go/service/voiceid"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/wafv2"
	"github.com/aws/aws-sdk-go/service/wellarchitected"
	"github.com/aws/aws-sdk-go/service/workdocs"
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/workmail"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/workspacesweb"
	"github.com/aws/aws-sdk-go/service/xray"
	"github.com/hashicorp/terraform-provider-aws/internal/experimental/intf"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

type AWSClient struct {
	AccountID                 string
	DefaultTagsConfig         *tftags.DefaultConfig
	DNSSuffix                 string
	IgnoreTagsConfig          *tftags.IgnoreConfig
	MediaConvertAccountConn   *mediaconvert.MediaConvert
	Partition                 string
	Region                    string
	ReverseDNSPrefix          string
	S3ConnURICleaningDisabled *s3.S3
	ServicePackages           []intf.ServicePackageData
	Session                   *session.Session
	SupportedPlatforms        []string
	TerraformVersion          string

	ssmClient lazyClient[*ssm_sdkv2.Client]

	ACMConn                          *acm.ACM
	ACMPCAConn                       *acmpca.ACMPCA
	AMPConn                          *prometheusservice.PrometheusService
	APIGatewayConn                   *apigateway.APIGateway
	APIGatewayManagementAPIConn      *apigatewaymanagementapi.ApiGatewayManagementApi
	APIGatewayV2Conn                 *apigatewayv2.ApiGatewayV2
	AccessAnalyzerConn               *accessanalyzer.AccessAnalyzer
	AccountConn                      *account.Account
	AlexaForBusinessConn             *alexaforbusiness.AlexaForBusiness
	AmplifyConn                      *amplify.Amplify
	AmplifyBackendConn               *amplifybackend.AmplifyBackend
	AmplifyUIBuilderConn             *amplifyuibuilder.AmplifyUIBuilder
	AppAutoScalingConn               *applicationautoscaling.ApplicationAutoScaling
	AppConfigConn                    *appconfig.AppConfig
	AppConfigDataConn                *appconfigdata.AppConfigData
	AppFlowConn                      *appflow.Appflow
	AppIntegrationsConn              *appintegrationsservice.AppIntegrationsService
	AppMeshConn                      *appmesh.AppMesh
	AppRunnerConn                    *apprunner.AppRunner
	AppStreamConn                    *appstream.AppStream
	AppSyncConn                      *appsync.AppSync
	ApplicationCostProfilerConn      *applicationcostprofiler.ApplicationCostProfiler
	ApplicationInsightsConn          *applicationinsights.ApplicationInsights
	AthenaConn                       *athena.Athena
	AuditManagerConn                 *auditmanager.AuditManager
	AutoScalingConn                  *autoscaling.AutoScaling
	AutoScalingPlansConn             *autoscalingplans.AutoScalingPlans
	BackupConn                       *backup.Backup
	BackupGatewayConn                *backupgateway.BackupGateway
	BatchConn                        *batch.Batch
	BillingConductorConn             *billingconductor.BillingConductor
	BraketConn                       *braket.Braket
	BudgetsConn                      *budgets.Budgets
	CEConn                           *costexplorer.CostExplorer
	CURConn                          *costandusagereportservice.CostandUsageReportService
	ChimeConn                        *chime.Chime
	ChimeSDKIdentityConn             *chimesdkidentity.ChimeSDKIdentity
	ChimeSDKMeetingsConn             *chimesdkmeetings.ChimeSDKMeetings
	ChimeSDKMessagingConn            *chimesdkmessaging.ChimeSDKMessaging
	Cloud9Conn                       *cloud9.Cloud9
	CloudControlConn                 *cloudcontrolapi.CloudControlApi
	CloudDirectoryConn               *clouddirectory.CloudDirectory
	CloudFormationConn               *cloudformation.CloudFormation
	CloudFrontConn                   *cloudfront.CloudFront
	CloudHSMV2Conn                   *cloudhsmv2.CloudHSMV2
	CloudSearchConn                  *cloudsearch.CloudSearch
	CloudSearchDomainConn            *cloudsearchdomain.CloudSearchDomain
	CloudTrailConn                   *cloudtrail.CloudTrail
	CloudWatchConn                   *cloudwatch.CloudWatch
	CodeArtifactConn                 *codeartifact.CodeArtifact
	CodeBuildConn                    *codebuild.CodeBuild
	CodeCommitConn                   *codecommit.CodeCommit
	CodeGuruProfilerConn             *codeguruprofiler.CodeGuruProfiler
	CodeGuruReviewerConn             *codegurureviewer.CodeGuruReviewer
	CodePipelineConn                 *codepipeline.CodePipeline
	CodeStarConn                     *codestar.CodeStar
	CodeStarConnectionsConn          *codestarconnections.CodeStarConnections
	CodeStarNotificationsConn        *codestarnotifications.CodeStarNotifications
	CognitoIDPConn                   *cognitoidentityprovider.CognitoIdentityProvider
	CognitoIdentityConn              *cognitoidentity.CognitoIdentity
	CognitoSyncConn                  *cognitosync.CognitoSync
	ComprehendClient                 *comprehend.Client
	ComprehendMedicalConn            *comprehendmedical.ComprehendMedical
	ComputeOptimizerClient           *computeoptimizer.Client
	ConfigServiceConn                *configservice.ConfigService
	ConnectConn                      *connect.Connect
	ConnectContactLensConn           *connectcontactlens.ConnectContactLens
	ConnectParticipantConn           *connectparticipant.ConnectParticipant
	ControlTowerConn                 *controltower.ControlTower
	CustomerProfilesConn             *customerprofiles.CustomerProfiles
	DAXConn                          *dax.DAX
	DLMConn                          *dlm.DLM
	DMSConn                          *databasemigrationservice.DatabaseMigrationService
	DRSConn                          *drs.Drs
	DSConn                           *directoryservice.DirectoryService
	DataBrewConn                     *gluedatabrew.GlueDataBrew
	DataExchangeConn                 *dataexchange.DataExchange
	DataPipelineConn                 *datapipeline.DataPipeline
	DataSyncConn                     *datasync.DataSync
	DeployConn                       *codedeploy.CodeDeploy
	DetectiveConn                    *detective.Detective
	DevOpsGuruConn                   *devopsguru.DevOpsGuru
	DeviceFarmConn                   *devicefarm.DeviceFarm
	DirectConnectConn                *directconnect.DirectConnect
	DiscoveryConn                    *applicationdiscoveryservice.ApplicationDiscoveryService
	DocDBConn                        *docdb.DocDB
	DynamoDBConn                     *dynamodb.DynamoDB
	DynamoDBStreamsConn              *dynamodbstreams.DynamoDBStreams
	EBSConn                          *ebs.EBS
	EC2Conn                          *ec2.EC2
	EC2InstanceConnectConn           *ec2instanceconnect.EC2InstanceConnect
	ECRConn                          *ecr.ECR
	ECRPublicConn                    *ecrpublic.ECRPublic
	ECSConn                          *ecs.ECS
	EFSConn                          *efs.EFS
	EKSConn                          *eks.EKS
	ELBConn                          *elb.ELB
	ELBV2Conn                        *elbv2.ELBV2
	EMRConn                          *emr.EMR
	EMRContainersConn                *emrcontainers.EMRContainers
	EMRServerlessConn                *emrserverless.EMRServerless
	ElastiCacheConn                  *elasticache.ElastiCache
	ElasticBeanstalkConn             *elasticbeanstalk.ElasticBeanstalk
	ElasticInferenceConn             *elasticinference.ElasticInference
	ElasticTranscoderConn            *elastictranscoder.ElasticTranscoder
	ElasticsearchConn                *elasticsearchservice.ElasticsearchService
	EventsConn                       *eventbridge.EventBridge
	EvidentlyConn                    *cloudwatchevidently.CloudWatchEvidently
	FISClient                        *fis.Client
	FMSConn                          *fms.FMS
	FSxConn                          *fsx.FSx
	FinSpaceConn                     *finspace.Finspace
	FinSpaceDataConn                 *finspacedata.FinSpaceData
	FirehoseConn                     *firehose.Firehose
	ForecastConn                     *forecastservice.ForecastService
	ForecastQueryConn                *forecastqueryservice.ForecastQueryService
	FraudDetectorConn                *frauddetector.FraudDetector
	GameLiftConn                     *gamelift.GameLift
	GlacierConn                      *glacier.Glacier
	GlobalAcceleratorConn            *globalaccelerator.GlobalAccelerator
	GlueConn                         *glue.Glue
	GrafanaConn                      *managedgrafana.ManagedGrafana
	GreengrassConn                   *greengrass.Greengrass
	GreengrassV2Conn                 *greengrassv2.GreengrassV2
	GroundStationConn                *groundstation.GroundStation
	GuardDutyConn                    *guardduty.GuardDuty
	HealthConn                       *health.Health
	HealthLakeConn                   *healthlake.HealthLake
	HoneycodeConn                    *honeycode.Honeycode
	IAMConn                          *iam.IAM
	IVSConn                          *ivs.IVS
	IVSChatClient                    *ivschat.Client
	IdentityStoreClient              *identitystore.Client
	ImageBuilderConn                 *imagebuilder.Imagebuilder
	InspectorConn                    *inspector.Inspector
	Inspector2Client                 *inspector2.Client
	IoTConn                          *iot.IoT
	IoT1ClickDevicesConn             *iot1clickdevicesservice.IoT1ClickDevicesService
	IoT1ClickProjectsConn            *iot1clickprojects.IoT1ClickProjects
	IoTAnalyticsConn                 *iotanalytics.IoTAnalytics
	IoTDataConn                      *iotdataplane.IoTDataPlane
	IoTDeviceAdvisorConn             *iotdeviceadvisor.IoTDeviceAdvisor
	IoTEventsConn                    *iotevents.IoTEvents
	IoTEventsDataConn                *ioteventsdata.IoTEventsData
	IoTFleetHubConn                  *iotfleethub.IoTFleetHub
	IoTJobsDataConn                  *iotjobsdataplane.IoTJobsDataPlane
	IoTSecureTunnelingConn           *iotsecuretunneling.IoTSecureTunneling
	IoTSiteWiseConn                  *iotsitewise.IoTSiteWise
	IoTThingsGraphConn               *iotthingsgraph.IoTThingsGraph
	IoTTwinMakerConn                 *iottwinmaker.IoTTwinMaker
	IoTWirelessConn                  *iotwireless.IoTWireless
	KMSConn                          *kms.KMS
	KafkaConn                        *kafka.Kafka
	KafkaConnectConn                 *kafkaconnect.KafkaConnect
	KendraClient                     *kendra.Client
	KeyspacesConn                    *keyspaces.Keyspaces
	KinesisConn                      *kinesis.Kinesis
	KinesisAnalyticsConn             *kinesisanalytics.KinesisAnalytics
	KinesisAnalyticsV2Conn           *kinesisanalyticsv2.KinesisAnalyticsV2
	KinesisVideoConn                 *kinesisvideo.KinesisVideo
	KinesisVideoArchivedMediaConn    *kinesisvideoarchivedmedia.KinesisVideoArchivedMedia
	KinesisVideoMediaConn            *kinesisvideomedia.KinesisVideoMedia
	KinesisVideoSignalingConn        *kinesisvideosignalingchannels.KinesisVideoSignalingChannels
	LakeFormationConn                *lakeformation.LakeFormation
	LambdaConn                       *lambda.Lambda
	LexModelsConn                    *lexmodelbuildingservice.LexModelBuildingService
	LexModelsV2Conn                  *lexmodelsv2.LexModelsV2
	LexRuntimeConn                   *lexruntimeservice.LexRuntimeService
	LexRuntimeV2Conn                 *lexruntimev2.LexRuntimeV2
	LicenseManagerConn               *licensemanager.LicenseManager
	LightsailConn                    *lightsail.Lightsail
	LocationConn                     *locationservice.LocationService
	LogsConn                         *cloudwatchlogs.CloudWatchLogs
	LookoutEquipmentConn             *lookoutequipment.LookoutEquipment
	LookoutMetricsConn               *lookoutmetrics.LookoutMetrics
	LookoutVisionConn                *lookoutforvision.LookoutForVision
	MQConn                           *mq.MQ
	MTurkConn                        *mturk.MTurk
	MWAAConn                         *mwaa.MWAA
	MachineLearningConn              *machinelearning.MachineLearning
	MacieConn                        *macie.Macie
	Macie2Conn                       *macie2.Macie2
	ManagedBlockchainConn            *managedblockchain.ManagedBlockchain
	MarketplaceCatalogConn           *marketplacecatalog.MarketplaceCatalog
	MarketplaceCommerceAnalyticsConn *marketplacecommerceanalytics.MarketplaceCommerceAnalytics
	MarketplaceEntitlementConn       *marketplaceentitlementservice.MarketplaceEntitlementService
	MarketplaceMeteringConn          *marketplacemetering.MarketplaceMetering
	MediaConnectConn                 *mediaconnect.MediaConnect
	MediaConvertConn                 *mediaconvert.MediaConvert
	MediaLiveClient                  *medialive.Client
	MediaPackageConn                 *mediapackage.MediaPackage
	MediaPackageVODConn              *mediapackagevod.MediaPackageVod
	MediaStoreConn                   *mediastore.MediaStore
	MediaStoreDataConn               *mediastoredata.MediaStoreData
	MediaTailorConn                  *mediatailor.MediaTailor
	MemoryDBConn                     *memorydb.MemoryDB
	MgHConn                          *migrationhub.MigrationHub
	MgnConn                          *mgn.Mgn
	MigrationHubConfigConn           *migrationhubconfig.MigrationHubConfig
	MigrationHubRefactorSpacesConn   *migrationhubrefactorspaces.MigrationHubRefactorSpaces
	MigrationHubStrategyConn         *migrationhubstrategyrecommendations.MigrationHubStrategyRecommendations
	MobileConn                       *mobile.Mobile
	NeptuneConn                      *neptune.Neptune
	NetworkFirewallConn              *networkfirewall.NetworkFirewall
	NetworkManagerConn               *networkmanager.NetworkManager
	NimbleConn                       *nimblestudio.NimbleStudio
	OpenSearchConn                   *opensearchservice.OpenSearchService
	OpsWorksConn                     *opsworks.OpsWorks
	OpsWorksCMConn                   *opsworkscm.OpsWorksCM
	OrganizationsConn                *organizations.Organizations
	OutpostsConn                     *outposts.Outposts
	PIConn                           *pi.PI
	PanoramaConn                     *panorama.Panorama
	PersonalizeConn                  *personalize.Personalize
	PersonalizeEventsConn            *personalizeevents.PersonalizeEvents
	PersonalizeRuntimeConn           *personalizeruntime.PersonalizeRuntime
	PinpointConn                     *pinpoint.Pinpoint
	PinpointEmailConn                *pinpointemail.PinpointEmail
	PinpointSMSVoiceConn             *pinpointsmsvoice.PinpointSMSVoice
	PollyConn                        *polly.Polly
	PricingConn                      *pricing.Pricing
	ProtonConn                       *proton.Proton
	QLDBConn                         *qldb.QLDB
	QLDBSessionConn                  *qldbsession.QLDBSession
	QuickSightConn                   *quicksight.QuickSight
	RAMConn                          *ram.RAM
	RBinConn                         *recyclebin.RecycleBin
	RDSConn                          *rds.RDS
	RDSDataConn                      *rdsdataservice.RDSDataService
	RUMConn                          *cloudwatchrum.CloudWatchRUM
	RedshiftConn                     *redshift.Redshift
	RedshiftDataConn                 *redshiftdataapiservice.RedshiftDataAPIService
	RedshiftServerlessConn           *redshiftserverless.RedshiftServerless
	RekognitionConn                  *rekognition.Rekognition
	ResilienceHubConn                *resiliencehub.ResilienceHub
	ResourceGroupsConn               *resourcegroups.ResourceGroups
	ResourceGroupsTaggingAPIConn     *resourcegroupstaggingapi.ResourceGroupsTaggingAPI
	RoboMakerConn                    *robomaker.RoboMaker
	RolesAnywhereClient              *rolesanywhere.Client
	Route53Conn                      *route53.Route53
	Route53DomainsClient             *route53domains.Client
	Route53RecoveryClusterConn       *route53recoverycluster.Route53RecoveryCluster
	Route53RecoveryControlConfigConn *route53recoverycontrolconfig.Route53RecoveryControlConfig
	Route53RecoveryReadinessConn     *route53recoveryreadiness.Route53RecoveryReadiness
	Route53ResolverConn              *route53resolver.Route53Resolver
	S3Conn                           *s3.S3
	S3ControlConn                    *s3control.S3Control
	S3ControlClient                  *s3control_sdkv2.Client
	S3OutpostsConn                   *s3outposts.S3Outposts
	SESConn                          *ses.SES
	SESV2Client                      *sesv2.Client
	SFNConn                          *sfn.SFN
	SMSConn                          *sms.SMS
	SNSConn                          *sns.SNS
	SQSConn                          *sqs.SQS
	SSMConn                          *ssm.SSM
	SSMContactsConn                  *ssmcontacts.SSMContacts
	SSMIncidentsConn                 *ssmincidents.SSMIncidents
	SSOConn                          *sso.SSO
	SSOAdminConn                     *ssoadmin.SSOAdmin
	SSOOIDCConn                      *ssooidc.SSOOIDC
	STSConn                          *sts.STS
	SWFConn                          *swf.SWF
	SageMakerConn                    *sagemaker.SageMaker
	SageMakerA2IRuntimeConn          *augmentedairuntime.AugmentedAIRuntime
	SageMakerEdgeConn                *sagemakeredgemanager.SagemakerEdgeManager
	SageMakerFeatureStoreRuntimeConn *sagemakerfeaturestoreruntime.SageMakerFeatureStoreRuntime
	SageMakerRuntimeConn             *sagemakerruntime.SageMakerRuntime
	SavingsPlansConn                 *savingsplans.SavingsPlans
	SchedulerClient                  *scheduler.Client
	SchemasConn                      *schemas.Schemas
	SecretsManagerConn               *secretsmanager.SecretsManager
	SecurityHubConn                  *securityhub.SecurityHub
	ServerlessRepoConn               *serverlessapplicationrepository.ServerlessApplicationRepository
	ServiceCatalogConn               *servicecatalog.ServiceCatalog
	ServiceCatalogAppRegistryConn    *appregistry.AppRegistry
	ServiceDiscoveryConn             *servicediscovery.ServiceDiscovery
	ServiceQuotasConn                *servicequotas.ServiceQuotas
	ShieldConn                       *shield.Shield
	SignerConn                       *signer.Signer
	SimpleDBConn                     *simpledb.SimpleDB
	SnowDeviceManagementConn         *snowdevicemanagement.SnowDeviceManagement
	SnowballConn                     *snowball.Snowball
	StorageGatewayConn               *storagegateway.StorageGateway
	SupportConn                      *support.Support
	SyntheticsConn                   *synthetics.Synthetics
	TextractConn                     *textract.Textract
	TimestreamQueryConn              *timestreamquery.TimestreamQuery
	TimestreamWriteConn              *timestreamwrite.TimestreamWrite
	TranscribeClient                 *transcribe.Client
	TranscribeStreamingConn          *transcribestreamingservice.TranscribeStreamingService
	TransferConn                     *transfer.Transfer
	TranslateConn                    *translate.Translate
	VoiceIDConn                      *voiceid.VoiceID
	WAFConn                          *waf.WAF
	WAFRegionalConn                  *wafregional.WAFRegional
	WAFV2Conn                        *wafv2.WAFV2
	WellArchitectedConn              *wellarchitected.WellArchitected
	WisdomConn                       *connectwisdomservice.ConnectWisdomService
	WorkDocsConn                     *workdocs.WorkDocs
	WorkLinkConn                     *worklink.WorkLink
	WorkMailConn                     *workmail.WorkMail
	WorkMailMessageFlowConn          *workmailmessageflow.WorkMailMessageFlow
	WorkSpacesConn                   *workspaces.WorkSpaces
	WorkSpacesWebConn                *workspacesweb.WorkSpacesWeb
	XRayConn                         *xray.XRay
}
