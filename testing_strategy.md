# My Testing Strategy

## Unit tests
- i test my helper functions, those in json_helpers.go, conversion.go and ensure they cover all edge cases and return expected values.
- Test my Middlewares such as the rate limiter to ensure that it works as expected
- unit test each layer differently and mocking any external dependency using gomock or any other testing libraries, and ensure they bahve expectdly.

## Integration
- Verify that SQL queries, repository methods, and database transactions work as intended, i would swap out my real implementation for a stub external db implementation or use a test db seeded with data to ensure that all layers work as expected and edge cases are handled properly

## API Testing
- Here i test my api layers and ensure they return the proposed api contract.
- Test all HTTP methods (GET, POST, PUT, DELETE) for success and failure scenarios.
- Confirm that proper status codes, headers, and error messages are returned.

## End 2 End
- Test the full application workflow from a user’s perspective using tools like thunder client(in vscode) 
- This time would use a test prod db and test things like network realted issues, performance and all.

## Key Considerations
- `Test Coverage`: I aim for meaningful coverage without chasing 100%, focusing on high-priority and high-risk areas.
- `Error Handling`: Include tests for various edge cases, unexpected inputs, and failure scenarios across all my layers.
- `Automation`: Automate tests where possible using tools build for such and create CI/CD pipelines.
- `Performance Tests`: Optionally test for performance bottlenecks in APIs and database queries, to this end testing how my app would bahave in prod under massive load.
- `Security(Authentication and Authorization)`: This app does not have this feature, but it's a critical part of my testing as i hold security as utmost priority, i usual ue paseto tokens so i ensure that my app allows and restricts access when neccessary.