package entities

type Phone string

func (p Phone) Uniform() Phone {
	if len(p) > 3 && p[:3] == "+98" {
		return Phone("0" + p[3:])
	}
	return p
}

func (p Phone) Validate() bool {
	if len(p) != 11 && p[:2] != "09" {
		return false
	}
	return true
}

type CredentialMethod string

const (
	CredentialMethodPhone CredentialMethod = "phone"
	CredentialMethodEmail CredentialMethod = "email"
	CredentialMethodOAuth CredentialMethod = "oauth"
)

type CredentialIdentifier string
