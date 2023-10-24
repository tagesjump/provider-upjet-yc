/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	backendgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/alb/backendgroup"
	httprouter "github.com/tagesjump/provider-upjet-yc/internal/controller/alb/httprouter"
	loadbalancer "github.com/tagesjump/provider-upjet-yc/internal/controller/alb/loadbalancer"
	targetgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/alb/targetgroup"
	virtualhost "github.com/tagesjump/provider-upjet-yc/internal/controller/alb/virtualhost"
	disk "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/disk"
	diskplacementgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/diskplacementgroup"
	filesystem "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/filesystem"
	gpucluster "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/gpucluster"
	image "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/image"
	instance "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/instance"
	instancegroup "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/instancegroup"
	placementgroup "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/placementgroup"
	snapshot "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/snapshot"
	snapshotschedule "github.com/tagesjump/provider-upjet-yc/internal/controller/compute/snapshotschedule"
	registry "github.com/tagesjump/provider-upjet-yc/internal/controller/container/registry"
	registryiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/container/registryiambinding"
	registryippermission "github.com/tagesjump/provider-upjet-yc/internal/controller/container/registryippermission"
	repository "github.com/tagesjump/provider-upjet-yc/internal/controller/container/repository"
	repositoryiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/container/repositoryiambinding"
	repositorylifecyclepolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/container/repositorylifecyclepolicy"
	recordset "github.com/tagesjump/provider-upjet-yc/internal/controller/dns/recordset"
	zone "github.com/tagesjump/provider-upjet-yc/internal/controller/dns/zone"
	serviceaccount "github.com/tagesjump/provider-upjet-yc/internal/controller/iam/serviceaccount"
	serviceaccountapikey "github.com/tagesjump/provider-upjet-yc/internal/controller/iam/serviceaccountapikey"
	serviceaccountiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/iam/serviceaccountiambinding"
	serviceaccountiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/iam/serviceaccountiammember"
	serviceaccountiampolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/iam/serviceaccountiampolicy"
	serviceaccountkey "github.com/tagesjump/provider-upjet-yc/internal/controller/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "github.com/tagesjump/provider-upjet-yc/internal/controller/iam/serviceaccountstaticaccesskey"
	asymmetricencryptionkey "github.com/tagesjump/provider-upjet-yc/internal/controller/kms/asymmetricencryptionkey"
	asymmetricencryptionkeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/kms/asymmetricencryptionkeyiambinding"
	asymmetricsignaturekey "github.com/tagesjump/provider-upjet-yc/internal/controller/kms/asymmetricsignaturekey"
	asymmetricsignaturekeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/kms/asymmetricsignaturekeyiambinding"
	secretciphertext "github.com/tagesjump/provider-upjet-yc/internal/controller/kms/secretciphertext"
	symmetrickey "github.com/tagesjump/provider-upjet-yc/internal/controller/kms/symmetrickey"
	symmetrickeyiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/kms/symmetrickeyiambinding"
	cluster "github.com/tagesjump/provider-upjet-yc/internal/controller/kubernetes/cluster"
	nodegroup "github.com/tagesjump/provider-upjet-yc/internal/controller/kubernetes/nodegroup"
	networkloadbalancer "github.com/tagesjump/provider-upjet-yc/internal/controller/lb/networkloadbalancer"
	targetgrouplb "github.com/tagesjump/provider-upjet-yc/internal/controller/lb/targetgroup"
	clickhousecluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/clickhousecluster"
	elasticsearchcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/elasticsearchcluster"
	greenplumcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/greenplumcluster"
	kafkacluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/kafkacluster"
	kafkaconnector "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/kafkaconnector"
	kafkatopic "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/kafkatopic"
	kafkauser "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/kafkauser"
	mongodbcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/mongodbcluster"
	mysqlcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/mysqlcluster"
	mysqldatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/mysqldatabase"
	mysqluser "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/mysqluser"
	postgresqlcluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/postgresqlcluster"
	postgresqldatabase "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/postgresqldatabase"
	postgresqluser "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/postgresqluser"
	rediscluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/rediscluster"
	sqlservercluster "github.com/tagesjump/provider-upjet-yc/internal/controller/mdb/sqlservercluster"
	group "github.com/tagesjump/provider-upjet-yc/internal/controller/organizationmanager/group"
	groupiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/organizationmanager/groupiammember"
	groupmembership "github.com/tagesjump/provider-upjet-yc/internal/controller/organizationmanager/groupmembership"
	organizationiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/organizationmanager/organizationiambinding"
	organizationiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/organizationmanager/organizationiammember"
	samlfederation "github.com/tagesjump/provider-upjet-yc/internal/controller/organizationmanager/samlfederation"
	samlfederationuseraccount "github.com/tagesjump/provider-upjet-yc/internal/controller/organizationmanager/samlfederationuseraccount"
	providerconfig "github.com/tagesjump/provider-upjet-yc/internal/controller/providerconfig"
	cloud "github.com/tagesjump/provider-upjet-yc/internal/controller/resourcemanager/cloud"
	cloudiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/resourcemanager/cloudiambinding"
	cloudiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/resourcemanager/cloudiammember"
	folder "github.com/tagesjump/provider-upjet-yc/internal/controller/resourcemanager/folder"
	folderiambinding "github.com/tagesjump/provider-upjet-yc/internal/controller/resourcemanager/folderiambinding"
	folderiammember "github.com/tagesjump/provider-upjet-yc/internal/controller/resourcemanager/folderiammember"
	folderiampolicy "github.com/tagesjump/provider-upjet-yc/internal/controller/resourcemanager/folderiampolicy"
	bucket "github.com/tagesjump/provider-upjet-yc/internal/controller/storage/bucket"
	object "github.com/tagesjump/provider-upjet-yc/internal/controller/storage/object"
	address "github.com/tagesjump/provider-upjet-yc/internal/controller/vpc/address"
	defaultsecuritygroup "github.com/tagesjump/provider-upjet-yc/internal/controller/vpc/defaultsecuritygroup"
	gateway "github.com/tagesjump/provider-upjet-yc/internal/controller/vpc/gateway"
	network "github.com/tagesjump/provider-upjet-yc/internal/controller/vpc/network"
	routetable "github.com/tagesjump/provider-upjet-yc/internal/controller/vpc/routetable"
	securitygroup "github.com/tagesjump/provider-upjet-yc/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/tagesjump/provider-upjet-yc/internal/controller/vpc/securitygrouprule"
	subnet "github.com/tagesjump/provider-upjet-yc/internal/controller/vpc/subnet"
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
		disk.Setup,
		diskplacementgroup.Setup,
		filesystem.Setup,
		gpucluster.Setup,
		image.Setup,
		instance.Setup,
		instancegroup.Setup,
		placementgroup.Setup,
		snapshot.Setup,
		snapshotschedule.Setup,
		registry.Setup,
		registryiambinding.Setup,
		registryippermission.Setup,
		repository.Setup,
		repositoryiambinding.Setup,
		repositorylifecyclepolicy.Setup,
		recordset.Setup,
		zone.Setup,
		serviceaccount.Setup,
		serviceaccountapikey.Setup,
		serviceaccountiambinding.Setup,
		serviceaccountiammember.Setup,
		serviceaccountiampolicy.Setup,
		serviceaccountkey.Setup,
		serviceaccountstaticaccesskey.Setup,
		asymmetricencryptionkey.Setup,
		asymmetricencryptionkeyiambinding.Setup,
		asymmetricsignaturekey.Setup,
		asymmetricsignaturekeyiambinding.Setup,
		secretciphertext.Setup,
		symmetrickey.Setup,
		symmetrickeyiambinding.Setup,
		cluster.Setup,
		nodegroup.Setup,
		networkloadbalancer.Setup,
		targetgrouplb.Setup,
		clickhousecluster.Setup,
		elasticsearchcluster.Setup,
		greenplumcluster.Setup,
		kafkacluster.Setup,
		kafkaconnector.Setup,
		kafkatopic.Setup,
		kafkauser.Setup,
		mongodbcluster.Setup,
		mysqlcluster.Setup,
		mysqldatabase.Setup,
		mysqluser.Setup,
		postgresqlcluster.Setup,
		postgresqldatabase.Setup,
		postgresqluser.Setup,
		rediscluster.Setup,
		sqlservercluster.Setup,
		group.Setup,
		groupiammember.Setup,
		groupmembership.Setup,
		organizationiambinding.Setup,
		organizationiammember.Setup,
		samlfederation.Setup,
		samlfederationuseraccount.Setup,
		providerconfig.Setup,
		cloud.Setup,
		cloudiambinding.Setup,
		cloudiammember.Setup,
		folder.Setup,
		folderiambinding.Setup,
		folderiammember.Setup,
		folderiampolicy.Setup,
		bucket.Setup,
		object.Setup,
		address.Setup,
		defaultsecuritygroup.Setup,
		gateway.Setup,
		network.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
