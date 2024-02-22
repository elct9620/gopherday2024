Feature: Probe
  Scenario: When get the "/livez" then it should be success
    When I make a GET request to "/livez"
    Then the response body should be
      """
      {"status":"ok"}
      """
    Then the response status code should be 200
