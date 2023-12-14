// Code generated for package lgraphql by go-bindata DO NOT EDIT. (@generated)
// sources:
// _lgraphql/addEnvVariable.graphql
// _lgraphql/addNotificationEmail.graphql
// _lgraphql/addNotificationMicrosoftTeams.graphql
// _lgraphql/addNotificationRocketChat.graphql
// _lgraphql/addNotificationSlack.graphql
// _lgraphql/addNotificationToProject.graphql
// _lgraphql/lagoonSchema.graphql
// _lgraphql/lagoonVersion.graphql
// _lgraphql/sshEndpointsByProject.graphql
// _lgraphql/taskByID.graphql
// _lgraphql/deployments/deployEnvironmentBranch.graphql
// _lgraphql/deployments/deployEnvironmentLatest.graphql
// _lgraphql/deployments/deployEnvironmentPromote.graphql
// _lgraphql/deployments/deployEnvironmentPullrequest.graphql
// _lgraphql/deployments/deploymentByName.graphql
// _lgraphql/deployments/deploymentByNameWithLog.graphql
// _lgraphql/deployments/deploymentByRemoteID.graphql
// _lgraphql/deployments/getDeploymentsByBulkID.graphql
// _lgraphql/deployments/getDeploymentsForEnvironment.graphql
// _lgraphql/deployments/updateDeployment.graphql
// _lgraphql/deploytargets/addDeployTarget.graphql
// _lgraphql/deploytargets/addDeployTargetToOrganization.graphql
// _lgraphql/deploytargets/deleteDeployTarget.graphql
// _lgraphql/deploytargets/deployTargetsByOrganizationId.graphql
// _lgraphql/deploytargets/deployTargetsByOrganizationName.graphql
// _lgraphql/deploytargets/listDeployTargets.graphql
// _lgraphql/deploytargets/removeDeployTargetFromOrganization.graphql
// _lgraphql/deploytargets/updateDeployTarget.graphql
// _lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql
// _lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql
// _lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql
// _lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql
// _lgraphql/projects/addProject.graphql
// _lgraphql/projects/minimalProjectByName.graphql
// _lgraphql/projects/projectByName.graphql
// _lgraphql/projects/projectByNameExtended.graphql
// _lgraphql/projects/projectByNameMetadata.graphql
// _lgraphql/projects/projectGroups.graphql
// _lgraphql/projects/projectsByMetadata.graphql
// _lgraphql/projects/projectsByOrganizationId.graphql
// _lgraphql/projects/removeProjectFromOrganization.graphql
// _lgraphql/projects/removeProjectMetadataByKey.graphql
// _lgraphql/projects/updateProject.graphql
// _lgraphql/projects/updateProjectMetadata.graphql
// _lgraphql/environments/addOrUpdateEnvironment.graphql
// _lgraphql/environments/addOrUpdateEnvironmentStorage.graphql
// _lgraphql/environments/addRestore.graphql
// _lgraphql/environments/backupsForEnvironmentByName.graphql
// _lgraphql/environments/deleteEnvironment.graphql
// _lgraphql/environments/environmentByName.graphql
// _lgraphql/environments/environmentByNamespace.graphql
// _lgraphql/environments/environmentsByProjectName.graphql
// _lgraphql/environments/setEnvironmentServices.graphql
// _lgraphql/environments/sshEndpointByEnvironment.graphql
// _lgraphql/environments/updateEnvironment.graphql
// _lgraphql/tasks/switchActiveStandby.graphql
// _lgraphql/tasks/updateTask.graphql
// _lgraphql/tasks/uploadFilesForTask.graphql
// _lgraphql/usergroups/addGroup.graphql
// _lgraphql/usergroups/addGroupToOrganization.graphql
// _lgraphql/usergroups/addGroupsToProject.graphql
// _lgraphql/usergroups/addSshKey.graphql
// _lgraphql/usergroups/addUser.graphql
// _lgraphql/usergroups/addUserToGroup.graphql
// _lgraphql/usergroups/addUserToOrganization.graphql
// _lgraphql/usergroups/allGroupMembers.graphql
// _lgraphql/usergroups/allGroupMembersWithKeys.graphql
// _lgraphql/usergroups/allUsers.graphql
// _lgraphql/usergroups/getGroupMembers.graphql
// _lgraphql/usergroups/groupsByOrganizationId.graphql
// _lgraphql/usergroups/me.graphql
// _lgraphql/usergroups/removeGroupsFromProject.graphql
// _lgraphql/usergroups/removeSSHKeyById.graphql
// _lgraphql/usergroups/removeUserFromGroup.graphql
// _lgraphql/usergroups/removeUserFromOrganization.graphql
// _lgraphql/usergroups/userByEmail.graphql
// _lgraphql/usergroups/userByEmailSSHKeys.graphql
// _lgraphql/usergroups/userBySSHFingerprint.graphql
// _lgraphql/usergroups/userBySSHKey.graphql
// _lgraphql/usergroups/userCanSSHToEnvironment.graphql
// _lgraphql/usergroups/usersByOrganization.graphql
// _lgraphql/usergroups/usersByOrganizationName.graphql
// _lgraphql/organizations/addOrganization.graphql
// _lgraphql/organizations/allOrganizations.graphql
// _lgraphql/organizations/deleteOrganization.graphql
// _lgraphql/organizations/organizationById.graphql
// _lgraphql/organizations/organizationByName.graphql
// _lgraphql/organizations/updateOrganization.graphql
// _lgraphql/notifications/addNotificationEmail.graphql
// _lgraphql/notifications/addNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/addNotificationRocketChat.graphql
// _lgraphql/notifications/addNotificationSlack.graphql
// _lgraphql/notifications/addNotificationToProject.graphql
// _lgraphql/notifications/addNotificationWebhook.graphql
// _lgraphql/notifications/deleteNotificationEmail.graphql
// _lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/deleteNotificationRocketChat.graphql
// _lgraphql/notifications/deleteNotificationSlack.graphql
// _lgraphql/notifications/deleteNotificationWebhook.graphql
// _lgraphql/notifications/listAllNotificationEmail.graphql
// _lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/listAllNotificationRocketChat.graphql
// _lgraphql/notifications/listAllNotificationSlack.graphql
// _lgraphql/notifications/listAllNotificationWebhook.graphql
// _lgraphql/notifications/projectNotificationEmail.graphql
// _lgraphql/notifications/projectNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/projectNotificationRocketChat.graphql
// _lgraphql/notifications/projectNotificationSlack.graphql
// _lgraphql/notifications/projectNotificationWebhook.graphql
// _lgraphql/notifications/removeNotificationFromProject.graphql
// _lgraphql/notifications/updateNotificationEmail.graphql
// _lgraphql/notifications/updateNotificationMicrosoftTeams.graphql
// _lgraphql/notifications/updateNotificationRocketChat.graphql
// _lgraphql/notifications/updateNotificationSlack.graphql
// _lgraphql/notifications/updateNotificationWebhook.graphql
package lgraphql

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __lgraphqlAddenvvariableGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xb1\xca\xc3\x20\x14\x46\xf7\x3c\xc5\xf7\xc3\x3f\x24\x90\x27\x70\xef\x90\x39\xa5\xfb\x6d\x95\x22\x24\x37\x21\xbd\x06\x42\xe9\xbb\x17\xbd\xda\x4a\x07\x97\x73\xd4\xef\xcc\x41\x48\xfc\xc2\x68\x1b\xe0\x5f\x8e\xd5\x19\x9c\x78\xbf\xd0\xe6\xe9\x3a\xb9\xf3\xb1\xba\xbf\xbe\xa8\xc1\x1a\x0c\x2c\x0a\x1e\xb7\xe5\xe7\xf2\x18\x89\x4a\xa6\xd9\x19\x8c\xb2\x79\xbe\x2b\xd9\x69\x0a\x5f\xd4\xe1\xd9\x00\x00\x59\x5b\x7d\xd0\x7a\x5e\x83\x98\xec\x00\xcd\x49\xd3\x7d\x85\x62\x46\xee\x29\x38\xc7\x68\x54\x81\x1a\x91\x5a\x0a\xca\x15\x5a\x93\xd8\xab\xfb\xcc\x79\x5b\x3d\x54\xd9\xc4\xf3\x0e\x00\x00\xff\xff\x1d\x4f\x13\xd7\x24\x01\x00\x00")

func _lgraphqlAddenvvariableGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddenvvariableGraphql,
		"_lgraphql/addEnvVariable.graphql",
	)
}

func _lgraphqlAddenvvariableGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddenvvariableGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addEnvVariable.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\xa4\xe6\x26\x66\xe6\x38\xa6\xa4\x14\xa5\x16\x17\xc3\x65\x34\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\x66\xb8\x82\xd4\x6a\x64\xe6\x15\x94\x96\x58\x41\x55\x28\x28\x40\x8c\x04\x9b\xac\x03\x15\x42\x35\x13\xc5\x0a\xb0\x8a\x5a\x4d\xb8\xee\xcc\x14\x24\x63\x20\x92\x5c\x20\x0c\x08\x00\x00\xff\xff\x7d\x3b\x64\xad\xb7\x00\x00\x00")

func _lgraphqlAddnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationemailGraphql,
		"_lgraphql/addNotificationEmail.graphql",
	)
}

func _lgraphqlAddnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\x94\xa7\x26\x65\xe4\xe7\x67\xc3\x05\x35\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\xfc\xf2\x4b\x32\xd3\x32\x93\xc1\xda\x7d\x33\x93\x8b\xf2\x8b\xf3\xd3\x4a\x42\x52\x13\x73\x8b\x35\x32\xf3\x0a\x4a\x4b\xac\xa0\x4a\x15\x14\x20\xc6\x82\x4d\xd7\x81\x0a\xc1\xcd\x85\xd9\x00\x16\xaf\xd5\x84\xeb\xc9\x4c\x41\xd2\x0c\x91\xe4\x02\x61\x40\x00\x00\x00\xff\xff\x97\x56\x46\x93\xb1\x00\x00\x00")

func _lgraphqlAddnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationmicrosoftteamsGraphql,
		"_lgraphql/addNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlAddnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x31\x0e\xc2\x30\x0c\x45\xf7\x9c\xe2\x23\x31\xb4\x52\x4f\xd0\x95\x9d\x01\x4e\x60\x92\x40\xac\x52\x1b\x21\x57\x0c\x88\xbb\xa3\xa6\x49\x04\x83\x97\xf7\xad\xf7\xe6\xc5\xc8\x58\x05\x9d\x03\xf6\x42\x73\x1c\x71\xb6\x27\xcb\x6d\x37\xac\xc4\x27\x12\x89\xf7\x7f\xf8\x8a\x97\xa4\x3a\x35\xd8\xe3\xed\x00\x80\x42\x38\xaa\xf1\x95\x7d\x76\x9e\xd4\x4f\xd1\x0e\x89\xac\x63\x79\x2c\x36\x96\x37\x60\xeb\xe4\xdc\x50\x50\x0b\xd5\x64\x1d\x5a\xac\x66\x33\xff\xf4\x4d\xc6\xe1\xc7\xba\x8d\x6e\xbd\x6f\x00\x00\x00\xff\xff\x53\x0f\xa2\x30\xdb\x00\x00\x00")

func _lgraphqlAddnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationrocketchatGraphql,
		"_lgraphql/addNotificationRocketChat.graphql",
	)
}

func _lgraphqlAddnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x41\x0e\x82\x40\x0c\x45\xf7\x73\x8a\x6f\xe2\x02\x12\x4e\xc0\x21\xdc\x70\x82\x3a\x33\x4a\x03\xb4\xc6\x94\xb8\x30\xde\xdd\x30\x30\x8d\x2e\xba\x79\xbf\x79\x6f\x59\x8d\x8c\x55\xd0\x04\xe0\x2c\xb4\xe4\x1e\x83\x3d\x59\xee\xa7\x6e\x23\x71\x24\x91\x3c\xff\xc3\x57\xbe\x8e\xaa\x93\xc3\x16\xef\x00\x00\x94\xd2\x45\x8d\x6f\x1c\x8b\x73\x98\x29\x4e\x0d\xcb\x63\xb5\xfe\xf8\x00\xf6\x44\x29\x75\x07\xf2\x46\xad\xd5\xc1\x3b\xb5\x58\xf8\xa7\x75\x19\xa7\x1f\xeb\x3e\x86\xed\xbe\x01\x00\x00\xff\xff\xd9\x6e\xb4\xda\xd6\x00\x00\x00")

func _lgraphqlAddnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationslackGraphql,
		"_lgraphql/addNotificationSlack.graphql",
	)
}

func _lgraphqlAddnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlAddnotificationtoprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x09\xe6\xe5\x97\x64\xa6\x65\x26\x83\x95\x85\x54\x16\xa4\x5a\x29\xf8\xa1\x89\x60\xaa\xf3\x4b\xcc\x4d\x85\x9b\xa2\xa9\x50\xcd\xa5\xa0\xa0\xa0\x90\x98\x92\x82\xa2\x33\x3f\x00\x62\x9d\x46\x66\x5e\x41\x69\x89\x15\x54\x95\x82\x02\xdc\x15\x30\xf7\x40\xc5\x31\x1d\x82\xe1\x36\x2c\x2a\x21\x4e\xc1\x70\x1d\x58\x65\xad\x26\xdc\xd2\xcc\x14\x98\x5e\xb8\x24\x17\x08\x03\x02\x00\x00\xff\xff\x5d\xba\xcd\xcd\x21\x01\x00\x00")

func _lgraphqlAddnotificationtoprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlAddnotificationtoprojectGraphql,
		"_lgraphql/addNotificationToProject.graphql",
	)
}

func _lgraphqlAddnotificationtoprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlAddnotificationtoprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/addNotificationToProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlLagoonschemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x88\x8f\x2f\x4e\xce\x48\xcd\x4d\x04\x73\x14\x14\x4a\x2a\x0b\x52\x8b\xa1\x6c\x05\x85\xec\xcc\xbc\x14\x28\x33\x2f\x31\x37\x15\xca\x4c\xcb\x4c\xcd\x49\x29\xd6\xc8\xcc\x4b\xce\x29\x4d\x49\x75\x49\x2d\x28\x4a\x4d\x4e\x2c\x49\x4d\xb1\x2a\x29\x2a\x4d\xd5\x84\x6b\x46\xd1\x53\xcb\x05\x23\x6b\xb9\x6a\x01\x01\x00\x00\xff\xff\x29\x07\x39\xef\x7e\x00\x00\x00")

func _lgraphqlLagoonschemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlLagoonschemaGraphql,
		"_lgraphql/lagoonSchema.graphql",
	)
}

func _lgraphqlLagoonschemaGraphql() (*asset, error) {
	bytes, err := _lgraphqlLagoonschemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/lagoonSchema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlLagoonversionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\xc8\x49\x4c\xcf\xcf\xcf\x0b\x4b\x2d\x2a\xce\xcc\xcf\xe3\xaa\x05\x04\x00\x00\xff\xff\x42\xb4\x77\x45\x19\x00\x00\x00")

func _lgraphqlLagoonversionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlLagoonversionGraphql,
		"_lgraphql/lagoonVersion.graphql",
	)
}

func _lgraphqlLagoonversionGraphql() (*asset, error) {
	bytes, err := _lgraphqlLagoonversionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/lagoonVersion.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlSshendpointsbyprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x00\xab\x83\x48\x2b\x28\x64\xa6\x80\x29\x90\x10\x98\x91\x9a\x57\x96\x59\x94\x9f\x97\x9b\x9a\x57\x52\x0c\x55\x02\x57\x84\xa4\x4c\x41\x21\xbf\x20\x35\xaf\x38\x23\x33\xad\x24\x00\x62\x85\x1f\x16\x29\xb8\x01\x48\x46\x28\x28\x14\x17\x67\x78\xe4\x17\x97\x20\xf3\x03\xf2\x8b\x60\xfc\x5a\x2e\x18\x59\xcb\x55\xcb\x05\x08\x00\x00\xff\xff\x1f\x6e\x6a\x1a\xdb\x00\x00\x00")

func _lgraphqlSshendpointsbyprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlSshendpointsbyprojectGraphql,
		"_lgraphql/sshEndpointsByProject.graphql",
	)
}

func _lgraphqlSshendpointsbyprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlSshendpointsbyprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/sshEndpointsByProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlTaskbyidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8c\x41\x0e\xc2\x30\x0c\x04\xef\x79\xc5\x22\xf5\x50\xbe\xc0\x91\x5b\x9e\x61\xd5\x16\x8a\xa8\x13\x70\x5c\xa4\x0a\xf5\xef\xa8\x29\xe9\xc9\xe3\xdd\xd5\xbc\x17\xb1\x15\x63\x00\x86\xc4\x37\xc4\xec\x97\x00\x5c\xbf\x01\x70\xaa\xcf\xfb\x1a\x79\xdc\x8b\x21\x71\x0b\x81\xc4\xed\x64\x52\x69\x30\x15\x55\xca\x47\x58\x9d\x7c\xa9\x1d\xcd\x85\xfb\xe4\x35\xcb\xf9\x99\x50\xe7\xb9\x3c\x8e\xb9\x89\x16\x97\xf8\xd7\x88\x7d\xd2\xb4\xeb\xb7\xb0\xfd\x02\x00\x00\xff\xff\x74\xe4\xc2\x1e\xa2\x00\x00\x00")

func _lgraphqlTaskbyidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlTaskbyidGraphql,
		"_lgraphql/taskByID.graphql",
	)
}

func _lgraphqlTaskbyidGraphql() (*asset, error) {
	bytes, err := _lgraphqlTaskbyidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/taskByID.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeployenvironmentbranchGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x2f\xb8\x68\xa1\x27\xc8\xb2\xe8\xd6\x85\x9e\x20\xd6\x51\x23\xed\xa4\x0c\x53\x41\xa4\x77\x97\x36\x89\x56\x70\x35\xcc\xfb\x9f\xcf\xeb\x47\x75\xea\x03\xa3\x34\xc0\x66\x90\x70\xa7\x56\x2d\x8e\x2a\x9e\xaf\x45\x8d\x99\x9e\xc4\x71\x7b\xfb\x0b\x0f\x74\xc9\xbc\x9e\xa9\x90\x8e\xc2\x5b\xa7\xce\xa2\x09\xa1\x23\xc7\x45\x85\x97\x01\x80\x33\x0d\x5d\x78\xee\xf8\xe1\x25\x70\x4f\xac\xcd\x32\x51\x7a\x1e\x46\xb5\xa9\x04\x64\x87\xfc\x03\xec\x7a\xb2\x1f\xb9\x84\xa7\x74\xa3\xc7\x3e\x56\xe2\xf3\x93\x2c\x86\x5f\xdb\x94\xad\x3d\x57\xd2\x26\x2f\x57\x66\x7a\x07\x00\x00\xff\xff\xb1\xcb\x5e\x73\x1a\x01\x00\x00")

func _lgraphqlDeploymentsDeployenvironmentbranchGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeployenvironmentbranchGraphql,
		"_lgraphql/deployments/deployEnvironmentBranch.graphql",
	)
}

func _lgraphqlDeploymentsDeployenvironmentbranchGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeployenvironmentbranchGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deployEnvironmentBranch.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeployenvironmentlatestGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\xb1\x6a\xc4\x30\x0c\x86\x77\x3f\x85\x0a\x1d\x7a\xaf\xe0\xb9\x37\x84\x96\x2e\x85\x5b\x4a\x07\x1d\x11\x45\xd4\x91\x83\x22\x07\x42\xc9\xbb\x17\x27\x4e\xe2\x9c\x37\xff\xfa\x2c\x7f\x7f\x97\x0c\x8d\xa3\xc0\x8b\x03\x78\x26\x19\x59\xa3\x74\x24\xe6\xe1\x7a\x5c\x1a\xe9\x93\x3d\x65\xe2\x9e\xc2\x6f\xd3\x7a\xf8\x34\x65\xf9\xd9\x92\x0f\xec\xa8\xce\x7a\xe5\xa8\x6c\x93\x87\x46\x6c\x85\x38\xb4\x37\x54\xc6\x7b\xa0\xc1\xc3\xd7\x55\xc6\x37\x9a\x6e\x18\x12\x2d\xcb\xbf\xdd\x05\xfe\x1c\x40\x4b\x7d\x88\x53\xf5\xf5\x3b\x1a\x0d\x96\xed\x00\x38\x93\x7e\xe1\xf2\x39\xd9\xd6\xee\x65\xbe\xb9\x16\xe9\x2a\x5d\x7d\x77\xf5\x32\x39\xac\xf7\x02\xfb\x9b\xb3\xfe\x43\x9f\x42\x29\x59\x52\x79\x45\x43\x0f\xa6\x69\x5d\x3b\x3b\x80\x8b\x9b\xdd\x7f\x00\x00\x00\xff\xff\xf3\x83\xe7\x9d\x69\x01\x00\x00")

func _lgraphqlDeploymentsDeployenvironmentlatestGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeployenvironmentlatestGraphql,
		"_lgraphql/deployments/deployEnvironmentLatest.graphql",
	)
}

func _lgraphqlDeploymentsDeployenvironmentlatestGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeployenvironmentlatestGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deployEnvironmentLatest.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeployenvironmentpromoteGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\xaa\x83\x30\x10\x45\xf7\xf9\x8a\x2b\xbc\x85\xc2\xfb\x82\x2c\x1f\xaf\xfb\x42\xbf\x20\xe8\x50\x52\xcc\x8c\xc4\x49\xa1\x14\xff\xbd\xd8\x1a\x6b\xab\xae\x02\xb9\xe7\x0c\x33\x37\x24\x75\xea\x85\x51\x1a\xe0\xa7\x8b\x72\xa1\x5a\x2d\x4e\x1a\x3d\x9f\x8b\x5f\x8c\xbf\xbd\xa4\x58\xd3\x81\xaf\x3e\x0a\x07\xe2\xef\xbc\xa1\x5e\x3d\x3f\xc7\x6c\x42\x23\x13\x49\x53\xe4\x7f\xa7\xce\xe2\x4f\xa4\x25\xc7\x45\x85\xbb\x01\x80\x86\xba\x56\x6e\x0b\xf5\x18\x25\x88\x52\xe9\xb9\x4b\x6a\x5f\x10\xb0\x5e\x23\x27\x00\xbb\x40\x76\x63\xd3\x19\xc8\x97\xbd\x95\x59\x9a\xa2\x39\x18\xcc\xe7\xbb\x56\x37\xc5\x8c\xef\xb5\xb1\x53\xd3\x64\x2d\xfb\x59\x94\x65\xf2\xe4\xca\x0c\x8f\x00\x00\x00\xff\xff\xcd\xc9\x92\xe1\xab\x01\x00\x00")

func _lgraphqlDeploymentsDeployenvironmentpromoteGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeployenvironmentpromoteGraphql,
		"_lgraphql/deployments/deployEnvironmentPromote.graphql",
	)
}

func _lgraphqlDeploymentsDeployenvironmentpromoteGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeployenvironmentpromoteGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deployEnvironmentPromote.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeployenvironmentpullrequestGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x41\x4e\xc3\x30\x10\x45\xf7\x39\xc5\x8f\xc4\xa2\x95\x7a\x02\x2f\x2b\x58\x74\x83\x2a\x38\xc1\xb4\x1d\xa8\x91\x33\x0e\xd3\x31\x12\x42\xbd\x3b\x4a\x33\x85\xc6\x51\x77\xc9\xff\x7f\x2c\xbd\xd7\x15\x23\x8b\x59\xb0\x68\x80\x87\x5e\xf3\x07\xef\x2d\x60\x3b\x7e\x6c\xa4\x2f\xd6\xae\x30\x74\x52\xba\x1d\x6b\xc0\x46\xac\x5d\x0d\x81\x45\x4b\x1c\xf0\x6a\x1a\xe5\x7d\x8c\x76\x74\xe2\xb5\x92\xec\x8f\xcf\xd4\xdd\xed\x5e\xf8\x6d\x5a\x1d\x99\x0e\xf7\xce\xfe\xbb\xd9\x99\xb2\x15\x95\x47\x32\x0a\x58\xe7\x9c\x98\xa4\x5d\xe2\xa7\x01\x80\x03\xf7\x29\x7f\x3f\xc9\x57\xd4\x2c\x1d\x8b\x6d\x4b\x4a\xca\x9f\x85\x4f\xb6\x88\x03\x56\xf0\x25\xf0\x47\x7d\xe5\xf7\xfc\x4a\xec\xe8\x9e\x3a\xf6\x88\xef\x59\xcd\x5d\x89\x98\xad\x2e\x28\x53\x23\xbe\xa9\x55\x54\x6e\x66\xab\xf1\xa5\xc9\xbf\x6f\x6e\xed\xdc\xa8\xba\xb4\xe7\x06\x58\x36\xe7\xdf\x00\x00\x00\xff\xff\x49\x97\xf5\x67\xfd\x01\x00\x00")

func _lgraphqlDeploymentsDeployenvironmentpullrequestGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeployenvironmentpullrequestGraphql,
		"_lgraphql/deployments/deployEnvironmentPullrequest.graphql",
	)
}

