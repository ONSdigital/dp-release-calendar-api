Feature: Legacy endpoint

  Scenario: Call the legacy endpoint without an url parameter
    When I GET "/v1/releases/legacy"
    Then the HTTP status code should be "404"

 Scenario: Call the legacy endpoint with a valid url parameter
    When I GET "/v1/releases/legacy?url=/release/economy"
    Then I should receive the following JSON response with status "200":
        """
        {
            "uri":"/release/economy",
            "markdown":["markdown 1", "markdown 2"],
            "related_documents": [
                {
                "title": "Document 1",
                "summary": "This is document 1",
                "uri": "/doc/1"
                }
            ],
            "related_api_datasets": [
                {
                "title": "API dataset 1",
                "summary": "This is api dataset 1",
                "uri": "/api-dataset/1"
                }
            ],
            "related_datasets": [
                {
                "title": "Dataset 1",
                "summary": "This is dataset 1",
                "uri": "/dataset/1"
                }
            ],
            "related_methodology": [
                {
                "title": "Methodology",
                "summary": "This is methodology 1",
                "uri": "/methodology/1"
                }
            ],
            "related_methodology_article": [
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
            "date_changes": [
                {
                "previous_date": "2022-02-15T11:12:05.592Z",
                "change_notice": "This release has changed"
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
              "national_statistic":true,
              "welsh_statistic": true,
              "survey": "census",
              "release_date":"2020-07-08T23:00:00.000Z",
              "next_release":"January 2021",
              "published":true,
              "finalised":true,
              "cancelled":true,
              "cancellation_notice":["cancelled for a reason"],
              "provisional_date":"July 2020"
            }
        }
        """

  Scenario: Call the legacy endpoint with an invalid url parameter
    When I GET "/v1/releases/legacy?url=/release/invalid"
    Then the HTTP status code should be "500"