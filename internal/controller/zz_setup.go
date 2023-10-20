/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

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