func _lgraphqlDeploymentsDeployenvironmentpullrequestGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeployenvironmentpullrequestGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deployEnvironmentPullrequest.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeploymentbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x4d\x4e\x80\x30\x10\x85\xf7\x9c\x62\x4c\x58\x68\xc2\x09\xba\x74\x67\x62\xd8\x70\x02\xa0\x13\x33\xd2\x3f\x87\xa9\x49\x43\xb8\xbb\x29\x6a\x9d\x08\x2b\xbe\xf7\x26\xef\xeb\x47\x46\x2e\x60\x31\xb9\x58\x3c\x06\x79\x2e\xe3\xec\xf1\xb1\x4f\x1c\xdf\x71\x95\x0a\x06\x26\x61\x0a\x6f\x0f\x03\xf4\x18\x3e\x89\x63\xa8\x97\xff\xab\xa0\xf9\x09\x8e\x0e\x00\xee\xc3\x14\x52\x16\x73\xc0\xcf\xbe\x01\x6d\x1a\x40\xed\x9b\x9b\x6d\x80\x6f\xc7\xa5\x3a\x7f\x15\xf5\x23\xdb\x7e\x6b\xd7\x60\x97\x59\xf2\xde\x70\x65\x9c\x05\xad\xae\x59\xf3\x1a\x7d\x72\xa8\x13\x46\x1f\x05\x5f\xfe\x82\x4c\xaf\x14\xb6\x86\x89\x29\x32\x49\x69\xc1\x92\xdd\xa6\xce\x2b\x8e\xfa\x45\x4b\x26\x67\x27\xc1\x74\x25\x67\x77\x7e\x05\x00\x00\xff\xff\xdc\xfc\xa7\xda\x81\x01\x00\x00")

func _lgraphqlDeploymentsDeploymentbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeploymentbynameGraphql,
		"_lgraphql/deployments/deploymentByName.graphql",
	)
}

func _lgraphqlDeploymentsDeploymentbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeploymentbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deploymentByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeploymentbynamewithlogGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x4d\x6a\xc4\x30\x0c\x85\xf7\x73\x0a\x15\x66\xd1\x42\x4e\xe0\x65\x77\x85\x61\x36\x39\x41\x12\x8b\xa0\xc6\x7f\x55\xe4\x82\x09\xb9\x7b\x71\xda\xba\xea\x24\xab\x7c\xef\x89\xf7\xf9\x23\x23\x17\xb0\x98\x5c\x2c\x1e\x83\xbc\x96\xfb\xe0\xf1\xf9\x9a\x38\xbe\xe3\x24\x15\x0c\xf4\xc2\x14\xe6\xa7\x0e\xae\x18\x3e\x89\x63\xa8\x97\x8f\x55\xd0\xfc\x02\xdb\x05\x00\xce\xc3\x14\x52\x16\xb3\xc1\xcf\xbe\x01\x6d\xea\x40\xed\x9b\x93\xad\x83\x6f\xc7\xa1\xda\x7f\x15\xf5\x23\xdb\x7e\x6b\xd7\x60\x95\x41\xf2\xda\x70\x62\x1c\x04\xad\xae\x59\xf3\x14\x7d\x72\xa8\x13\x46\x1f\x05\xdf\xfe\x82\x4c\x37\x0a\x4b\xc3\xc4\x14\x99\xa4\xb4\x60\xcc\x6e\x51\xe7\x15\xef\xfa\x45\x63\x26\x67\x7b\xc1\xf4\x3f\xb9\xc5\xf9\x08\xf6\xcb\xfe\x15\x00\x00\xff\xff\x74\xa3\x41\xb0\x92\x01\x00\x00")

func _lgraphqlDeploymentsDeploymentbynamewithlogGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeploymentbynamewithlogGraphql,
		"_lgraphql/deployments/deploymentByNameWithLog.graphql",
	)
}

func _lgraphqlDeploymentsDeploymentbynamewithlogGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeploymentbynamewithlogGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deploymentByNameWithLog.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsDeploymentbyremoteidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x4d\x0a\xc2\x30\x10\x85\xf7\x3d\xc5\x08\x2e\xf4\x0a\x5d\xba\x13\xc4\x85\x3d\x41\xdb\x0c\x32\x34\x7f\x4e\x27\x42\x90\xde\x5d\x8a\x76\x18\xd0\x2c\xc2\x7c\x5f\x5e\x48\xde\xa3\x20\x57\x70\x98\x7d\xaa\x01\xa3\x9c\xea\x0d\x43\x12\x3c\xbb\xc3\x9e\xbf\x53\x0b\x9d\x30\xc5\xfb\xee\x08\xaf\x06\x00\xfe\xc7\xc9\xb5\xa0\x57\xb6\xe4\xba\xc8\xe9\x18\xfb\x80\x0a\xb3\xf4\x52\x66\xc5\x91\xb1\x17\x74\xf6\x98\x2d\x8f\x29\x64\x8f\xd6\x6c\x8f\xa9\x28\x74\xa1\x38\x29\x66\xa6\xc4\x24\x55\xc5\x50\xfc\x64\xe2\x2b\x5e\xed\x8f\x86\x42\xde\x75\x82\x59\x0d\xc6\x27\x71\x8a\x6b\x55\xd3\xe8\xa7\xca\xd2\x7c\xf6\xe5\x1d\x00\x00\xff\xff\xfe\x8f\x07\x79\x4f\x01\x00\x00")

func _lgraphqlDeploymentsDeploymentbyremoteidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsDeploymentbyremoteidGraphql,
		"_lgraphql/deployments/deploymentByRemoteID.graphql",
	)
}

func _lgraphqlDeploymentsDeploymentbyremoteidGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsDeploymentbyremoteidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/deploymentByRemoteID.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsGetdeploymentsbybulkidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x4d\x0e\x82\x30\x10\x85\xf7\x9c\xa2\x26\x2e\xe0\x0a\x2e\xd9\x91\x18\x37\x9c\x00\xe8\xc4\x8c\xf4\xcf\x61\x6a\xd2\x18\xee\x6e\x00\x2d\x05\xe9\x6a\xbe\xd7\xf7\x32\xf3\x9e\x1e\x28\x88\xfc\xdc\x7a\xd5\x57\xf2\x22\x6a\x26\x34\xf7\x53\x21\xde\x99\x10\x42\x48\x70\xca\x06\x0d\x86\x87\x32\x94\xb3\x27\xff\x59\xbf\x99\x62\x36\x2e\xf6\xe9\xa1\x8c\xa3\x69\x34\x44\x18\xb8\x61\x3f\x44\xec\x08\x1a\x06\x99\x7e\x53\xca\x9d\xd5\x4e\x41\xaa\x10\x68\xcb\x50\xad\x82\xc7\x2b\x9a\x3e\xa2\x23\xb4\x84\x1c\xa2\xb0\xdc\xb7\xc1\x5b\x7a\x51\xeb\x51\xc9\x9a\xc1\x45\x05\xcc\x0b\xc9\x9a\xa9\xef\x5a\x68\x57\xea\xaf\xd8\xb2\xdb\x3e\xa0\xdb\x85\x0e\x82\x87\xe1\x31\xdb\x4e\x63\x36\x7e\x02\x00\x00\xff\xff\xfd\xd7\x46\x98\x97\x01\x00\x00")

func _lgraphqlDeploymentsGetdeploymentsbybulkidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsGetdeploymentsbybulkidGraphql,
		"_lgraphql/deployments/getDeploymentsByBulkID.graphql",
	)
}

func _lgraphqlDeploymentsGetdeploymentsbybulkidGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsGetdeploymentsbybulkidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/getDeploymentsByBulkID.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcf\x4f\x6a\x85\x30\x10\x06\xf0\xbd\xa7\x18\xc1\x85\x85\x9e\xc0\x65\x77\x42\x71\xe3\x09\xa2\x19\xca\xd4\xfc\xeb\x38\x29\x04\xf1\xee\x45\xfb\x94\x97\xbc\x59\x64\xf1\xfb\x32\x1f\xc9\x4f\x44\x4e\xd0\x36\x81\xfd\x37\xce\xd2\x41\xef\xa4\x7e\x87\x06\xdd\x2f\xb1\x77\x16\x9d\x74\x30\x0a\x93\xfb\xaa\xdf\xb6\x0a\x00\xe0\x29\xfa\x48\x83\xb2\xd8\x9e\x7c\xcd\xdd\x74\x75\x66\xa9\x53\x16\xbb\xac\xfe\x8c\x1f\xd5\xc7\x68\x0c\xc6\xa7\x23\x59\xb7\x6c\x95\xf4\x4b\x53\x06\xab\x28\x89\x6b\x46\x33\xa3\x12\xd4\xe5\x35\x2e\x6d\xf6\x36\x18\x2c\x95\xd1\x7a\xc1\x3e\xc7\x48\x9f\xe4\x96\xe2\xc7\xe4\x99\x24\x65\x38\x45\xb3\x14\xab\x07\x0d\xe5\xab\xa7\x48\x46\x8f\x82\xe1\xd6\xbd\xfa\x3f\xf7\xbf\x00\x00\x00\xff\xff\x94\xe5\x53\xa7\x9c\x01\x00\x00")

func _lgraphqlDeploymentsGetdeploymentsforenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql,
		"_lgraphql/deployments/getDeploymentsForEnvironment.graphql",
	)
}

func _lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsGetdeploymentsforenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/getDeploymentsForEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploymentsUpdatedeploymentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x31\x8e\x84\x30\x0c\x45\x7b\x4e\x61\xa4\x2d\x96\x2b\xa4\xde\x06\x69\x8b\x69\xe6\x00\x19\x62\x69\x2c\x88\x13\x05\x67\x24\x84\x72\xf7\x51\x00\x85\x40\x2a\xff\xe7\x6f\xff\xd8\x46\xd1\x42\x8e\xe1\xb7\x01\x00\xf8\x21\xa3\xa0\x67\x69\x77\xe5\xb5\x0c\x6f\x05\x4f\x6f\xb4\xe0\x1f\xfa\xc9\x2d\x16\x59\x1e\x19\xf7\xec\xa3\xb4\xdd\xba\x39\xe3\xcd\xb1\x6f\xcb\x8f\xb2\x4d\xc1\x5a\xc0\x06\x8d\xca\x51\x17\x76\x64\xed\x99\xa5\x93\xb6\xaa\x3b\xc7\xab\x29\xd6\x16\x8b\x98\x45\x4b\x9c\x8b\x1c\x02\x6a\x41\x53\xb7\x43\xad\x07\x67\xfd\x84\x35\x09\x68\x9d\x60\x7f\x82\x48\xff\xc4\x63\x91\xc8\x1f\x0a\x8e\xf3\x75\xb7\x6b\x2e\xff\x48\xa5\xf2\x81\x5c\x20\x59\x0a\x78\xc5\x69\x3c\xf6\xa7\x26\x7d\x03\x00\x00\xff\xff\xdd\x19\x4b\x1e\x7a\x01\x00\x00")

func _lgraphqlDeploymentsUpdatedeploymentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploymentsUpdatedeploymentGraphql,
		"_lgraphql/deployments/updateDeployment.graphql",
	)
}

func _lgraphqlDeploymentsUpdatedeploymentGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploymentsUpdatedeploymentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployments/updateDeployment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsAdddeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xd1\xcf\x4e\xc3\x30\x0c\x06\xf0\x7b\x9f\xc2\x48\x1c\x36\x69\x4f\xd0\x2b\x1c\x80\xc3\x56\x31\x78\x80\x40\xdc\x2e\x22\x8d\x2b\xd7\x45\x9a\x10\xef\x8e\xfa\x27\x55\xbc\x94\x43\x0f\xfd\xbe\x44\xb2\x7f\x69\x07\x31\xe2\x28\xc0\xae\x00\xb8\x0f\xa6\xc5\x12\xce\xc2\x2e\x34\x77\x87\x31\xf9\xa4\xd0\x93\xc7\x77\xf6\x3a\x17\xfa\xc2\x10\xa3\x29\x61\x1a\x04\xb9\x32\x22\xc8\xba\xe9\xfb\xcb\x13\xf5\x72\x9b\x55\xc4\x3a\xfb\x18\x9c\xb7\xcf\xad\x69\x50\xc5\x2d\x05\x27\x34\xfe\x3e\x50\xa8\x5d\x53\xc2\xcb\xf9\x74\x9c\xaa\x9a\x1d\x06\xeb\xaf\xc7\x64\xee\x79\x6c\x4f\x83\xad\x98\xbe\x9d\x45\xce\x9b\x57\x6c\x1c\xad\x43\xee\xe1\xa7\x00\x00\x30\xd6\x3e\x62\xe7\xe9\xfa\x66\xb8\x41\x29\xc7\xe0\xd4\x61\xe8\x2f\xae\x96\x9d\x0b\xdd\x20\xe5\x72\x14\x60\xa6\x9a\xc4\x0e\x4b\x94\x5a\x25\x70\xb1\x5e\xc8\x66\xba\x25\xbb\x41\xd3\x88\xf1\xe2\xea\x17\x25\x93\x62\x46\x8c\x9c\xb1\x48\x25\x13\xd6\x58\xe7\xa2\x19\x72\x3c\xaa\x85\x15\xf8\xba\xb6\xb6\xd6\xf6\xea\x50\x64\x4f\x1f\x61\xea\x7f\xf7\xab\xab\xb3\x09\x70\xbc\xcc\x68\x04\x6d\xea\x98\x91\x6b\x2b\x0d\xb4\x35\x68\x3e\xd7\xc6\xc6\x5b\xaf\x94\x11\xff\x83\x3a\x2f\x56\x8c\xdf\x5f\x00\x00\x00\xff\xff\x0d\x3b\x84\x12\x65\x03\x00\x00")

func _lgraphqlDeploytargetsAdddeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsAdddeploytargetGraphql,
		"_lgraphql/deploytargets/addDeployTarget.graphql",
	)
}

func _lgraphqlDeploytargetsAdddeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsAdddeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/addDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsAdddeploytargettoorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x49\x49\x2d\xc8\xc9\xaf\x0c\x49\x2c\x4a\x4f\x2d\xb1\x52\xf0\xcc\x2b\x51\xd4\x01\x09\xe7\x17\xa5\x27\xe6\x65\x56\x81\x15\x42\x85\x35\xab\xb9\x14\x14\x12\x53\x52\x5c\x90\x74\x84\xe4\xfb\x23\x29\xd4\xc8\xcc\x2b\x28\x2d\xb1\x52\x00\x29\x54\x50\x40\x35\x19\xc5\x22\x1d\xb0\x02\x54\x3b\x50\xac\x04\x29\xa8\xd5\x84\x1a\x94\x99\x02\xa6\xf2\x12\x73\x53\x41\xc2\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\xec\x80\xbb\x8c\xc3\x00\x00\x00")

func _lgraphqlDeploytargetsAdddeploytargettoorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsAdddeploytargettoorganizationGraphql,
		"_lgraphql/deploytargets/addDeployTargetToOrganization.graphql",
	)
}

func _lgraphqlDeploytargetsAdddeploytargettoorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsAdddeploytargettoorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/addDeployTargetToOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsDeletedeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\xd2\x54\xa8\xe6\x52\x50\x48\x49\xcd\x49\x2d\x49\x75\x49\x2d\xc8\xc9\xaf\x0c\x49\x2c\x4a\x4f\x2d\xb1\x82\x8a\xf9\x17\xa4\xe6\x15\x67\x64\xa6\x95\x68\x64\xe6\x15\x94\x96\x58\x81\x95\x2b\x28\x40\x0c\x01\x9b\xc5\xa5\xa0\x50\xab\xc9\x55\xcb\x05\x08\x00\x00\xff\xff\x8e\x88\xf8\xcc\x66\x00\x00\x00")

func _lgraphqlDeploytargetsDeletedeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsDeletedeploytargetGraphql,
		"_lgraphql/deploytargets/deleteDeployTarget.graphql",
	)
}

func _lgraphqlDeploytargetsDeletedeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsDeletedeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/deleteDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsDeploytargetsbyorganizationidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x41\x8e\xc2\x30\x0c\x45\xf7\x39\x85\x47\xea\xa2\x73\x85\x59\xce\x6a\xba\x19\x55\x88\x0b\x84\xda\x04\x4b\x69\x0c\x8e\x83\x14\x50\xef\x8e\x58\x34\x2d\x1b\xcb\xef\x7d\xeb\xfb\x56\x48\x2b\xf4\x0e\xa0\x63\xfc\x81\x21\xd9\xd7\x37\x3c\x1d\x00\x80\x68\xf0\x89\x1f\xde\x58\xd2\x6f\x1d\xb0\x7f\x1f\x74\x8c\x6b\x0e\x80\x74\x8d\x52\x8f\x5e\x03\x59\x6e\x16\x80\xb1\xad\xc9\xcf\xd4\x60\x52\xf2\x46\x5b\x38\x45\x29\x78\xa0\xc0\x92\x3e\xdd\xa8\x72\x67\x24\x6d\xf6\xac\x4c\x09\x63\xfd\xdf\xd7\xe5\x7c\xf9\x93\x6c\x7b\x1e\x45\x37\x56\x29\x46\x3a\x7a\x33\xd2\xed\xc1\xa9\x70\xc4\x61\xf6\x61\x2d\x5a\xdc\x3a\x17\xf7\x0a\x00\x00\xff\xff\xe2\x3c\x45\xaa\x0f\x01\x00\x00")

func _lgraphqlDeploytargetsDeploytargetsbyorganizationidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsDeploytargetsbyorganizationidGraphql,
		"_lgraphql/deploytargets/deployTargetsByOrganizationId.graphql",
	)
}

func _lgraphqlDeploytargetsDeploytargetsbyorganizationidGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsDeploytargetsbyorganizationidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/deployTargetsByOrganizationId.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsDeploytargetsbyorganizationnameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4d\x8e\xc2\x30\x0c\x46\xf7\x39\x85\x47\x9a\x45\xe7\x0a\xb3\x9c\xd5\xb0\x41\x15\x70\x01\xd3\x98\x60\x29\x8d\xc1\x71\x90\x0a\xea\xdd\x11\xad\x9a\xc2\x26\xca\xf7\xfc\xfc\x73\x2d\xa4\x03\x34\x0e\xe0\x3b\x61\x4f\xbf\xb0\x37\xe5\x14\xbe\x7e\xe0\xe1\x00\x00\x44\x03\x26\xbe\xa3\xb1\xa4\xbf\x61\x8b\x3d\x35\xb3\x37\xe9\x8b\x05\xe0\xe9\x12\x65\x38\xa0\x06\xb2\x5c\x29\x00\xfb\xfa\x7d\x35\xd4\xd0\x29\xa1\xd1\x5a\xec\xa2\x14\xbf\xa3\xc0\x92\x3e\x59\xab\x72\x63\x4f\x5a\xe9\x49\x99\x92\x8f\xd3\x2d\x15\xe6\x7c\xfe\x97\x6c\xef\xb9\x15\x5d\xb3\x4a\x31\xd2\x16\xcd\x48\xd7\x05\xc7\xc2\xd1\x6f\x7a\x0c\xcb\xa0\xd1\xcd\xef\xe8\x9e\x01\x00\x00\xff\xff\x51\x79\xa0\x01\x18\x01\x00\x00")

func _lgraphqlDeploytargetsDeploytargetsbyorganizationnameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsDeploytargetsbyorganizationnameGraphql,
		"_lgraphql/deploytargets/deployTargetsByOrganizationName.graphql",
	)
}

func _lgraphqlDeploytargetsDeploytargetsbyorganizationnameGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsDeploytargetsbyorganizationnameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/deployTargetsByOrganizationName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsListdeploytargetsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\x51\x4a\xc4\x40\x10\x44\xff\x73\x8a\x3e\x87\xbf\xfa\xa1\x08\x12\x44\x0f\x30\xc9\x54\xc6\xc6\x4e\xb7\xf6\xf4\x2c\x84\x25\x77\x5f\x36\x43\xd8\xaf\x7a\x55\x3c\xea\xbf\xc1\x37\xba\x0e\x44\xc2\x35\x5e\xf0\x27\xb6\x7d\x25\x2f\x88\xfa\x44\x49\xe4\xbd\x4d\x70\x45\xa0\xde\x1d\x22\xce\x47\x68\x5a\x71\xc0\xec\x48\x81\x3e\xce\xa6\xd5\x04\xdf\x2e\xbd\x8a\xb5\xfc\x89\xc2\xa6\x8f\x3e\xba\x5d\x38\xc3\x8f\x65\x71\x86\x66\xd9\x3e\xce\xb7\xb0\x5f\x74\xb9\xd6\x9f\x57\xab\x71\xf2\x68\xde\xd9\xad\x05\x7c\x4c\x11\xf0\x6e\x4e\x8d\x25\xbf\xad\xa9\xf4\x8b\xd5\x94\xc3\x9c\xb5\x3c\x9b\x2e\x5c\x06\xa2\x7d\xd8\x6f\x01\x00\x00\xff\xff\x7c\xcb\xe1\x0e\xe7\x00\x00\x00")

func _lgraphqlDeploytargetsListdeploytargetsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsListdeploytargetsGraphql,
		"_lgraphql/deploytargets/listDeployTargets.graphql",
	)
}

func _lgraphqlDeploytargetsListdeploytargetsGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsListdeploytargetsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/listDeployTargets.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsRemovedeploytargetfromorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x49\x49\x2d\xc8\xc9\xaf\x0c\x49\x2c\x4a\x4f\x2d\xb1\x52\xf0\xcc\x2b\x51\xd4\x01\x09\xe7\x17\xa5\x27\xe6\x65\x56\x81\x15\x42\x85\x35\xab\xb9\x14\x14\x8a\x52\x73\xf3\xcb\x52\x5d\x90\x34\xb9\x15\xe5\xe7\xfa\x23\xa9\xd6\xc8\xcc\x2b\x28\x2d\xb1\x52\x00\xa9\x56\x50\x40\x35\x1e\xc5\x36\x1d\xb0\x02\x54\x8b\x50\xec\x05\x29\xa8\xd5\x84\x1a\x94\x99\x02\xa6\xf2\x12\x73\x53\x41\xc2\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\x56\xb2\x5d\x48\xc8\x00\x00\x00")

func _lgraphqlDeploytargetsRemovedeploytargetfromorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsRemovedeploytargetfromorganizationGraphql,
		"_lgraphql/deploytargets/removeDeployTargetFromOrganization.graphql",
	)
}

func _lgraphqlDeploytargetsRemovedeploytargetfromorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsRemovedeploytargetfromorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/removeDeployTargetFromOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetsUpdatedeploytargetGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x41\x6e\xea\x30\x10\x86\xf7\x73\x8a\x79\x12\x0b\x90\x38\x41\xb6\xaf\x8b\xd2\x05\xa0\xd2\x1e\xc0\xc5\x43\xb0\x9a\x78\x22\x67\x52\x09\x55\xdc\xbd\x0a\x24\xb1\x27\xce\x2a\xf1\xff\x8f\xa5\xf1\xf7\xd5\x9d\x18\x71\xec\x71\x0d\x88\x2b\x67\x0b\xdc\x79\xf9\xd7\xff\x7b\x53\x53\x81\x27\x09\xce\x97\xdb\x3e\x38\xb3\x6f\xb9\xa2\xcf\x50\xa9\x58\xf8\x9b\xbc\x4a\x02\x77\x42\xe1\x68\x44\x28\xe8\xa6\x6d\xaf\xaf\xdc\xca\x3c\x3b\x72\xd0\xd9\x57\xe7\x2a\xbb\xab\x4d\xa9\x37\xa8\xd9\x3b\xe1\xfe\xf8\x9f\xfd\xc5\x95\x05\xbe\x9d\x0e\xfb\x47\x75\x09\x8e\xbc\xad\x6e\xfb\x6c\xeb\x8a\x3b\x7b\x0c\xfc\xe3\x2c\x85\xbc\x79\xa7\xd2\xf1\xb4\xe4\x06\x7f\x01\x11\xb1\x6b\xac\x11\x7a\xa1\xa6\xe2\xdb\x87\x09\x25\x49\x31\x64\x87\x86\x7c\x7b\x75\x17\x59\x3b\xdf\x74\x52\x0c\x17\x10\x7b\x74\x2b\x67\x87\x53\x63\xe4\x7c\x8d\x25\xe2\x93\xe6\x03\xea\x76\x0a\x53\xa2\x09\xde\x38\x30\xb0\x7d\x32\x9e\xd2\x19\x5f\xcd\x3b\x5e\x9e\x60\x8f\xd8\x55\xf5\x64\x3e\xd2\x8f\x55\x8a\x3e\xf1\x10\x07\x72\x09\x99\x97\x38\xac\xb5\x28\x4b\x09\x06\xad\x48\x2b\x9b\x8d\x8d\xbe\x52\x7b\xc3\xc4\xfd\xf1\xbd\x6f\x12\x25\x10\xe1\x0f\xbf\xe7\x40\x46\x68\x2c\x52\xae\x91\x3f\x28\x7e\xa0\x90\xc1\xc2\xca\x90\xed\x07\xf9\xeb\x61\x8e\x17\x16\x64\xc2\x32\x62\x18\x9f\x77\x87\xbf\x00\x00\x00\xff\xff\x63\x54\xb3\x44\xb0\x03\x00\x00")

