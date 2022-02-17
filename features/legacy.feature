Feature: Legacy endpoint

  Scenario: Call the legacy endpoint without an url parameter
    When I GET "/releasecalendar/legacy"
    Then the HTTP status code should be "404"

 Scenario: Call the legacy endpoint with a valid url parameter
    When I GET "/releasecalendar/legacy?url=/release/economy"
    And I should receive the following JSON response with status "200":
        """
        {
            "uri":"/release/economy",
            "markdown":["markdown 1", "markdown 2"],
            "relatedDocuments": [
                {
                "title": "Document 1",
                "summary": "This is document 1",
                "uri": "/doc/1"
                }
            ],
            "relatedDatasets": [
                {
                "title": "Dataset 1",
                "summary": "This is dataset 1",
                "uri": "/dataset/1"
                }
            ],
            "relatedMethodology": [
                {
                "title": "Methodology",
                "summary": "This is methodology 1",
                "uri": "/methodology/1"
                }
            ],
            "relatedMethodologyArticle": [
                {
                "title": "Methodology Article",
                "summary": "This is methodology article 1",
                "uri": "/methodology/article/1"
                }
            ],
            "links": [
                {
                "title": "Link 1",
                "summary": "This is link 1",
                "uri": "/link/1"
                }
            ],
            "dateChanges": [
                {
                "previousDate": "2022-02-15T11:12:05.592Z",
                "changeNotice": "This release has changed"
                }
            ],
            "description": {
              "title":"Release title", 
              "summary":"Release summary",
              "contact": {
                "email":"contact@ons.gov.uk", 
                "name":"Contact name", 
                "telephone":"029"
              }, 
              "nationalStatistic":true, 
              "releaseDate":"2020-07-08T23:00:00.000Z",
              "nextRelease":"January 2021",
              "published":true,
              "finalised":true,
              "cancelled":true,
              "cancellationNotice":["cancelled for a reason"],
              "provisionalDate":"July 2020"
            }
        }
        """

  Scenario: Call the legacy endpoint with an invalid url parameter
    When I GET "/releasecalendar/legacy?url=/release/invalid"
    Then the HTTP status code should be "500"