package constants

const (
	DSMAppTypeLabel         = "dsm.daocloud.io/type"
	DSMServiceImportedLabel = "dsm.daocloud.io/service-imported"

	DSMServiceImportTimeAnno         = "anno.dsm.daocloud.io/imported-at"
	DSMServiceImportStatusAnno       = "anno.dsm.daocloud.io/import-status"
	DSMServiceImportFailedReasonAnno = "anno.dsm.daocloud.io/import-failed-reason"

	// DsmManageLabel
	DsmManageLabel = "dsm.daocloud.io/managed"

	// DXOriginOwnerAnnotationName store the origin owner name
	DXOriginOwnerAnnotationName = "sm.dx.daocloud.io/origin-name"

	// CustomWorkloadNameAnno defines the workload type and name
	// if this specified, controller will not try to create this workload.
	// but controller always try to fetch the status of the workload.
	// Format {workloadType}/{workloadName}
	CustomWorkloadNameAnno = "sm.dx.daocloud.io/custom-workload"
)

var (
	// avoid unused warnings
	_ = DSMAppTypeLabel
	_ = DSMServiceImportedLabel
	_ = DSMServiceImportTimeAnno
	_ = DSMServiceImportStatusAnno
	_ = DSMServiceImportFailedReasonAnno
	_ = DsmManageLabel
	_ = DXOriginOwnerAnnotationName
	_ = CustomWorkloadNameAnno
)
