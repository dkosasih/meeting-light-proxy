package openhab

import (
	"reflect"
	"testing"

	"net/http/httptest"

	v1 "github.com/dkosasih/meeting-light-proxy/openhab/v1"
	"github.com/gin-gonic/gin"
)

func TestRegisterEndpoints(t *testing.T) {
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterEndpoints(tt.args.r)
		})
	}
}

func Test_updateOpenHab(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateOpenHab(tt.args.c)
		})
	}
}

func Test_createHandler(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		version string
		args    args
		want    OpenhabHandlerInterface
	}{
		{"should create v1 handler when context is v1", "2", args{c}, &v1.OpenhabHandler{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Set("version", tt.version)
			got := createHandler(tt.args.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
