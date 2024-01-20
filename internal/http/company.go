package http

import (
	"context"
	"net/http"
	"xm-challenge/internal/domain"
)

type CompanyRepo interface {
	CreateCompany(ctx context.Context, company domain.Company) error
	UpsertCompany(ctx context.Context, company domain.Company) (*domain.Company, error)
	DeleteCompany(ctx context.Context, company domain.Company) error
	GetCompany(ctx context.Context, id string) (*domain.Company, error)
}

func GetCompanyHandler(w http.ResponseWriter, r *http.Request)

func PatchCompanyHandler(w http.ResponseWriter, r *http.Request)

func DeleteCompanyHandler(w http.ResponseWriter, r *http.Request)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request)
