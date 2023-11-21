package schema

// SSHKeyType .
type SSHKeyType string

// . .
const (
	SSHRsa      SSHKeyType = "SSH_RSA"
	SSHEd25519  SSHKeyType = "SSH_ED25519"
	SSHECDSA256 SSHKeyType = "ECDSA_SHA2_NISTP256"
	SSHECDSA384 SSHKeyType = "ECDSA_SHA2_NISTP384"
	SSHECDSA521 SSHKeyType = "ECDSA_SHA2_NISTP521"
)

// SSHKey is the basic SSH key information, used by both config and API data.
// @TODO: once Lagoon API returns proper TZ, fix up `Created` to time.Time.
type SSHKey struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	KeyValue       string     `json:"keyValue"`
	Created        string     `json:"created"`
	KeyType        SSHKeyType `json:"keyType"`
	KeyFingerprint string     `json:"keyFingerprint"`
}

// AddSSHKeyInput is based on the Lagoon API type.
type AddSSHKeyInput struct {
	SSHKey
	UserEmail string `json:"userEmail"`
}

type DeleteSSHKeyByIDInput struct {
	ID uint `json:"id"`
}
