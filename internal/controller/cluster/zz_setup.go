package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/airflow/cluster"
	backendgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/alb/backendgroup"
	httprouter "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/alb/httprouter"
	loadbalancer "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/alb/loadbalancer"
	targetgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/alb/targetgroup"
	virtualhost "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/alb/virtualhost"
	gateway "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/api/gateway"
	trailstrail "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/audit/trailstrail"
	policy "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/backup/policy"
	policybindings "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/backup/policybindings"
	cloudbinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/billing/cloudbinding"
	origingroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cdn/origingroup"
	resource "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cdn/resource"
	desktopsdesktop "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cloud/desktopsdesktop"
	desktopsdesktopgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cloud/desktopsdesktopgroup"
	registry "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cloudregistry/registry"
	registryiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cloudregistry/registryiambinding"
	registryippermission "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cloudregistry/registryippermission"
	certificate "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cm/certificate"
	certificateiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cm/certificateiambinding"
	certificateiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/cm/certificateiammember"
	disk "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/disk"
	diskiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/diskiambinding"
	diskplacementgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/diskplacementgroup"
	diskplacementgroupiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/diskplacementgroupiambinding"
	filesystem "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/filesystem"
	filesystemiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/filesystemiambinding"
	gpucluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/gpucluster"
	gpuclusteriambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/gpuclusteriambinding"
	image "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/image"
	imageiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/imageiambinding"
	instance "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/instance"
	instancegroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/instancegroup"
	instanceiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/instanceiambinding"
	placementgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/placementgroup"
	placementgroupiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/placementgroupiambinding"
	snapshot "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/snapshot"
	snapshotiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/snapshotiambinding"
	snapshotschedule "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/snapshotschedule"
	snapshotscheduleiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/compute/snapshotscheduleiambinding"
	connection "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/connectionmanager/connection"
	registrycontainer "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/container/registry"
	registryiambindingcontainer "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/container/registryiambinding"
	registryippermissioncontainer "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/container/registryippermission"
	repository "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/container/repository"
	repositoryiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/container/repositoryiambinding"
	repositorylifecyclepolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/container/repositorylifecyclepolicy"
	clusterdataproc "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/dataproc/cluster"
	community "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/datasphere/community"
	communityiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/datasphere/communityiambinding"
	project "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/datasphere/project"
	projectiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/datasphere/projectiambinding"
	endpoint "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/datatransfer/endpoint"
	transfer "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/datatransfer/transfer"
	recordset "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/dns/recordset"
	zone "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/dns/zone"
	zoneiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/dns/zoneiambinding"
	iambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/function/iambinding"
	scalingpolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/function/scalingpolicy"
	trigger "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/function/trigger"
	instancegitlab "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/gitlab/instance"
	oauthclient "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/oauthclient"
	oauthclientsecret "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/oauthclientsecret"
	serviceaccount "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/serviceaccount"
	serviceaccountapikey "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/serviceaccountapikey"
	serviceaccountiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/serviceaccountiambinding"
	serviceaccountiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/serviceaccountiammember"
	serviceaccountiampolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/serviceaccountiampolicy"
	serviceaccountkey "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/serviceaccountstaticaccesskey"
	workloadidentityfederatedcredential "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/workloadidentityfederatedcredential"
	workloadidentityoidcfederation "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/workloadidentityoidcfederation"
	workloadidentityoidcfederationiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iam/workloadidentityoidcfederationiambinding"
	corebroker "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iot/corebroker"
	coredevice "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iot/coredevice"
	coreregistry "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/iot/coreregistry"
	asymmetricencryptionkey "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/asymmetricencryptionkey"
	asymmetricencryptionkeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/asymmetricencryptionkeyiambinding"
	asymmetricencryptionkeyiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/asymmetricencryptionkeyiammember"
	asymmetricsignaturekey "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/asymmetricsignaturekey"
	asymmetricsignaturekeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/asymmetricsignaturekeyiambinding"
	asymmetricsignaturekeyiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/asymmetricsignaturekeyiammember"
	secretciphertext "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/secretciphertext"
	symmetrickey "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/symmetrickey"
	symmetrickeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/symmetrickeyiambinding"
	symmetrickeyiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kms/symmetrickeyiammember"
	clusterkubernetes "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kubernetes/cluster"
	clusteriambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kubernetes/clusteriambinding"
	clusteriammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kubernetes/clusteriammember"
	marketplacehelmrelease "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kubernetes/marketplacehelmrelease"
	nodegroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/kubernetes/nodegroup"
	networkloadbalancer "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/lb/networkloadbalancer"
	targetgrouplb "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/lb/targetgroup"
	agent "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/loadtesting/agent"
	secret "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/lockbox/secret"
	secretiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/lockbox/secretiambinding"
	secretiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/lockbox/secretiammember"
	secretversion "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/lockbox/secretversion"
	secretversionhashed "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/lockbox/secretversionhashed"
	group "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/logging/group"
	clickhousecluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/clickhousecluster"
	clickhouseclusterv2 "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/clickhouseclusterv2"
	clickhousedatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/clickhousedatabase"
	clickhouseuser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/clickhouseuser"
	greenplumcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/greenplumcluster"
	greenplumclusterv2 "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/greenplumclusterv2"
	greenplumresourcegroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/greenplumresourcegroup"
	greenplumuser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/greenplumuser"
	kafkacluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/kafkacluster"
	kafkaconnector "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/kafkaconnector"
	kafkatopic "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/kafkatopic"
	kafkauser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/kafkauser"
	mongodbcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/mongodbcluster"
	mongodbdatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/mongodbdatabase"
	mongodbuser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/mongodbuser"
	mysqlcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/mysqlcluster"
	mysqlclusterv2 "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/mysqlclusterv2"
	mysqldatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/mysqldatabase"
	mysqluser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/mysqluser"
	opensearchcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/opensearchcluster"
	postgresqlcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/postgresqlcluster"
	postgresqlclusterv2 "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/postgresqlclusterv2"
	postgresqldatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/postgresqldatabase"
	postgresqluser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/postgresqluser"
	rediscluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/rediscluster"
	redisclusterv2 "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/redisclusterv2"
	redisuser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/redisuser"
	shardedpostgresqlcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/shardedpostgresqlcluster"
	shardedpostgresqldatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/shardedpostgresqldatabase"
	shardedpostgresqlshard "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/shardedpostgresqlshard"
	shardedpostgresqluser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/shardedpostgresqluser"
	sqlservercluster "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/mdb/sqlservercluster"
	queue "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/message/queue"
	clustermetastore "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/metastore/cluster"
	dashboard "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/monitoring/dashboard"
	grouporganizationmanager "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/group"
	groupiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/groupiammember"
	groupmapping "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/groupmapping"
	groupmappingitem "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/groupmappingitem"
	groupmembership "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/groupmembership"
	idpapplicationoauthapplication "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/idpapplicationoauthapplication"
	idpapplicationoauthapplicationassignment "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/idpapplicationoauthapplicationassignment"
	idpapplicationsamlapplication "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/idpapplicationsamlapplication"
	idpapplicationsamlapplicationassignment "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/idpapplicationsamlapplicationassignment"
	idpapplicationsamlsignaturecertificate "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/idpapplicationsamlsignaturecertificate"
	idpuser "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/idpuser"
	idpuserpool "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/idpuserpool"
	idpuserpooldomain "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/idpuserpooldomain"
	mfaenforcement "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/mfaenforcement"
	mfaenforcementaudience "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/mfaenforcementaudience"
	organizationiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/organizationiambinding"
	organizationiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/organizationiammember"
	osloginsettings "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/osloginsettings"
	samlfederation "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/samlfederation"
	samlfederationuseraccount "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/samlfederationuseraccount"
	usersshkey "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/organizationmanager/usersshkey"
	providerconfig "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/providerconfig"
	cloud "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/resourcemanager/cloud"
	cloudiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/resourcemanager/cloudiambinding"
	cloudiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/resourcemanager/cloudiammember"
	folder "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/resourcemanager/folder"
	folderiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/resourcemanager/folderiambinding"
	folderiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/resourcemanager/folderiammember"
	folderiampolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/resourcemanager/folderiampolicy"
	container "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/serverless/container"
	containeriambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/serverless/containeriambinding"
	eventrouterbus "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/serverless/eventrouterbus"
	eventrouterconnector "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/serverless/eventrouterconnector"
	eventrouterrule "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/serverless/eventrouterrule"
	captcha "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/smartcaptcha/captcha"
	clusterspark "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/spark/cluster"
	bucket "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/storage/bucket"
	bucketgrant "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/storage/bucketgrant"
	bucketiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/storage/bucketiambinding"
	bucketpolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/storage/bucketpolicy"
	object "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/storage/object"
	advancedratelimiterprofile "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/sws/advancedratelimiterprofile"
	securityprofile "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/sws/securityprofile"
	wafprofile "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/sws/wafprofile"
	accesscontrol "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/trino/accesscontrol"
	catalog "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/trino/catalog"
	clustertrino "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/trino/cluster"
	address "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/address"
	defaultsecuritygroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/defaultsecuritygroup"
	gatewayvpc "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/gateway"
	network "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/network"
	privateendpoint "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/privateendpoint"
	routetable "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/routetable"
	securitygroup "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/securitygroup"
	securitygrouprule "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/securitygrouprule"
	subnet "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/vpc/subnet"
	function "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/yandex/function"
	databasededicated "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/ydb/databasededicated"
	databaseiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/ydb/databaseiambinding"
	databaseserverless "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/ydb/databaseserverless"
	table "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/ydb/table"
	tablechangefeed "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/ydb/tablechangefeed"
	tableindex "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/ydb/tableindex"
	topic "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/ydb/topic"
	monitoringconnection "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/yq/monitoringconnection"
	objectstoragebinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/yq/objectstoragebinding"
	objectstorageconnection "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/yq/objectstorageconnection"
	ydbconnection "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/yq/ydbconnection"
	ydsbinding "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/yq/ydsbinding"
	ydsconnection "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/yq/ydsconnection"
	clusterytsaurus "github.com/tagesjump/provider-upjet-yc/internal/controller/cluster/ytsaurus/cluster"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		backendgroup.Setup,
		httprouter.Setup,
		loadbalancer.Setup,
		targetgroup.Setup,
		virtualhost.Setup,
		gateway.Setup,
		trailstrail.Setup,
		policy.Setup,
		policybindings.Setup,
		cloudbinding.Setup,
		origingroup.Setup,
		resource.Setup,
		desktopsdesktop.Setup,
		desktopsdesktopgroup.Setup,
		registry.Setup,
		registryiambinding.Setup,
		registryippermission.Setup,
		certificate.Setup,
		certificateiambinding.Setup,
		certificateiammember.Setup,
		disk.Setup,
		diskiambinding.Setup,
		diskplacementgroup.Setup,
		diskplacementgroupiambinding.Setup,
		filesystem.Setup,
		filesystemiambinding.Setup,
		gpucluster.Setup,
		gpuclusteriambinding.Setup,
		image.Setup,
		imageiambinding.Setup,
		instance.Setup,
		instancegroup.Setup,
		instanceiambinding.Setup,
		placementgroup.Setup,
		placementgroupiambinding.Setup,
		snapshot.Setup,
		snapshotiambinding.Setup,
		snapshotschedule.Setup,
		snapshotscheduleiambinding.Setup,
		connection.Setup,
		registrycontainer.Setup,
		registryiambindingcontainer.Setup,
		registryippermissioncontainer.Setup,
		repository.Setup,
		repositoryiambinding.Setup,
		repositorylifecyclepolicy.Setup,
		clusterdataproc.Setup,
		community.Setup,
		communityiambinding.Setup,
		project.Setup,
		projectiambinding.Setup,
		endpoint.Setup,
		transfer.Setup,
		recordset.Setup,
		zone.Setup,
		zoneiambinding.Setup,
		iambinding.Setup,
		scalingpolicy.Setup,
		trigger.Setup,
		instancegitlab.Setup,
		oauthclient.Setup,
		oauthclientsecret.Setup,
		serviceaccount.Setup,
		serviceaccountapikey.Setup,
		serviceaccountiambinding.Setup,
		serviceaccountiammember.Setup,
		serviceaccountiampolicy.Setup,
		serviceaccountkey.Setup,
		serviceaccountstaticaccesskey.Setup,
		workloadidentityfederatedcredential.Setup,
		workloadidentityoidcfederation.Setup,
		workloadidentityoidcfederationiambinding.Setup,
		corebroker.Setup,
		coredevice.Setup,
		coreregistry.Setup,
		asymmetricencryptionkey.Setup,
		asymmetricencryptionkeyiambinding.Setup,
		asymmetricencryptionkeyiammember.Setup,
		asymmetricsignaturekey.Setup,
		asymmetricsignaturekeyiambinding.Setup,
		asymmetricsignaturekeyiammember.Setup,
		secretciphertext.Setup,
		symmetrickey.Setup,
		symmetrickeyiambinding.Setup,
		symmetrickeyiammember.Setup,
		clusterkubernetes.Setup,
		clusteriambinding.Setup,
		clusteriammember.Setup,
		marketplacehelmrelease.Setup,
		nodegroup.Setup,
		networkloadbalancer.Setup,
		targetgrouplb.Setup,
		agent.Setup,
		secret.Setup,
		secretiambinding.Setup,
		secretiammember.Setup,
		secretversion.Setup,
		secretversionhashed.Setup,
		group.Setup,
		clickhousecluster.Setup,
		clickhouseclusterv2.Setup,
		clickhousedatabase.Setup,
		clickhouseuser.Setup,
		greenplumcluster.Setup,
		greenplumclusterv2.Setup,
		greenplumresourcegroup.Setup,
		greenplumuser.Setup,
		kafkacluster.Setup,
		kafkaconnector.Setup,
		kafkatopic.Setup,
		kafkauser.Setup,
		mongodbcluster.Setup,
		mongodbdatabase.Setup,
		mongodbuser.Setup,
		mysqlcluster.Setup,
		mysqlclusterv2.Setup,
		mysqldatabase.Setup,
		mysqluser.Setup,
		opensearchcluster.Setup,
		postgresqlcluster.Setup,
		postgresqlclusterv2.Setup,
		postgresqldatabase.Setup,
		postgresqluser.Setup,
		rediscluster.Setup,
		redisclusterv2.Setup,
		redisuser.Setup,
		shardedpostgresqlcluster.Setup,
		shardedpostgresqldatabase.Setup,
		shardedpostgresqlshard.Setup,
		shardedpostgresqluser.Setup,
		sqlservercluster.Setup,
		queue.Setup,
		clustermetastore.Setup,
		dashboard.Setup,
		grouporganizationmanager.Setup,
		groupiammember.Setup,
		groupmapping.Setup,
		groupmappingitem.Setup,
		groupmembership.Setup,
		idpapplicationoauthapplication.Setup,
		idpapplicationoauthapplicationassignment.Setup,
		idpapplicationsamlapplication.Setup,
		idpapplicationsamlapplicationassignment.Setup,
		idpapplicationsamlsignaturecertificate.Setup,
		idpuser.Setup,
		idpuserpool.Setup,
		idpuserpooldomain.Setup,
		mfaenforcement.Setup,
		mfaenforcementaudience.Setup,
		organizationiambinding.Setup,
		organizationiammember.Setup,
		osloginsettings.Setup,
		samlfederation.Setup,
		samlfederationuseraccount.Setup,
		usersshkey.Setup,
		providerconfig.Setup,
		cloud.Setup,
		cloudiambinding.Setup,
		cloudiammember.Setup,
		folder.Setup,
		folderiambinding.Setup,
		folderiammember.Setup,
		folderiampolicy.Setup,
		container.Setup,
		containeriambinding.Setup,
		eventrouterbus.Setup,
		eventrouterconnector.Setup,
		eventrouterrule.Setup,
		captcha.Setup,
		clusterspark.Setup,
		bucket.Setup,
		bucketgrant.Setup,
		bucketiambinding.Setup,
		bucketpolicy.Setup,
		object.Setup,
		advancedratelimiterprofile.Setup,
		securityprofile.Setup,
		wafprofile.Setup,
		accesscontrol.Setup,
		catalog.Setup,
		clustertrino.Setup,
		address.Setup,
		defaultsecuritygroup.Setup,
		gatewayvpc.Setup,
		network.Setup,
		privateendpoint.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		function.Setup,
		databasededicated.Setup,
		databaseiambinding.Setup,
		databaseserverless.Setup,
		table.Setup,
		tablechangefeed.Setup,
		tableindex.Setup,
		topic.Setup,
		monitoringconnection.Setup,
		objectstoragebinding.Setup,
		objectstorageconnection.Setup,
		ydbconnection.Setup,
		ydsbinding.Setup,
		ydsconnection.Setup,
		clusterytsaurus.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.SetupGated,
		backendgroup.SetupGated,
		httprouter.SetupGated,
		loadbalancer.SetupGated,
		targetgroup.SetupGated,
		virtualhost.SetupGated,
		gateway.SetupGated,
		trailstrail.SetupGated,
		policy.SetupGated,
		policybindings.SetupGated,
		cloudbinding.SetupGated,
		origingroup.SetupGated,
		resource.SetupGated,
		desktopsdesktop.SetupGated,
		desktopsdesktopgroup.SetupGated,
		registry.SetupGated,
		registryiambinding.SetupGated,
		registryippermission.SetupGated,
		certificate.SetupGated,
		certificateiambinding.SetupGated,
		certificateiammember.SetupGated,
		disk.SetupGated,
		diskiambinding.SetupGated,
		diskplacementgroup.SetupGated,
		diskplacementgroupiambinding.SetupGated,
		filesystem.SetupGated,
		filesystemiambinding.SetupGated,
		gpucluster.SetupGated,
		gpuclusteriambinding.SetupGated,
		image.SetupGated,
		imageiambinding.SetupGated,
		instance.SetupGated,
		instancegroup.SetupGated,
		instanceiambinding.SetupGated,
		placementgroup.SetupGated,
		placementgroupiambinding.SetupGated,
		snapshot.SetupGated,
		snapshotiambinding.SetupGated,
		snapshotschedule.SetupGated,
		snapshotscheduleiambinding.SetupGated,
		connection.SetupGated,
		registrycontainer.SetupGated,
		registryiambindingcontainer.SetupGated,
		registryippermissioncontainer.SetupGated,
		repository.SetupGated,
		repositoryiambinding.SetupGated,
		repositorylifecyclepolicy.SetupGated,
		clusterdataproc.SetupGated,
		community.SetupGated,
		communityiambinding.SetupGated,
		project.SetupGated,
		projectiambinding.SetupGated,
		endpoint.SetupGated,
		transfer.SetupGated,
		recordset.SetupGated,
		zone.SetupGated,
		zoneiambinding.SetupGated,
		iambinding.SetupGated,
		scalingpolicy.SetupGated,
		trigger.SetupGated,
		instancegitlab.SetupGated,
		oauthclient.SetupGated,
		oauthclientsecret.SetupGated,
		serviceaccount.SetupGated,
		serviceaccountapikey.SetupGated,
		serviceaccountiambinding.SetupGated,
		serviceaccountiammember.SetupGated,
		serviceaccountiampolicy.SetupGated,
		serviceaccountkey.SetupGated,
		serviceaccountstaticaccesskey.SetupGated,
		workloadidentityfederatedcredential.SetupGated,
		workloadidentityoidcfederation.SetupGated,
		workloadidentityoidcfederationiambinding.SetupGated,
		corebroker.SetupGated,
		coredevice.SetupGated,
		coreregistry.SetupGated,
		asymmetricencryptionkey.SetupGated,
		asymmetricencryptionkeyiambinding.SetupGated,
		asymmetricencryptionkeyiammember.SetupGated,
		asymmetricsignaturekey.SetupGated,
		asymmetricsignaturekeyiambinding.SetupGated,
		asymmetricsignaturekeyiammember.SetupGated,
		secretciphertext.SetupGated,
		symmetrickey.SetupGated,
		symmetrickeyiambinding.SetupGated,
		symmetrickeyiammember.SetupGated,
		clusterkubernetes.SetupGated,
		clusteriambinding.SetupGated,
		clusteriammember.SetupGated,
		marketplacehelmrelease.SetupGated,
		nodegroup.SetupGated,
		networkloadbalancer.SetupGated,
		targetgrouplb.SetupGated,
		agent.SetupGated,
		secret.SetupGated,
		secretiambinding.SetupGated,
		secretiammember.SetupGated,
		secretversion.SetupGated,
		secretversionhashed.SetupGated,
		group.SetupGated,
		clickhousecluster.SetupGated,
		clickhouseclusterv2.SetupGated,
		clickhousedatabase.SetupGated,
		clickhouseuser.SetupGated,
		greenplumcluster.SetupGated,
		greenplumclusterv2.SetupGated,
		greenplumresourcegroup.SetupGated,
		greenplumuser.SetupGated,
		kafkacluster.SetupGated,
		kafkaconnector.SetupGated,
		kafkatopic.SetupGated,
		kafkauser.SetupGated,
		mongodbcluster.SetupGated,
		mongodbdatabase.SetupGated,
		mongodbuser.SetupGated,
		mysqlcluster.SetupGated,
		mysqlclusterv2.SetupGated,
		mysqldatabase.SetupGated,
		mysqluser.SetupGated,
		opensearchcluster.SetupGated,
		postgresqlcluster.SetupGated,
		postgresqlclusterv2.SetupGated,
		postgresqldatabase.SetupGated,
		postgresqluser.SetupGated,
		rediscluster.SetupGated,
		redisclusterv2.SetupGated,
		redisuser.SetupGated,
		shardedpostgresqlcluster.SetupGated,
		shardedpostgresqldatabase.SetupGated,
		shardedpostgresqlshard.SetupGated,
		shardedpostgresqluser.SetupGated,
		sqlservercluster.SetupGated,
		queue.SetupGated,
		clustermetastore.SetupGated,
		dashboard.SetupGated,
		grouporganizationmanager.SetupGated,
		groupiammember.SetupGated,
		groupmapping.SetupGated,
		groupmappingitem.SetupGated,
		groupmembership.SetupGated,
		idpapplicationoauthapplication.SetupGated,
		idpapplicationoauthapplicationassignment.SetupGated,
		idpapplicationsamlapplication.SetupGated,
		idpapplicationsamlapplicationassignment.SetupGated,
		idpapplicationsamlsignaturecertificate.SetupGated,
		idpuser.SetupGated,
		idpuserpool.SetupGated,
		idpuserpooldomain.SetupGated,
		mfaenforcement.SetupGated,
		mfaenforcementaudience.SetupGated,
		organizationiambinding.SetupGated,
		organizationiammember.SetupGated,
		osloginsettings.SetupGated,
		samlfederation.SetupGated,
		samlfederationuseraccount.SetupGated,
		usersshkey.SetupGated,
		providerconfig.SetupGated,
		cloud.SetupGated,
		cloudiambinding.SetupGated,
		cloudiammember.SetupGated,
		folder.SetupGated,
		folderiambinding.SetupGated,
		folderiammember.SetupGated,
		folderiampolicy.SetupGated,
		container.SetupGated,
		containeriambinding.SetupGated,
		eventrouterbus.SetupGated,
		eventrouterconnector.SetupGated,
		eventrouterrule.SetupGated,
		captcha.SetupGated,
		clusterspark.SetupGated,
		bucket.SetupGated,
		bucketgrant.SetupGated,
		bucketiambinding.SetupGated,
		bucketpolicy.SetupGated,
		object.SetupGated,
		advancedratelimiterprofile.SetupGated,
		securityprofile.SetupGated,
		wafprofile.SetupGated,
		accesscontrol.SetupGated,
		catalog.SetupGated,
		clustertrino.SetupGated,
		address.SetupGated,
		defaultsecuritygroup.SetupGated,
		gatewayvpc.SetupGated,
		network.SetupGated,
		privateendpoint.SetupGated,
		routetable.SetupGated,
		securitygroup.SetupGated,
		securitygrouprule.SetupGated,
		subnet.SetupGated,
		function.SetupGated,
		databasededicated.SetupGated,
		databaseiambinding.SetupGated,
		databaseserverless.SetupGated,
		table.SetupGated,
		tablechangefeed.SetupGated,
		tableindex.SetupGated,
		topic.SetupGated,
		monitoringconnection.SetupGated,
		objectstoragebinding.SetupGated,
		objectstorageconnection.SetupGated,
		ydbconnection.SetupGated,
		ydsbinding.SetupGated,
		ydsconnection.SetupGated,
		clusterytsaurus.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
