package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCreateNewCustomerIssue(t *testing.T) {
	type args struct {
		ctx   *gin.Context
		store db.Store
		req   db.CreateCustomerIssueParams
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateNewCustomerIssue(tt.args.ctx, tt.args.store, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNewCustomerIssue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateNewCustomerIssue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
