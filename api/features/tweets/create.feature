Feature: Create a tweet
  As an API user
  I want to create tweet
  So I can broadcast a message

  Scenario: Create tweet
    When I send a "POST" request to "/v1/tweets" with:
    """
    {
      "data": {
        "type": "Tweets",
        "attributes": {
          "message": "hello my friend"
        },
        "tag": "abc"
      }
    }
    """
    Then the response code should be 201
    And the JSON response should contain tweet data

  Scenario: Fail to create tweet
    When I send a "POST" request to "/v1/tweets" with:
    """
        {
        }
    """
    Then the response code should be 400
    And the JSON response should contain an error
