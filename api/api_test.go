package api

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSetup(t *testing.T) {
	Convey("Given an API instance", t, func() {
		r := mux.NewRouter()
		ctx := context.Background()

		mockZebedeeClient := &ZebedeeClientMock{
			GetReleaseFunc: func(ctx context.Context, userAccessToken, collectionID, lang, uri string) (zebedee.Release, error) {
				return zebedee.Release{}, nil
			},
		}

		api := Setup(ctx, r, mockZebedeeClient)

		Convey("When created the following routes should have been added", func() {
			So(hasRoute(api.Router, "/releasecalendar/legacy", "GET"), ShouldBeTrue)
		})
	})
}

func hasRoute(r *mux.Router, path, method string) bool {
	req := httptest.NewRequest(method, path, nil)
	match := &mux.RouteMatch{}
	return r.Match(req, match)
}
