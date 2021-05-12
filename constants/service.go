package constants

type ImportService struct {
	DefaultVersion string              `json:"default_version"` // 默认版本号
	Ports          []ImportServicePort `json:"ports"`           // 端口和协议
}

type ImportServicePort struct {
	Number   int32  `json:"number"`   // 监听端口
	Protocol string `json:"protocol"` // 协议，如TCP、UDP
}

func NewImportService(version string, ports []ImportServicePort) *ImportService {
	v := DSMSVCDefaultVersion
	if version != "" {
		v = version
	}
	return &ImportService{
		DefaultVersion: v,
		Ports:          ports,
	}
}

// avoid unused warnings
var _ = NewImportService("", nil)
