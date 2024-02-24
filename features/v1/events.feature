Feature: Events
  Scenario: When get the "/v1/events" then it should return empty array
    When I make a GET request to "/v1/events"
    Then the response body should be
      """
      []
      """
    Then the response status code should be 200

  Scenario: When post the "/v1/events" then it should return ok
    When I make a POST request to "/v1/events" with the body
      """
      {}
      """
    Then the response json should have "id"
    And the response status code should be 200

  Scenario: When get the "/v1/events" then it should return array with one element
    When I make a POST request to "/v1/events" with the body
      """
      {}
      """
    And I make a GET request to "/v1/events"
    Then the response json should have "[].id"
    And the response status code should be 200
