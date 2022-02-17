package mapper

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/releasecalendar"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
)

func CreateReleaseFromZebedee(zr zebedee.Release) releasecalendar.Release {
	release := releasecalendar.Release{
		URI:      zr.URI,
		Markdown: zr.Markdown,
		Description: releasecalendar.ReleaseDescription{
			Title:   zr.Description.Title,
			Summary: zr.Description.Summary,
			Contact: releasecalendar.Contact{
				Name:      zr.Description.Contact.Name,
				Email:     zr.Description.Contact.Email,
				Telephone: zr.Description.Contact.Telephone,
			},
			NationalStatistic:  zr.Description.NationalStatistic,
			ReleaseDate:        zr.Description.ReleaseDate,
			NextRelease:        zr.Description.NextRelease,
			Published:          zr.Description.Published,
			Finalised:          zr.Description.Finalised,
			Cancelled:          zr.Description.Cancelled,
			CancellationNotice: zr.Description.CancellationNotice,
			ProvisionalDate:    zr.Description.ProvisionalDate,
		},
	}
	release.RelatedDocuments = mapZebedeeLink(zr.RelatedDocuments)
	release.RelatedDatasets = mapZebedeeLink(zr.RelatedDatasets)
	release.RelatedMethodology = mapZebedeeLink(zr.RelatedMethodology)
	release.RelatedMethodologyArticle = mapZebedeeLink(zr.RelatedMethodologyArticle)
	release.Links = mapZebedeeLink(zr.Links)

	release.DateChanges = []releasecalendar.ReleaseDateChange{}
	for _, changes := range zr.DateChanges {
		release.DateChanges = append(release.DateChanges, releasecalendar.ReleaseDateChange{
			Date:         changes.Date,
			ChangeNotice: changes.ChangeNotice,
		})
	}

	return release
}

func mapZebedeeLink(origin []zebedee.Link) []releasecalendar.Link {
	res := []releasecalendar.Link{}
	for _, related := range origin {
		res = append(res, releasecalendar.Link{
			Title:   related.Title,
			Summary: related.Summary,
			URI:     related.URI,
		})
	}
	return res
}
