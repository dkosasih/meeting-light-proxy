package albums

import (
	"net/http/httptest"
	"reflect"
	"testing"

	v1 "github.com/dkosasih/meeting-light-proxy/albums/v1"
	v2 "github.com/dkosasih/meeting-light-proxy/albums/v2"
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

func Test_getAlbums(t *testing.T) {
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
			getAlbums(tt.args.c)
		})
	}
}

func Test_getAlbumByID(t *testing.T) {
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
			getAlbumByID(tt.args.c)
		})
	}
}

func Test_postAlbums(t *testing.T) {
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
			postAlbums(tt.args.c)
		})
	}
}

func Test_createHandler(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name                string
		version             string
		shouldVersionExists bool
		args                args
		want                AlbumHandlerInterface
	}{
		{"should create v1 handler when context version is v1", "1", true, args{c}, &v1.AlbumHandler{}},
		{"should create v2 handler when context version is v2", "2", true, args{c}, &v2.AlbumHandler{}},
		{"should create v2 as default handler when context version is not v1 or v2", "5", true, args{c}, &v2.AlbumHandler{}},
		{"should create v2 as default handler when context version is not v1 or v2", "", true, args{c}, &v2.AlbumHandler{}},
		{"should create v2 as default handler when context version not exists", "", false, args{c}, &v2.AlbumHandler{}},
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
