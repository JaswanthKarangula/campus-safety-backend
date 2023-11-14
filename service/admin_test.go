package service

import (
	db "dronesaefty-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestAddNewAdmin(t *testing.T) {
	type args struct {
		ctx   *gin.Context
		store db.Store
		arg   db.CreateNewAdminParams
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
			got, err := AddNewAdmin(tt.args.ctx, tt.args.store, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddNewAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNewAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
