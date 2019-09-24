Feature: Check health
  As an API user
  I want to check API health
  So I can be sure I can interact with the API

  Background:
    When I send a "GET" request to "/v1/health"

  Scenario: Success
    Then the response code should be 200
    And the response should have health status up