func _lgraphqlDeploytargetsUpdatedeploytargetGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetsUpdatedeploytargetGraphql,
		"_lgraphql/deploytargets/updateDeployTarget.graphql",
	)
}

func _lgraphqlDeploytargetsUpdatedeploytargetGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetsUpdatedeploytargetGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargets/updateDeployTarget.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x4e\xc3\x30\x0c\x86\xef\x7d\x8a\x20\x71\xd8\x5e\xa1\x57\xb8\x70\x41\x15\xf0\x02\xa1\xf1\x32\xa3\xcc\x29\x9e\x03\x9a\x50\xdf\x1d\x6d\x59\x92\x25\xad\x96\xe3\xef\xaf\x76\xbf\xff\x10\x44\x0b\x7a\x52\x9b\x4e\x29\xa5\x1e\x27\xf6\x5f\x30\x4a\xaf\x5e\x48\x1e\x62\xf4\x0b\x68\xf7\x31\x89\xc1\x27\x6b\x1a\xf7\x70\xec\xd5\xbb\x30\x92\xbd\x72\x53\x70\x8e\xe1\x3b\xc0\x51\xda\x91\x81\xc9\xf9\xd3\x87\x66\x0b\xd5\xea\xdb\x7c\x88\x97\x07\x2d\x02\x4c\x69\x41\xb7\xfd\xbb\xa0\xda\x98\xe7\x1b\xf8\xc9\xd3\x0e\xed\x06\x69\x0a\xd2\x47\xe2\xfc\xea\x3b\xd5\xfa\xcc\x24\x9d\xab\x57\xce\x8b\x55\x16\xcc\xb3\x5a\xad\x32\x2d\x4c\x6a\x2e\x75\xb8\xfa\x57\xad\xe5\x9d\x0a\x2e\xdf\xcf\xdb\x62\x87\x66\x75\x65\x01\x1a\xe8\xfc\x48\x1f\xa0\x0a\x76\x8c\x40\xc6\x9d\x5e\xdb\xc1\xe8\x7c\x30\x03\xfb\x1f\x34\xc0\xcb\xc9\x1b\x58\xf4\x94\xf3\x79\xd1\xdb\x6a\x59\x4d\xeb\x51\xa9\x9b\xff\x03\x00\x00\xff\xff\x2a\xf3\x98\x73\x76\x02\x00\x00")

func _lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql,
		"_lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql",
	)
}

func _lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\x0a\x8a\xf2\xb3\x52\x93\x4b\xa0\x42\x9a\xd5\x60\xd1\x94\xd4\x9c\xd4\x92\x54\x97\xd4\x82\x9c\xfc\xca\x90\xc4\xa2\xf4\xd4\x12\xe7\xfc\xbc\xb4\xcc\x74\x8d\xcc\xbc\x82\xd2\x12\x2b\x88\x22\x10\x00\x99\xa5\x92\x99\x02\xe7\xc3\x4d\x83\x99\x0b\x96\xa9\xd5\xe4\xaa\x05\x04\x00\x00\xff\xff\x66\x8a\x2f\xa2\x86\x00\x00\x00")

func _lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql,
		"_lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql",
	)
}

func _lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\xca\xc2\x30\x10\x84\xef\x79\x8a\xfd\xe1\x3f\xe8\x2b\x78\xd4\x53\x2f\x52\xc4\x17\xa8\xcd\x34\x8d\xc4\x6c\xbb\x4d\x94\x20\x7d\x77\x69\x82\xa5\xe6\x90\x65\xbf\xd9\x19\x66\x8c\x90\x44\x3b\x45\xf4\x3f\x08\xdf\xd1\x86\x03\x55\x3e\xfc\xa9\xfd\x5b\x11\x69\x0c\x8e\xd3\xb5\x11\x83\x70\x62\xdf\x59\x33\x1d\x53\x5d\xee\x2a\xbd\xb8\x88\x56\xdb\x37\x40\x11\x65\x33\x91\xd5\x79\x6c\x53\x8a\xb0\x11\x97\xe7\x9b\x07\xd6\xa5\x13\x0b\xaf\x5d\x3a\x6f\x61\xeb\x38\xea\x5a\xf8\x69\x35\xe4\x97\x5e\x60\x2c\xfb\xcc\xe6\xfc\xdf\xa4\xf1\x6d\x8f\xa9\xb4\x8b\xce\x09\xc6\x88\x29\x14\xf0\x82\x35\xfd\xd2\x71\x56\xf3\x27\x00\x00\xff\xff\xba\x92\xa1\x8b\xfd\x00\x00\x00")

func _lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql,
		"_lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql",
	)
}

func _lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x4e\xf4\x30\x0c\x85\xf7\x39\x45\x7e\x69\x16\x33\x57\xe8\xf6\x67\xc3\x06\x55\xc0\x05\x42\xe3\x49\x8d\x32\x4e\xf0\x38\xa0\x11\xea\xdd\x51\x1b\x28\x71\x19\xf0\xaa\x79\xef\xd5\x7a\xfe\x4e\x45\x9c\x60\x22\xbb\x37\xd6\x5a\xbb\x43\xdf\xd9\x5b\x92\x7f\xf5\xf5\x06\x18\x46\x59\x94\x2a\x3c\xb1\xa3\x61\x84\x73\x67\x1f\x84\x91\x42\x55\x73\x89\x91\xe1\xa5\xc0\x59\x36\x8e\x87\x1c\xd3\xe5\xd1\x71\x80\x76\x4d\x2b\xf7\x9c\x9e\x61\x90\xde\x89\x00\xd3\xfa\xfb\xe1\x7d\x89\x96\xec\x9d\xc0\x4d\x93\xff\x9f\xe8\x88\x61\x8f\x94\x8b\x74\x35\x34\xcf\x5c\x7c\x87\x7e\x7d\x67\x27\xc3\xd8\xf8\xf3\xe8\x36\xaa\x85\xca\x7d\x9d\xfd\x79\xbf\xf2\xbe\x09\xac\x30\x94\xaf\x59\x28\x34\xbf\x76\xd9\x22\xf8\x83\xcf\xba\x63\x5a\xbe\xa6\x43\x4b\xc0\x5c\x5b\xae\x11\x34\xa1\x79\xc8\x9d\x40\x09\x47\x46\x20\x1f\x2f\x77\x5b\x63\x88\xa9\xf8\x9e\xd3\x2b\x7a\xe0\x9f\xce\x3d\x04\x4c\xdb\x76\x2d\x2f\x73\x0d\x90\xd1\xc4\xeb\x49\x66\xfa\x08\x00\x00\xff\xff\x82\xe7\x64\x80\x95\x02\x00\x00")

func _lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql,
		"_lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql",
	)
}

func _lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsAddprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xcf\x4e\xc3\x30\x0c\xc6\xef\x7b\x8a\x4c\xe2\xb0\x49\x7b\x82\x1d\x41\x1c\x26\x90\x36\x09\xf1\x00\x5e\xe3\x75\x41\xa9\x53\x1c\x67\x08\x10\xef\x8e\xda\x35\xff\xc6\xd8\xa1\x97\xcf\xdf\xd7\xd8\xfe\xb9\x0b\x02\x62\x1c\xa9\xc5\x4c\xa9\x3b\x82\x0e\xd7\xea\x45\xd8\x50\x3b\x5f\x0d\x8a\xe3\x16\xc8\x7c\x8d\x9e\xb5\xda\x90\x8c\x6a\x6b\xe4\x95\x6d\xed\xf4\x61\x7f\x70\x56\x23\x47\xf9\x9c\xef\x91\xfc\xd1\x1c\x64\x0c\xcf\x6b\x6d\xc7\xee\x0d\x1b\xd9\x81\x08\x32\x55\xb9\x3d\x03\x35\x47\xf4\x95\xd8\x07\x6b\x19\xdf\x03\x7a\xb9\x28\xb0\xd3\xa1\x19\x7a\x7c\xa4\x93\x61\x47\x1d\x92\xd4\xed\x41\x10\xb7\xd1\x16\xf3\x10\x5e\x1c\x43\x8b\x0f\x60\x9b\x2c\x6a\x3c\xa1\x75\xfd\x90\x2f\x7e\xe5\x9f\x4d\x67\x24\xbb\x40\xeb\x2d\xb7\xdb\x0f\x1a\xa6\xbd\x77\xce\x22\xd0\xd4\x88\x39\x81\xe0\x13\x7e\xc6\xd7\x97\xea\x7b\xa6\x94\x52\xa0\xf5\x34\xee\xc2\x50\x1f\x64\x3d\xe9\x4a\x9d\x97\x3e\xee\x7e\x35\x49\xf5\xd6\x2b\x08\xd1\x12\x11\x4c\x2c\xa2\x5c\x50\xc8\x44\xd2\x6f\x33\x8c\x0c\xe1\x4f\xf1\x92\xca\x7f\xbc\x62\x30\xa3\x4a\xd4\x62\xa9\x06\x56\xf1\x4b\x96\xeb\xe8\xae\x23\x8d\xa1\x0c\x33\x71\x4d\x0b\x28\xa9\x96\x8c\xa3\xe1\x36\xe1\x9b\x07\x90\x9e\x2f\xf1\x97\xc7\x90\x87\xca\x67\x50\xdc\xc4\x58\xfd\x59\x26\xf2\x46\x17\x27\x70\x2e\xce\x86\xef\x37\x00\x00\xff\xff\xc5\xe9\x48\x5a\x95\x03\x00\x00")

func _lgraphqlProjectsAddprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsAddprojectGraphql,
		"_lgraphql/projects/addProject.graphql",
	)
}

func _lgraphqlProjectsAddprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsAddprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/addProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsMinimalprojectbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x4a\x04\x31\x10\x45\xf7\x39\x45\x09\x2e\xc6\x2b\xb8\x14\x5c\x08\x22\x03\xe2\x01\x62\xa7\xec\x29\x49\xaa\x32\x95\xca\xc0\x28\x73\x77\x49\x37\x26\x2d\x32\xd9\x24\x79\xbf\x7e\xf2\xff\xb1\xa2\x9e\x61\xe7\x00\x6e\xd9\x27\xbc\x87\x57\x53\xe2\xf9\xe6\x0e\xbe\x1d\x00\x40\x56\xf9\xc4\xc9\x1e\xce\x2f\x3e\xe1\x6e\x41\x00\xeb\xe4\x62\xf8\x9d\x6b\x8b\x42\x3f\x36\xa9\x5f\x7c\x35\x79\x0a\x71\x80\x77\xf5\x3c\x1d\xb0\x74\x90\x6b\x8c\x8a\xc7\x8a\xc5\x36\x50\x25\xd4\xc9\x48\xf8\x91\x4f\xa4\xc2\x09\xd9\xba\x2a\x19\xb9\x1c\xe8\xc3\xf6\x6b\xc2\xbd\x37\x43\xe5\xae\x07\x3c\x61\x94\xdc\x3c\x1b\x7b\x79\xa6\x44\xe3\x91\x99\xec\x4d\xe3\xf5\xa0\xa2\xb3\x67\xfa\xf2\x2d\xc4\xff\x9f\x47\xf3\x3f\xdd\x2f\x6e\xec\x17\xf7\x13\x00\x00\xff\xff\x0b\xea\x82\x53\x61\x01\x00\x00")

func _lgraphqlProjectsMinimalprojectbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsMinimalprojectbynameGraphql,
		"_lgraphql/projects/minimalProjectByName.graphql",
	)
}

func _lgraphqlProjectsMinimalprojectbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsMinimalprojectbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/minimalProjectByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\xcd\x6e\xdb\x30\x0c\xbe\xe7\x29\x38\x60\x87\xee\x92\x07\xd8\x6d\x2d\x8a\x6d\xe8\x16\x14\x6d\xd6\x6b\xc1\xc8\x74\xac\x45\x16\x5d\x92\xf6\x60\x0c\x7d\xf7\xc1\x09\xe6\x5f\x35\xeb\x61\x40\x75\x89\xf2\x7d\xf4\x47\x8a\x7f\x4f\x35\x49\x0b\x17\x2b\x80\xf7\x11\x4b\xfa\x08\xf7\x26\x3e\xee\xdf\x7d\x80\xdf\x2b\x00\x80\x4a\xf8\x27\x39\xbb\x6c\x37\x58\xd2\xc5\x11\x02\x38\x59\x1e\x3f\xf8\x6b\xd7\x1d\x9f\xf5\xd7\x8e\xea\xff\x60\x6d\xfc\x35\x0b\x03\xb0\x13\x8c\xae\x20\xed\x81\xaa\x0e\x41\xe8\xa9\x26\xb5\x11\x28\xbe\x41\xa3\x1b\x6a\x47\x10\x67\xb5\x33\xcf\xf1\x3a\x36\x5e\x38\x96\x14\xad\x67\xd5\x58\x70\x4f\x57\x18\x5c\x8f\x71\x45\x51\x0b\x9f\xdb\xed\xe9\x21\xb7\x68\x46\x12\x7b\x3e\xa3\x86\x02\x57\x9d\xce\x48\x52\xbf\xf9\xd2\x0f\xc2\x7b\x6f\x3f\x24\xbc\xfc\x9e\xbd\x70\x5d\xe9\x90\x08\x80\xf5\x7a\x0d\x1c\xe1\x73\x47\xc0\x98\x00\x78\x7c\xb4\xb6\xa2\x49\x82\x16\x19\xeb\x4e\x49\xe5\x8e\x44\xa7\x1f\x03\xd4\x4a\x32\xc7\x00\xa8\x44\x1f\x16\xa8\x6a\x71\x43\xed\x42\x22\xe9\xee\x74\x0e\xd4\x6e\xdb\xea\x05\xe6\x01\x43\xbd\xa4\x9e\x17\x48\xee\x45\x6d\x93\x72\x10\x30\x49\xcc\x25\x84\xc3\xd4\x64\x6c\x30\xdc\x87\x5b\x64\xf3\xb9\x77\xd8\x75\x86\x42\xa2\x10\x9b\x91\xc1\x7d\x40\x77\x78\x5d\x51\x7e\xd1\xae\x60\x3e\x9c\x2f\x94\x2b\x30\x46\x0a\xc9\x10\x93\x01\xdc\xb1\x3b\x90\x5d\x15\x68\x6f\x19\xc5\x75\xd7\x32\xaf\x0b\xe0\xd8\x5d\x9f\xb2\x4c\x48\xf5\x5c\x14\xff\xf0\xf8\xdd\x3b\x61\xe5\xdc\xb6\x84\xa5\xfe\xb7\xb7\xa7\x1a\xa2\x1f\xfb\xb1\x93\xd1\x7e\x1a\x2c\x29\x36\x0f\x28\x1e\x77\x81\xa6\x21\xcd\xbc\xa8\xe3\xc9\x5c\x34\x93\x61\x98\xe8\xf5\x5b\xe4\x9c\x5e\x46\x55\xe0\xf9\xb0\x9d\xc0\x4b\x54\xba\xa3\x7c\x81\x7f\x21\xcc\x52\xf8\xd6\xdb\x64\x5e\x46\x21\xcc\xf4\xe7\xdb\x70\x36\x8b\x8b\xc5\x76\x26\x3d\x89\x16\x9c\xa7\x68\x9e\xa4\x54\xa9\xba\xdf\xe7\xd5\x9f\x00\x00\x00\xff\xff\x0a\x9a\xaa\x65\x86\x06\x00\x00")

func _lgraphqlProjectsProjectbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectbynameGraphql,
		"_lgraphql/projects/projectByName.graphql",
	)
}

func _lgraphqlProjectsProjectbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectbynameextendedGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\x4a\xc5\x30\x10\x45\xf7\xfd\x8a\x11\x5c\xe8\x2f\xb8\x14\x5d\x08\x22\x82\xbe\x0f\x48\xdb\xeb\x73\x24\xc9\xe4\x4d\x26\x0f\x8a\xf4\xdf\xa5\xd3\xfa\xac\x9b\x70\xef\x9d\x03\x39\xa7\x06\x9d\xe8\xa6\x23\xba\xce\x21\xe1\x8e\xde\x4c\x39\x1f\xaf\x6e\xe9\xbb\x23\x2a\x2a\x5f\x18\xec\x7e\x7a\x09\x09\x0b\x44\xb4\x52\x0e\xaf\x0c\x11\x8f\x97\x8b\x87\x23\xdb\x41\xa3\x47\x95\x66\xd0\xd7\x60\x06\xcd\xbe\xf4\x1a\xf2\xf0\x89\xea\xa5\xb4\x18\x15\xa7\x86\x6a\xdb\xa0\xd2\x47\xa4\x7a\x60\xaf\x1f\x61\xb0\xdf\x5c\x54\xc6\x36\x18\x4b\x7e\xcc\x67\x56\xc9\x09\xd9\xfc\x32\xa2\x44\x99\x96\x5a\x1f\xb8\x86\x3e\x62\x15\xc2\x1f\x57\x37\xd5\x8b\xec\x4e\xf7\x1f\xf8\x3e\x95\x75\x9d\xfd\x0d\xcd\xe4\x69\x8c\xd8\xfe\x39\x23\x4a\x59\xb0\x9d\x42\x7d\xe6\xc4\x8b\xc8\xdc\xcd\xdd\x4f\x00\x00\x00\xff\xff\xf2\x22\xc4\x88\x4f\x01\x00\x00")

func _lgraphqlProjectsProjectbynameextendedGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectbynameextendedGraphql,
		"_lgraphql/projects/projectByNameExtended.graphql",
	)
}

func _lgraphqlProjectsProjectbynameextendedGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectbynameextendedGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectByNameExtended.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectbynamemetadataGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x00\x0b\x29\x28\x40\x54\x82\x35\xc0\xd4\x81\x40\x66\x0a\x9c\x09\x92\x82\x73\x72\x53\x4b\x12\x53\x12\x4b\x12\xa1\x02\xb5\x5c\x20\x0c\x08\x00\x00\xff\xff\xd4\x75\x13\x1d\x79\x00\x00\x00")

func _lgraphqlProjectsProjectbynamemetadataGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectbynamemetadataGraphql,
		"_lgraphql/projects/projectByNameMetadata.graphql",
	)
}

func _lgraphqlProjectsProjectbynamemetadataGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectbynamemetadataGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectByNameMetadata.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectgroupsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x00\x0b\x29\x28\x40\x54\x82\x35\xc0\xd4\x81\x40\x66\x0a\x9c\x09\x92\x82\x73\xd2\x8b\xf2\x4b\x0b\x8a\x91\x14\xa2\x28\x45\x53\xac\xa0\x50\x52\x59\x80\xcc\xcd\x2f\x4a\x4f\xcc\xcb\xac\x4a\x2c\xc9\xcc\xcf\x83\x0b\xd7\x72\x21\xe8\x5a\x2e\x40\x00\x00\x00\xff\xff\xef\xa5\xe7\x17\xc5\x00\x00\x00")

func _lgraphqlProjectsProjectgroupsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectgroupsGraphql,
		"_lgraphql/projects/projectGroups.graphql",
	)
}

func _lgraphqlProjectsProjectgroupsGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectgroupsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectGroups.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectsbymetadataGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xcf\x4a\x04\x31\x0c\xc6\xef\xfb\x14\x15\x3c\xb8\xaf\xb0\x47\xc1\x83\xa0\xb0\x20\x3e\x40\x9c\xc6\xdd\x68\x9b\x76\xd3\x74\x60\x58\xe6\xdd\xa5\xd3\xb1\x1d\x91\xcd\x29\xfc\x92\xef\xcb\x9f\x4b\x46\x99\xcc\xc3\xce\x18\x63\xee\xbf\x71\x3a\x98\x37\x15\xe2\xd3\x5d\x25\x23\xb8\x8c\xbf\x6c\x7f\x5d\x60\x94\xf0\x85\x83\xa6\xc7\xe9\x15\x15\x2c\x28\x54\x79\x09\xbf\x92\x83\xb9\x36\x56\x62\x71\x2e\xfe\x7f\xe8\xea\x5e\xa7\xb4\xca\xbc\x64\xfb\x6e\x40\xb6\xa5\x0c\xbe\x37\x42\xd6\xf0\x6c\x5d\x07\x1f\x02\x3c\x9c\x31\x35\x10\xb3\x73\x82\x97\x8c\x49\xd3\xbf\x1d\x7b\x97\x04\x9b\x07\xa5\xc0\x4f\x3c\x92\x04\xf6\xc8\xda\xaa\x21\x22\xa7\x33\x7d\xea\xb1\xde\x7d\x04\x55\x14\x6e\x75\x8b\x23\xba\x10\x8b\x66\x23\x4f\x2f\xe4\xa9\x9b\x9c\x48\xdf\xc5\xdd\xde\xbc\x0d\xd9\xbe\x6d\x73\x77\xfd\xc9\xbc\x9b\x7f\x02\x00\x00\xff\xff\xf0\x75\x45\x7d\xb0\x01\x00\x00")

func _lgraphqlProjectsProjectsbymetadataGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectsbymetadataGraphql,
		"_lgraphql/projects/projectsByMetadata.graphql",
	)
}

func _lgraphqlProjectsProjectsbymetadataGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectsbymetadataGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectsByMetadata.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsProjectsbyorganizationidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\xe0\x52\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\xd4\x54\xa8\xe6\x52\x50\xc8\x2f\x4a\x4f\xcc\xcb\xac\x4a\x2c\xc9\xcc\xcf\x73\xaa\xf4\x4c\x01\x29\x51\x50\x00\xa9\x51\xc9\x4c\x81\x28\x51\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x72\x41\xf2\x50\x46\x5e\x62\x6e\x2a\x94\x89\x6c\x12\x54\x28\xbd\x28\xbf\xb4\xc0\x39\xbf\x34\xaf\x04\x2c\x50\xcb\x05\xc2\xb5\x5c\x80\x00\x00\x00\xff\xff\xf8\x1b\xb1\x64\x8d\x00\x00\x00")

func _lgraphqlProjectsProjectsbyorganizationidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsProjectsbyorganizationidGraphql,
		"_lgraphql/projects/projectsByOrganizationId.graphql",
	)
}

func _lgraphqlProjectsProjectsbyorganizationidGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsProjectsbyorganizationidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/projectsByOrganizationId.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsRemoveprojectfromorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\xf0\xcc\x2b\x51\xd4\x01\x89\xe4\x17\xa5\x27\xe6\x65\x56\x81\xd5\x40\x85\x35\xab\xb9\x14\x14\x8a\x52\x73\xf3\xcb\x52\x03\x20\xea\xdd\x8a\xf2\x73\xfd\x91\x14\x6a\x64\xe6\x15\x94\x96\x58\x29\x80\x14\x2a\x28\xc0\x0d\x85\x19\xaf\x03\x16\x46\x35\x19\xc5\x22\x90\x82\x5a\x4d\xa8\xf6\xcc\x14\x30\x95\x97\x98\x9b\x0a\x12\xe6\xaa\xe5\x02\x04\x00\x00\xff\xff\x0f\x9a\x0e\x27\xb4\x00\x00\x00")

func _lgraphqlProjectsRemoveprojectfromorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsRemoveprojectfromorganizationGraphql,
		"_lgraphql/projects/removeProjectFromOrganization.graphql",
	)
}

func _lgraphqlProjectsRemoveprojectfromorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsRemoveprojectfromorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/removeProjectFromOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsRemoveprojectmetadatabykeyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\xb2\x53\x2b\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x15\x35\xab\xc1\x62\x45\xa9\xb9\xf9\x65\xa9\x01\x45\xf9\x59\xa9\xc9\x25\xbe\xa9\x25\x89\x29\x89\x25\x89\x4e\x95\xde\xa9\x95\x10\x13\x40\x20\x33\xaf\xa0\xb4\xc4\xaa\x1a\xce\x07\x8b\xa5\x58\x81\x4c\x47\x11\x03\x1b\x0f\xb2\x04\x2e\x5a\x0b\x66\x69\x22\xb4\x22\xe9\xc8\x4b\xcc\x4d\x85\x73\x72\xa1\x36\x73\x41\x74\xd5\x02\x02\x00\x00\xff\xff\xfc\x0e\xa7\xd1\xc8\x00\x00\x00")

func _lgraphqlProjectsRemoveprojectmetadatabykeyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsRemoveprojectmetadatabykeyGraphql,
		"_lgraphql/projects/removeProjectMetadataByKey.graphql",
	)
}

