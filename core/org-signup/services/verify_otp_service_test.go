package org_signup_services

import (
	"context"
	org_signup_dtos "sharpnr-engine/core/org-signup/dtos"
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	"sharpnr-engine/sere"
	"testing"
)

func Test_orgSignUpService_VerifyOtp(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		sere              *sere.SERE
		tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface
		mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface
		// Named input parameters for target function.
		orgSignUpSlug string
		otp           org_signup_dtos.VerifyOtpRequest_Dto
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewOrgSignUpService(tt.sere, tt.tempOrgSignupRepo, tt.mainOrgSignUpRepo)
			gotErr := s.VerifyOtp(context.Background(), tt.orgSignUpSlug, tt.otp)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("VerifyOtp() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("VerifyOtp() succeeded unexpectedly")
			}
		})
	}
}
