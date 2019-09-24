Feature: Fetch a tweet
  As an API user
  I want to get a tweet
  So I can get details of the tweet

  Scenario: Get tweet
    When I send a "GET" request to "/v1/tweets/a8dfdf10-33fa-4301-b859-e19853641655"
    And the JSON response should contain tweet data
    Then the response code should be 200

  Scenario: Fail to get a tweet
    When I send a "GET" request to "/v1/tweets/f3c5f34a-3985-44b2-bb1d-a51ffda32baf"
    Then the response code should be 404
    And the JSON response should contain an error

  Scenario: Fail to get a tweet due to non uuid
    When I send a "GET" request to "/v1/tweets/abc"
    Then the response code should be 422