func _lgraphqlProjectsRemoveprojectmetadatabykeyGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsRemoveprojectmetadatabykeyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/removeProjectMetadataByKey.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsUpdateprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\x0a\x12\x4b\x92\x33\xac\x14\x42\x0b\x52\x12\x4b\x52\x03\x8a\xf2\xb3\x52\x93\x4b\x02\x40\x62\x9e\x79\x05\xa5\x25\x8a\x9a\xd5\x60\x65\xa5\xc8\xd2\x10\x73\x40\x20\x13\xa4\xc6\x4a\xa1\x1a\x2e\x00\x16\x4c\xb1\x02\x59\x82\x22\x06\xb5\x05\x62\x1b\x5c\xa6\x16\xcc\xd2\x44\x68\x47\xd2\x95\x97\x98\x9b\xca\x05\x51\x54\x0b\x08\x00\x00\xff\xff\x58\x7f\xbf\xd7\xc2\x00\x00\x00")

func _lgraphqlProjectsUpdateprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsUpdateprojectGraphql,
		"_lgraphql/projects/updateProject.graphql",
	)
}

func _lgraphqlProjectsUpdateprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsUpdateprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/updateProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlProjectsUpdateprojectmetadataGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x4d\xca\x02\x31\x0c\x86\xf7\x3d\x45\x3e\xf8\x16\xce\x15\x7a\x03\x17\x82\xe0\x09\xc2\x34\x68\xd5\x66\xca\xf0\x56\x90\x61\xee\x2e\xfd\xd1\xb1\x98\x55\xf3\xa4\xc9\xfb\x84\x04\x86\x9f\x94\x76\x86\x88\xe8\xdf\x3b\x4b\x7b\xc5\x5f\xed\x6e\xf2\xb4\x74\xc2\xec\xf5\xdc\xc8\x83\xef\x49\xde\x6c\x58\x0a\x4c\xd1\x31\xe4\x38\x4f\x57\x19\x71\x10\xb0\x63\x70\xbd\x97\xcb\x6b\x4c\xb0\xcb\xa7\x2f\xcc\xd9\x9c\xd5\xb1\xc8\x18\x2f\x96\xfa\x8f\xb9\x8a\x45\x76\xf9\x99\x34\x9b\x6a\xd5\x4d\x57\xd3\xbf\x86\xed\xec\x57\xac\x72\xd8\xd6\x42\x33\x37\x75\x6b\x7d\x05\x00\x00\xff\xff\xd8\xa2\xd3\x0b\x1b\x01\x00\x00")

func _lgraphqlProjectsUpdateprojectmetadataGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlProjectsUpdateprojectmetadataGraphql,
		"_lgraphql/projects/updateProjectMetadata.graphql",
	)
}

func _lgraphqlProjectsUpdateprojectmetadataGraphql() (*asset, error) {
	bytes, err := _lgraphqlProjectsUpdateprojectmetadataGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/projects/updateProjectMetadata.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsAddorupdateenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4d\x6a\xc4\x30\x0c\x85\xf7\x39\x85\x0a\x5d\xcc\xc0\x9c\x20\xcb\xd2\x81\x76\xd3\x96\xfe\x1c\xc0\xc4\x9a\xd6\x25\x91\x8d\xad\x04\x86\x92\xbb\x97\xc6\xb1\x23\x27\xa1\x9b\x20\xe9\x29\xd2\xd3\xe7\xae\x67\xc5\xc6\x12\x1c\x2a\x80\x5b\x52\x1d\xd6\xf0\xc6\xde\xd0\xe7\xcd\xe9\xaf\xe2\xbc\xfd\xc6\x86\x6b\x78\x24\x8e\x15\x8d\xae\xb5\xd7\xf7\xab\xc3\x1a\xee\x73\x2c\xb5\x3b\x15\xf0\x15\x2f\xe5\xa0\x28\x3d\xa0\xd2\x42\x92\x03\x0d\xb7\x58\xd4\x91\x06\xe3\x2d\x75\x48\x1c\xb7\x9d\x69\x58\x56\x59\x87\x14\xbe\xcc\x85\x5f\xa2\xc3\x27\x69\xfd\x08\x3f\x15\x00\x80\xd2\xfa\xd9\x7f\x38\xad\x18\xcf\xcb\xb4\x83\x21\xd7\x73\x3d\xf7\x00\xc4\xab\xa7\xe3\x4f\x73\x29\x9f\x9d\x00\x24\x41\x5e\x2f\x50\x94\x72\x06\x50\x02\x29\x9b\x32\x8a\x12\xcd\x6a\x51\xa4\x22\x19\xa5\x86\x0d\x9e\x35\xb0\xd4\xb8\x4f\x6a\x17\xe0\xf4\xc7\x78\xcc\x64\x8c\x16\x88\xe6\xd0\xdb\x9e\x8b\x38\x6c\xd0\xec\x3b\xfc\xc7\xce\x2c\xf5\xd3\x4b\xa5\x9d\x8d\x47\x91\x69\x6c\x71\xc9\x02\xfa\xc1\x34\x18\x92\xd1\xc2\xec\xca\xf0\x58\xa5\xef\x58\xfd\x06\x00\x00\xff\xff\x6d\xb2\xb5\x91\xef\x02\x00\x00")

func _lgraphqlEnvironmentsAddorupdateenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsAddorupdateenvironmentGraphql,
		"_lgraphql/environments/addOrUpdateEnvironment.graphql",
	)
}

func _lgraphqlEnvironmentsAddorupdateenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsAddorupdateenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/addOrUpdateEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x3f\x0a\xc2\x30\x14\xc6\xf7\x9e\xe2\x13\x1c\xea\x15\xb2\x8a\x83\x93\x43\xe9\x01\x22\x79\x94\x07\xe6\x25\x24\xaf\x82\x48\xee\x2e\xa9\xc1\xd6\xa1\xeb\xf7\xf7\xe7\x67\xb5\xca\x41\xd0\x77\xc0\x91\xe4\xc9\x29\x88\x27\x51\x83\xab\xe8\xa1\x8a\x91\x52\xe6\xac\x24\x3a\x68\x48\x76\xa2\xf3\xc3\xb2\x37\x18\x34\xb1\x4c\x4b\xe4\xfe\x52\xca\x63\x26\xf7\x6b\x9d\xf0\xee\x00\xc0\x3a\x77\x4b\x63\x74\x56\xe9\xb2\x8e\xb7\xa1\x9e\x25\xce\x6a\x5a\x14\xf8\xbb\xdf\xc2\x34\x7f\x8f\x64\x07\xb1\xb5\x36\x70\x2b\xe8\xe2\x95\x0a\xfa\x3d\xe7\xaa\x94\xae\x7c\x02\x00\x00\xff\xff\x68\x1d\xbe\x72\x10\x01\x00\x00")

func _lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql,
		"_lgraphql/environments/addOrUpdateEnvironmentStorage.graphql",
	)
}

func _lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/addOrUpdateEnvironmentStorage.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsAddrestoreGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x49\x4a\x4c\xce\x2e\x2d\xc8\x4c\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x4c\x49\x09\x4a\x2d\x2e\xc9\x2f\x4a\xd5\xc8\xcc\x2b\x28\x2d\xb1\x52\xa8\x86\xa8\xf5\x4c\xb1\x42\x68\xab\x85\x29\x57\x50\xc8\x4c\x01\x33\x6a\xb9\x40\x18\x10\x00\x00\xff\xff\xdc\x18\x6c\xe7\x65\x00\x00\x00")

func _lgraphqlEnvironmentsAddrestoreGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsAddrestoreGraphql,
		"_lgraphql/environments/addRestore.graphql",
	)
}

func _lgraphqlEnvironmentsAddrestoreGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsAddrestoreGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/addRestore.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x6e\x84\x30\x0c\x85\xf7\x39\xc5\x43\xea\x02\x24\x4e\xc0\xb2\x3b\xa4\xaa\x9b\x9e\x20\x0d\x56\x95\xb6\x24\x8c\xe3\x8c\x84\x10\x77\x1f\xf1\x33\x99\xc0\xcc\xca\xce\x67\xbf\xcf\xb9\x44\xe2\x11\xa5\x02\xde\x9c\xee\xa9\xc1\x97\xb0\x75\x3f\x45\xbd\x90\x81\xfd\x2f\x19\x69\xd0\x3a\x29\x2a\x4c\x0a\x00\xc8\x5d\x2d\x7b\xd7\x93\x93\xf7\xf1\x53\xf7\x54\xae\x18\xd8\xf2\xab\xa6\xde\x51\x12\xdc\x55\xd5\x3e\x98\xf6\x0a\xd8\x2e\xb5\x4b\x32\x3d\xbe\xb5\xf9\x8b\x43\x78\x2c\x1e\x56\x81\xe0\x23\x1b\xca\xc0\x16\x68\xf3\x1d\xc3\xa4\x85\x72\xd2\xd1\x3f\x1d\x09\x53\x10\xcf\x94\xdf\x39\x5d\x02\x82\x68\x89\xe1\x80\x9e\xd5\x2f\x7f\x90\xfc\x1f\xde\x68\xb1\xde\x65\xb3\x59\x9d\xbb\xa5\xce\xea\x16\x00\x00\xff\xff\x0c\x15\x4e\x54\x93\x01\x00\x00")

func _lgraphqlEnvironmentsBackupsforenvironmentbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql,
		"_lgraphql/environments/backupsForEnvironmentByName.graphql",
	)
}

func _lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsBackupsforenvironmentbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/backupsForEnvironmentByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsDeleteenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x08\x2e\xec\x15\xb2\x14\x3c\x81\x27\x08\xf5\x23\x91\x66\xa6\x84\x1f\x11\x4a\xef\x2e\x4d\x68\xa0\xfd\xab\xf0\x1e\xe4\x4d\x2a\x0c\x8c\xa6\x72\x73\x22\x22\x57\x0d\x09\x5e\x9e\xcc\x51\xdf\x97\x86\xe6\x6c\x1f\x8c\x3c\x51\xfc\x30\x16\xc2\xcb\xdd\x6c\x42\x50\x37\x2c\x55\xbc\x30\x81\x78\xe8\x37\x66\xd3\x04\x65\xfb\x78\x5b\xd4\xb9\xd0\xcb\xd2\xc1\xb6\x16\xac\xdd\x03\xef\xd5\xbd\x7f\xb0\xbd\xbe\xdf\xd1\xed\x5a\x5f\x83\x5b\xff\x01\x00\x00\xff\xff\x02\x42\x6d\x97\xda\x00\x00\x00")

func _lgraphqlEnvironmentsDeleteenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsDeleteenvironmentGraphql,
		"_lgraphql/environments/deleteEnvironment.graphql",
	)
}

func _lgraphqlEnvironmentsDeleteenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsDeleteenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/deleteEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsEnvironmentbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4d\x0a\xc2\x30\x10\x85\xf7\x39\xc5\x14\x5c\x28\xf4\x04\x5d\xba\x73\x53\x04\xbd\x40\x69\x9e\x1a\x69\x93\x38\x99\x16\x82\xf4\xee\xd2\x84\x86\xba\x18\xf8\x78\xc3\xfb\xf9\x4c\xe0\x48\x47\x45\x74\xb0\xdd\x88\x86\x6e\xc2\xc6\x3e\xab\x7a\x55\x3c\xbb\x37\x7a\x69\xe8\x62\xa5\x52\xa7\xaf\x22\x82\x9d\x0d\x3b\x3b\xc2\xca\x39\xb6\xdd\x88\xd5\x4a\x94\xbd\x29\xa2\x4e\x42\xb1\x6e\x21\x8a\x28\x05\x10\x19\x5d\x2c\x09\xd8\x4d\xb2\xa3\x90\x50\xc3\x0f\x2e\xde\xa3\xcf\x9f\x5d\x6d\xd1\x9c\x87\x0d\x2f\xf3\x90\x6b\x6e\x68\xb7\xc0\xc9\xeb\x4e\x90\x5b\x7a\x46\x61\x8d\x01\x1b\x07\xf0\x6c\x7a\x84\x3c\x69\x37\xeb\x6f\xda\xa2\xd6\x5b\xd4\x2f\x00\x00\xff\xff\x5b\x2b\x68\x2b\x26\x01\x00\x00")

func _lgraphqlEnvironmentsEnvironmentbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsEnvironmentbynameGraphql,
		"_lgraphql/environments/environmentByName.graphql",
	)
}

func _lgraphqlEnvironmentsEnvironmentbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsEnvironmentbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/environmentByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsEnvironmentbynamespaceGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xc1\x8a\xc3\x20\x10\x86\xef\x3e\xc5\x2c\xec\x61\xf7\x15\x72\xdc\xeb\x42\x28\xb4\x2f\x90\xea\xdf\xd6\x36\x51\x3b\x6a\x40\x4a\xde\xbd\x68\x88\x18\xe8\x41\xfc\x98\x7f\x98\xff\x7b\x46\x70\xa2\x1f\x41\xf4\x6d\x86\x09\xde\x0d\x12\x1d\x1d\x03\x6b\x73\xfd\x12\xbf\x2f\x41\x04\x33\x6b\xb6\x66\x82\x09\x7f\xa9\xaf\x4b\xbb\xf1\x7f\x3c\x83\x0d\x02\x7c\x5d\xc8\x90\xef\x12\x3d\x3e\x87\x5d\x53\x29\x88\x4a\x17\x91\x56\xe5\xcb\x49\x01\xb6\x31\x34\xe4\x0b\x2a\xb8\xd1\xa6\x53\x72\x6b\xd2\xa8\xd4\x99\x75\x30\xfe\xa6\x2f\xe1\xc0\xf6\x0e\x19\xfa\xed\x60\x74\x6a\x08\x58\x5b\x24\xa3\xb2\xc2\x88\x8d\x3d\x78\xd6\x12\x7e\x55\x6a\xb4\x76\x6a\x8b\xc8\x6f\x11\xef\x00\x00\x00\xff\xff\x81\x48\x8a\x83\x44\x01\x00\x00")

func _lgraphqlEnvironmentsEnvironmentbynamespaceGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsEnvironmentbynamespaceGraphql,
		"_lgraphql/environments/environmentByNamespace.graphql",
	)
}

func _lgraphqlEnvironmentsEnvironmentbynamespaceGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsEnvironmentbynamespaceGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/environmentByNamespace.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsEnvironmentsbyprojectnameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x6e\xc3\x20\x10\x45\xf7\x9c\x62\x2a\x75\x91\x5e\x21\xcb\xac\x92\x8d\x65\xa5\xb9\x00\x35\x63\x87\x0a\x33\xee\x30\x44\x42\x95\xef\x5e\x25\x6d\x6d\xc0\x61\x03\x3c\xe6\x7f\xfe\xff\x8a\xc8\x09\x76\x0a\xe0\x75\x62\xfa\xc4\x4e\xf6\xf0\x2e\x6c\xfd\xf0\xf2\x06\xdf\x0a\x00\xe0\x8f\x1f\x52\xa3\x47\xdc\x3d\x10\x80\xd7\x23\xee\x17\xcd\xff\x28\x00\xfa\x9b\x65\xf2\x23\x7a\x09\x0b\x04\xb0\x66\x39\xde\x95\xcb\xc5\xe0\xe4\x28\x5d\xd2\x54\xa3\x83\x0e\x78\xc6\xbe\xa2\x47\xd4\x66\x4b\x2f\x56\xdc\xaa\xcf\x12\x14\xbe\x34\xa1\x0f\x57\xdb\x4b\xfb\x9b\xb9\xc9\x73\xe8\x28\x74\x32\x99\x0b\x53\x94\x27\xda\xac\x52\x55\x6b\x53\xed\xbe\x3a\x46\x2d\x58\x0e\x75\x8e\xa2\x39\xe3\x60\xc9\x6f\x79\xcb\x74\xb3\x06\xb9\x78\xe9\xd9\xa2\x37\x2e\x35\xb5\x7d\x08\xd7\x23\x05\xa9\x59\x4b\x5c\xb2\x47\x19\x6e\xb5\x08\x72\xf9\xe9\x47\xb4\xce\x9c\x46\x3d\xac\xc6\xb3\xca\xf7\x59\xcd\xea\x27\x00\x00\xff\xff\xa3\x30\xe3\xff\x24\x02\x00\x00")

func _lgraphqlEnvironmentsEnvironmentsbyprojectnameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsEnvironmentsbyprojectnameGraphql,
		"_lgraphql/environments/environmentsByProjectName.graphql",
	)
}

func _lgraphqlEnvironmentsEnvironmentsbyprojectnameGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsEnvironmentsbyprojectnameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/environmentsByProjectName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsSetenvironmentservicesGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\x84\xf0\x8a\x53\x8b\xca\x32\x93\x53\x8b\xad\x14\xa2\x83\x4b\x8a\x32\xf3\xd2\x63\x15\x35\xab\xc1\x52\xc5\xa9\x25\xae\x79\x65\x99\x45\xf9\x79\xb9\xa9\x79\x25\xc1\x50\x75\x10\x43\x40\x20\x33\xaf\xa0\xb4\xc4\x4a\xa1\x1a\x2e\x00\x02\xa9\x08\x1d\x56\x20\xab\x50\x24\x11\x76\xc1\xad\x85\xcb\xd7\x82\x59\x9a\x08\xd3\x90\xf4\xe6\x25\xe6\xa6\x72\x41\x14\xd5\x02\x02\x00\x00\xff\xff\xfc\x33\x0b\xcb\xce\x00\x00\x00")

func _lgraphqlEnvironmentsSetenvironmentservicesGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsSetenvironmentservicesGraphql,
		"_lgraphql/environments/setEnvironmentServices.graphql",
	)
}

func _lgraphqlEnvironmentsSetenvironmentservicesGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsSetenvironmentservicesGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/setEnvironmentServices.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsSshendpointbyenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4b\xcc\x4d\x2d\x2e\x48\x4c\x4e\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\xd2\xac\x06\x4b\xa5\xe6\x95\x65\x16\xe5\xe7\xe5\xa6\xe6\x95\x38\x55\xfa\xc1\x95\xa1\x08\x7b\x97\x26\xa5\x16\xe5\xa5\x96\xa4\x16\xc3\x15\x80\x18\x10\xb3\x41\x20\x1b\xbb\x02\x2b\x24\x8b\xc1\x6a\xa1\x76\x82\x40\x66\x0a\x9c\x09\x52\x03\xe7\xe4\x17\xa4\xe6\x15\x67\x64\xa6\x95\x04\x14\xe5\x67\xa5\x26\x97\xf8\x61\x95\x54\x40\x18\x84\x66\x18\x08\x14\x17\x67\x78\xe4\x17\x97\xa0\x8b\x05\xe4\x17\x21\xc4\x6a\xb9\x20\x64\x2d\x17\x20\x00\x00\xff\xff\x7c\x08\x82\x7b\x29\x01\x00\x00")

func _lgraphqlEnvironmentsSshendpointbyenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsSshendpointbyenvironmentGraphql,
		"_lgraphql/environments/sshEndpointByEnvironment.graphql",
	)
}

func _lgraphqlEnvironmentsSshendpointbyenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsSshendpointbyenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/sshEndpointByEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlEnvironmentsUpdateenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x41\x6e\x86\x20\x10\x85\xf7\x9e\x62\x4c\xba\xa8\x57\x60\xdf\x85\x9b\xc6\x45\x7b\x00\x02\xd3\x48\xa3\x03\x81\xc1\xc4\x18\xef\xfe\x07\xf9\x03\x22\x2b\xbe\xc7\x1b\xde\xbc\x35\xb2\x64\x63\x09\x3e\x3b\x00\x80\x0f\xa3\x05\x8c\xc4\x7d\x26\x27\x59\xcd\x02\x7e\x9d\x96\x8c\x5f\xb4\x19\x6f\x69\x45\xe2\x29\xe9\x23\xb9\xc8\xfd\x70\x5c\xd6\xf8\xb4\xe4\xff\xd2\x31\xc9\x27\xe0\x28\xc2\x25\x6a\x91\xc2\x1a\xed\x9d\x96\x53\xcb\xcb\x79\xdd\x86\x3a\x7e\x9b\x22\xb9\x62\x01\x6f\x23\x3f\x28\x14\xd4\xe8\x16\xbb\xff\xec\xae\x3a\xb0\x2e\xdb\xe8\xd6\x21\x85\xd9\xfc\xf1\xe4\xed\x3f\x2a\xfe\xbe\x87\xe4\x9e\x75\x03\xe5\xb1\x61\x8d\x0b\xde\x39\xa0\xdf\x8c\xc2\xf0\x6c\xdf\x60\x53\x23\xf7\x3d\xbb\xf3\x15\x00\x00\xff\xff\x5a\xa6\xf7\x47\x9b\x01\x00\x00")

func _lgraphqlEnvironmentsUpdateenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlEnvironmentsUpdateenvironmentGraphql,
		"_lgraphql/environments/updateEnvironment.graphql",
	)
}

func _lgraphqlEnvironmentsUpdateenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlEnvironmentsUpdateenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/environments/updateEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlTasksSwitchactivestandbyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\x52\x50\xd0\x54\xa8\xe6\x52\x50\x28\x2e\xcf\x2c\x49\xce\x70\x4c\x2e\xc9\x2c\x4b\x0d\x2e\x49\xcc\x4b\x49\xaa\x04\x69\x50\x50\xc8\xcc\x2b\x28\x2d\xb1\x02\xab\x01\x01\x98\x01\x30\xbe\x82\x42\x5e\x62\x6e\xaa\x15\xdc\x64\xa8\x70\x2d\x17\x8c\xd4\x84\x6a\xcd\x4c\xe1\x02\x09\xd4\x02\x02\x00\x00\xff\xff\x76\xc4\x5e\xef\x90\x00\x00\x00")

func _lgraphqlTasksSwitchactivestandbyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlTasksSwitchactivestandbyGraphql,
		"_lgraphql/tasks/switchActiveStandby.graphql",
	)
}

func _lgraphqlTasksSwitchactivestandbyGraphql() (*asset, error) {
	bytes, err := _lgraphqlTasksSwitchactivestandbyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/tasks/switchActiveStandby.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlTasksUpdatetaskGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x4d\x0e\x02\x21\x0c\x85\xf7\x73\x8a\x4e\xe2\xc2\xb9\x02\x37\x60\xe7\x42\x0f\xd0\x40\x13\x89\xf2\x13\x28\xab\x09\x77\x37\x80\xa9\xc8\x8a\xef\x7b\xed\xab\xaf\x8c\xec\x62\x80\xeb\x06\x00\x70\x71\x56\x81\x0e\xbc\x4f\x4a\xc8\xe6\xa9\xe0\x91\x2c\x32\xdd\xb1\xbc\x6e\x5d\xe8\x90\x2a\xef\xc7\x39\x66\xaa\x64\xb3\xa1\x3f\xd7\x07\x14\x9c\x22\x86\xb4\xaa\xd7\xff\xb9\x6f\xff\xbc\x23\x49\x1b\xbf\xe3\xb7\xbe\x6c\x05\xf4\x24\x50\x18\xb9\x16\x41\x93\x09\x99\xec\x1a\xe7\x95\x4d\xf4\xe9\x4d\xab\xc9\xe4\x23\x93\x9e\xa2\x6d\xed\x13\x00\x00\xff\xff\x25\xa0\x63\x9d\x0e\x01\x00\x00")

func _lgraphqlTasksUpdatetaskGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlTasksUpdatetaskGraphql,
		"_lgraphql/tasks/updateTask.graphql",
	)
}

func _lgraphqlTasksUpdatetaskGraphql() (*asset, error) {
	bytes, err := _lgraphqlTasksUpdatetaskGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/tasks/updateTask.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlTasksUploadfilesfortaskGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\x8c\x41\x0a\x83\x40\x0c\x45\xaf\xf2\x07\x5c\x28\x78\x82\x39\x80\xd0\x7d\xbb\x2a\x5d\x04\xb4\x30\xa8\x19\x69\x92\x55\xf0\xee\x65\x32\x64\x11\x92\xc7\x7b\xa7\x29\x69\xa9\x8c\x11\x83\x92\xec\x19\x0f\xd6\x34\x63\xf8\x96\x63\x93\x8c\xf7\xeb\x3a\x2a\xad\xe9\x93\x26\x38\x2c\x8e\xa5\xa1\xa5\xfe\x9e\x24\xfb\x58\xf8\x32\xcd\x1e\x6e\x14\x66\x74\xb5\x17\xee\x09\x5e\x56\x30\x9d\x1b\x44\x49\x4d\x3a\x86\xc7\x8e\xff\xdd\xe6\x1f\x00\x00\xff\xff\x4f\x59\x28\x5d\x89\x00\x00\x00")

func _lgraphqlTasksUploadfilesfortaskGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlTasksUploadfilesfortaskGraphql,
		"_lgraphql/tasks/uploadFilesForTask.graphql",
	)
}

func _lgraphqlTasksUploadfilesfortaskGraphql() (*asset, error) {
	bytes, err := _lgraphqlTasksUploadfilesfortaskGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/tasks/uploadFilesForTask.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddgroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x4c\x49\x71\x2f\xca\x2f\x2d\xd0\xc8\xcc\x2b\x28\x2d\xb1\x82\x8a\x2a\x28\x40\x14\x83\xf5\x80\x45\x6a\x35\xe1\x52\x99\x29\x48\x6a\x20\x92\x5c\x20\x0c\x08\x00\x00\xff\xff\x33\xcc\xe8\xad\x6e\x00\x00\x00")

func _lgraphqlUsergroupsAddgroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddgroupGraphql,
		"_lgraphql/usergroups/addGroup.graphql",
	)
}

func _lgraphqlUsergroupsAddgroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddgroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddgrouptoorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\xe4\x17\xa5\x27\xe6\x65\x56\x81\xd5\x58\x29\x78\xe6\x95\x28\xea\x68\x56\x73\x29\x28\x24\xa6\xa4\xb8\x17\xe5\x97\x16\x84\xe4\xfb\x23\xa9\xd0\xc8\xcc\x2b\x28\x2d\xb1\x52\x00\xa9\x50\x50\x80\x98\x06\x36\x54\x07\x2c\x80\x6a\x18\x8a\xd9\x20\x05\xb5\x9a\x50\x8d\x99\x29\x70\xfd\x20\x61\xae\x5a\x2e\x40\x00\x00\x00\xff\xff\x92\x85\xef\xa4\xa7\x00\x00\x00")

func _lgraphqlUsergroupsAddgrouptoorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddgrouptoorganizationGraphql,
		"_lgraphql/usergroups/addGroupToOrganization.graphql",
	)
}

func _lgraphqlUsergroupsAddgrouptoorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddgrouptoorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addGroupToOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddgroupstoprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x80\x30\x3c\xf3\x0a\x4a\x4b\x14\x75\x40\x52\xe9\x45\xf9\xa5\x05\xc5\x56\x0a\xd1\xee\x20\x06\x44\x22\x56\x51\x53\xa1\x9a\x4b\x41\x41\x41\x21\x31\x25\x05\x2c\x5e\x1c\x92\x0f\xd5\xaa\x91\x09\x52\x62\x05\x95\x57\x50\x80\x1b\x0d\xb3\x04\x2a\x0e\x33\x17\x6a\x01\x58\xb4\x56\x13\xae\x2d\x33\x05\xca\xc8\x4b\xcc\x4d\x85\x48\x72\x81\x30\x20\x00\x00\xff\xff\x8b\x70\xd6\xd3\xb8\x00\x00\x00")

func _lgraphqlUsergroupsAddgroupstoprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddgroupstoprojectGraphql,
		"_lgraphql/usergroups/addGroupsToProject.graphql",
	)
}

func _lgraphqlUsergroupsAddgroupstoprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddgroupstoprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addGroupsToProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddsshkeyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xc1\x0a\xc2\x30\x10\x44\xef\xf9\x8a\x11\x3c\xb4\xe0\x17\xf4\xee\xc9\xa3\xe2\x7d\x21\x41\x17\x9b\x58\xda\xe4\x50\x24\xff\x2e\xe9\xa6\x59\xc1\x53\xc2\x9b\x61\xdf\xf8\x14\x29\xf2\x3b\xa0\x33\xc0\x31\x90\x77\x03\xae\x71\xe6\xf0\x38\x9c\x0a\x79\xb9\xf5\x4e\x63\xfa\xa7\xb7\x75\x2a\x70\x79\x5e\xe4\x2f\x41\x5a\xdc\x7c\xf6\xc4\x63\xeb\xf7\xf8\x18\x00\x20\x6b\xa5\xdc\x71\x98\x52\x1c\x2a\x06\xc4\xb9\xa9\x2b\x51\x67\xd3\x6b\x22\xde\x7d\x41\xe5\x45\xab\x17\x01\x27\x13\x74\x4e\x4d\xf2\xf6\xe6\xbe\x55\xd9\xfe\xac\x30\x7b\x25\x9b\x6f\x00\x00\x00\xff\xff\xbf\xff\xb7\x89\x17\x01\x00\x00")

func _lgraphqlUsergroupsAddsshkeyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddsshkeyGraphql,
		"_lgraphql/usergroups/addSshKey.graphql",
	)
}

func _lgraphqlUsergroupsAddsshkeyGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddsshkeyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addSshKey.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAdduserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x0e\xc2\x30\x0c\x85\xe1\x3d\xa7\x78\x48\x0c\xad\xc4\x09\x7a\x08\x16\xc4\x01\x2c\x12\x90\xa5\xda\x45\xa9\x3b\x21\xee\x8e\xda\x3a\x69\x06\x06\x2f\xff\x1b\x3e\xcb\x62\x64\x3c\x29\xba\x00\x9c\x93\x10\x8f\x03\x6e\x96\x59\x5f\xa7\xcb\x9a\x9e\x9c\x67\xbb\x92\xa4\x92\xb7\x3a\xd2\x9f\xf8\x98\x44\x92\x5a\x69\x3d\x3e\x01\x00\x28\xc6\xfb\x9c\x72\xc7\xfa\x5e\x6c\xf0\x08\xb8\xb5\x9b\xde\x1a\xec\x80\x7d\x3b\xc8\xaa\xfb\x52\xdd\xf2\xc1\xd6\xbf\x7d\xa5\x38\xb6\xe6\xbe\x86\xf5\x7e\x01\x00\x00\xff\xff\xf7\xf8\x9f\x0d\xfe\x00\x00\x00")

func _lgraphqlUsergroupsAdduserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAdduserGraphql,
		"_lgraphql/usergroups/addUser.graphql",
	)
}

func _lgraphqlUsergroupsAdduserGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAdduserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addUser.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddusertogroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xbd\x0a\xc3\x30\x0c\x84\xf7\x3c\xc5\x05\x3a\x24\xd0\x27\xf0\x5e\xba\x75\xe8\xcf\x03\x18\x6c\x82\x20\xb6\x83\x62\x4f\x21\xef\x5e\xa4\xb8\xf1\xd2\x41\x20\xee\x8e\xfb\x2e\x94\x6c\x33\xa5\x88\xa1\x03\x2e\x65\xf5\x7c\x0b\x96\x66\x83\x57\x66\x8a\x53\x7f\x15\x79\xe2\x54\x96\x87\x0d\xfe\x8f\xfc\x4c\xb3\x37\xb8\xff\xde\x7e\xc4\xd6\x01\x80\x75\xee\xb3\x7a\x7e\x27\xb5\x06\x8a\x4b\xc9\xa6\x7a\x80\x80\x0c\x36\xf8\x03\xd6\xc0\xd8\x6b\x42\xcb\x25\x12\x95\xdb\x36\x9c\x09\x56\x72\x5b\xa1\xf2\x3e\x9e\x0c\x72\xf5\x91\x86\xc3\xec\xe4\xbe\x01\x00\x00\xff\xff\x00\xb0\x61\x57\xf3\x00\x00\x00")

func _lgraphqlUsergroupsAddusertogroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddusertogroupGraphql,
		"_lgraphql/usergroups/addUserToGroup.graphql",
	)
}

func _lgraphqlUsergroupsAddusertogroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddusertogroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addUserToGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAddusertoorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x2d\x4e\x2d\xb2\x52\x08\x2d\x4e\x2d\xf2\xcc\x2b\x28\x2d\x51\xd4\x01\x09\xe6\x17\xa5\x27\xe6\x65\x56\x81\x95\x59\x29\x78\xe6\xc1\x84\xcb\xf3\x40\x8a\x9d\xf2\xf3\x73\x52\x13\xf3\x34\x15\xaa\xb9\x14\x14\x14\x14\x12\x53\x52\x40\xda\x43\xf2\xfd\x91\x74\x69\x64\x82\x4c\xb3\x82\x2a\x51\x50\x80\x58\x03\xb6\x0d\x2a\x82\x6a\x07\x8a\x95\x30\x15\x10\xeb\x20\xd6\x82\xc5\x6a\x35\xe1\x06\x66\xa6\x40\x19\x79\x89\xb9\xa9\x10\x49\x2e\x10\x06\x04\x00\x00\xff\xff\x31\x25\xe0\x75\xde\x00\x00\x00")

func _lgraphqlUsergroupsAddusertoorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAddusertoorganizationGraphql,
		"_lgraphql/usergroups/addUserToOrganization.graphql",
	)
}

func _lgraphqlUsergroupsAddusertoorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAddusertoorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/addUserToOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAllgroupmembersGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8c\x41\xae\xc2\x30\x0c\x44\xf7\x39\x85\x17\x7f\xf1\xb9\x42\x2f\x80\x58\xc0\x86\x13\xa4\xd4\x54\x96\xec\x04\x9c\x64\x81\x50\xee\x8e\x42\xd4\xb4\x51\xc5\xac\xec\x79\x33\xf3\x4c\xa8\x2f\xb0\xcc\x47\xf5\xe9\x11\xe0\xff\xcf\x59\xc1\x01\xae\x51\xc9\xcd\x07\x78\x1b\x00\x68\xfc\x8c\x32\xa2\x86\x61\x5b\xa8\xf9\x6f\x6d\x89\x17\xd1\xd4\xce\x82\xda\x23\x75\x62\x0d\x16\xa5\x80\xda\x3b\x45\x28\x96\x78\xe7\xde\x49\x43\xbc\x6c\x27\x17\xb1\xfd\x01\x6e\x5e\x04\x5d\xdc\xf9\x33\x45\xb6\xe3\x69\xea\x40\xee\x3e\xf5\xbc\xee\x55\x94\x4d\x36\x9f\x00\x00\x00\xff\xff\x82\x01\xf0\x6f\x37\x01\x00\x00")

func _lgraphqlUsergroupsAllgroupmembersGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAllgroupmembersGraphql,
		"_lgraphql/usergroups/allGroupMembers.graphql",
	)
}

func _lgraphqlUsergroupsAllgroupmembersGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAllgroupmembersGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/allGroupMembers.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAllgroupmemberswithkeysGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xb1\x0e\xc2\x30\x0c\x44\xf7\x7c\x85\x07\x06\xf8\x05\x7e\x00\x21\x24\x16\x10\xbb\x2b\x4c\x89\x70\xda\x62\x37\x43\x84\xf2\xef\x28\x89\xda\xb4\xbd\xc5\x97\x7b\xa7\x53\xbe\x9e\x24\x00\x32\x9f\xa4\xf7\x83\xc2\x7e\xd7\xa1\xa3\x23\xdc\x46\xb1\x5d\x7b\x80\x9f\x81\x25\x2d\x30\x77\x0a\x03\x48\x3e\x1b\xfb\xcc\xc7\x91\x6b\x48\xb4\x40\x00\xaf\x24\x93\x9f\x3b\x49\xe4\xd0\xf2\xfc\x7a\x59\xd1\xf1\x3a\x4d\x25\x31\x6e\x82\xd6\x8e\x8c\xcd\xb9\x2e\xa8\xbe\x2f\x14\xb4\xae\xaf\xf6\x17\x3f\x2b\xfa\x50\xb8\x87\x61\x93\x3c\x90\x7d\x8d\xa2\x59\x5f\xe9\xb9\xc0\x14\x44\x13\xcd\x3f\x00\x00\xff\xff\xd9\x20\x19\x28\x2f\x01\x00\x00")

func _lgraphqlUsergroupsAllgroupmemberswithkeysGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAllgroupmemberswithkeysGraphql,
		"_lgraphql/usergroups/allGroupMembersWithKeys.graphql",
	)
}

func _lgraphqlUsergroupsAllgroupmemberswithkeysGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAllgroupmemberswithkeysGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/allGroupMembersWithKeys.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsAllusersGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\xb1\x0d\x02\x31\x0c\x85\xe1\x3e\x53\xb8\xb8\x02\x56\xb8\x0d\xae\xa1\x41\x0c\x60\x88\x39\x59\x72\x82\x70\x4c\x81\xd0\xed\x8e\x92\x80\x51\x48\xf7\xfe\xc8\xdf\xfd\x41\xfa\x04\x14\x39\x15\xd2\xb2\x0b\x00\x00\x13\x25\x64\x99\xe1\x68\xca\x79\xed\x89\xe3\xb8\x57\x36\xc1\xf3\x12\x67\x58\xb2\x85\x3d\xbc\x5a\x1e\x9d\xfa\x3e\x54\x27\xbd\x56\x6d\xe2\xe8\xfb\xa7\x39\xdc\xfe\xbe\x6e\xbf\x19\x51\x5f\x57\xd6\x62\x07\x4c\xe4\x45\xf0\x2f\x5c\x6e\x29\x51\xb6\xb6\xb7\xb0\xbd\x03\x00\x00\xff\xff\xe8\xc3\x51\xc9\xf5\x00\x00\x00")

func _lgraphqlUsergroupsAllusersGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsAllusersGraphql,
		"_lgraphql/usergroups/allUsers.graphql",
	)
}

func _lgraphqlUsergroupsAllusersGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsAllusersGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/allUsers.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsGetgroupmembersGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8c\x31\x0e\xc2\x30\x0c\x45\xf7\x9c\xc2\x48\x0c\x70\x05\x46\x36\x96\x2e\x9c\xc0\xa5\x26\xb2\x14\x27\xe0\x24\x43\x85\x72\x77\x14\xaa\xb6\x44\x55\xff\x64\xff\xf7\xf4\xdf\x99\x74\x04\xab\x21\xbf\xae\x63\x87\x42\x70\x3a\x7a\x14\xba\xc0\x3d\x29\x7b\x7b\x38\xc3\xc7\x00\x40\xab\x4c\xc6\x4f\x9c\x79\x0d\x0f\xcb\x59\xd1\xf2\x08\x49\x4f\x1a\x57\xb1\x26\x47\xd2\xb6\xa9\x21\x41\x76\x9b\xf6\xc9\x1a\x53\xf7\x3f\x39\xc7\xe1\x0e\x78\x04\x11\xf2\x69\xd3\x5b\x4e\x0e\xfb\xdb\xd0\x80\xd2\x7c\x1a\xdc\xba\x37\xa1\x62\x8a\xf9\x06\x00\x00\xff\xff\x00\x11\x27\x4c\x2b\x01\x00\x00")

func _lgraphqlUsergroupsGetgroupmembersGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsGetgroupmembersGraphql,
		"_lgraphql/usergroups/getGroupMembers.graphql",
	)
}

func _lgraphqlUsergroupsGetgroupmembersGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsGetgroupmembersGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/getGroupMembers.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsGroupsbyorganizationidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\xe0\x52\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\xd4\x54\xa8\xe6\x52\x50\xc8\x2f\x4a\x4f\xcc\xcb\xac\x4a\x2c\xc9\xcc\xcf\x73\xaa\xf4\x4c\x01\x29\x51\x50\x00\xa9\x51\xc9\x4c\x81\x28\x51\x50\x48\x2f\xca\x2f\x2d\x28\x86\x72\x40\xb2\x50\x46\x5e\x62\x6e\x2a\x94\x59\x52\x59\x00\x63\xe6\xa6\xe6\x26\xa5\x16\x39\xe7\x97\xe6\x95\x40\x45\x90\x2d\x01\x0b\xd5\x72\x81\x70\x2d\x17\x20\x00\x00\xff\xff\xed\x26\x0a\x8d\x97\x00\x00\x00")

func _lgraphqlUsergroupsGroupsbyorganizationidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsGroupsbyorganizationidGraphql,
		"_lgraphql/usergroups/groupsByOrganizationId.graphql",
	)
}

func _lgraphqlUsergroupsGroupsbyorganizationidGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsGroupsbyorganizationidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/groupsByOrganizationId.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsMeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xb1\xae\x82\x21\x14\x83\xf7\xff\x29\x78\x9a\xbb\xdc\xc4\xc1\x18\x77\xa2\xf5\x97\x78\x0e\xe0\x01\x06\x62\x78\x77\x83\x31\x47\x02\x9d\xda\x6f\x68\xfb\x2c\x90\x6a\x5e\x9b\x31\xc6\x30\xbe\xa6\xcb\x5d\xd5\x82\xad\x23\x4d\x37\x27\x29\x1f\x2c\x43\x09\xd9\x09\x5c\x02\x33\x7c\xd6\x9c\xd2\xfd\x1f\x35\x0d\xed\xd3\x42\x97\x1f\x1b\xba\x1e\xa8\xa7\x1a\x17\x76\xb6\x54\x16\xf8\xe7\xfc\x0e\x89\xe2\x86\xd5\xcf\x13\x81\xcd\xf8\x0d\x35\x75\xbb\x84\x12\x8f\x81\x30\xdf\x5a\x7e\x48\x20\x4c\x05\x6d\x6b\xef\x00\x00\x00\xff\xff\x80\x58\xa8\xdd\x39\x01\x00\x00")

func _lgraphqlUsergroupsMeGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsMeGraphql,
		"_lgraphql/usergroups/me.graphql",
	)
}

func _lgraphqlUsergroupsMeGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsMeGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/me.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsRemovegroupsfromprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x29\x28\xca\xcf\x4a\x4d\x2e\xb1\x52\x08\x80\x30\x3c\xf3\x0a\x4a\x4b\x14\x75\x40\x52\xe9\x45\xf9\xa5\x05\xc5\x56\x0a\xd1\xee\x20\x06\x44\x22\x56\x51\x53\xa1\x9a\x4b\x41\x41\x41\xa1\x28\x35\x37\xbf\x2c\x15\x2c\x55\xec\x56\x94\x9f\x0b\xd5\xaf\x91\x09\x52\x67\x05\x55\xa4\xa0\x00\x37\x1f\x66\x13\x54\x1c\x66\x38\xd4\x16\xb0\x68\xad\x26\x5c\x5b\x66\x0a\x94\x91\x97\x98\x9b\x0a\x91\xe4\x02\x61\x40\x00\x00\x00\xff\xff\x38\x23\xd0\x35\xbd\x00\x00\x00")

func _lgraphqlUsergroupsRemovegroupsfromprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsRemovegroupsfromprojectGraphql,
		"_lgraphql/usergroups/removeGroupsFromProject.graphql",
	)
}

func _lgraphqlUsergroupsRemovegroupsfromprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsRemovegroupsfromprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/removeGroupsFromProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsRemovesshkeybyidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\xd4\xd1\x54\xa8\xe6\x52\x50\x48\x49\xcd\x49\x2d\x49\x0d\x2e\xce\xf0\x4e\xad\x74\xaa\xf4\x4c\xd1\xc8\xcc\x2b\x28\x2d\xb1\x02\xcb\x29\x28\x80\x14\xab\x64\xa6\x70\x29\x28\xd4\x6a\x72\xd5\x72\x01\x02\x00\x00\xff\xff\x75\x2f\xbe\x0a\x4a\x00\x00\x00")

func _lgraphqlUsergroupsRemovesshkeybyidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsRemovesshkeybyidGraphql,
		"_lgraphql/usergroups/removeSSHKeyById.graphql",
	)
}

func _lgraphqlUsergroupsRemovesshkeybyidGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsRemovesshkeybyidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/removeSSHKeyById.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsRemoveuserfromgroupGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x31\x0a\xc3\x30\x0c\x45\x77\x9f\xe2\x17\x3a\x24\xd0\x13\x78\x6f\xbb\x75\x29\x3d\x80\xa1\x22\x08\x2a\x3b\x28\x76\x97\x90\xbb\x07\x25\x22\x19\x0c\xe6\xeb\xf1\x9e\xb4\x9a\x2a\x97\x8c\x2e\x00\xd7\x36\x91\xde\x25\xf1\x2f\xe2\x5d\x95\xf3\x70\xb9\xd9\x3c\x68\x69\xe3\x2b\x09\x1d\x73\x8f\x39\x00\x80\x92\x94\x3f\x7d\x26\xd2\x87\x16\x79\x1a\xd7\x71\x1e\x5b\x8d\x0e\x00\xe6\x8c\x98\x41\xbb\xf7\x6c\x60\x71\x62\xd3\x1b\x92\xb7\xc4\x99\x73\x62\xe9\x0f\x19\x7f\xfd\x63\xe8\x7e\x0c\xf6\xd6\x00\x00\x00\xff\xff\x8a\xee\x6f\x4c\xc7\x00\x00\x00")

func _lgraphqlUsergroupsRemoveuserfromgroupGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsRemoveuserfromgroupGraphql,
		"_lgraphql/usergroups/removeUserFromGroup.graphql",
	)
}

func _lgraphqlUsergroupsRemoveuserfromgroupGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsRemoveuserfromgroupGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/removeUserFromGroup.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsRemoveuserfromorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\x31\x0a\x02\x31\x10\x45\xfb\x9c\xe2\x0b\x16\xbb\xe0\x09\x52\x5a\x08\x5b\x59\x79\x80\x80\x83\x04\xcc\x8c\xcc\x26\x0a\xca\xde\x5d\x66\xa2\xb0\x5b\x85\xfc\x79\xff\xbf\xd2\x6a\xaa\x59\x18\x43\x00\xf6\x6d\x26\x8d\xb8\xcc\xa4\x13\x3f\x5a\xdd\x1d\x2c\x14\xbd\x25\xce\x6f\xc7\x22\x26\xfe\xc7\x2f\x36\xf8\x28\x72\xa7\xc4\x23\x3e\x01\x50\x2a\xf2\x24\xeb\x9f\x54\xca\x79\x55\x1c\xb2\x0d\x46\xa7\x80\xee\x71\x9d\xff\xb7\x8a\x8d\xb1\xdf\xbb\xab\x3b\x03\xb0\x8c\xbf\xa1\x7c\xf5\x87\x53\x21\x8b\xc3\x12\xbe\x01\x00\x00\xff\xff\x7e\xd2\x27\x2a\xd1\x00\x00\x00")

func _lgraphqlUsergroupsRemoveuserfromorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsRemoveuserfromorganizationGraphql,
		"_lgraphql/usergroups/removeUserFromOrganization.graphql",
	)
}

func _lgraphqlUsergroupsRemoveuserfromorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsRemoveuserfromorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/removeUserFromOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsUserbyemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x4e\xbd\xce\xc2\x30\x10\xdb\xfb\x14\xf7\x49\xdf\x50\x5e\x81\x11\x09\x16\x24\x06\x40\xec\x11\x98\x72\x22\x3f\xe5\x92\x0c\x11\xea\xbb\xa3\x14\x08\x55\xea\xe9\xce\xb6\x6c\x3f\x22\x24\x51\xfb\x0f\xa3\x58\x2f\xe9\x10\x84\x6d\xf7\xb7\xa0\x67\x43\x44\x14\x3d\x64\x95\xd6\x59\x6b\x3f\x8e\xb7\xf3\x6b\xc8\xe0\x4b\x39\x47\xad\x7c\x57\x16\x1f\x76\xca\xa0\x30\x5a\x55\xc4\xd9\x19\x03\x1b\xca\xef\xfd\x6d\x8b\xe4\x27\xe9\x55\x43\x86\x9d\x26\x64\xdc\x91\x8e\xa9\x9f\x71\x27\xa5\xe3\x8c\xdc\xb0\xed\x20\xbd\xf0\xa4\x75\x5c\x22\x50\x01\xbf\xa2\xa1\x5c\x9d\xb8\xd8\xef\x9d\x46\x3d\x6b\xb6\x43\x9c\x46\x15\x30\x34\xc3\x2b\x00\x00\xff\xff\xa1\x66\xca\x90\x63\x01\x00\x00")

func _lgraphqlUsergroupsUserbyemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsUserbyemailGraphql,
		"_lgraphql/usergroups/userByEmail.graphql",
	)
}

func _lgraphqlUsergroupsUserbyemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsUserbyemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/userByEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsUserbyemailsshkeysGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\x49\xcd\x4d\xcc\xcc\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x28\x2d\x4e\x2d\x72\xaa\x74\x05\xc9\x68\x40\xe5\x21\xea\x20\xd2\x0a\x0a\x60\x0e\x98\x55\x5c\x9c\xe1\x9d\x5a\x59\x0c\x15\x57\x50\xc8\x4c\x81\x32\xf2\x12\x73\x53\xa1\xcc\xec\xd4\xca\x90\xca\x02\x24\x5e\x58\x62\x4e\x29\x84\x5b\xcb\x05\xc2\xb5\x5c\x80\x00\x00\x00\xff\xff\x6c\x86\xee\xf9\x8e\x00\x00\x00")

func _lgraphqlUsergroupsUserbyemailsshkeysGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsUserbyemailsshkeysGraphql,
		"_lgraphql/usergroups/userByEmailSSHKeys.graphql",
	)
}

func _lgraphqlUsergroupsUserbyemailsshkeysGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsUserbyemailsshkeysGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/userByEmailSSHKeys.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsUserbysshfingerprintGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xbf\xce\xc2\x30\x0c\xc4\xf7\x3e\x85\x3f\xe9\x1b\xca\x2b\x30\x32\xb0\x20\x31\x50\xc4\x1e\x95\x6b\x1b\x91\x3f\xc5\x49\x86\x08\xf5\xdd\x51\x07\xd2\x34\xb9\xc9\xfe\x25\x77\xb6\xdf\x01\x1c\xa9\xfd\x1f\xa4\x19\xc1\x33\x4b\xe3\x8f\xd4\x79\x96\x66\xfc\x3b\xd0\xa7\x21\x22\x0a\x0e\x7c\x8a\x9d\x9b\xce\xdb\xa7\x76\x67\xc8\xed\x3f\xd7\x2a\xf9\x4c\x25\xb4\x90\x2a\x75\x83\x64\xe7\xaf\x42\x23\x11\x25\x0a\xd0\x5b\xad\x61\x7c\xea\x9d\x9b\x2e\x88\x2e\x4b\x2f\x26\xac\x32\x79\xc2\xaa\x17\xe2\x3d\xce\x15\x7b\x08\x15\x2a\x98\x9d\xb7\x7b\xea\x19\xc2\x63\x1b\xb4\xa4\x6a\x64\x1b\xe6\x9b\x55\x28\xd7\xaa\xf6\x60\xab\x50\x04\x2c\xcd\xf2\x0d\x00\x00\xff\xff\x5e\x93\x94\x37\x7e\x01\x00\x00")

