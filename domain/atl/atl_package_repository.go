package domain

type Reader interface {
	GetPackage(request ATLRequest) (*ATLPackageResult, error)
	FindPackage(id string) (*ATLPackageResponse, error)
	GetRegion() (*RegionResult, error)
}

type ATLRepository interface {
	Reader
}
