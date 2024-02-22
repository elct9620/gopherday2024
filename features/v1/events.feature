Feature: Events
  Scenario: When get the "/v1/events" then it should return empty array
    When I make a GET request to "/v1/events"
    Then the response body should be
      """
      []
      """
    Then the response status code should be 200
