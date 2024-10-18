// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ExecuteByTimeInitParameters struct {

	// — If true, schedule will be applied on the last day of month.
	// See day_type for available values.
	IncludeLastDayOfMonth *bool `json:"includeLastDayOfMonth,omitempty" tf:"include_last_day_of_month,omitempty"`

	// — List of days when schedule applies. Used in "MONTHLY" type.
	Monthdays []*float64 `json:"monthdays,omitempty" tf:"monthdays,omitempty"`

	// — seconds
	Months []*float64 `json:"months,omitempty" tf:"months,omitempty"`

	// hours format), when the schedule applies.
	RepeatAt []*string `json:"repeatAt,omitempty" tf:"repeat_at,omitempty"`

	// — Frequency of backup repetition. See interval_type for available values.
	RepeatEvery *string `json:"repeatEvery,omitempty" tf:"repeat_every,omitempty"`

	// — Type of the scheduling. Available values are: "HOURLY", "DAILY", "WEEKLY", "MONTHLY".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// — List of weekdays when the backup will be applied. Used in "WEEKLY" type.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type ExecuteByTimeObservation struct {

	// — If true, schedule will be applied on the last day of month.
	// See day_type for available values.
	IncludeLastDayOfMonth *bool `json:"includeLastDayOfMonth,omitempty" tf:"include_last_day_of_month,omitempty"`

	// — List of days when schedule applies. Used in "MONTHLY" type.
	Monthdays []*float64 `json:"monthdays,omitempty" tf:"monthdays,omitempty"`

	// — seconds
	Months []*float64 `json:"months,omitempty" tf:"months,omitempty"`

	// hours format), when the schedule applies.
	RepeatAt []*string `json:"repeatAt,omitempty" tf:"repeat_at,omitempty"`

	// — Frequency of backup repetition. See interval_type for available values.
	RepeatEvery *string `json:"repeatEvery,omitempty" tf:"repeat_every,omitempty"`

	// — Type of the scheduling. Available values are: "HOURLY", "DAILY", "WEEKLY", "MONTHLY".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// — List of weekdays when the backup will be applied. Used in "WEEKLY" type.
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type ExecuteByTimeParameters struct {

	// — If true, schedule will be applied on the last day of month.
	// See day_type for available values.
	// +kubebuilder:validation:Optional
	IncludeLastDayOfMonth *bool `json:"includeLastDayOfMonth,omitempty" tf:"include_last_day_of_month,omitempty"`

	// — List of days when schedule applies. Used in "MONTHLY" type.
	// +kubebuilder:validation:Optional
	Monthdays []*float64 `json:"monthdays,omitempty" tf:"monthdays,omitempty"`

	// — seconds
	// +kubebuilder:validation:Optional
	Months []*float64 `json:"months,omitempty" tf:"months,omitempty"`

	// hours format), when the schedule applies.
	// +kubebuilder:validation:Optional
	RepeatAt []*string `json:"repeatAt,omitempty" tf:"repeat_at,omitempty"`

	// — Frequency of backup repetition. See interval_type for available values.
	// +kubebuilder:validation:Optional
	RepeatEvery *string `json:"repeatEvery,omitempty" tf:"repeat_every,omitempty"`

	// — Type of the scheduling. Available values are: "HOURLY", "DAILY", "WEEKLY", "MONTHLY".
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// — List of weekdays when the backup will be applied. Used in "WEEKLY" type.
	// +kubebuilder:validation:Optional
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type PolicyInitParameters struct {

	// [Plan ID]-[Unique ID]a) — The name of generated archives.
	ArchiveName *string `json:"archiveName,omitempty" tf:"archive_name,omitempty"`

	// — Configuration of Changed Block Tracking.
	// Available values are: "USE_IF_ENABLED", "ENABLED_AND_USE", "DO_NOT_USE".
	Cbt *string `json:"cbt,omitempty" tf:"cbt,omitempty"`

	// — Archive compression level. Affects CPU.
	// Available values: "NORMAL", "HIGH", "MAX", "OFF".
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// — Enable flag
	FastBackupEnabled *bool `json:"fastBackupEnabled,omitempty" tf:"fast_backup_enabled,omitempty"`

	// — days
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// — Format of the backup. It's strongly recommend to leave this option empty or "AUTO".
	// Available values: "AUTO", "VERSION_11", "VERSION_12".
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// — If true, snapshots of multiple volumes will be taken simultaneously.
	MultiVolumeSnapshottingEnabled *bool `json:"multiVolumeSnapshottingEnabled,omitempty" tf:"multi_volume_snapshotting_enabled,omitempty"`

	// — Name of the policy
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// — Time windows for performance limitations of backup.
	PerformanceWindowEnabled *bool `json:"performanceWindowEnabled,omitempty" tf:"performance_window_enabled,omitempty"`

	// — Preserves file security settings. It's better to set this option to true.
	PreserveFileSecuritySettings *bool `json:"preserveFileSecuritySettings,omitempty" tf:"preserve_file_security_settings,omitempty"`

	// — If true, a quiesced snapshot of the virtual machine will be taken.
	QuiesceSnapshottingEnabled *bool `json:"quiesceSnapshottingEnabled,omitempty" tf:"quiesce_snapshotting_enabled,omitempty"`

	// — Amount of reattempts that should be performed while trying to make backup at the host.
	// This attribute consists of the following parameters:
	Reattempts []ReattemptsInitParameters `json:"reattempts,omitempty" tf:"reattempts,omitempty"`

	// — Retention policy for backups. Allows to setup backups lifecycle.
	// This attribute consists of the following parameters:
	Retention []RetentionInitParameters `json:"retention,omitempty" tf:"retention,omitempty"`

	// — Schedule settings for creating backups on the host.
	Scheduling []SchedulingInitParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// — if true, a user interaction will be avoided when possible.
	SilentModeEnabled *bool `json:"silentModeEnabled,omitempty" tf:"silent_mode_enabled,omitempty"`

	// — determines the size to split backups. It's better to leave this option unchanged.
	SplittingBytes *string `json:"splittingBytes,omitempty" tf:"splitting_bytes,omitempty"`

	// (Requied) — Amount of reattempts that should be performed while trying to make snapshot.
	// This attribute consists of the following parameters:
	VMSnapshotReattempts []VMSnapshotReattemptsInitParameters `json:"vmSnapshotReattempts,omitempty" tf:"vm_snapshot_reattempts,omitempty"`

	// — Settings for the volume shadow copy service.
	// Available values are: "NATIVE", "TARGET_SYSTEM_DEFINED"
	VssProvider *string `json:"vssProvider,omitempty" tf:"vss_provider,omitempty"`
}

type PolicyObservation struct {

	// [Plan ID]-[Unique ID]a) — The name of generated archives.
	ArchiveName *string `json:"archiveName,omitempty" tf:"archive_name,omitempty"`

	// — Configuration of Changed Block Tracking.
	// Available values are: "USE_IF_ENABLED", "ENABLED_AND_USE", "DO_NOT_USE".
	Cbt *string `json:"cbt,omitempty" tf:"cbt,omitempty"`

	// — Archive compression level. Affects CPU.
	// Available values: "NORMAL", "HIGH", "MAX", "OFF".
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// — Enable flag
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Enable flag
	FastBackupEnabled *bool `json:"fastBackupEnabled,omitempty" tf:"fast_backup_enabled,omitempty"`

	// — days
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// — Format of the backup. It's strongly recommend to leave this option empty or "AUTO".
	// Available values: "AUTO", "VERSION_11", "VERSION_12".
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// — days
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// — If true, snapshots of multiple volumes will be taken simultaneously.
	MultiVolumeSnapshottingEnabled *bool `json:"multiVolumeSnapshottingEnabled,omitempty" tf:"multi_volume_snapshotting_enabled,omitempty"`

	// — Name of the policy
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// — Time windows for performance limitations of backup.
	PerformanceWindowEnabled *bool `json:"performanceWindowEnabled,omitempty" tf:"performance_window_enabled,omitempty"`

	// — Preserves file security settings. It's better to set this option to true.
	PreserveFileSecuritySettings *bool `json:"preserveFileSecuritySettings,omitempty" tf:"preserve_file_security_settings,omitempty"`

	// — If true, a quiesced snapshot of the virtual machine will be taken.
	QuiesceSnapshottingEnabled *bool `json:"quiesceSnapshottingEnabled,omitempty" tf:"quiesce_snapshotting_enabled,omitempty"`

	// — Amount of reattempts that should be performed while trying to make backup at the host.
	// This attribute consists of the following parameters:
	Reattempts []ReattemptsObservation `json:"reattempts,omitempty" tf:"reattempts,omitempty"`

	// — Retention policy for backups. Allows to setup backups lifecycle.
	// This attribute consists of the following parameters:
	Retention []RetentionObservation `json:"retention,omitempty" tf:"retention,omitempty"`

	// — Schedule settings for creating backups on the host.
	Scheduling []SchedulingObservation `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// — if true, a user interaction will be avoided when possible.
	SilentModeEnabled *bool `json:"silentModeEnabled,omitempty" tf:"silent_mode_enabled,omitempty"`

	// — determines the size to split backups. It's better to leave this option unchanged.
	SplittingBytes *string `json:"splittingBytes,omitempty" tf:"splitting_bytes,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// (Requied) — Amount of reattempts that should be performed while trying to make snapshot.
	// This attribute consists of the following parameters:
	VMSnapshotReattempts []VMSnapshotReattemptsObservation `json:"vmSnapshotReattempts,omitempty" tf:"vm_snapshot_reattempts,omitempty"`

	// — Settings for the volume shadow copy service.
	// Available values are: "NATIVE", "TARGET_SYSTEM_DEFINED"
	VssProvider *string `json:"vssProvider,omitempty" tf:"vss_provider,omitempty"`
}

type PolicyParameters struct {

	// [Plan ID]-[Unique ID]a) — The name of generated archives.
	// +kubebuilder:validation:Optional
	ArchiveName *string `json:"archiveName,omitempty" tf:"archive_name,omitempty"`

	// — Configuration of Changed Block Tracking.
	// Available values are: "USE_IF_ENABLED", "ENABLED_AND_USE", "DO_NOT_USE".
	// +kubebuilder:validation:Optional
	Cbt *string `json:"cbt,omitempty" tf:"cbt,omitempty"`

	// — Archive compression level. Affects CPU.
	// Available values: "NORMAL", "HIGH", "MAX", "OFF".
	// +kubebuilder:validation:Optional
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// — Enable flag
	// +kubebuilder:validation:Optional
	FastBackupEnabled *bool `json:"fastBackupEnabled,omitempty" tf:"fast_backup_enabled,omitempty"`

	// — days
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// — Format of the backup. It's strongly recommend to leave this option empty or "AUTO".
	// Available values: "AUTO", "VERSION_11", "VERSION_12".
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// — If true, snapshots of multiple volumes will be taken simultaneously.
	// +kubebuilder:validation:Optional
	MultiVolumeSnapshottingEnabled *bool `json:"multiVolumeSnapshottingEnabled,omitempty" tf:"multi_volume_snapshotting_enabled,omitempty"`

	// — Name of the policy
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// — Time windows for performance limitations of backup.
	// +kubebuilder:validation:Optional
	PerformanceWindowEnabled *bool `json:"performanceWindowEnabled,omitempty" tf:"performance_window_enabled,omitempty"`

	// — Preserves file security settings. It's better to set this option to true.
	// +kubebuilder:validation:Optional
	PreserveFileSecuritySettings *bool `json:"preserveFileSecuritySettings,omitempty" tf:"preserve_file_security_settings,omitempty"`

	// — If true, a quiesced snapshot of the virtual machine will be taken.
	// +kubebuilder:validation:Optional
	QuiesceSnapshottingEnabled *bool `json:"quiesceSnapshottingEnabled,omitempty" tf:"quiesce_snapshotting_enabled,omitempty"`

	// — Amount of reattempts that should be performed while trying to make backup at the host.
	// This attribute consists of the following parameters:
	// +kubebuilder:validation:Optional
	Reattempts []ReattemptsParameters `json:"reattempts,omitempty" tf:"reattempts,omitempty"`

	// — Retention policy for backups. Allows to setup backups lifecycle.
	// This attribute consists of the following parameters:
	// +kubebuilder:validation:Optional
	Retention []RetentionParameters `json:"retention,omitempty" tf:"retention,omitempty"`

	// — Schedule settings for creating backups on the host.
	// +kubebuilder:validation:Optional
	Scheduling []SchedulingParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// — if true, a user interaction will be avoided when possible.
	// +kubebuilder:validation:Optional
	SilentModeEnabled *bool `json:"silentModeEnabled,omitempty" tf:"silent_mode_enabled,omitempty"`

	// — determines the size to split backups. It's better to leave this option unchanged.
	// +kubebuilder:validation:Optional
	SplittingBytes *string `json:"splittingBytes,omitempty" tf:"splitting_bytes,omitempty"`

	// (Requied) — Amount of reattempts that should be performed while trying to make snapshot.
	// This attribute consists of the following parameters:
	// +kubebuilder:validation:Optional
	VMSnapshotReattempts []VMSnapshotReattemptsParameters `json:"vmSnapshotReattempts,omitempty" tf:"vm_snapshot_reattempts,omitempty"`

	// — Settings for the volume shadow copy service.
	// Available values are: "NATIVE", "TARGET_SYSTEM_DEFINED"
	// +kubebuilder:validation:Optional
	VssProvider *string `json:"vssProvider,omitempty" tf:"vss_provider,omitempty"`
}

type ReattemptsInitParameters struct {

	// — Enable flag
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Retry interval. See interval_type for available values
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// — Maximum number of attempts before throwing an error
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`
}

type ReattemptsObservation struct {

	// — Enable flag
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Retry interval. See interval_type for available values
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// — Maximum number of attempts before throwing an error
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`
}

type ReattemptsParameters struct {

	// — Enable flag
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Retry interval. See interval_type for available values
	// +kubebuilder:validation:Optional
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// — Maximum number of attempts before throwing an error
	// +kubebuilder:validation:Optional
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`
}

type RetentionInitParameters struct {

	// — Defines whether retention rule applies after creating backup or before.
	AfterBackup *bool `json:"afterBackup,omitempty" tf:"after_backup,omitempty"`

	// — seconds
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type RetentionObservation struct {

	// — Defines whether retention rule applies after creating backup or before.
	AfterBackup *bool `json:"afterBackup,omitempty" tf:"after_backup,omitempty"`

	// — seconds
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`
}

type RetentionParameters struct {

	// — Defines whether retention rule applies after creating backup or before.
	// +kubebuilder:validation:Optional
	AfterBackup *bool `json:"afterBackup,omitempty" tf:"after_backup,omitempty"`

	// — seconds
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

type RulesInitParameters struct {

	// (Conflicts with max_count) — Deletes backups that older than max_age. Exactly one of max_count or max_age should be set.
	MaxAge *string `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (Conflicts with max_age) — Deletes backups if it's count exceeds max_count. Exactly one of max_count or max_age should be set.
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// — days
	RepeatPeriod []*string `json:"repeatPeriod,omitempty" tf:"repeat_period,omitempty"`
}

type RulesObservation struct {

	// (Conflicts with max_count) — Deletes backups that older than max_age. Exactly one of max_count or max_age should be set.
	MaxAge *string `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (Conflicts with max_age) — Deletes backups if it's count exceeds max_count. Exactly one of max_count or max_age should be set.
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// — days
	RepeatPeriod []*string `json:"repeatPeriod,omitempty" tf:"repeat_period,omitempty"`
}

type RulesParameters struct {

	// (Conflicts with max_count) — Deletes backups that older than max_age. Exactly one of max_count or max_age should be set.
	// +kubebuilder:validation:Optional
	MaxAge *string `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// (Conflicts with max_age) — Deletes backups if it's count exceeds max_count. Exactly one of max_count or max_age should be set.
	// +kubebuilder:validation:Optional
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// — days
	// +kubebuilder:validation:Optional
	RepeatPeriod []*string `json:"repeatPeriod,omitempty" tf:"repeat_period,omitempty"`
}

type SchedulingInitParameters struct {

	// — Enable flag
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Perform backup by interval, since last backup of the host. Maximum value is: 9999 days.
	// See interval_type for available values. Exactly on of options should be set: execute_by_interval or execute_by_time.
	ExecuteByInterval *float64 `json:"executeByInterval,omitempty" tf:"execute_by_interval,omitempty"`

	// — Perform backup periodically at specific time. Exactly on of options should be set: execute_by_interval or execute_by_time.
	ExecuteByTime []ExecuteByTimeInitParameters `json:"executeByTime,omitempty" tf:"execute_by_time,omitempty"`

	// — Maximum number of backup processes allowed to run in parallel. 0 for unlimited.
	MaxParallelBackups *float64 `json:"maxParallelBackups,omitempty" tf:"max_parallel_backups,omitempty"`

	// — Configuration of the random delay between the execution of parallel tasks.
	// See interval_type for available values.
	RandomMaxDelay *string `json:"randomMaxDelay,omitempty" tf:"random_max_delay,omitempty"`

	// — Scheme of the backups.
	// Available values are: "ALWAYS_INCREMENTAL", "ALWAYS_FULL", "WEEKLY_FULL_DAILY_INCREMENTAL", 'WEEKLY_INCREMENTAL".
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// — A day of week to start weekly backups.
	// See day_type for available values.
	WeeklyBackupDay *string `json:"weeklyBackupDay,omitempty" tf:"weekly_backup_day,omitempty"`
}

type SchedulingObservation struct {

	// — Enable flag
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Perform backup by interval, since last backup of the host. Maximum value is: 9999 days.
	// See interval_type for available values. Exactly on of options should be set: execute_by_interval or execute_by_time.
	ExecuteByInterval *float64 `json:"executeByInterval,omitempty" tf:"execute_by_interval,omitempty"`

	// — Perform backup periodically at specific time. Exactly on of options should be set: execute_by_interval or execute_by_time.
	ExecuteByTime []ExecuteByTimeObservation `json:"executeByTime,omitempty" tf:"execute_by_time,omitempty"`

	// — Maximum number of backup processes allowed to run in parallel. 0 for unlimited.
	MaxParallelBackups *float64 `json:"maxParallelBackups,omitempty" tf:"max_parallel_backups,omitempty"`

	// — Configuration of the random delay between the execution of parallel tasks.
	// See interval_type for available values.
	RandomMaxDelay *string `json:"randomMaxDelay,omitempty" tf:"random_max_delay,omitempty"`

	// — Scheme of the backups.
	// Available values are: "ALWAYS_INCREMENTAL", "ALWAYS_FULL", "WEEKLY_FULL_DAILY_INCREMENTAL", 'WEEKLY_INCREMENTAL".
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// — A day of week to start weekly backups.
	// See day_type for available values.
	WeeklyBackupDay *string `json:"weeklyBackupDay,omitempty" tf:"weekly_backup_day,omitempty"`
}

type SchedulingParameters struct {

	// — Enable flag
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Perform backup by interval, since last backup of the host. Maximum value is: 9999 days.
	// See interval_type for available values. Exactly on of options should be set: execute_by_interval or execute_by_time.
	// +kubebuilder:validation:Optional
	ExecuteByInterval *float64 `json:"executeByInterval,omitempty" tf:"execute_by_interval,omitempty"`

	// — Perform backup periodically at specific time. Exactly on of options should be set: execute_by_interval or execute_by_time.
	// +kubebuilder:validation:Optional
	ExecuteByTime []ExecuteByTimeParameters `json:"executeByTime,omitempty" tf:"execute_by_time,omitempty"`

	// — Maximum number of backup processes allowed to run in parallel. 0 for unlimited.
	// +kubebuilder:validation:Optional
	MaxParallelBackups *float64 `json:"maxParallelBackups,omitempty" tf:"max_parallel_backups,omitempty"`

	// — Configuration of the random delay between the execution of parallel tasks.
	// See interval_type for available values.
	// +kubebuilder:validation:Optional
	RandomMaxDelay *string `json:"randomMaxDelay,omitempty" tf:"random_max_delay,omitempty"`

	// — Scheme of the backups.
	// Available values are: "ALWAYS_INCREMENTAL", "ALWAYS_FULL", "WEEKLY_FULL_DAILY_INCREMENTAL", 'WEEKLY_INCREMENTAL".
	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// — A day of week to start weekly backups.
	// See day_type for available values.
	// +kubebuilder:validation:Optional
	WeeklyBackupDay *string `json:"weeklyBackupDay,omitempty" tf:"weekly_backup_day,omitempty"`
}

type VMSnapshotReattemptsInitParameters struct {

	// — Enable flag
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Retry interval. See interval_type for available values
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// — Maximum number of attempts before throwing an error
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`
}

type VMSnapshotReattemptsObservation struct {

	// — Enable flag
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Retry interval. See interval_type for available values
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// — Maximum number of attempts before throwing an error
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`
}

type VMSnapshotReattemptsParameters struct {

	// — Enable flag
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// — Retry interval. See interval_type for available values
	// +kubebuilder:validation:Optional
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// — Maximum number of attempts before throwing an error
	// +kubebuilder:validation:Optional
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PolicyInitParameters `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Policy is the Schema for the Policys API. Allows management of Yandex.Cloud Backup Policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.reattempts) || (has(self.initProvider) && has(self.initProvider.reattempts))",message="spec.forProvider.reattempts is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.retention) || (has(self.initProvider) && has(self.initProvider.retention))",message="spec.forProvider.retention is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scheduling) || (has(self.initProvider) && has(self.initProvider.scheduling))",message="spec.forProvider.scheduling is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vmSnapshotReattempts) || (has(self.initProvider) && has(self.initProvider.vmSnapshotReattempts))",message="spec.forProvider.vmSnapshotReattempts is a required parameter"
	Spec   PolicySpec   `json:"spec"`
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}