package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestCreateNewOfficerIssue(t *testing.T) {
	type args struct {
		ctx   *gin.Context
		store db.Store
		req   db.CreateSecurityOfficerIssueParams
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
			got, err := CreateNewOfficerIssue(tt.args.ctx, tt.args.store, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNewOfficerIssue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateNewOfficerIssue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNewSecurityOfficer(t *testing.T) {
	type args struct {
		ctx   *gin.Context
		store db.Store
		arg   db.CreateNewSecurityOfficerParams
	}
	tests := []struct {
		name    string
		args    args
		want    db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateNewSecurityOfficer(tt.args.ctx, tt.args.store, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNewSecurityOfficer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNewSecurityOfficer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSchedules(t *testing.T) {
	type args struct {
		ctx      *gin.Context
		store    db.Store
		schedule []db.UpdateScheduleParams
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
			got, err := UpdateSchedules(tt.args.ctx, tt.args.store, tt.args.schedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateSchedules() got = %v, want %v", got, tt.want)
			}
		})
	}
}
