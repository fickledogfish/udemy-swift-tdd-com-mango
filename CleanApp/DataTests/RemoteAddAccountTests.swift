import XCTest

protocol HttpPostClient {
    func post(url: URL)
}

class RemoteAddAccount {
    private let url: URL
    private let httpClient: HttpPostClient

    init(url: URL, httpClient: HttpPostClient) {
        self.url = url
        self.httpClient = httpClient
    }

    func add() {
        httpClient.post(url: url)
    }
}

class RemoteAddAccountTests: XCTestCase {
    func test_add_should_call_httpClient_with_correct_url() {
        // Arrange
        let url = URL(string: "https://xkcd.com")!
        let httpClientSpy = HttpClientSpy()
        let sut = RemoteAddAccount(url: url, httpClient: httpClientSpy)

        // Act
        sut.add()

        // Assert
        XCTAssertEqual(httpClientSpy.url, url)
    }
}

// MARK: - RemoteAddAccountTests helpers

extension RemoteAddAccountTests {
    class HttpClientSpy: HttpPostClient {
        var url: URL?

        func post(url: URL) {
            self.url = url
        }
    }
}
