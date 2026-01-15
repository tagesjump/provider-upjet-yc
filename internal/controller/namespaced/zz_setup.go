package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backendgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/alb/backendgroup"
	httprouter "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/alb/httprouter"
	loadbalancer "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/alb/loadbalancer"
	targetgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/alb/targetgroup"
	virtualhost "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/alb/virtualhost"
	gateway "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/api/gateway"
	trailstrail "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/audit/trailstrail"
	policy "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/backup/policy"
	cloudbinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/billing/cloudbinding"
	origingroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/cdn/origingroup"
	resource "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/cdn/resource"
	certificate "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/cm/certificate"
	disk "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/disk"
	diskiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/diskiambinding"
	diskplacementgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/diskplacementgroup"
	diskplacementgroupiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/diskplacementgroupiambinding"
	filesystem "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/filesystem"
	filesystemiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/filesystemiambinding"
	gpucluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/gpucluster"
	gpuclusteriambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/gpuclusteriambinding"
	image "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/image"
	imageiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/imageiambinding"
	instance "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/instance"
	instancegroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/instancegroup"
	instanceiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/instanceiambinding"
	placementgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/placementgroup"
	placementgroupiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/placementgroupiambinding"
	snapshot "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/snapshot"
	snapshotiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/snapshotiambinding"
	snapshotschedule "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/snapshotschedule"
	snapshotscheduleiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/compute/snapshotscheduleiambinding"
	registry "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/container/registry"
	registryiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/container/registryiambinding"
	registryippermission "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/container/registryippermission"
	repository "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/container/repository"
	repositoryiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/container/repositoryiambinding"
	repositorylifecyclepolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/container/repositorylifecyclepolicy"
	cluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/dataproc/cluster"
	endpoint "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/datatransfer/endpoint"
	transfer "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/datatransfer/transfer"
	recordset "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/dns/recordset"
	zone "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/dns/zone"
	zoneiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/dns/zoneiambinding"
	iambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/function/iambinding"
	scalingpolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/function/scalingpolicy"
	trigger "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/function/trigger"
	serviceaccount "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iam/serviceaccount"
	serviceaccountapikey "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iam/serviceaccountapikey"
	serviceaccountiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iam/serviceaccountiambinding"
	serviceaccountiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iam/serviceaccountiammember"
	serviceaccountiampolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iam/serviceaccountiampolicy"
	serviceaccountkey "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iam/serviceaccountstaticaccesskey"
	corebroker "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iot/corebroker"
	coredevice "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iot/coredevice"
	coreregistry "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/iot/coreregistry"
	asymmetricencryptionkey "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kms/asymmetricencryptionkey"
	asymmetricencryptionkeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kms/asymmetricencryptionkeyiambinding"
	asymmetricsignaturekey "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kms/asymmetricsignaturekey"
	asymmetricsignaturekeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kms/asymmetricsignaturekeyiambinding"
	secretciphertext "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kms/secretciphertext"
	symmetrickey "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kms/symmetrickey"
	symmetrickeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kms/symmetrickeyiambinding"
	clusterkubernetes "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kubernetes/cluster"
	nodegroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/kubernetes/nodegroup"
	networkloadbalancer "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/lb/networkloadbalancer"
	targetgrouplb "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/lb/targetgroup"
	agent "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/loadtesting/agent"
	secret "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/lockbox/secret"
	secretiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/lockbox/secretiambinding"
	secretversion "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/lockbox/secretversion"
	group "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/logging/group"
	clickhousecluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/clickhousecluster"
	elasticsearchcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/elasticsearchcluster"
	greenplumcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/greenplumcluster"
	kafkacluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/kafkacluster"
	kafkaconnector "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/kafkaconnector"
	kafkatopic "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/kafkatopic"
	kafkauser "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/kafkauser"
	mongodbcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/mongodbcluster"
	mongodbdatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/mongodbdatabase"
	mongodbuser "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/mongodbuser"
	mysqlcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/mysqlcluster"
	mysqldatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/mysqldatabase"
	mysqluser "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/mysqluser"
	opensearchcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/opensearchcluster"
	postgresqlcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/postgresqlcluster"
	postgresqldatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/postgresqldatabase"
	postgresqluser "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/postgresqluser"
	rediscluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/rediscluster"
	sqlservercluster "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/mdb/sqlservercluster"
	queue "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/message/queue"
	dashboard "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/monitoring/dashboard"
	grouporganizationmanager "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/group"
	groupiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/groupiammember"
	groupmembership "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/groupmembership"
	organizationiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/organizationiambinding"
	organizationiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/organizationiammember"
	osloginsettings "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/osloginsettings"
	samlfederation "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/samlfederation"
	samlfederationuseraccount "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/samlfederationuseraccount"
	usersshkey "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/organizationmanager/usersshkey"
	providerconfig "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/providerconfig"
	cloud "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/resourcemanager/cloud"
	cloudiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/resourcemanager/cloudiambinding"
	cloudiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/resourcemanager/cloudiammember"
	folder "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/resourcemanager/folder"
	folderiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/resourcemanager/folderiambinding"
	folderiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/resourcemanager/folderiammember"
	folderiampolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/resourcemanager/folderiampolicy"
	container "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/serverless/container"
	containeriambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/serverless/containeriambinding"
	captcha "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/smartcaptcha/captcha"
	bucket "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/storage/bucket"
	object "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/storage/object"
	securityprofile "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/sws/securityprofile"
	address "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/vpc/address"
	defaultsecuritygroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/vpc/defaultsecuritygroup"
	gatewayvpc "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/vpc/gateway"
	network "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/vpc/network"
	routetable "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/vpc/routetable"
	securitygroup "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/vpc/securitygroup"
	securitygrouprule "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/vpc/securitygrouprule"
	subnet "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/vpc/subnet"
	function "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/yandex/function"
	databasededicated "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/ydb/databasededicated"
	databaseiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/ydb/databaseiambinding"
	databaseserverless "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/ydb/databaseserverless"
	table "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/ydb/table"
	tablechangefeed "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/ydb/tablechangefeed"
	tableindex "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/ydb/tableindex"
	topic "github.com/tagesjump/provider-upjet-yc/internal/controller/namespaced/ydb/topic"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backendgroup.Setup,
		httprouter.Setup,
		loadbalancer.Setup,
		targetgroup.Setup,
		virtualhost.Setup,
		gateway.Setup,
		trailstrail.Setup,
		policy.Setup,
		cloudbinding.Setup,
		origingroup.Setup,
		resource.Setup,
		certificate.Setup,
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
		registry.Setup,
		registryiambinding.Setup,
		registryippermission.Setup,
		repository.Setup,
		repositoryiambinding.Setup,
		repositorylifecyclepolicy.Setup,
		cluster.Setup,
		endpoint.Setup,
		transfer.Setup,
		recordset.Setup,
		zone.Setup,
		zoneiambinding.Setup,
		iambinding.Setup,
		scalingpolicy.Setup,
		trigger.Setup,
		serviceaccount.Setup,
		serviceaccountapikey.Setup,
		serviceaccountiambinding.Setup,
		serviceaccountiammember.Setup,
		serviceaccountiampolicy.Setup,
		serviceaccountkey.Setup,
		serviceaccountstaticaccesskey.Setup,
		corebroker.Setup,
		coredevice.Setup,
		coreregistry.Setup,
		asymmetricencryptionkey.Setup,
		asymmetricencryptionkeyiambinding.Setup,
		asymmetricsignaturekey.Setup,
		asymmetricsignaturekeyiambinding.Setup,
		secretciphertext.Setup,
		symmetrickey.Setup,
		symmetrickeyiambinding.Setup,
		clusterkubernetes.Setup,
		nodegroup.Setup,
		networkloadbalancer.Setup,
		targetgrouplb.Setup,
		agent.Setup,
		secret.Setup,
		secretiambinding.Setup,
		secretversion.Setup,
		group.Setup,
		clickhousecluster.Setup,
		elasticsearchcluster.Setup,
		greenplumcluster.Setup,
		kafkacluster.Setup,
		kafkaconnector.Setup,
		kafkatopic.Setup,
		kafkauser.Setup,
		mongodbcluster.Setup,
		mongodbdatabase.Setup,
		mongodbuser.Setup,
		mysqlcluster.Setup,
		mysqldatabase.Setup,
		mysqluser.Setup,
		opensearchcluster.Setup,
		postgresqlcluster.Setup,
		postgresqldatabase.Setup,
		postgresqluser.Setup,
		rediscluster.Setup,
		sqlservercluster.Setup,
		queue.Setup,
		dashboard.Setup,
		grouporganizationmanager.Setup,
		groupiammember.Setup,
		groupmembership.Setup,
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
		captcha.Setup,
		bucket.Setup,
		object.Setup,
		securityprofile.Setup,
		address.Setup,
		defaultsecuritygroup.Setup,
		gatewayvpc.Setup,
		network.Setup,
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
		backendgroup.SetupGated,
		httprouter.SetupGated,
		loadbalancer.SetupGated,
		targetgroup.SetupGated,
		virtualhost.SetupGated,
		gateway.SetupGated,
		trailstrail.SetupGated,
		policy.SetupGated,
		cloudbinding.SetupGated,
		origingroup.SetupGated,
		resource.SetupGated,
		certificate.SetupGated,
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
		registry.SetupGated,
		registryiambinding.SetupGated,
		registryippermission.SetupGated,
		repository.SetupGated,
		repositoryiambinding.SetupGated,
		repositorylifecyclepolicy.SetupGated,
		cluster.SetupGated,
		endpoint.SetupGated,
		transfer.SetupGated,
		recordset.SetupGated,
		zone.SetupGated,
		zoneiambinding.SetupGated,
		iambinding.SetupGated,
		scalingpolicy.SetupGated,
		trigger.SetupGated,
		serviceaccount.SetupGated,
		serviceaccountapikey.SetupGated,
		serviceaccountiambinding.SetupGated,
		serviceaccountiammember.SetupGated,
		serviceaccountiampolicy.SetupGated,
		serviceaccountkey.SetupGated,
		serviceaccountstaticaccesskey.SetupGated,
		corebroker.SetupGated,
		coredevice.SetupGated,
		coreregistry.SetupGated,
		asymmetricencryptionkey.SetupGated,
		asymmetricencryptionkeyiambinding.SetupGated,
		asymmetricsignaturekey.SetupGated,
		asymmetricsignaturekeyiambinding.SetupGated,
		secretciphertext.SetupGated,
		symmetrickey.SetupGated,
		symmetrickeyiambinding.SetupGated,
		clusterkubernetes.SetupGated,
		nodegroup.SetupGated,
		networkloadbalancer.SetupGated,
		targetgrouplb.SetupGated,
		agent.SetupGated,
		secret.SetupGated,
		secretiambinding.SetupGated,
		secretversion.SetupGated,
		group.SetupGated,
		clickhousecluster.SetupGated,
		elasticsearchcluster.SetupGated,
		greenplumcluster.SetupGated,
		kafkacluster.SetupGated,
		kafkaconnector.SetupGated,
		kafkatopic.SetupGated,
		kafkauser.SetupGated,
		mongodbcluster.SetupGated,
		mongodbdatabase.SetupGated,
		mongodbuser.SetupGated,
		mysqlcluster.SetupGated,
		mysqldatabase.SetupGated,
		mysqluser.SetupGated,
		opensearchcluster.SetupGated,
		postgresqlcluster.SetupGated,
		postgresqldatabase.SetupGated,
		postgresqluser.SetupGated,
		rediscluster.SetupGated,
		sqlservercluster.SetupGated,
		queue.SetupGated,
		dashboard.SetupGated,
		grouporganizationmanager.SetupGated,
		groupiammember.SetupGated,
		groupmembership.SetupGated,
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
		captcha.SetupGated,
		bucket.SetupGated,
		object.SetupGated,
		securityprofile.SetupGated,
		address.SetupGated,
		defaultsecuritygroup.SetupGated,
		gatewayvpc.SetupGated,
		network.SetupGated,
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