func _lgraphqlUsergroupsUserbysshfingerprintGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsUserbysshfingerprintGraphql,
		"_lgraphql/usergroups/userBySSHFingerprint.graphql",
	)
}

func _lgraphqlUsergroupsUserbysshfingerprintGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsUserbysshfingerprintGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/userBySSHFingerprint.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsUserbysshkeyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x4e\xbb\x0e\xc2\x30\x10\xdb\xfb\x15\x87\xc4\x50\x7e\x81\x91\x81\x05\x89\x81\x22\xf6\x08\x4c\x39\x91\x47\xb9\x24\x43\x84\xfa\xef\xa8\x85\x86\x2a\xf5\x74\xf6\x59\xb6\x5f\x11\x92\xa8\x5e\x7b\xff\x38\x20\x6d\xa9\x09\xc2\xb6\x5d\x6d\xe8\x5d\x11\x11\x45\x0f\xd9\xa5\x66\x7c\xd6\x93\xe7\x67\x9e\x3c\x03\xf8\x96\x4f\x18\xc5\x3a\xb3\x3b\x8b\x0f\x47\x65\x90\x15\xad\x0a\xe1\xea\x8c\x81\x0d\x99\x7f\xd3\xfd\x2c\xbd\x68\x18\x60\xe7\x09\x03\x9e\x48\xe7\xd4\x2d\xb4\x8b\xd2\x71\x21\xee\xd9\xb6\x90\x4e\x78\xd6\x3a\x2e\x11\xa8\x80\x7f\x51\x9f\xaf\x56\x5c\xec\x4e\x4e\xa3\x9c\xb5\xd8\x21\x4e\xa3\x08\xe8\xab\xfe\x13\x00\x00\xff\xff\xd2\x6c\xcd\xa1\x67\x01\x00\x00")

func _lgraphqlUsergroupsUserbysshkeyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsUserbysshkeyGraphql,
		"_lgraphql/usergroups/userBySSHKey.graphql",
	)
}

func _lgraphqlUsergroupsUserbysshkeyGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsUserbysshkeyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/userBySSHKey.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsUsercansshtoenvironmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\x2d\x2e\x48\x4c\x4e\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x2d\x4e\x2d\x72\x4e\xcc\x0b\x2e\xce\x08\xc9\x77\xcd\x2b\xcb\x2c\xca\xcf\xcb\x4d\xcd\x2b\xd1\xc8\x2e\x4d\x4a\x2d\xca\x4b\x2d\x49\x2d\xf6\x83\xe9\x05\x31\xac\x14\x10\x66\xc1\x8c\x00\x01\x1c\xca\xc1\xf2\xb5\x5c\xb5\x80\x00\x00\x00\xff\xff\xb9\x87\xd6\x3a\x88\x00\x00\x00")

func _lgraphqlUsergroupsUsercansshtoenvironmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsUsercansshtoenvironmentGraphql,
		"_lgraphql/usergroups/userCanSSHToEnvironment.graphql",
	)
}

func _lgraphqlUsergroupsUsercansshtoenvironmentGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsUsercansshtoenvironmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/userCanSSHToEnvironment.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsUsersbyorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8b\x41\x0a\xc2\x40\x0c\x45\xf7\x39\xc5\x17\x5c\xe8\x15\xba\x74\xe7\x46\xcf\x10\x6c\x94\x40\x27\x83\x99\x14\xa9\xd2\xbb\x4b\xd3\xcd\xac\xf2\x5e\x78\xff\x3d\x8b\x2f\x38\x11\x70\xd4\x71\xc0\xd5\xe2\x70\xc6\x8f\x80\xb9\x89\xb7\xcb\x72\xf7\x17\x9b\x7e\x39\xb4\xda\x56\x01\xb5\xfb\x0c\xdb\x6a\xef\x01\x1d\xf3\x48\x61\x9d\x92\x9e\xea\x2d\x6e\x5c\x24\x6d\xe2\x4e\xea\xc7\xc4\x93\x1e\xb5\x14\xb1\x20\x60\xa5\x95\xfe\x01\x00\x00\xff\xff\x56\xbe\xa2\x65\x8f\x00\x00\x00")

func _lgraphqlUsergroupsUsersbyorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsUsersbyorganizationGraphql,
		"_lgraphql/usergroups/usersByOrganization.graphql",
	)
}

func _lgraphqlUsergroupsUsersbyorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsUsersbyorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/usersByOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlUsergroupsUsersbyorganizationnameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8c\x31\x0a\x02\x41\x0c\x45\xfb\x39\xc5\x17\x2c\xf4\x0a\x96\x1e\xc0\xc6\x13\x04\x8d\x4b\x60\x93\xc1\x6c\x44\x56\xd9\xbb\xcb\xcc\xa4\xb1\x08\xbc\x07\x2f\xff\xf9\x62\x5f\x71\x28\xc0\xde\x48\xf9\x84\x6b\xb8\xd8\xb4\x3b\xe2\x5b\x80\xea\x13\x99\x7c\x28\xa4\xda\x79\xbd\x90\x72\x2b\x81\x91\xf6\x8f\x11\x02\xf5\x6d\xec\x4b\x0a\x20\xf7\x04\x56\x92\x39\xf9\x21\xbe\x44\x5b\x49\x9f\xe9\x4f\xfb\x44\xf2\xad\xaa\xb2\x45\xb7\xad\xb4\xdb\xca\x2f\x00\x00\xff\xff\x3b\xd8\x0c\x59\xac\x00\x00\x00")

func _lgraphqlUsergroupsUsersbyorganizationnameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlUsergroupsUsersbyorganizationnameGraphql,
		"_lgraphql/usergroups/usersByOrganizationName.graphql",
	)
}

func _lgraphqlUsergroupsUsersbyorganizationnameGraphql() (*asset, error) {
	bytes, err := _lgraphqlUsergroupsUsersbyorganizationnameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/usergroups/usersByOrganizationName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlOrganizationsAddorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xc1\x4a\xc5\x30\x10\x45\xf7\xf9\x8a\x11\x5c\xb4\xd0\x2f\xe8\x5e\xc4\x4d\x15\xfd\x82\xd0\xa4\x65\xc4\x4e\x6a\x9c\x08\xfa\xe8\xbf\x3f\x92\x94\xf7\x66\xda\xed\xbd\x77\x06\xce\x59\x12\x5b\xc6\x40\xd0\x18\x80\x47\xb2\x8b\xef\xe1\x83\x23\xd2\xfc\xd0\xe5\x64\x8a\xe8\xc9\x7d\xfd\x0d\xa2\x29\x85\xf3\x3f\x63\xc4\x35\xdf\xaa\xfc\x3b\x05\xb6\x6f\x31\x7c\xfa\x91\x7b\x78\x21\xbe\xa7\xcf\x31\xa4\xf5\x90\x0d\x81\x71\xc2\xd1\xd6\x3f\xaa\x7a\xa2\x5f\x8c\x81\x16\x4f\xc7\x47\xef\x21\xb1\xaf\x59\x0b\x17\x03\x60\x9d\x7b\x8d\xb3\x25\xfc\x2f\x8f\x1a\xa4\x35\x71\x5f\x2a\x80\x0a\x55\xd8\xba\x12\x68\x26\x85\x58\x07\x8a\x4d\x92\xd6\x5a\x23\x2a\x62\x31\xd8\x69\x05\xba\x28\x35\xf6\x59\x85\x98\x2a\x0d\x27\x33\x62\xb8\x5b\x11\x8a\x72\xb9\xb5\xbb\x06\x74\x37\x1b\x39\x36\x9b\xb9\x06\x00\x00\xff\xff\x80\x81\xe4\x11\xfc\x01\x00\x00")

func _lgraphqlOrganizationsAddorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlOrganizationsAddorganizationGraphql,
		"_lgraphql/organizations/addOrganization.graphql",
	)
}

func _lgraphqlOrganizationsAddorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlOrganizationsAddorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/organizations/addOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlOrganizationsAllorganizationsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8c\x41\x0a\xc2\x40\x0c\x45\xf7\x73\x8a\x5c\x46\xdc\x55\xf1\x06\x61\x26\x95\xc8\x34\x69\xd3\x8c\x50\xa5\x77\x97\x06\x74\x5c\x25\xff\xf1\xff\x5b\x1a\xd9\x06\xef\x04\x80\xb5\x5e\xec\x8e\xc2\x2f\x74\x56\x59\x03\x02\x70\x89\x23\x38\x51\x3c\x85\xd6\x6c\x3c\x1f\x95\xc8\xa3\x31\x49\xa9\xdb\xf0\x2d\x2c\x4d\x1d\xaf\xa6\x0f\xca\xde\xc1\xd9\xb4\xcd\x3d\x0e\xea\x3c\x72\xc6\x9f\x26\xe8\x49\x9e\x6c\x2a\x13\xc9\xdf\xf2\xa6\xcd\x0f\xf3\x9e\xf6\xf4\x09\x00\x00\xff\xff\xf9\xa7\x53\xe6\xae\x00\x00\x00")

func _lgraphqlOrganizationsAllorganizationsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlOrganizationsAllorganizationsGraphql,
		"_lgraphql/organizations/allOrganizations.graphql",
	)
}

func _lgraphqlOrganizationsAllorganizationsGraphql() (*asset, error) {
	bytes, err := _lgraphqlOrganizationsAllorganizationsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/organizations/allOrganizations.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlOrganizationsDeleteorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4c\xb1\x52\xf0\xcc\x2b\x51\xe4\xd2\xac\xe6\x52\x50\x48\x49\xcd\x49\x2d\x49\xf5\x2f\x4a\x4f\xcc\xcb\xac\x02\xab\x02\x29\x52\x50\xc8\xcc\x2b\x28\x2d\xb1\x52\xa8\x06\x73\x14\x14\x40\x9a\x54\x32\x53\xc0\xbc\x5a\x2e\x05\x05\x4d\xae\x5a\x2e\x40\x00\x00\x00\xff\xff\x54\xd8\xea\x66\x57\x00\x00\x00")

func _lgraphqlOrganizationsDeleteorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlOrganizationsDeleteorganizationGraphql,
		"_lgraphql/organizations/deleteOrganization.graphql",
	)
}

func _lgraphqlOrganizationsDeleteorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlOrganizationsDeleteorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/organizations/deleteOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlOrganizationsOrganizationbyidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcd\x41\xca\xc2\x30\x10\x05\xe0\x7d\x4f\x31\x3f\x74\xd1\xff\x0a\x2e\x05\x91\x6e\x8a\x78\x83\x90\x4c\x65\xc4\xce\xb4\xe3\x44\x88\xe2\xdd\x25\x35\x62\x5c\x04\x1e\x1f\xef\x65\x96\x88\x9a\xa0\x6b\x00\x5a\x0a\x1b\xe8\xd9\xfe\xfe\xe1\xd1\x00\x00\x88\x9e\x1c\xd3\xdd\x19\x09\x6f\x53\x1f\xba\x55\x01\x72\xaf\xa5\xf0\xa9\x65\x28\x81\xdd\x84\x25\x06\xbc\x7a\xa5\x39\x6f\x8b\x8c\x4a\xc8\xe1\x92\x86\x6f\x69\x89\x62\xee\xa0\x72\x46\x6f\x35\xed\x55\xe2\x5c\xc3\x20\x46\x23\x79\x57\x7d\xb7\xfa\x8e\x6f\xa4\xc2\x13\xf2\xcf\xfe\x28\xd1\xde\x37\x9e\x4d\x7e\xaf\x00\x00\x00\xff\xff\x19\xdb\x0e\xd3\xe5\x00\x00\x00")

func _lgraphqlOrganizationsOrganizationbyidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlOrganizationsOrganizationbyidGraphql,
		"_lgraphql/organizations/organizationById.graphql",
	)
}

func _lgraphqlOrganizationsOrganizationbyidGraphql() (*asset, error) {
	bytes, err := _lgraphqlOrganizationsOrganizationbyidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/organizations/organizationById.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlOrganizationsOrganizationbynameGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcd\x41\x0a\xc2\x30\x10\x05\xd0\x7d\x4e\xf1\x05\x17\xf5\x0a\x2e\x05\x71\x57\x44\x4f\x10\xda\x69\x19\xb1\x33\xed\x98\x08\x55\x7a\x77\x31\x8d\x10\x77\xe1\xe5\xff\x3f\x53\x24\x9b\x51\x39\x60\x2b\x7e\xa0\x3d\xae\xc1\x58\xfa\xcd\x0e\x6f\x07\x00\x6a\xbd\x17\x7e\xf9\xc0\x2a\x87\xb9\xf6\x03\x55\xc9\x81\x35\x9e\x5a\xbf\x30\xc0\x6d\xf1\x9b\x9f\x2d\x3d\x1a\xe3\xf1\xbb\x90\xa5\x33\x26\x69\xef\x69\x2e\xd3\x14\x35\xf8\xb3\xe9\x8d\x9a\x50\xd2\xc9\x34\x8e\x25\xd4\x1a\xb8\xe3\xc6\x17\x73\xc9\x8f\xf2\x64\x53\x19\x48\xfe\xfa\x17\x8d\x61\xbd\xb1\xb8\xc5\x7d\x02\x00\x00\xff\xff\xc3\xe8\x9e\x1a\xee\x00\x00\x00")

func _lgraphqlOrganizationsOrganizationbynameGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlOrganizationsOrganizationbynameGraphql,
		"_lgraphql/organizations/organizationByName.graphql",
	)
}

func _lgraphqlOrganizationsOrganizationbynameGraphql() (*asset, error) {
	bytes, err := _lgraphqlOrganizationsOrganizationbynameGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/organizations/organizationByName.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlOrganizationsUpdateorganizationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\x31\x0e\x02\x31\x0c\x04\x7b\xbf\x62\x4f\xba\x82\xfb\x42\x7e\x70\x15\x34\x3c\xc0\x22\x11\xb8\x88\x89\x90\xd3\x10\xe5\xef\xc8\x49\x3a\x3a\xaf\x77\x34\x9b\xab\xb1\xc9\x5b\x71\x21\x60\x97\x18\x70\xaa\x6d\x7e\x17\xb6\xc7\x2b\xe0\x5e\x22\x5b\xba\x7e\x9e\xac\xf2\x1d\xe8\xcd\x8b\x53\x4b\xb5\xed\x68\x04\xd4\x3f\xc2\x5d\x80\x38\x12\xd0\x46\x00\xdc\xbd\x4b\x5c\x69\xd9\xe7\xca\xf8\x75\x02\x8e\x09\x2f\x4a\x39\x27\xf2\xa2\xd3\x2f\x00\x00\xff\xff\x70\xcd\x28\x48\xa7\x00\x00\x00")

func _lgraphqlOrganizationsUpdateorganizationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlOrganizationsUpdateorganizationGraphql,
		"_lgraphql/organizations/updateOrganization.graphql",
	)
}

func _lgraphqlOrganizationsUpdateorganizationGraphql() (*asset, error) {
	bytes, err := _lgraphqlOrganizationsUpdateorganizationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/organizations/updateOrganization.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\xa4\xe6\x26\x66\xe6\x38\xa6\xa4\x14\xa5\x16\x17\xc3\x65\x40\x12\xf9\x45\xe9\x89\x79\x99\x55\x60\xcd\x56\x0a\x9e\x79\x25\x9a\x0a\xd5\x5c\x0a\x0a\x0a\x0a\x89\x29\x29\x7e\xf9\x25\x99\x69\x99\xc9\x60\x39\x57\x90\x01\x1a\x99\x79\x05\xa5\x25\x56\x50\x15\x0a\x0a\x10\x7b\xc0\xd6\xe9\x40\x85\x50\x2d\x42\xb1\x17\xa6\x04\xd5\x4a\x14\x17\x80\x55\xd4\x6a\xc2\x2d\xc8\x4c\x41\xb2\x09\x8b\x0d\x58\x4c\x84\x18\xc1\x05\xc2\x80\x00\x00\x00\xff\xff\xc5\xbd\xd3\x53\x15\x01\x00\x00")

func _lgraphqlNotificationsAddnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationemailGraphql,
		"_lgraphql/notifications/addNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x0a\xc2\x30\x10\x85\xf7\x39\xc5\x13\x5c\xb4\xd0\x13\xe4\x06\x2e\x74\xa3\x17\x88\x4d\x5a\x07\xc9\x8c\xa4\x53\x04\xc5\xbb\x4b\xd3\x1f\x0c\xb8\xc8\xe6\xcd\x7b\xdf\x97\x38\xaa\x53\x12\x46\x65\x80\x3d\xbb\x18\x2c\xce\x9a\x88\xfb\x5d\x33\x25\xcf\x70\xbd\x89\xdc\xcb\x50\x52\xef\x98\x5e\x79\x68\x71\x60\xad\xf1\x36\x00\xe0\xbc\x3f\x89\x52\x47\x6d\xbe\x1d\xa9\x4d\x32\x48\xa7\x97\xe0\xe2\x50\x11\x3f\x46\xb5\x4b\x15\x98\x65\xd9\xd9\x2c\xd1\x66\x5b\xbd\xeb\xa1\x34\x16\x1f\xc8\x8d\x4f\xbd\x61\xc9\xff\xf0\x4b\xee\x1f\xd8\xbc\x36\xd3\xfb\x06\x00\x00\xff\xff\x53\x9c\x13\xf7\x0b\x01\x00\x00")

func _lgraphqlNotificationsAddnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/addNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xc1\x0a\x82\x50\x10\x45\xf7\x7e\xc5\x0d\x5a\x28\xf8\x05\x6e\x5b\xb5\x69\x51\x5f\x30\xe9\x4b\x07\x73\x26\x64\x24\x28\xfa\xf7\xf0\xe9\x93\x06\x5a\xcc\xe6\xde\xe1\x9e\x33\x4c\x46\xc6\x2a\xc8\x33\x60\x2f\x34\x84\x0a\x17\x1b\x59\xda\x5d\x39\x27\x75\x47\x22\xe1\xee\xc3\x67\xb8\x76\xaa\xbd\x0f\x75\x6c\x49\xf8\x15\xd7\x2a\x1c\xc5\x0a\xbc\x33\x00\xa0\xa6\x39\xa9\xf1\x8d\xeb\xd8\x9d\xb5\xee\x83\x1d\x3a\xb2\x9c\xe5\x31\x59\xb5\xbe\x01\x0b\x3d\x4a\x94\x6b\xb4\xe1\x93\x48\x2a\x36\x85\x24\x93\x0a\xaf\xe1\xac\xe2\xc7\xa7\xd8\x78\xdc\xfc\x80\x3d\xd0\x53\xfe\x4c\x2f\x5b\xd9\x7c\xdf\x00\x00\x00\xff\xff\xea\x3a\x5d\xd7\x43\x01\x00\x00")

func _lgraphqlNotificationsAddnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationrocketchatGraphql,
		"_lgraphql/notifications/addNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x41\x0a\x83\x40\x0c\x45\xf7\x9e\xe2\x17\xba\x50\xf0\x04\xde\xa0\x9b\x6e\x3c\x41\xea\x4c\x35\xa8\x99\x22\x91\x42\x4b\xef\x5e\x1c\x9d\xa1\x81\x2e\xb2\xf9\x3f\xfc\xf7\xe6\x55\x49\x39\x08\xca\x02\x38\x0b\xcd\xbe\x41\xab\x0b\x4b\x7f\xaa\xb7\xa4\x1b\x48\xc4\x4f\x36\x7c\xfa\xdb\x10\xc2\x68\xc3\xb0\xf4\x24\xfc\x8a\x6b\x0d\x2e\xa2\x15\xde\x05\x00\x90\x73\xd7\xa0\x7c\xe7\x2e\x76\xed\x44\xdd\x58\xb2\x3c\x56\x6d\x8e\x0f\x60\x07\x47\x7e\x7d\x44\x99\x9c\x1c\x52\x91\xe9\xc9\x23\x15\xd6\xc0\x08\xc5\x8f\x4f\x95\x79\xec\x7e\xc0\x16\x68\x29\x7f\xa6\xf7\xad\x62\xbb\x6f\x00\x00\x00\xff\xff\xa8\xe5\x1d\xa5\x3e\x01\x00\x00")

func _lgraphqlNotificationsAddnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationslackGraphql,
		"_lgraphql/notifications/addNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationtoprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x0a\x02\x31\x0c\x85\xf7\x39\xc5\x13\x5c\x38\x57\xe8\x21\x06\x41\x2f\x50\xa6\x55\x22\x4c\x5a\x86\xcc\x42\x86\xde\x5d\xda\xfa\x83\xc6\x55\xcb\xcb\x97\xbc\x6f\x5e\xd5\x2b\x27\xc1\x81\x80\x7d\x5e\xd2\x2d\x4e\xea\x70\xd2\x85\xe5\xba\xab\x99\x24\xe5\x0b\x4f\x8d\x3a\xdf\x73\x74\x18\x7f\x12\x83\x8d\x7e\x8e\x9f\x1b\xc3\x46\x80\x0f\xe1\x6b\x2d\x1d\x7b\x55\xad\x05\x58\xf2\xaa\x6e\x6b\x7f\xe0\x6d\xf1\xf2\x79\xe6\xd6\xc4\xc8\xfd\x21\xbb\x8c\xf1\x6b\x64\x21\x60\xe8\xb5\x1c\xda\x23\x7d\x54\xa8\xd0\x23\x00\x00\xff\xff\x6a\x63\xe5\x98\x1b\x01\x00\x00")

func _lgraphqlNotificationsAddnotificationtoprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationtoprojectGraphql,
		"_lgraphql/notifications/addNotificationToProject.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationtoprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationtoprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationToProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsAddnotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x89\x94\xa7\x26\x65\xe4\xe7\x67\xa3\x0a\xe6\x17\xa5\x27\xe6\x65\x56\x81\x35\x5a\x29\x78\xe6\x95\x68\x2a\x54\x73\x29\x28\x28\x28\x24\xa6\xa4\xf8\xe5\x97\x64\xa6\x65\x26\x83\xe5\xc2\x21\x9a\x35\x32\xf3\x0a\x4a\x4b\xac\xa0\x6a\x14\x14\x20\xb6\x80\x2d\xd3\x81\x0a\xc1\xad\x81\x59\x08\x93\x40\xb5\x0a\xc5\x66\xb0\x8a\x5a\x4d\xb8\xb1\x99\x29\x48\xe6\xa3\x9a\x8b\xc5\x30\x88\x6e\x2e\x10\x06\x04\x00\x00\xff\xff\xd4\x9d\x98\x54\x04\x01\x00\x00")

func _lgraphqlNotificationsAddnotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsAddnotificationwebhookGraphql,
		"_lgraphql/notifications/addNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsAddnotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsAddnotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/addNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\xe6\x9a\x9b\x98\x99\xa3\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x07\x6c\x5c\xad\x26\x57\x2d\x20\x00\x00\xff\xff\xac\x83\xb7\xaa\x62\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationemailGraphql,
		"_lgraphql/notifications/deleteNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\xe6\x9b\x99\x5c\x94\x5f\x9c\x9f\x56\x12\x92\x9a\x98\x5b\xac\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x10\x6c\x6e\xad\x26\x57\x2d\x20\x00\x00\xff\xff\x54\xf6\x18\xff\x6b\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\x16\x94\x9f\x9c\x9d\x5a\xe2\x9c\x91\x58\xa2\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x0c\x6c\x66\xad\x26\x57\x2d\x20\x00\x00\xff\xff\x36\x65\x00\x24\x67\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationrocketchatGraphql,
		"_lgraphql/notifications/deleteNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\x16\x9c\x93\x98\x9c\xad\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x07\x6c\x5c\xad\x26\x57\x2d\x20\x00\x00\xff\xff\xee\xc5\xad\x57\x62\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationslackGraphql,
		"_lgraphql/notifications/deleteNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsDeletenotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\xcd\x49\x2d\x49\xf5\xcb\x2f\xc9\x4c\xcb\x4c\x06\xab\xb4\xc2\x22\x16\x9e\x9a\x94\x91\x9f\x9f\xad\x91\x99\x57\x50\x5a\x62\x55\x0d\x31\x09\x6c\x60\xad\x26\x57\x2d\x20\x00\x00\xff\xff\x8a\x89\x7f\x55\x64\x00\x00\x00")

func _lgraphqlNotificationsDeletenotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsDeletenotificationwebhookGraphql,
		"_lgraphql/notifications/deleteNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsDeletenotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsDeletenotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/deleteNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\x70\xf5\x75\xf4\xf4\xd1\x44\xd2\x0c\x02\x7a\x7a\x7a\x0a\xf9\x79\x0a\x7e\x48\xea\x5d\x73\x13\x33\x73\xd0\x94\x81\x40\x7c\x3c\xc8\x1c\x14\x1b\x61\x20\x15\xa4\xc5\x31\x25\xa5\x28\xb5\xb8\x18\x43\x12\x43\x47\x2d\x17\x2a\xab\x96\xab\x16\x10\x00\x00\xff\xff\x24\xd5\x22\x2e\xea\x00\x00\x00")

func _lgraphqlNotificationsListallnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationemailGraphql,
		"_lgraphql/notifications/listAllNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\xf0\xf5\x74\x0e\xf2\x0f\xf6\x77\x0b\x09\x71\x75\xf4\x0d\xd6\x44\x32\x05\x04\xf4\xf4\xf4\x14\xf2\xf3\x14\xfc\x90\x34\xfa\x66\x26\x17\xe5\x17\xe7\xa7\x95\x84\xa4\x26\xe6\x16\xa3\xa9\x07\x81\xf8\x78\x90\xc9\x28\x6e\x80\x81\xf2\xd4\xa4\x8c\xfc\xfc\x6c\x0c\x71\x0c\xc5\xb5\x5c\xa8\xac\x5a\xae\x5a\x40\x00\x00\x00\xff\xff\x3c\x1a\x87\x3b\xf7\x00\x00\x00")

func _lgraphqlNotificationsListallnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\x08\xf2\x77\xf6\x76\x0d\x71\xf6\x70\x0c\xd1\x44\x32\x01\x04\xf4\xf4\xf4\x14\xf2\xf3\x14\xfc\x90\x34\x05\xe5\x27\x67\xa7\x96\x38\x67\x24\x96\xa0\xa9\x05\x81\xf8\x78\x90\x89\x28\x76\xc3\x40\x79\x6a\x52\x46\x7e\x7e\x36\x86\x78\x72\x46\x62\x5e\x5e\x6a\x0e\x86\x38\x86\x21\xb5\x5c\xa8\xac\x5a\xae\x5a\x40\x00\x00\x00\xff\xff\x4a\x62\x23\x28\x07\x01\x00\x00")

func _lgraphqlNotificationsListallnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationrocketchatGraphql,
		"_lgraphql/notifications/listAllNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\x08\xf6\x71\x74\xf6\xd6\x44\xd2\x0c\x02\x7a\x7a\x7a\x0a\xf9\x79\x0a\x7e\x48\xea\x83\x73\x12\x93\xb3\xd1\x94\x81\x40\x7c\x3c\xc8\x1c\x14\x1b\x61\xa0\x3c\x35\x29\x23\x3f\x3f\x1b\x43\x3c\x39\x23\x31\x2f\x2f\x35\x07\x43\x1c\xc3\x90\x5a\x2e\x54\x56\x2d\x57\x2d\x20\x00\x00\xff\xff\xff\x6b\xa2\xcc\xfd\x00\x00\x00")

func _lgraphqlNotificationsListallnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationslackGraphql,
		"_lgraphql/notifications/listAllNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsListallnotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\x48\xcc\xc9\x09\x28\xca\xcf\x4a\x4d\x2e\x29\x86\x8a\x80\x40\x5e\x62\x6e\x2a\x9c\x93\x99\x82\x10\xcf\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\xd6\x28\xa9\x2c\x48\xb5\x52\x08\x77\x75\xf2\xf0\xf7\xf7\xd6\x44\xd2\x0e\x02\x7a\x7a\x7a\x0a\xf9\x79\x0a\x7e\x48\x3a\xc2\x53\x93\x32\xf2\xf3\xb3\xd1\x14\x82\x40\x7c\x3c\xc8\x2c\x14\x5b\x61\xa0\x1c\xa2\x09\x43\x1c\x43\x71\x2d\x17\x2a\xab\x96\xab\x16\x10\x00\x00\xff\xff\x92\xfc\x32\xab\xe9\x00\x00\x00")

func _lgraphqlNotificationsListallnotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsListallnotificationwebhookGraphql,
		"_lgraphql/notifications/listAllNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsListallnotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsListallnotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/listAllNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x01\xab\x84\x29\x00\x81\xbc\xfc\x92\xcc\xb4\xcc\xe4\xc4\x92\xcc\xfc\xbc\x62\x8d\x92\xca\x82\x54\x2b\x05\x57\x5f\x47\x4f\x1f\x64\x45\x20\xa0\xa7\xa7\xa7\x90\x9f\xa7\xe0\x87\xa4\xde\x35\x37\x31\x33\x07\x4d\x19\x08\xc4\xc7\x83\xcc\x01\xd9\x84\x21\x95\x0a\xd2\xe2\x98\x92\x52\x94\x5a\x5c\x8c\x21\x89\xa1\xa3\x96\x0b\x95\x55\xcb\x55\x0b\x08\x00\x00\xff\xff\x65\xff\x3e\x1e\xf2\x00\x00\x00")

func _lgraphqlNotificationsProjectnotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationemailGraphql,
		"_lgraphql/notifications/projectNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x01\xab\x84\x29\x00\x81\xbc\xfc\x92\xcc\xb4\xcc\xe4\xc4\x92\xcc\xfc\xbc\x62\x8d\x92\xca\x82\x54\x2b\x05\x5f\x4f\xe7\x20\xff\x60\x7f\xb7\x90\x10\x57\x47\xdf\x60\x64\xd5\x20\xa0\xa7\xa7\xa7\x90\x9f\xa7\xe0\x87\xa4\xd1\x37\x33\xb9\x28\xbf\x38\x3f\xad\x24\x24\x35\x31\xb7\x18\x4d\x3d\x08\xc4\xc7\x83\x4c\x06\xd9\x8d\x21\x55\x9e\x9a\x94\x91\x9f\x9f\x8d\x21\x8e\xa1\xb8\x96\x0b\x95\x55\xcb\x55\x0b\x08\x00\x00\xff\xff\xbf\xb5\x85\x65\xff\x00\x00\x00")

func _lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/projectNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xc1\x0a\xc2\x30\x0c\x86\xef\x7b\x8a\x08\x1e\xe6\x65\x0f\xb0\x9b\x0e\x41\x10\x26\xcc\xdd\x47\x2d\xd1\xd6\xb9\x64\xd6\x88\x14\xe9\xbb\x4b\x15\x61\xda\xef\x14\xfe\x7c\xf9\x73\xbd\xa3\xf3\x90\xcf\x49\x0d\x58\xc2\x5e\x9c\xa5\xd3\x6c\x01\xcf\x0c\x00\x60\x74\x7c\x46\x2d\x2b\x5f\xab\x01\xf3\x8f\xf2\x36\xbf\x42\x84\x58\xec\xd1\x6a\x25\x96\xe9\x96\x8b\x1f\xb1\x84\x66\x57\x6d\xd7\x6d\xb5\x59\xb6\x53\x33\x52\x14\x05\x30\x41\x3d\x39\x6a\x58\xf7\x28\x95\x51\xf2\xe7\x46\xba\x2e\x36\xc6\x9f\xc9\xea\x81\x07\xc3\xdc\x27\xb9\x36\x8a\x08\x2f\x49\x9e\x94\x84\xec\x77\x0a\x59\x78\x05\x00\x00\xff\xff\x98\xc9\xde\xdd\x0f\x01\x00\x00")

func _lgraphqlNotificationsProjectnotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationrocketchatGraphql,
		"_lgraphql/notifications/projectNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x01\xab\x84\x29\x00\x81\xbc\xfc\x92\xcc\xb4\xcc\xe4\xc4\x92\xcc\xfc\xbc\x62\x8d\x92\xca\x02\x90\x39\x3e\x8e\xce\xde\xc8\x8a\x40\x40\x4f\x4f\x4f\x21\x3f\x4f\xc1\x0f\x49\x7d\x70\x4e\x62\x72\x36\x9a\x32\x10\x88\x8f\x07\x99\x03\xb2\x09\x43\xaa\x3c\x35\x29\x23\x3f\x3f\x1b\x43\x3c\x39\x23\x31\x2f\x2f\x35\x07\x43\x1c\xc3\x90\x5a\x2e\x54\x56\x2d\x57\x2d\x20\x00\x00\xff\xff\xff\xce\xda\x6a\x05\x01\x00\x00")

func _lgraphqlNotificationsProjectnotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationslackGraphql,
		"_lgraphql/notifications/projectNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsProjectnotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x28\x28\xca\xcf\x4a\x4d\x2e\x71\xaa\xf4\x4b\xcc\x4d\xd5\x80\x28\x01\xab\x84\x29\x00\x81\xbc\xfc\x92\xcc\xb4\xcc\xe4\xc4\x92\xcc\xfc\xbc\x62\x8d\x92\xca\x82\x54\x2b\x85\x70\x57\x27\x0f\x7f\x7f\x6f\x64\x65\x20\xa0\xa7\xa7\xa7\x90\x9f\xa7\xe0\x87\xa4\x23\x3c\x35\x29\x23\x3f\x3f\x1b\x4d\x21\x08\xc4\xc7\x83\xcc\x02\xd9\x86\x21\x55\x0e\xd1\x84\x21\x8e\xa1\xb8\x96\x0b\x95\x55\xcb\x55\x0b\x08\x00\x00\xff\xff\x7d\x54\xe7\x02\xf1\x00\x00\x00")

func _lgraphqlNotificationsProjectnotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsProjectnotificationwebhookGraphql,
		"_lgraphql/notifications/projectNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsProjectnotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsProjectnotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/projectNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsRemovenotificationfromprojectGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x0a\x02\x31\x0c\x85\xf7\x39\xc5\x13\x5c\x38\x57\xe8\x01\x5c\x0e\x82\x5e\x60\x18\xa3\x44\x68\x53\x4a\x46\x90\x61\xee\x2e\x6d\xfd\xc3\xce\x2e\x79\x7c\xc9\xfb\xfc\x64\x83\x89\x06\xec\x08\xd8\xc6\xa4\x37\x1e\xcd\xe1\x68\x49\xc2\x75\x93\xb3\xa0\x26\x17\x19\x0b\x75\x7a\x44\x76\xe8\xff\x92\x06\xeb\x07\xcf\xdf\x1f\xdd\x4c\x40\x62\xaf\x77\xfe\xbd\xdc\x27\xf5\x87\xda\x97\xbb\x01\x09\x71\x32\x37\x97\x19\xf8\xa8\xbc\xa5\x5e\x79\xab\xd3\x18\xae\x90\xd5\xa8\x91\x2c\xe4\x42\x40\x57\x6b\xe5\x4c\x79\x5f\xe8\x19\x00\x00\xff\xff\x17\x0d\x7b\xbb\x17\x01\x00\x00")

func _lgraphqlNotificationsRemovenotificationfromprojectGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsRemovenotificationfromprojectGraphql,
		"_lgraphql/notifications/removeNotificationFromProject.graphql",
	)
}

func _lgraphqlNotificationsRemovenotificationfromprojectGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsRemovenotificationfromprojectGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/removeNotificationFromProject.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationemailGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x3d\x0a\x02\x31\x10\x85\xeb\xcc\x29\x66\xc1\xc2\xbd\x42\x3a\x0b\x0b\x1b\x11\xc4\x03\x0c\x9b\xa8\x03\x26\x1b\x76\x27\x55\xc8\xdd\x65\x36\x12\x14\x9c\xee\xfd\xe4\x7d\x09\x59\x48\x78\x8e\xb8\x07\x44\xc4\x5d\xa4\xe0\x2d\x5e\x65\xe1\xf8\x18\x9a\x95\x48\xa6\xa7\xc5\x5b\x72\x24\xfe\x3c\x0b\xdf\x79\xda\xde\x1c\x03\xf1\xeb\xa2\xe9\x29\xa6\x2c\x03\x8c\x58\xc0\xe4\xff\xbd\xb6\xaf\xc7\x5a\xb6\xa5\x6b\xbd\x86\xdd\xe8\x3f\xfe\x87\xdd\xfe\xd0\x93\x0a\x66\x2c\x60\x0c\x3b\xf8\x1e\xe8\xc2\x2b\xf0\xe0\xdc\xe2\xd7\x15\x4c\x85\xfa\x0e\x00\x00\xff\xff\x53\xd5\xfd\x0c\xe6\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationemailGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationemailGraphql,
		"_lgraphql/notifications/updateNotificationEmail.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationemailGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationemailGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationEmail.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x31\x0e\xc2\x30\x0c\x45\xe7\xf8\x14\xae\xc4\x40\xaf\xd0\x1b\x30\x80\x90\x80\x03\x84\x34\xa5\x16\x4a\x1c\xb5\x8e\x18\xa2\xdc\x1d\xa5\x41\x11\x4c\xfc\xed\x7f\xfb\xdb\xcf\x45\xd1\x42\xec\x71\x0f\x88\x88\x3b\xaf\x9d\x1d\xf0\x22\x0b\xf9\x47\x57\xa3\xa0\xc5\xcc\x03\xde\xc2\xa8\xc5\x9e\x58\x68\x22\xb3\x75\x8e\x64\x16\x5e\x79\x92\xab\xd5\x6e\x3d\x97\xb5\x83\x0f\x51\x3a\xe8\x31\x81\x8a\x7f\x0a\xf5\x63\x11\x95\xd6\x90\x9a\x2f\xaa\x20\x1b\xcf\x4f\xfe\xa1\xa9\x54\x6d\x92\x41\xf5\x09\x94\xa2\x11\xbe\x0f\x34\xf3\xb2\xf7\x99\xf9\x09\x2a\x43\x7e\x07\x00\x00\xff\xff\x71\xce\xd3\xe8\xf3\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql,
		"_lgraphql/notifications/updateNotificationMicrosoftTeams.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationMicrosoftTeams.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationrocketchatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x0e\xc2\x30\x0c\x44\xe7\xf8\x2b\x5c\x89\x81\xfe\x42\x57\x26\x16\x84\x40\x7c\x80\x49\x03\xb1\x4a\x9d\x0a\x39\x62\x88\xf2\xef\x28\x0d\x8a\x60\xe1\xb6\xbb\xdc\xc5\x6f\x8e\x4a\xca\x41\x70\x0b\x88\x88\x1b\xa1\xd9\x0d\x78\xd6\x27\xcb\xbd\xab\xd1\x42\x6a\xfd\x80\x97\x65\x24\x75\x87\xa0\x7c\x63\xbb\x6e\x4e\xc1\x4e\x4e\x77\x9e\xf4\x58\x2a\x7b\x59\xa2\x76\xd0\x63\x02\x13\xff\x94\xeb\xa5\x22\x2e\x8b\x21\x35\x5f\x54\x01\x56\x8e\x9f\xfc\x43\x51\x69\xda\x4b\x06\xd3\x27\x30\x86\x47\xf8\xfe\xa0\x19\xeb\x49\xc4\x3d\x9a\x7f\xb9\xab\x0f\x61\x02\x93\x21\xbf\x03\x00\x00\xff\xff\x5d\x2a\x5a\x5e\xfb\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationrocketchatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationrocketchatGraphql,
		"_lgraphql/notifications/updateNotificationRocketChat.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationrocketchatGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationrocketchatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationRocketChat.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationslackGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x4e\x4b\x0a\xc2\x30\x10\x5d\x67\x4e\x31\x05\x17\xf6\x0a\xbd\x81\x1b\x11\x8a\x07\x18\xd3\x68\x87\xb6\x93\x20\x13\x5c\x84\xdc\x5d\xd2\x48\x50\xf0\xed\xde\x67\xde\x9b\x2d\x2a\x29\x7b\xc1\x23\x20\x22\x1e\x84\x36\x37\xe0\xa8\x4f\x96\x47\x57\xa5\x40\x6a\xe7\x01\xaf\x61\x22\x75\x67\xaf\x7c\x67\xbb\xdf\x8c\x2b\xd9\xe5\x52\xdc\x93\x84\xa8\x1d\xf4\x98\xc0\xc4\xff\xb9\xda\x5f\xc0\x25\x3c\xa4\xc6\x0b\xea\xec\xbe\xfe\xa3\x7f\xb6\xeb\x0f\xcd\xc9\x60\xfa\x04\xc6\xf0\x04\xdf\x05\x8d\xd8\x99\x44\xdc\xda\xf8\xcb\xdd\x66\xef\x17\x30\x19\xf2\x3b\x00\x00\xff\xff\xd4\x02\x8b\x34\xf1\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationslackGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationslackGraphql,
		"_lgraphql/notifications/updateNotificationSlack.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationslackGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationslackGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationSlack.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlNotificationsUpdatenotificationwebhookGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xd0\xe0\x52\x50\x50\x50\x50\xc9\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x57\x84\x08\x15\x24\x96\x24\x67\x58\x29\x84\x16\xa4\x24\x96\xa4\xfa\xe5\x97\x64\xa6\x65\x26\x83\xf5\x84\xa7\x26\x65\xe4\xe7\x67\x07\x80\xe4\x3d\xf3\x0a\x4a\x4b\x14\xb9\x34\x15\xaa\xb9\x38\x4b\x71\xa9\x84\xd8\x01\x02\x99\x20\xe5\x56\xd5\x70\x3e\x08\x40\xac\x06\xbb\x00\x45\x1c\x6a\x3f\xc4\x1d\x70\x99\x5a\x2e\x4e\xcd\x6a\x2e\x4e\xce\xcc\x14\x2e\x64\x03\xe0\x9c\x72\x88\x95\x5c\x9c\xb5\x5c\xb5\x80\x00\x00\x00\xff\xff\x93\xb3\xe1\xf1\xe5\x00\x00\x00")

func _lgraphqlNotificationsUpdatenotificationwebhookGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlNotificationsUpdatenotificationwebhookGraphql,
		"_lgraphql/notifications/updateNotificationWebhook.graphql",
	)
}

