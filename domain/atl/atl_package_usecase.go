package domain

type ATLPackageUseCase interface {
	GetPackage(domain ATLRequest) (*ATLPackageResult, error)
	FindPackage(id string) (*ATLPackageResponse, error)
	GetRegion() (*RegionResult, error)
}
