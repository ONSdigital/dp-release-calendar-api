package mapper

import (
	"testing"

	"github.com/ONSdigital/dp-api-clients-go/v2/releasecalendar"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitMapper(t *testing.T) {

	Convey("Given a zebedee release", t, func() {
		zr := zebedee.Release{
			URI:      "/release/example",
			Markdown: []string{"markdown1", "markdown 2"},
			RelatedDocuments: []zebedee.Link{
				{
					Title:   "Document 1",
					Summary: "This is document 1",
					URI:     "/doc/1",
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
				ReleaseDate:        "2020-07-08T23:00:00.000Z",
				NextRelease:        "January 2021",
				Published:          true,
				Finalised:          true,
				Cancelled:          true,
				CancellationNotice: []string{"cancelled for a reason"},
				ProvisionalDate:    "July 2020",
				Edition:            "edition is not used",
				Keywords:           []string{"unused"},
				MetaDescription:    "meta description is not used",
				LatestRelease:      true,
				DatasetID:          "none",
				Unit:               "u",
				PreUnit:            "pre",
				Source:             "not in use",
				VersionLabel:       "none",
			},
		}

		Convey("CreateReleaseFromZebedee maps correctly", func() {
			release := CreateReleaseFromZebedee(zr)
			So(release.URI, ShouldEqual, zr.URI)
			So(release.Markdown, ShouldResemble, zr.Markdown)
			So(len(release.DateChanges), ShouldEqual, len(zr.DateChanges))
			So(release.DateChanges[0].ChangeNotice, ShouldEqual, zr.DateChanges[0].ChangeNotice)
			So(release.DateChanges[0].Date, ShouldEqual, zr.DateChanges[0].Date)
			assertLinks(zr.RelatedDocuments, release.RelatedDocuments)
			assertLinks(zr.RelatedDocuments, release.RelatedDocuments)
			assertLinks(zr.RelatedDatasets, release.RelatedDatasets)
			assertLinks(zr.RelatedMethodology, release.RelatedMethodology)
			assertLinks(zr.RelatedMethodologyArticle, release.RelatedMethodologyArticle)
			assertLinks(zr.Links, release.Links)
			So(release.Description.Title, ShouldEqual, zr.Description.Title)
			So(release.Description.Summary, ShouldEqual, zr.Description.Summary)
			So(release.Description.NationalStatistic, ShouldEqual, zr.Description.NationalStatistic)
			So(release.Description.ReleaseDate, ShouldEqual, zr.Description.ReleaseDate)
			So(release.Description.NextRelease, ShouldEqual, zr.Description.NextRelease)
			So(release.Description.Published, ShouldEqual, zr.Description.Published)
			So(release.Description.Finalised, ShouldEqual, zr.Description.Finalised)
			So(release.Description.Cancelled, ShouldEqual, zr.Description.Cancelled)
			So(release.Description.CancellationNotice, ShouldResemble, zr.Description.CancellationNotice)
			So(release.Description.ProvisionalDate, ShouldEqual, zr.Description.ProvisionalDate)
			So(release.Description.Contact.Name, ShouldEqual, zr.Description.Contact.Name)
			So(release.Description.Contact.Email, ShouldEqual, zr.Description.Contact.Email)
			So(release.Description.Contact.Telephone, ShouldEqual, zr.Description.Contact.Telephone)
		})
	})
}

// assertLinks checks that the actual release Link content is equal to the expected zebedee Link
func assertLinks(expected []zebedee.Link, actual []releasecalendar.Link) {
	So(len(actual), ShouldEqual, len(expected))
	for i := range expected {
		So(actual[i].URI, ShouldEqual, expected[i].URI)
		So(actual[i].Title, ShouldEqual, expected[i].Title)
		So(actual[i].Summary, ShouldEqual, expected[i].Summary)
	}
}
