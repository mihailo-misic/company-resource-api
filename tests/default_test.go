package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/mihailo-misic/company-resource-api/routers"
	. "github.com/smartystreets/goconvey/convey"

	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "..")))
	appPath = filepath.Join(appPath, "..")
	beego.TestBeegoInit(appPath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", fmt.Sprintf("Code[%d]\n%s", w.Code, w.Body.String()), fmt.Sprintf("\nHeadMap: %+v\n URL: %s", w.HeaderMap, r.URL))

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})

		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
