package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-api-clients-go/v2/headers"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-release-calendar-api/mapper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLegacyHandler(t *testing.T) {

	Convey("Given a Legacy handler ", t, func() {
		url := "/release/adoption"
		r := zebedee.Release{URI: url, Description: zebedee.Description{Title: "Test"}}
		mockZebedeeClient := &ZebedeeClientMock{
			GetReleaseFunc: func(ctx context.Context, userAccessToken, collectionId, lang, uri string) (zebedee.Release, error) {
				return r, nil
			},
		}
		handler := LegacyHandler(context.Background(), mockZebedeeClient)

		Convey("when a valid request without headers is received (web)", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:8080/releasecalendar/legacy?url=%s", url), nil)
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Convey("Then the call to Zebedee is correct", func() {
				So(len(mockZebedeeClient.GetReleaseCalls()), ShouldEqual, 1)
				So(mockZebedeeClient.GetReleaseCalls()[0].UserAccessToken, ShouldEqual, "")
				So(mockZebedeeClient.GetReleaseCalls()[0].CollectionID, ShouldEqual, "")
				So(mockZebedeeClient.GetReleaseCalls()[0].Lang, ShouldEqual, "en")
				So(mockZebedeeClient.GetReleaseCalls()[0].URI, ShouldEqual, url)
			})
			Convey("And the response is correct", func() {
				So(resp.Code, ShouldEqual, http.StatusOK)
				expectedJson, _ := json.Marshal(mapper.CreateReleaseFromZebedee(r))
				So(resp.Body.Bytes(), ShouldResemble, expectedJson)
				So(len(mockZebedeeClient.GetReleaseCalls()), ShouldEqual, 1)
			})
		})

		Convey("when a valid request with headers is received (publishing)", func() {
			accessToken := "user-access-token"
			collectionID := "my-collection"

			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:8080/releasecalendar/legacy?url=%s", url), nil)
			headers.SetAuthToken(req, accessToken)
			headers.SetCollectionID(req, collectionID)

			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Convey("Then the call to Zebedee is correct", func() {
				So(len(mockZebedeeClient.GetReleaseCalls()), ShouldEqual, 1)
				So(mockZebedeeClient.GetReleaseCalls()[0].UserAccessToken, ShouldEqual, accessToken)
				So(mockZebedeeClient.GetReleaseCalls()[0].CollectionID, ShouldEqual, collectionID)
				So(mockZebedeeClient.GetReleaseCalls()[0].Lang, ShouldEqual, "en")
				So(mockZebedeeClient.GetReleaseCalls()[0].URI, ShouldEqual, url)
			})
			Convey("And the response is correct", func() {
				So(resp.Code, ShouldEqual, http.StatusOK)
				expectedJson, _ := json.Marshal(mapper.CreateReleaseFromZebedee(r))
				So(resp.Body.Bytes(), ShouldResemble, expectedJson)
				So(len(mockZebedeeClient.GetReleaseCalls()), ShouldEqual, 1)
			})
		})

		Convey("when a request without an url parameter is received", func() {
			req := httptest.NewRequest("GET", "http://localhost:8080/releasecalendar/legacy", nil)
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Convey("Then an error is returned", func() {
				So(resp.Code, ShouldEqual, http.StatusNotFound)
				So(resp.Body.String(), ShouldResemble, "URL not found\n")
			})
		})

		Convey("when a request with an empty url parameter is received", func() {
			req := httptest.NewRequest("GET", "http://localhost:8080/releasecalendar/legacy?url=", nil)
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Convey("Then an error is returned", func() {
				So(resp.Code, ShouldEqual, http.StatusNotFound)
				So(resp.Body.String(), ShouldResemble, "URL not found\n")
			})
		})
	})
}
