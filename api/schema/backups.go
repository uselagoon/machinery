package schema

// RestoreStatusType .
type RestoreStatusType string

// Guest .
const (
	RestorePending    RestoreStatusType = "PENDING"
	RestoreSuccessful RestoreStatusType = "SUCCESSFUL"
	RestoreFailed     RestoreStatusType = "FAILED"
)

// Backup is the Lagoon API Backup object.
type Backup struct {
	ID       int     `json:"id"`
	Source   string  `json:"source"`
	BackupID string  `json:"backupId"`
	Created  string  `json:"created"`
	Deleted  string  `json:"deleted"`
	Restore  Restore `json:"restore"`
}

// Restore is the Lagoon API Restore object.
type Restore struct {
	ID              int    `json:"id"`
	Status          string `json:"status"`
	Created         string `json:"created"`
	BackupID        string `json:"backupId"`
	RestoreLocation string `json:"restoreLocation"`
}

// AddRestoreInput is based on the input to
// addRestore.
type AddRestoreInput struct {
	ID              uint   `json:"id,omitempty"`
	Status          string `json:"status"`
	Created         string `json:"created"`
	BackupID        string `json:"backupId"`
	RestoreLocation string `json:"restoreLocation"`
	Execute         bool   `json:"execute"`
}

// DeleteBackup is the response.
type DeleteBackup struct {
	DeleteBackup string `json:"deleteBackup"`
}

// AddBackupInput is based on the input to
// addBackup.
type AddBackupInput struct {
	ID          uint   `json:"id,omitempty"`
	Environment uint   `json:"environment"`
	Source      string `json:"source"`
	BackupID    string `json:"backupId"`
	Created     string `json:"created"`
}

type UpdateRestoreInput struct {
	BackupID string                  `json:"backupId"`
	Patch    UpdateRestorePatchInput `json:"patch: UpdateRestorePatchInput"`
}

type UpdateRestorePatchInput struct {
	Status          RestoreStatusType `json:"status,omitempty"`
	Created         string            `json:"created,omitempty"`
	RestoreLocation string            `json:"restoreLocation,omitempty"`
}
