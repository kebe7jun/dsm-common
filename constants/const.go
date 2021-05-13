package constants

const (
	DSMAppTypeLabel         = "dsm.daocloud.io/type"
	DSMServiceImportedLabel = "dsm.daocloud.io/service-imported"
	// istio 1.5 use app as source_app label, and it is immutable.
	// so, we can olny use this label name.
	MeshServiceLabelName     = "app"               // label k8s service名称
	IstioGatewayService      = "mesh"              // service间请求的默认网关
	MeshGatewayHostLabelName = "mesh-gateway-host" // 网关访问域名

	DSMServiceImportTimeAnno         = "anno.dsm.daocloud.io/imported-at"
	DSMServiceImportStatusAnno       = "anno.dsm.daocloud.io/import-status"
	DSMServiceImportFailedReasonAnno = "anno.dsm.daocloud.io/import-failed-reason"
	DSMServiceImportConfigAnno       = "anno.dsm.daocloud.io/import-config"

	// DsmManageLabel
	DsmManageLabel = "dsm.daocloud.io/managed"

	// DXOriginOwnerAnnotationName store the origin owner name
	DXOriginOwnerAnnotationName = "sm.dx.daocloud.io/origin-name"

	// CustomWorkloadNameAnno defines the workload type and name
	// if this specified, controller will not try to create this workload.
	// but controller always try to fetch the status of the workload.
	// Format {workloadType}/{workloadName}
	CustomWorkloadNameAnno = "sm.dx.daocloud.io/custom-workload"

	DSMSVCTypeNormal      = "normal"
	DSMSVCTypeSpringCloud = "springcloud"
	DSMSVCTypeVM          = "vm"
	DSMSVCTypeDubbo       = "dubbo"

	ProtocolTCP   Protocol = "tcp"
	ProtocolHTTP  Protocol = "http"
	ProtocolHTTPS Protocol = "https"
	ProtocolUDP   Protocol = "udp"

	// DSMSVCDefaultVersion defines the default version of DSM imported service
	// when users specify none service version
	DSMSVCDefaultVersion = "v1"
)

var (
	// avoid unused warnings
	_ = DSMAppTypeLabel
	_ = DSMServiceImportedLabel
	_ = DSMServiceImportTimeAnno
	_ = DSMServiceImportStatusAnno
	_ = DSMServiceImportFailedReasonAnno
	_ = DSMServiceImportConfigAnno
	_ = DsmManageLabel
	_ = DXOriginOwnerAnnotationName
	_ = CustomWorkloadNameAnno

	_ = DSMSVCTypeNormal
	_ = DSMSVCTypeSpringCloud
	_ = DSMSVCTypeVM
	_ = DSMSVCTypeDubbo

	_ = ProtocolTCP
	_ = ProtocolHTTP
	_ = ProtocolHTTPS
	_ = ProtocolUDP
)
