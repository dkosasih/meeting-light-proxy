package openhab

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	assert := assert.New(t)

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
			endpoint := &Endpoints{
				r:       r,
				creator: tt.args.creator,
			}
			endpoint.Register()

			assert.Equal(len(r.Routes()), tt.xpectLength, "Expect only one endpoint being created")
		})
	}
}

func Test_updateOpenHab(t *testing.T) {
	assert := assert.New(t)

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
			r := gin.Default()
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			endpoints := &Endpoints{
				r:       r,
				creator: tt.args.creator,
			}
			got := endpoints.updateOpenHab()
			got(c)

			value, exists := c.Get("UpdateOpenHabCalled")
			assert.True(exists)
			assert.NotEmpty(value)
			assert.True(value.(bool))
		})
	}
}
