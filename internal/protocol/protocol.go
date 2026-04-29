package protocol

import "encoding/json"

const (
	RoleAgent  = "agent"
	RoleClient = "client"

	TypeHello   = "hello"
	TypeRequest = "request"
	TypeResponse = "response"
	TypeError   = "error"

	CommandStatus  = "status"
	CommandListDir = "list_dir"
)

type Message struct {
	Type      string          `json:"type"`
	Role      string          `json:"role,omitempty"`
	DeviceID  string          `json:"device_id,omitempty"`
	RequestID string          `json:"request_id,omitempty"`
	Command   string          `json:"command,omitempty"`
	Payload   json.RawMessage `json:"payload,omitempty"`
	Error     string          `json:"error,omitempty"`
}

type HelloPayload struct {
	Token string `json:"token"`
}

type StatusResponse struct {
	Hostname string `json:"hostname"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	User     string `json:"user"`
}

type ListDirRequest struct {
	Path string `json:"path"`
}

type FileEntry struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	IsDir bool   `json:"is_dir"`
	Size  int64  `json:"size"`
}

type ListDirResponse struct {
	Path    string      `json:"path"`
	Entries []FileEntry `json:"entries"`
}

func MustJSON(v any) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
