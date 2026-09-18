package resource_remote_storage_location

import (
	"encoding/json"
	"fmt"
)

type customNDFCRemoteStorageLocationModel NDFCRemoteStorageLocationModel

type remoteStorageLocationAuthenticationPayload struct {
	NDFCAuthenticationValue
	AuthenticationType string `json:"type,omitempty"`
}

type remoteStorageLocationPayload struct {
	customNDFCRemoteStorageLocationModel
	StorageLocationType string                                      `json:"type,omitempty"`
	Port                *int64                                      `json:"port,omitempty"`
	AlertThreshold      *int64                                      `json:"alertThreshold,omitempty"`
	Limit               string                                      `json:"limit,omitempty"`
	ReadWrite           *bool                                       `json:"readWrite,omitempty"`
	Authentication      *remoteStorageLocationAuthenticationPayload `json:"authentication,omitempty"`
}

type remoteStorageLocationStatusPayload struct {
	HealthState string `json:"healthState,omitempty"`
	Message     string `json:"message,omitempty"`
}

type remoteStorageLocationWrappedResponse struct {
	Spec   *remoteStorageLocationPayload       `json:"spec"`
	Status *remoteStorageLocationStatusPayload `json:"status,omitempty"`
}

func hasNfsConfiguration(value NDFCNfsValue) bool {
	return value.Port != nil ||
		value.Limit != "" ||
		value.ReadWrite != nil ||
		value.AlertThreshold != nil
}

func hasScpSftpConfiguration(value NDFCScpSftpValue) bool {
	return value.Protocol != "" ||
		value.Port != nil ||
		value.AcceptHostKey ||
		value.Authentication.Username != "" ||
		value.Authentication.Password != "" ||
		value.Authentication.SshKey != "" ||
		value.Authentication.Passphrase != "" ||
		value.Authentication.IgnoreHostKeyValidation != nil
}

func (m *NDFCRemoteStorageLocationModel) UnmarshalJSON(data []byte) error {
	var wrapped remoteStorageLocationWrappedResponse
	if err := json.Unmarshal(data, &wrapped); err != nil {
		return err
	}

	payload := wrapped.Spec
	if payload == nil {
		payload = new(remoteStorageLocationPayload)
		if err := json.Unmarshal(data, payload); err != nil {
			return fmt.Errorf("unmarshal remote storage payload: %w", err)
		}
	}

	decoded := NDFCRemoteStorageLocationModel(payload.customNDFCRemoteStorageLocationModel)

	switch payload.StorageLocationType {
	case "nfs":
		decoded.Nfs = NDFCNfsValue{
			Port:           payload.Port,
			Limit:          payload.Limit,
			ReadWrite:      payload.ReadWrite,
			AlertThreshold: payload.AlertThreshold,
		}
	case "scp", "sftp":
		if payload.Authentication == nil {
			return fmt.Errorf("%s remote storage response is missing authentication", payload.StorageLocationType)
		}
		if payload.Authentication.AuthenticationType != "" &&
			payload.Authentication.AuthenticationType != "password" &&
			payload.Authentication.AuthenticationType != "key" {
			return fmt.Errorf("unsupported remote storage authentication type %q", payload.Authentication.AuthenticationType)
		}
		decoded.ScpSftp = NDFCScpSftpValue{
			Protocol:       payload.StorageLocationType,
			Port:           payload.Port,
			Authentication: payload.Authentication.NDFCAuthenticationValue,
		}
	default:
		return fmt.Errorf("unsupported or missing remote storage type %q", payload.StorageLocationType)
	}

	if wrapped.Status != nil {
		decoded.HealthState = wrapped.Status.HealthState
		decoded.HealthStateMessage = wrapped.Status.Message
	}

	*m = decoded
	return nil
}

func (m NDFCRemoteStorageLocationModel) MarshalJSON() ([]byte, error) {
	payload := remoteStorageLocationPayload{
		customNDFCRemoteStorageLocationModel: customNDFCRemoteStorageLocationModel(m),
	}

	if hasNfsConfiguration(m.Nfs) {
		payload.StorageLocationType = "nfs"
		payload.Port = m.Nfs.Port
		payload.Limit = m.Nfs.Limit
		payload.ReadWrite = m.Nfs.ReadWrite
		payload.AlertThreshold = m.Nfs.AlertThreshold
	} else if hasScpSftpConfiguration(m.ScpSftp) {
		hasPassword := m.ScpSftp.Authentication.Password != ""
		hasSSHKey := m.ScpSftp.Authentication.SshKey != ""
		if hasPassword == hasSSHKey {
			return nil, fmt.Errorf("SCP/SFTP payload requires either password or ssh_key authentication")
		}
		if m.ScpSftp.Authentication.Passphrase != "" && !hasSSHKey {
			return nil, fmt.Errorf("SCP/SFTP payload passphrase requires ssh_key authentication")
		}

		authenticationType := "password"
		if hasSSHKey {
			authenticationType = "key"
		}

		payload.StorageLocationType = m.ScpSftp.Protocol
		payload.Port = m.ScpSftp.Port
		payload.Authentication = &remoteStorageLocationAuthenticationPayload{
			NDFCAuthenticationValue: m.ScpSftp.Authentication,
			AuthenticationType:      authenticationType,
		}
	} else {
		return nil, fmt.Errorf("remote storage payload requires either an nfs or scp_sftp configuration")
	}

	return json.Marshal(payload)
}
