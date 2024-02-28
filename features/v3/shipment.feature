Feature: Shipments
  Scenario: When get the "/v3/shipments/a449b857-5e97-4bee-8ffe-1b544fa4ec5b" then it should return unknown state
    When I make a GET request to "/v3/shipments/a449b857-5e97-4bee-8ffe-1b544fa4ec5b"
    Then the response body should be
      """
      {
        "id": "a449b857-5e97-4bee-8ffe-1b544fa4ec5b",
        "state": "unknown",
        "items": [],
        "updated_at": null
      }
      """
    Then the response status code should be 200

  Scenario: When post the "/v3/shipments" then it should return creation success
    When I make a POST request to "/v3/shipments" with the body
      """
      {
        "id": "a449b857-5e97-4bee-8ffe-1b544fa4ec5b"
      }
      """
    Then the response json should have "id" with value "a449b857-5e97-4bee-8ffe-1b544fa4ec5b"
    And the response json should have "state" with value "pending"
    And the response json should have "items"
    And the response json should have "updated_at"
    And the response status code should be 200
