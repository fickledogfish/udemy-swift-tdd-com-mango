import XCTest

class RemoteAddAccountTests: XCTestCase {
    func test() throws {
        // Arrange
        let url = URL(string: "https://xkcd.com")!
        let httpClientSpy = HttpClientSpy()
        let sut = RemoteAddAccount(url: url, httpClient: httpClientSpy)

        // Act
        sut.add()

        // Assert
        XCTAssertEqual(httpClientSpy.url, url)
    }

    class HttpClientSpy: HttpClient {
        var url: URL?

        func post(url: URL) {
            self.url = url
        }
    }
}

protocol HttpClient {
    func post(url: URL)
}

class RemoteAddAccount {
    private let url: URL
    private let httpClient: HttpClient

    init(url: URL, httpClient: HttpClient) {
        self.url = url
        self.httpClient = httpClient
    }

    func add() {
        httpClient.post(url: url)
    }
}