func _lgraphqlNotificationsUpdatenotificationwebhookGraphql() (*asset, error) {
	bytes, err := _lgraphqlNotificationsUpdatenotificationwebhookGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/notifications/updateNotificationWebhook.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"_lgraphql/addEnvVariable.graphql":                                     _lgraphqlAddenvvariableGraphql,
	"_lgraphql/addNotificationEmail.graphql":                               _lgraphqlAddnotificationemailGraphql,
	"_lgraphql/addNotificationMicrosoftTeams.graphql":                      _lgraphqlAddnotificationmicrosoftteamsGraphql,
	"_lgraphql/addNotificationRocketChat.graphql":                          _lgraphqlAddnotificationrocketchatGraphql,
	"_lgraphql/addNotificationSlack.graphql":                               _lgraphqlAddnotificationslackGraphql,
	"_lgraphql/addNotificationToProject.graphql":                           _lgraphqlAddnotificationtoprojectGraphql,
	"_lgraphql/lagoonSchema.graphql":                                       _lgraphqlLagoonschemaGraphql,
	"_lgraphql/lagoonVersion.graphql":                                      _lgraphqlLagoonversionGraphql,
	"_lgraphql/sshEndpointsByProject.graphql":                              _lgraphqlSshendpointsbyprojectGraphql,
	"_lgraphql/taskByID.graphql":                                           _lgraphqlTaskbyidGraphql,
	"_lgraphql/deployments/deployEnvironmentBranch.graphql":                _lgraphqlDeploymentsDeployenvironmentbranchGraphql,
	"_lgraphql/deployments/deployEnvironmentLatest.graphql":                _lgraphqlDeploymentsDeployenvironmentlatestGraphql,
	"_lgraphql/deployments/deployEnvironmentPromote.graphql":               _lgraphqlDeploymentsDeployenvironmentpromoteGraphql,
	"_lgraphql/deployments/deployEnvironmentPullrequest.graphql":           _lgraphqlDeploymentsDeployenvironmentpullrequestGraphql,
	"_lgraphql/deployments/deploymentByName.graphql":                       _lgraphqlDeploymentsDeploymentbynameGraphql,
	"_lgraphql/deployments/deploymentByNameWithLog.graphql":                _lgraphqlDeploymentsDeploymentbynamewithlogGraphql,
	"_lgraphql/deployments/deploymentByRemoteID.graphql":                   _lgraphqlDeploymentsDeploymentbyremoteidGraphql,
	"_lgraphql/deployments/getDeploymentsByBulkID.graphql":                 _lgraphqlDeploymentsGetdeploymentsbybulkidGraphql,
	"_lgraphql/deployments/getDeploymentsForEnvironment.graphql":           _lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql,
	"_lgraphql/deployments/updateDeployment.graphql":                       _lgraphqlDeploymentsUpdatedeploymentGraphql,
	"_lgraphql/deploytargets/addDeployTarget.graphql":                      _lgraphqlDeploytargetsAdddeploytargetGraphql,
	"_lgraphql/deploytargets/addDeployTargetToOrganization.graphql":        _lgraphqlDeploytargetsAdddeploytargettoorganizationGraphql,
	"_lgraphql/deploytargets/deleteDeployTarget.graphql":                   _lgraphqlDeploytargetsDeletedeploytargetGraphql,
	"_lgraphql/deploytargets/deployTargetsByOrganizationId.graphql":        _lgraphqlDeploytargetsDeploytargetsbyorganizationidGraphql,
	"_lgraphql/deploytargets/deployTargetsByOrganizationName.graphql":      _lgraphqlDeploytargetsDeploytargetsbyorganizationnameGraphql,
	"_lgraphql/deploytargets/listDeployTargets.graphql":                    _lgraphqlDeploytargetsListdeploytargetsGraphql,
	"_lgraphql/deploytargets/removeDeployTargetFromOrganization.graphql":   _lgraphqlDeploytargetsRemovedeploytargetfromorganizationGraphql,
	"_lgraphql/deploytargets/updateDeployTarget.graphql":                   _lgraphqlDeploytargetsUpdatedeploytargetGraphql,
	"_lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql":          _lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql,
	"_lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql":       _lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql,
	"_lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql": _lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql,
	"_lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql":       _lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql,
	"_lgraphql/projects/addProject.graphql":                                _lgraphqlProjectsAddprojectGraphql,
	"_lgraphql/projects/minimalProjectByName.graphql":                      _lgraphqlProjectsMinimalprojectbynameGraphql,
	"_lgraphql/projects/projectByName.graphql":                             _lgraphqlProjectsProjectbynameGraphql,
	"_lgraphql/projects/projectByNameExtended.graphql":                     _lgraphqlProjectsProjectbynameextendedGraphql,
	"_lgraphql/projects/projectByNameMetadata.graphql":                     _lgraphqlProjectsProjectbynamemetadataGraphql,
	"_lgraphql/projects/projectGroups.graphql":                             _lgraphqlProjectsProjectgroupsGraphql,
	"_lgraphql/projects/projectsByMetadata.graphql":                        _lgraphqlProjectsProjectsbymetadataGraphql,
	"_lgraphql/projects/projectsByOrganizationId.graphql":                  _lgraphqlProjectsProjectsbyorganizationidGraphql,
	"_lgraphql/projects/removeProjectFromOrganization.graphql":             _lgraphqlProjectsRemoveprojectfromorganizationGraphql,
	"_lgraphql/projects/removeProjectMetadataByKey.graphql":                _lgraphqlProjectsRemoveprojectmetadatabykeyGraphql,
	"_lgraphql/projects/updateProject.graphql":                             _lgraphqlProjectsUpdateprojectGraphql,
	"_lgraphql/projects/updateProjectMetadata.graphql":                     _lgraphqlProjectsUpdateprojectmetadataGraphql,
	"_lgraphql/environments/addOrUpdateEnvironment.graphql":                _lgraphqlEnvironmentsAddorupdateenvironmentGraphql,
	"_lgraphql/environments/addOrUpdateEnvironmentStorage.graphql":         _lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql,
	"_lgraphql/environments/addRestore.graphql":                            _lgraphqlEnvironmentsAddrestoreGraphql,
	"_lgraphql/environments/backupsForEnvironmentByName.graphql":           _lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql,
	"_lgraphql/environments/deleteEnvironment.graphql":                     _lgraphqlEnvironmentsDeleteenvironmentGraphql,
	"_lgraphql/environments/environmentByName.graphql":                     _lgraphqlEnvironmentsEnvironmentbynameGraphql,
	"_lgraphql/environments/environmentByNamespace.graphql":                _lgraphqlEnvironmentsEnvironmentbynamespaceGraphql,
	"_lgraphql/environments/environmentsByProjectName.graphql":             _lgraphqlEnvironmentsEnvironmentsbyprojectnameGraphql,
	"_lgraphql/environments/setEnvironmentServices.graphql":                _lgraphqlEnvironmentsSetenvironmentservicesGraphql,
	"_lgraphql/environments/sshEndpointByEnvironment.graphql":              _lgraphqlEnvironmentsSshendpointbyenvironmentGraphql,
	"_lgraphql/environments/updateEnvironment.graphql":                     _lgraphqlEnvironmentsUpdateenvironmentGraphql,
	"_lgraphql/tasks/switchActiveStandby.graphql":                          _lgraphqlTasksSwitchactivestandbyGraphql,
	"_lgraphql/tasks/updateTask.graphql":                                   _lgraphqlTasksUpdatetaskGraphql,
	"_lgraphql/tasks/uploadFilesForTask.graphql":                           _lgraphqlTasksUploadfilesfortaskGraphql,
	"_lgraphql/usergroups/addGroup.graphql":                                _lgraphqlUsergroupsAddgroupGraphql,
	"_lgraphql/usergroups/addGroupToOrganization.graphql":                  _lgraphqlUsergroupsAddgrouptoorganizationGraphql,
	"_lgraphql/usergroups/addGroupsToProject.graphql":                      _lgraphqlUsergroupsAddgroupstoprojectGraphql,
	"_lgraphql/usergroups/addSshKey.graphql":                               _lgraphqlUsergroupsAddsshkeyGraphql,
	"_lgraphql/usergroups/addUser.graphql":                                 _lgraphqlUsergroupsAdduserGraphql,
	"_lgraphql/usergroups/addUserToGroup.graphql":                          _lgraphqlUsergroupsAddusertogroupGraphql,
	"_lgraphql/usergroups/addUserToOrganization.graphql":                   _lgraphqlUsergroupsAddusertoorganizationGraphql,
	"_lgraphql/usergroups/allGroupMembers.graphql":                         _lgraphqlUsergroupsAllgroupmembersGraphql,
	"_lgraphql/usergroups/allGroupMembersWithKeys.graphql":                 _lgraphqlUsergroupsAllgroupmemberswithkeysGraphql,
	"_lgraphql/usergroups/allUsers.graphql":                                _lgraphqlUsergroupsAllusersGraphql,
	"_lgraphql/usergroups/getGroupMembers.graphql":                         _lgraphqlUsergroupsGetgroupmembersGraphql,
	"_lgraphql/usergroups/groupsByOrganizationId.graphql":                  _lgraphqlUsergroupsGroupsbyorganizationidGraphql,
	"_lgraphql/usergroups/me.graphql":                                      _lgraphqlUsergroupsMeGraphql,
	"_lgraphql/usergroups/removeGroupsFromProject.graphql":                 _lgraphqlUsergroupsRemovegroupsfromprojectGraphql,
	"_lgraphql/usergroups/removeSSHKeyById.graphql":                        _lgraphqlUsergroupsRemovesshkeybyidGraphql,
	"_lgraphql/usergroups/removeUserFromGroup.graphql":                     _lgraphqlUsergroupsRemoveuserfromgroupGraphql,
	"_lgraphql/usergroups/removeUserFromOrganization.graphql":              _lgraphqlUsergroupsRemoveuserfromorganizationGraphql,
	"_lgraphql/usergroups/userByEmail.graphql":                             _lgraphqlUsergroupsUserbyemailGraphql,
	"_lgraphql/usergroups/userByEmailSSHKeys.graphql":                      _lgraphqlUsergroupsUserbyemailsshkeysGraphql,
	"_lgraphql/usergroups/userBySSHFingerprint.graphql":                    _lgraphqlUsergroupsUserbysshfingerprintGraphql,
	"_lgraphql/usergroups/userBySSHKey.graphql":                            _lgraphqlUsergroupsUserbysshkeyGraphql,
	"_lgraphql/usergroups/userCanSSHToEnvironment.graphql":                 _lgraphqlUsergroupsUsercansshtoenvironmentGraphql,
	"_lgraphql/usergroups/usersByOrganization.graphql":                     _lgraphqlUsergroupsUsersbyorganizationGraphql,
	"_lgraphql/usergroups/usersByOrganizationName.graphql":                 _lgraphqlUsergroupsUsersbyorganizationnameGraphql,
	"_lgraphql/organizations/addOrganization.graphql":                      _lgraphqlOrganizationsAddorganizationGraphql,
	"_lgraphql/organizations/allOrganizations.graphql":                     _lgraphqlOrganizationsAllorganizationsGraphql,
	"_lgraphql/organizations/deleteOrganization.graphql":                   _lgraphqlOrganizationsDeleteorganizationGraphql,
	"_lgraphql/organizations/organizationById.graphql":                     _lgraphqlOrganizationsOrganizationbyidGraphql,
	"_lgraphql/organizations/organizationByName.graphql":                   _lgraphqlOrganizationsOrganizationbynameGraphql,
	"_lgraphql/organizations/updateOrganization.graphql":                   _lgraphqlOrganizationsUpdateorganizationGraphql,
	"_lgraphql/notifications/addNotificationEmail.graphql":                 _lgraphqlNotificationsAddnotificationemailGraphql,
	"_lgraphql/notifications/addNotificationMicrosoftTeams.graphql":        _lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/addNotificationRocketChat.graphql":            _lgraphqlNotificationsAddnotificationrocketchatGraphql,
	"_lgraphql/notifications/addNotificationSlack.graphql":                 _lgraphqlNotificationsAddnotificationslackGraphql,
	"_lgraphql/notifications/addNotificationToProject.graphql":             _lgraphqlNotificationsAddnotificationtoprojectGraphql,
	"_lgraphql/notifications/addNotificationWebhook.graphql":               _lgraphqlNotificationsAddnotificationwebhookGraphql,
	"_lgraphql/notifications/deleteNotificationEmail.graphql":              _lgraphqlNotificationsDeletenotificationemailGraphql,
	"_lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql":     _lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/deleteNotificationRocketChat.graphql":         _lgraphqlNotificationsDeletenotificationrocketchatGraphql,
	"_lgraphql/notifications/deleteNotificationSlack.graphql":              _lgraphqlNotificationsDeletenotificationslackGraphql,
	"_lgraphql/notifications/deleteNotificationWebhook.graphql":            _lgraphqlNotificationsDeletenotificationwebhookGraphql,
	"_lgraphql/notifications/listAllNotificationEmail.graphql":             _lgraphqlNotificationsListallnotificationemailGraphql,
	"_lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql":    _lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/listAllNotificationRocketChat.graphql":        _lgraphqlNotificationsListallnotificationrocketchatGraphql,
	"_lgraphql/notifications/listAllNotificationSlack.graphql":             _lgraphqlNotificationsListallnotificationslackGraphql,
	"_lgraphql/notifications/listAllNotificationWebhook.graphql":           _lgraphqlNotificationsListallnotificationwebhookGraphql,
	"_lgraphql/notifications/projectNotificationEmail.graphql":             _lgraphqlNotificationsProjectnotificationemailGraphql,
	"_lgraphql/notifications/projectNotificationMicrosoftTeams.graphql":    _lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/projectNotificationRocketChat.graphql":        _lgraphqlNotificationsProjectnotificationrocketchatGraphql,
	"_lgraphql/notifications/projectNotificationSlack.graphql":             _lgraphqlNotificationsProjectnotificationslackGraphql,
	"_lgraphql/notifications/projectNotificationWebhook.graphql":           _lgraphqlNotificationsProjectnotificationwebhookGraphql,
	"_lgraphql/notifications/removeNotificationFromProject.graphql":        _lgraphqlNotificationsRemovenotificationfromprojectGraphql,
	"_lgraphql/notifications/updateNotificationEmail.graphql":              _lgraphqlNotificationsUpdatenotificationemailGraphql,
	"_lgraphql/notifications/updateNotificationMicrosoftTeams.graphql":     _lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql,
	"_lgraphql/notifications/updateNotificationRocketChat.graphql":         _lgraphqlNotificationsUpdatenotificationrocketchatGraphql,
	"_lgraphql/notifications/updateNotificationSlack.graphql":              _lgraphqlNotificationsUpdatenotificationslackGraphql,
	"_lgraphql/notifications/updateNotificationWebhook.graphql":            _lgraphqlNotificationsUpdatenotificationwebhookGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"_lgraphql": &bintree{nil, map[string]*bintree{
		"addEnvVariable.graphql":                &bintree{_lgraphqlAddenvvariableGraphql, map[string]*bintree{}},
		"addNotificationEmail.graphql":          &bintree{_lgraphqlAddnotificationemailGraphql, map[string]*bintree{}},
		"addNotificationMicrosoftTeams.graphql": &bintree{_lgraphqlAddnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
		"addNotificationRocketChat.graphql":     &bintree{_lgraphqlAddnotificationrocketchatGraphql, map[string]*bintree{}},
		"addNotificationSlack.graphql":          &bintree{_lgraphqlAddnotificationslackGraphql, map[string]*bintree{}},
		"addNotificationToProject.graphql":      &bintree{_lgraphqlAddnotificationtoprojectGraphql, map[string]*bintree{}},
		"deployments": &bintree{nil, map[string]*bintree{
			"deployEnvironmentBranch.graphql":      &bintree{_lgraphqlDeploymentsDeployenvironmentbranchGraphql, map[string]*bintree{}},
			"deployEnvironmentLatest.graphql":      &bintree{_lgraphqlDeploymentsDeployenvironmentlatestGraphql, map[string]*bintree{}},
			"deployEnvironmentPromote.graphql":     &bintree{_lgraphqlDeploymentsDeployenvironmentpromoteGraphql, map[string]*bintree{}},
			"deployEnvironmentPullrequest.graphql": &bintree{_lgraphqlDeploymentsDeployenvironmentpullrequestGraphql, map[string]*bintree{}},
			"deploymentByName.graphql":             &bintree{_lgraphqlDeploymentsDeploymentbynameGraphql, map[string]*bintree{}},
			"deploymentByNameWithLog.graphql":      &bintree{_lgraphqlDeploymentsDeploymentbynamewithlogGraphql, map[string]*bintree{}},
			"deploymentByRemoteID.graphql":         &bintree{_lgraphqlDeploymentsDeploymentbyremoteidGraphql, map[string]*bintree{}},
			"getDeploymentsByBulkID.graphql":       &bintree{_lgraphqlDeploymentsGetdeploymentsbybulkidGraphql, map[string]*bintree{}},
			"getDeploymentsForEnvironment.graphql": &bintree{_lgraphqlDeploymentsGetdeploymentsforenvironmentGraphql, map[string]*bintree{}},
			"updateDeployment.graphql":             &bintree{_lgraphqlDeploymentsUpdatedeploymentGraphql, map[string]*bintree{}},
		}},
		"deploytargetconfigs": &bintree{nil, map[string]*bintree{
			"addDeployTargetConfig.graphql":          &bintree{_lgraphqlDeploytargetconfigsAdddeploytargetconfigGraphql, map[string]*bintree{}},
			"deleteDeployTargetConfig.graphql":       &bintree{_lgraphqlDeploytargetconfigsDeletedeploytargetconfigGraphql, map[string]*bintree{}},
			"deployTargetConfigsByProjectId.graphql": &bintree{_lgraphqlDeploytargetconfigsDeploytargetconfigsbyprojectidGraphql, map[string]*bintree{}},
			"updateDeployTargetConfig.graphql":       &bintree{_lgraphqlDeploytargetconfigsUpdatedeploytargetconfigGraphql, map[string]*bintree{}},
		}},
		"deploytargets": &bintree{nil, map[string]*bintree{
			"addDeployTarget.graphql":                    &bintree{_lgraphqlDeploytargetsAdddeploytargetGraphql, map[string]*bintree{}},
			"addDeployTargetToOrganization.graphql":      &bintree{_lgraphqlDeploytargetsAdddeploytargettoorganizationGraphql, map[string]*bintree{}},
			"deleteDeployTarget.graphql":                 &bintree{_lgraphqlDeploytargetsDeletedeploytargetGraphql, map[string]*bintree{}},
			"deployTargetsByOrganizationId.graphql":      &bintree{_lgraphqlDeploytargetsDeploytargetsbyorganizationidGraphql, map[string]*bintree{}},
			"deployTargetsByOrganizationName.graphql":    &bintree{_lgraphqlDeploytargetsDeploytargetsbyorganizationnameGraphql, map[string]*bintree{}},
			"listDeployTargets.graphql":                  &bintree{_lgraphqlDeploytargetsListdeploytargetsGraphql, map[string]*bintree{}},
			"removeDeployTargetFromOrganization.graphql": &bintree{_lgraphqlDeploytargetsRemovedeploytargetfromorganizationGraphql, map[string]*bintree{}},
			"updateDeployTarget.graphql":                 &bintree{_lgraphqlDeploytargetsUpdatedeploytargetGraphql, map[string]*bintree{}},
		}},
		"environments": &bintree{nil, map[string]*bintree{
			"addOrUpdateEnvironment.graphql":        &bintree{_lgraphqlEnvironmentsAddorupdateenvironmentGraphql, map[string]*bintree{}},
			"addOrUpdateEnvironmentStorage.graphql": &bintree{_lgraphqlEnvironmentsAddorupdateenvironmentstorageGraphql, map[string]*bintree{}},
			"addRestore.graphql":                    &bintree{_lgraphqlEnvironmentsAddrestoreGraphql, map[string]*bintree{}},
			"backupsForEnvironmentByName.graphql":   &bintree{_lgraphqlEnvironmentsBackupsforenvironmentbynameGraphql, map[string]*bintree{}},
			"deleteEnvironment.graphql":             &bintree{_lgraphqlEnvironmentsDeleteenvironmentGraphql, map[string]*bintree{}},
			"environmentByName.graphql":             &bintree{_lgraphqlEnvironmentsEnvironmentbynameGraphql, map[string]*bintree{}},
			"environmentByNamespace.graphql":        &bintree{_lgraphqlEnvironmentsEnvironmentbynamespaceGraphql, map[string]*bintree{}},
			"environmentsByProjectName.graphql":     &bintree{_lgraphqlEnvironmentsEnvironmentsbyprojectnameGraphql, map[string]*bintree{}},
			"setEnvironmentServices.graphql":        &bintree{_lgraphqlEnvironmentsSetenvironmentservicesGraphql, map[string]*bintree{}},
			"sshEndpointByEnvironment.graphql":      &bintree{_lgraphqlEnvironmentsSshendpointbyenvironmentGraphql, map[string]*bintree{}},
			"updateEnvironment.graphql":             &bintree{_lgraphqlEnvironmentsUpdateenvironmentGraphql, map[string]*bintree{}},
		}},
		"lagoonSchema.graphql":  &bintree{_lgraphqlLagoonschemaGraphql, map[string]*bintree{}},
		"lagoonVersion.graphql": &bintree{_lgraphqlLagoonversionGraphql, map[string]*bintree{}},
		"notifications": &bintree{nil, map[string]*bintree{
			"addNotificationEmail.graphql":              &bintree{_lgraphqlNotificationsAddnotificationemailGraphql, map[string]*bintree{}},
			"addNotificationMicrosoftTeams.graphql":     &bintree{_lgraphqlNotificationsAddnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"addNotificationRocketChat.graphql":         &bintree{_lgraphqlNotificationsAddnotificationrocketchatGraphql, map[string]*bintree{}},
			"addNotificationSlack.graphql":              &bintree{_lgraphqlNotificationsAddnotificationslackGraphql, map[string]*bintree{}},
			"addNotificationToProject.graphql":          &bintree{_lgraphqlNotificationsAddnotificationtoprojectGraphql, map[string]*bintree{}},
			"addNotificationWebhook.graphql":            &bintree{_lgraphqlNotificationsAddnotificationwebhookGraphql, map[string]*bintree{}},
			"deleteNotificationEmail.graphql":           &bintree{_lgraphqlNotificationsDeletenotificationemailGraphql, map[string]*bintree{}},
			"deleteNotificationMicrosoftTeams.graphql":  &bintree{_lgraphqlNotificationsDeletenotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"deleteNotificationRocketChat.graphql":      &bintree{_lgraphqlNotificationsDeletenotificationrocketchatGraphql, map[string]*bintree{}},
			"deleteNotificationSlack.graphql":           &bintree{_lgraphqlNotificationsDeletenotificationslackGraphql, map[string]*bintree{}},
			"deleteNotificationWebhook.graphql":         &bintree{_lgraphqlNotificationsDeletenotificationwebhookGraphql, map[string]*bintree{}},
			"listAllNotificationEmail.graphql":          &bintree{_lgraphqlNotificationsListallnotificationemailGraphql, map[string]*bintree{}},
			"listAllNotificationMicrosoftTeams.graphql": &bintree{_lgraphqlNotificationsListallnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"listAllNotificationRocketChat.graphql":     &bintree{_lgraphqlNotificationsListallnotificationrocketchatGraphql, map[string]*bintree{}},
			"listAllNotificationSlack.graphql":          &bintree{_lgraphqlNotificationsListallnotificationslackGraphql, map[string]*bintree{}},
			"listAllNotificationWebhook.graphql":        &bintree{_lgraphqlNotificationsListallnotificationwebhookGraphql, map[string]*bintree{}},
			"projectNotificationEmail.graphql":          &bintree{_lgraphqlNotificationsProjectnotificationemailGraphql, map[string]*bintree{}},
			"projectNotificationMicrosoftTeams.graphql": &bintree{_lgraphqlNotificationsProjectnotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"projectNotificationRocketChat.graphql":     &bintree{_lgraphqlNotificationsProjectnotificationrocketchatGraphql, map[string]*bintree{}},
			"projectNotificationSlack.graphql":          &bintree{_lgraphqlNotificationsProjectnotificationslackGraphql, map[string]*bintree{}},
			"projectNotificationWebhook.graphql":        &bintree{_lgraphqlNotificationsProjectnotificationwebhookGraphql, map[string]*bintree{}},
			"removeNotificationFromProject.graphql":     &bintree{_lgraphqlNotificationsRemovenotificationfromprojectGraphql, map[string]*bintree{}},
			"updateNotificationEmail.graphql":           &bintree{_lgraphqlNotificationsUpdatenotificationemailGraphql, map[string]*bintree{}},
			"updateNotificationMicrosoftTeams.graphql":  &bintree{_lgraphqlNotificationsUpdatenotificationmicrosoftteamsGraphql, map[string]*bintree{}},
			"updateNotificationRocketChat.graphql":      &bintree{_lgraphqlNotificationsUpdatenotificationrocketchatGraphql, map[string]*bintree{}},
			"updateNotificationSlack.graphql":           &bintree{_lgraphqlNotificationsUpdatenotificationslackGraphql, map[string]*bintree{}},
			"updateNotificationWebhook.graphql":         &bintree{_lgraphqlNotificationsUpdatenotificationwebhookGraphql, map[string]*bintree{}},
		}},
		"organizations": &bintree{nil, map[string]*bintree{
			"addOrganization.graphql":    &bintree{_lgraphqlOrganizationsAddorganizationGraphql, map[string]*bintree{}},
			"allOrganizations.graphql":   &bintree{_lgraphqlOrganizationsAllorganizationsGraphql, map[string]*bintree{}},
			"deleteOrganization.graphql": &bintree{_lgraphqlOrganizationsDeleteorganizationGraphql, map[string]*bintree{}},
			"organizationById.graphql":   &bintree{_lgraphqlOrganizationsOrganizationbyidGraphql, map[string]*bintree{}},
			"organizationByName.graphql": &bintree{_lgraphqlOrganizationsOrganizationbynameGraphql, map[string]*bintree{}},
			"updateOrganization.graphql": &bintree{_lgraphqlOrganizationsUpdateorganizationGraphql, map[string]*bintree{}},
		}},
		"projects": &bintree{nil, map[string]*bintree{
			"addProject.graphql":                    &bintree{_lgraphqlProjectsAddprojectGraphql, map[string]*bintree{}},
			"minimalProjectByName.graphql":          &bintree{_lgraphqlProjectsMinimalprojectbynameGraphql, map[string]*bintree{}},
			"projectByName.graphql":                 &bintree{_lgraphqlProjectsProjectbynameGraphql, map[string]*bintree{}},
			"projectByNameExtended.graphql":         &bintree{_lgraphqlProjectsProjectbynameextendedGraphql, map[string]*bintree{}},
			"projectByNameMetadata.graphql":         &bintree{_lgraphqlProjectsProjectbynamemetadataGraphql, map[string]*bintree{}},
			"projectGroups.graphql":                 &bintree{_lgraphqlProjectsProjectgroupsGraphql, map[string]*bintree{}},
			"projectsByMetadata.graphql":            &bintree{_lgraphqlProjectsProjectsbymetadataGraphql, map[string]*bintree{}},
			"projectsByOrganizationId.graphql":      &bintree{_lgraphqlProjectsProjectsbyorganizationidGraphql, map[string]*bintree{}},
			"removeProjectFromOrganization.graphql": &bintree{_lgraphqlProjectsRemoveprojectfromorganizationGraphql, map[string]*bintree{}},
			"removeProjectMetadataByKey.graphql":    &bintree{_lgraphqlProjectsRemoveprojectmetadatabykeyGraphql, map[string]*bintree{}},
			"updateProject.graphql":                 &bintree{_lgraphqlProjectsUpdateprojectGraphql, map[string]*bintree{}},
			"updateProjectMetadata.graphql":         &bintree{_lgraphqlProjectsUpdateprojectmetadataGraphql, map[string]*bintree{}},
		}},
		"sshEndpointsByProject.graphql": &bintree{_lgraphqlSshendpointsbyprojectGraphql, map[string]*bintree{}},
		"taskByID.graphql":              &bintree{_lgraphqlTaskbyidGraphql, map[string]*bintree{}},
		"tasks": &bintree{nil, map[string]*bintree{
			"switchActiveStandby.graphql": &bintree{_lgraphqlTasksSwitchactivestandbyGraphql, map[string]*bintree{}},
			"updateTask.graphql":          &bintree{_lgraphqlTasksUpdatetaskGraphql, map[string]*bintree{}},
			"uploadFilesForTask.graphql":  &bintree{_lgraphqlTasksUploadfilesfortaskGraphql, map[string]*bintree{}},
		}},
		"usergroups": &bintree{nil, map[string]*bintree{
			"addGroup.graphql":                   &bintree{_lgraphqlUsergroupsAddgroupGraphql, map[string]*bintree{}},
			"addGroupToOrganization.graphql":     &bintree{_lgraphqlUsergroupsAddgrouptoorganizationGraphql, map[string]*bintree{}},
			"addGroupsToProject.graphql":         &bintree{_lgraphqlUsergroupsAddgroupstoprojectGraphql, map[string]*bintree{}},
			"addSshKey.graphql":                  &bintree{_lgraphqlUsergroupsAddsshkeyGraphql, map[string]*bintree{}},
			"addUser.graphql":                    &bintree{_lgraphqlUsergroupsAdduserGraphql, map[string]*bintree{}},
			"addUserToGroup.graphql":             &bintree{_lgraphqlUsergroupsAddusertogroupGraphql, map[string]*bintree{}},
			"addUserToOrganization.graphql":      &bintree{_lgraphqlUsergroupsAddusertoorganizationGraphql, map[string]*bintree{}},
			"allGroupMembers.graphql":            &bintree{_lgraphqlUsergroupsAllgroupmembersGraphql, map[string]*bintree{}},
			"allGroupMembersWithKeys.graphql":    &bintree{_lgraphqlUsergroupsAllgroupmemberswithkeysGraphql, map[string]*bintree{}},
			"allUsers.graphql":                   &bintree{_lgraphqlUsergroupsAllusersGraphql, map[string]*bintree{}},
			"getGroupMembers.graphql":            &bintree{_lgraphqlUsergroupsGetgroupmembersGraphql, map[string]*bintree{}},
			"groupsByOrganizationId.graphql":     &bintree{_lgraphqlUsergroupsGroupsbyorganizationidGraphql, map[string]*bintree{}},
			"me.graphql":                         &bintree{_lgraphqlUsergroupsMeGraphql, map[string]*bintree{}},
			"removeGroupsFromProject.graphql":    &bintree{_lgraphqlUsergroupsRemovegroupsfromprojectGraphql, map[string]*bintree{}},
			"removeSSHKeyById.graphql":           &bintree{_lgraphqlUsergroupsRemovesshkeybyidGraphql, map[string]*bintree{}},
			"removeUserFromGroup.graphql":        &bintree{_lgraphqlUsergroupsRemoveuserfromgroupGraphql, map[string]*bintree{}},
			"removeUserFromOrganization.graphql": &bintree{_lgraphqlUsergroupsRemoveuserfromorganizationGraphql, map[string]*bintree{}},
			"userByEmail.graphql":                &bintree{_lgraphqlUsergroupsUserbyemailGraphql, map[string]*bintree{}},
			"userByEmailSSHKeys.graphql":         &bintree{_lgraphqlUsergroupsUserbyemailsshkeysGraphql, map[string]*bintree{}},
			"userBySSHFingerprint.graphql":       &bintree{_lgraphqlUsergroupsUserbysshfingerprintGraphql, map[string]*bintree{}},
			"userBySSHKey.graphql":               &bintree{_lgraphqlUsergroupsUserbysshkeyGraphql, map[string]*bintree{}},
			"userCanSSHToEnvironment.graphql":    &bintree{_lgraphqlUsergroupsUsercansshtoenvironmentGraphql, map[string]*bintree{}},
			"usersByOrganization.graphql":        &bintree{_lgraphqlUsergroupsUsersbyorganizationGraphql, map[string]*bintree{}},
			"usersByOrganizationName.graphql":    &bintree{_lgraphqlUsergroupsUsersbyorganizationnameGraphql, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
