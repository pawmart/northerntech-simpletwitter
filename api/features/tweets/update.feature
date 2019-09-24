Feature: Update tweet
  As an API user
  I want to update a tweet
  So I can change the state of the tweet

  Scenario: Update tweet
    When I send a "PATCH" request to "/v1/tweets" with:
    """
    {
        "data": {
            "id": "a8dfdf10-33fa-4301-b859-e19853641655",
            "type": "Tweets",
            "tag": "abc",
            "attributes": {
                "message": "abc abc"
            }
        }
    }
    """
    And the response code should be 200