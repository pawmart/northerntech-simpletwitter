Feature: List tweets
  As an API user
  I want to get tweets
  So I can get a list of tweets

  Scenario: List filtered tweets
    When I send a "GET" request to "/v1/tweets?filter[tag]=abc&count=1&year=2019"
    Then the response code should be 200
    And the JSON response should contain tweet data collection

  Scenario: List filtered tweets but not found
    When I send a "GET" request to "/v1/tweets?filter[tag]=asdfasdfasdfasdf"
    Then the response code should be 200
    And the JSON response should contain no tweet data collection

  Scenario: List tweets
    When I send a "GET" request to "/v1/tweets"
    Then the response code should be 200
    And the JSON response should contain tweet data collection
