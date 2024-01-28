package usecase

import (
	domain "content-management/domain/atl"
)

type PromoUseCase struct {
	pr domain.ATLRepository
}

func NewATLPackageUseCase(pr domain.ATLRepository) *PromoUseCase {
	return &PromoUseCase{pr}
}

func (p *PromoUseCase) GetPackage(request domain.ATLRequest) (*domain.ATLPackageResult, error) {

	packageResult, _ := p.pr.GetPackage(request)

	return packageResult, nil
}

func (p *PromoUseCase) FindPackage(id string) (*domain.ATLPackageResponse, error) {

	packageResult, _ := p.pr.FindPackage(id)

	return packageResult, nil
}

func (p *PromoUseCase) GetRegion() (*domain.RegionResult, error) {

	regionResult, _ := p.pr.GetRegion()

	return regionResult, nil
}
