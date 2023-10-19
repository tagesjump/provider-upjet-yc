/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ServiceAccount
func (mg *ServiceAccount) GetTerraformResourceType() string {
	return "yandex_iam_service_account"
}

// GetConnectionDetailsMapping for this ServiceAccount
func (tr *ServiceAccount) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServiceAccount
func (tr *ServiceAccount) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccount
func (tr *ServiceAccount) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccount
func (tr *ServiceAccount) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccount
func (tr *ServiceAccount) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccount
func (tr *ServiceAccount) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ServiceAccount
func (tr *ServiceAccount) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ServiceAccount using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccount) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccount) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountAPIKey
func (mg *ServiceAccountAPIKey) GetTerraformResourceType() string {
	return "yandex_iam_service_account_api_key"
}

// GetConnectionDetailsMapping for this ServiceAccountAPIKey
func (tr *ServiceAccountAPIKey) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"secret_key": "status.atProvider.secretKey"}
}

// GetObservation of this ServiceAccountAPIKey
func (tr *ServiceAccountAPIKey) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountAPIKey
func (tr *ServiceAccountAPIKey) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountAPIKey
func (tr *ServiceAccountAPIKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountAPIKey
func (tr *ServiceAccountAPIKey) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountAPIKey
func (tr *ServiceAccountAPIKey) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ServiceAccountAPIKey
func (tr *ServiceAccountAPIKey) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ServiceAccountAPIKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountAPIKey) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountAPIKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountAPIKey) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountIAMBinding
func (mg *ServiceAccountIAMBinding) GetTerraformResourceType() string {
	return "yandex_iam_service_account_iam_binding"
}

// GetConnectionDetailsMapping for this ServiceAccountIAMBinding
func (tr *ServiceAccountIAMBinding) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServiceAccountIAMBinding
func (tr *ServiceAccountIAMBinding) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountIAMBinding
func (tr *ServiceAccountIAMBinding) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountIAMBinding
func (tr *ServiceAccountIAMBinding) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountIAMBinding
func (tr *ServiceAccountIAMBinding) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountIAMBinding
func (tr *ServiceAccountIAMBinding) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ServiceAccountIAMBinding
func (tr *ServiceAccountIAMBinding) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ServiceAccountIAMBinding using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountIAMBinding) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountIAMBindingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountIAMBinding) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountIAMMember
func (mg *ServiceAccountIAMMember) GetTerraformResourceType() string {
	return "yandex_iam_service_account_iam_member"
}

// GetConnectionDetailsMapping for this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ServiceAccountIAMMember using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountIAMMember) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountIAMMemberParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountIAMMember) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountIAMPolicy
func (mg *ServiceAccountIAMPolicy) GetTerraformResourceType() string {
	return "yandex_iam_service_account_iam_policy"
}

// GetConnectionDetailsMapping for this ServiceAccountIAMPolicy
func (tr *ServiceAccountIAMPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServiceAccountIAMPolicy
func (tr *ServiceAccountIAMPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountIAMPolicy
func (tr *ServiceAccountIAMPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountIAMPolicy
func (tr *ServiceAccountIAMPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountIAMPolicy
func (tr *ServiceAccountIAMPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountIAMPolicy
func (tr *ServiceAccountIAMPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ServiceAccountIAMPolicy
func (tr *ServiceAccountIAMPolicy) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ServiceAccountIAMPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountIAMPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountIAMPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountIAMPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountKey
func (mg *ServiceAccountKey) GetTerraformResourceType() string {
	return "yandex_iam_service_account_key"
}

// GetConnectionDetailsMapping for this ServiceAccountKey
func (tr *ServiceAccountKey) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"private_key": "status.atProvider.privateKey"}
}

// GetObservation of this ServiceAccountKey
func (tr *ServiceAccountKey) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountKey
func (tr *ServiceAccountKey) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountKey
func (tr *ServiceAccountKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountKey
func (tr *ServiceAccountKey) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountKey
func (tr *ServiceAccountKey) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ServiceAccountKey
func (tr *ServiceAccountKey) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ServiceAccountKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountKey) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountKey) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountStaticAccessKey
func (mg *ServiceAccountStaticAccessKey) GetTerraformResourceType() string {
	return "yandex_iam_service_account_static_access_key"
}

// GetConnectionDetailsMapping for this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"secret_key": "status.atProvider.secretKey"}
}

// GetObservation of this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ServiceAccountStaticAccessKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountStaticAccessKey) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountStaticAccessKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountStaticAccessKey) GetTerraformSchemaVersion() int {
	return 0
}
