package steps

import (
	"context"
	"errors"
	"net/http"

	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-release-calendar-api/api"
	"github.com/ONSdigital/dp-release-calendar-api/config"
	"github.com/ONSdigital/dp-release-calendar-api/service"
	"github.com/ONSdigital/dp-release-calendar-api/service/mock"

	componenttest "github.com/ONSdigital/dp-component-test"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
)

type Component struct {
	componenttest.ErrorFeature
	svcList        *service.ExternalServiceList
	svc            *service.Service
	errorChan      chan error
	Config         *config.Config
	HTTPServer     *http.Server
	ServiceRunning bool
	apiFeature     *componenttest.APIFeature
}

//nolint:gosec // component test
func NewComponent() (*Component, error) {
	c := &Component{
		HTTPServer:     &http.Server{},
		errorChan:      make(chan error),
		ServiceRunning: false,
	}

	var err error

	c.Config, err = config.Get()
	if err != nil {
		return nil, err
	}

	initMock := &mock.InitialiserMock{
		DoGetHealthCheckFunc:   c.DoGetHealthcheckOk,
		DoGetHTTPServerFunc:    c.DoGetHTTPServer,
		DoGetZebedeeClientFunc: c.DoGetZebedeeClient,
	}

	c.svcList = service.NewServiceList(initMock)

	c.apiFeature = componenttest.NewAPIFeature(c.InitialiseService)

	return c, nil
}

func (c *Component) Reset() *Component {
	c.apiFeature.Reset()
	return c
}

func (c *Component) Close() error {
	if c.svc != nil && c.ServiceRunning {
		c.svc.Close(context.Background())
		c.ServiceRunning = false
	}
	return nil
}

func (c *Component) InitialiseService() (http.Handler, error) {
	var err error
	c.svc, err = service.Run(context.Background(), c.Config, c.svcList, "1", "", "", c.errorChan)
	if err != nil {
		return nil, err
	}

	c.ServiceRunning = true
	return c.HTTPServer.Handler, nil
}

//nolint:revive // component test code
func (c *Component) DoGetHealthcheckOk(cfg *config.Config, buildTime, gitCommit, version string) (service.HealthChecker, error) {
	return &mock.HealthCheckerMock{
		AddCheckFunc: func(name string, checker healthcheck.Checker) error { return nil },
		StartFunc:    func(ctx context.Context) {},
		StopFunc:     func() {},
	}, nil
}

func (c *Component) DoGetHTTPServer(bindAddr string, router http.Handler) service.HTTPServer {
	c.HTTPServer.Addr = bindAddr
	c.HTTPServer.Handler = router
	return c.HTTPServer
}

//nolint:revive // component test code
func (c *Component) DoGetZebedeeClient(url string) api.ZebedeeClient {
	return &api.ZebedeeClientMock{
		GetReleaseFunc: func(ctx context.Context, userAccessToken, collectionID, lang, uri string) (zebedee.Release, error) {
			if uri == "/release/economy" {
				return zebedee.Release{
					URI:      uri,
					Markdown: []string{"markdown 1", "markdown 2"},
					RelatedDocuments: []zebedee.Link{
						{
							Title:   "Document 1",
							Summary: "This is document 1",
							URI:     "/doc/1",
						},
					},
					RelatedAPIDatasets: []zebedee.Link{
						{
							Title:   "API dataset 1",
							Summary: "This is api dataset 1",
							URI:     "/api-dataset/1",
						},
					},
					RelatedDatasets: []zebedee.Link{
						{
							Title:   "Dataset 1",
							Summary: "This is dataset 1",
							URI:     "/dataset/1",
						},
					},
					RelatedMethodology: []zebedee.Link{
						{
							Title:   "Methodology",
							Summary: "This is methodology 1",
							URI:     "/methodology/1",
						},
					},
					RelatedMethodologyArticle: []zebedee.Link{
						{
							Title:   "Methodology Article",
							Summary: "This is methodology article 1",
							URI:     "/methodology/article/1",
						},
					},
					Links: []zebedee.Link{
						{
							Title:   "Link 1",
							Summary: "This is link 1",
							URI:     "/link/1",
						},
					},
					DateChanges: []zebedee.ReleaseDateChange{
						{
							Date:         "2022-02-15T11:12:05.592Z",
							ChangeNotice: "This release has changed",
						},
					},
					Description: zebedee.Description{
						Title:   "Release title",
						Summary: "Release summary",
						Contact: zebedee.Contact{
							Email:     "contact@ons.gov.uk",
							Name:      "Contact name",
							Telephone: "029",
						},
						NationalStatistic:  true,
						WelshStatistic:     true,
						Survey:             "census",
						ReleaseDate:        "2020-07-08T23:00:00.000Z",
						NextRelease:        "January 2021",
						Published:          true,
						Finalised:          true,
						Cancelled:          true,
						CancellationNotice: []string{"cancelled for a reason"},
						ProvisionalDate:    "July 2020",
					},
				}, nil
			}
			return zebedee.Release{}, errors.New("unsupported endpoint")
		},
	}
}
