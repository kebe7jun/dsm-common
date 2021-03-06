package constants

const (
	DSMAppTypeLabel         = "dsm.daocloud.io/type"
	DSMDeployTypeLabel      = "dsm.daocloud.io/deploy-type"
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
	DSMSVCTypeDubbo       = "dubbo"

	DSMDeployTypeVM        = "vm"
	DSMDeployTypeContainer = "container"

	ProtocolTCP   Protocol = "tcp"
	ProtocolHTTP  Protocol = "http"
	ProtocolHTTP2 Protocol = "http2"
	ProtocolHTTPS Protocol = "https"

	// DSMSVCDefaultVersion defines the default version of DSM imported service
	// when users specify none service version
	DSMSVCDefaultVersion = "v1"
)

var (
	// avoid unused warnings
	_ = DSMAppTypeLabel
	_ = DSMDeployTypeLabel
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
	_ = DSMSVCTypeDubbo

	_ = DSMDeployTypeVM
	_ = DSMDeployTypeContainer

	_ = ProtocolTCP
	_ = ProtocolHTTP
	_ = ProtocolHTTP2
	_ = ProtocolHTTPS
)
