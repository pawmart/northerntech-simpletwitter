Feature: Delete tweet
  As an API user
  I want to delete a tweet
  So I can remove a tweet resource

  Scenario: Delete tweet success
    When I send a "DELETE" request to "/v1/tweets/b8dfdf10-33fa-4301-b859-e19853641651"
    Then the response code should be 204

  Scenario: Delete tweet fail when not exist
    When I send a "DELETE" request to "/v1/tweets/f3c5f34a-3985-44b2-bb1d-a51ffda32baf"
    Then the response code should be 404

  Scenario: Delete tweet fail due to non uuid
    When I send a "DELETE" request to "/v1/tweets/abc"
    Then the response code should be 422