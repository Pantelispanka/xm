package http

import (
	"context"
	"encoding/json"
	"net/http"
	"xm-challenge/internal/domain"
)

type CompanyRepo interface {
	CreateOrganization(ctx context.Context, companyGot domain.Company) (companyCreated domain.Company, err error)
	UpdateOrg(ctx context.Context, orgToUpdate domain.Company) (orgUpdated domain.Company, err error)
	DeleteOrg(ctx context.Context, id string) (orgsDeleted int64, err error)
	GetOrgByName(ctx context.Context, name string) (org domain.Company, err error)
}

type CompanyHandlers struct {
	companyRepo CompanyRepo
	context     context.Context
}

func (repo *CompanyHandlers) GetCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var org domain.Company
	name := r.FormValue("name")

	org, err := repo.companyRepo.GetOrgByName(repo.context, name)

	if err != nil {
		err := domain.ErrorReport{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(org)
}

func (repo *CompanyHandlers) PatchCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var org domain.Company
	json.NewDecoder(r.Body).Decode(&org)

	org, err := repo.companyRepo.UpdateOrg(repo.context, org)

	if err != nil {
		err := domain.ErrorReport{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(org)
}

func (repo *CompanyHandlers) DeleteCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var org domain.Company
	json.NewDecoder(r.Body).Decode(&org)

	_, err := repo.companyRepo.DeleteOrg(repo.context, org.ID)

	if err != nil {
		err := domain.ErrorReport{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(org)
}

func (repo *CompanyHandlers) CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var org domain.Company
	json.NewDecoder(r.Body).Decode(&org)

	org, err := repo.companyRepo.CreateOrganization(repo.context, org)

	if err != nil {
		err := domain.ErrorReport{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(org)
}
