package openhab

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type HandlerMock struct {
}

func (hm *HandlerMock) UpdateOpenHab(c *gin.Context) {
	c.Set("UpdateOpenHabCalled", true)
}

type HandlerFactoryMock struct {
}

func (hmf *HandlerFactoryMock) CreateHandler(*gin.Context) OpenhabHandlerInterface {
	var hm OpenhabHandlerInterface = &HandlerMock{}

	return hm
}

func TestRegisterEndpoints(t *testing.T) {
	type args struct {
		creator OpenhabHandlerCreator
	}
	tests := []struct {
		name        string
		xpectLength int
		args        args
	}{
		{"Should create 1 endpoint when register endpoint is called", 1, args{creator: &HandlerFactoryMock{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			RegisterEndpoints(r, tt.args.creator)

			if len(r.Routes()) != tt.xpectLength {
				t.Errorf("Expect only one endpoint being created")
			}
		})
	}
}

func Test_updateOpenHab(t *testing.T) {
	type args struct {
		creator OpenhabHandlerCreator
	}
	tests := []struct {
		name string
		args args
		want func(*gin.Context)
	}{
		{"Should call UpdateOpenHab func in Handler", args{creator: &HandlerFactoryMock{}}, func(*gin.Context) {}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			got := updateOpenHab(tt.args.creator)
			got(c)

			if v, e := c.Get("UpdateOpenHabCalled"); v == "" && !e {
				t.Errorf("Expect UpdateOpenHab() from versioned handler to be called")
			}
		})
	}
}
