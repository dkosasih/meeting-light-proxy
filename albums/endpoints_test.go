package albums

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterEndpoints(t *testing.T) {
	type args struct {
		r       *gin.Engine
		creator AlbumHandlerCreator
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterEndpoints(tt.args.r, tt.args.creator)
		})
	}
}

func Test_getAlbums(t *testing.T) {
	type args struct {
		creator AlbumHandlerCreator
	}
	tests := []struct {
		name string
		args args
		want func(*gin.Context)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAlbums(tt.args.creator); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("getAlbums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAlbumByID(t *testing.T) {
	type args struct {
		creator AlbumHandlerCreator
	}
	tests := []struct {
		name string
		args args
		want func(c *gin.Context)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAlbumByID(tt.args.creator); !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("getAlbumByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postAlbums(t *testing.T) {
	type args struct {
		creator AlbumHandlerCreator
	}
	tests := []struct {
		name string
		args args
		want func(c *gin.Context)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postAlbums(tt.args.creator); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("postAlbums() = %v, want %v", got, tt.want)
			}
		})
	}
}
