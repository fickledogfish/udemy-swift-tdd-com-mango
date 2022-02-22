import XCTest
import Domain

protocol HttpPostClient {
    func post(to url: URL, with data: Data?)
}

class RemoteAddAccount {
    private let url: URL
    private let httpClient: HttpPostClient

    init(url: URL, httpClient: HttpPostClient) {
        self.url = url
        self.httpClient = httpClient
    }

    func add(account: AddAccountModel) {
        let data = try? JSONEncoder().encode(account)
        httpClient.post(to: url, with: data)
    }
}

class RemoteAddAccountTests: XCTestCase {
    func test_add_should_call_httpClient_with_correct_url() {
        // Arrange
        let url = URL(string: "https://xkcd.com")!
        let httpClientSpy = HttpClientSpy()
        let sut = RemoteAddAccount(url: url, httpClient: httpClientSpy)
        let addAccountModel = AddAccountModel(
            name: "Any Name",
            email: "any.name@example.com",
            password: "pass",
            passwordConfirmation: "pass"
        )

        // Act
        sut.add(account: addAccountModel)

        // Assert
        XCTAssertEqual(httpClientSpy.url, url)
    }

    func test_add_should_call_httpClient_with_correct_data() {
        // Arrange
        let url = URL(string: "https://xkcd.com")!
        let httpClientSpy = HttpClientSpy()
        let sut = RemoteAddAccount(url: url, httpClient: httpClientSpy)
        let addAccountModel = AddAccountModel(
            name: "Any Name",
            email: "any.name@example.com",
            password: "pass",
            passwordConfirmation: "pass"
        )
        let data = try? JSONEncoder().encode(addAccountModel)

        // Act
        sut.add(account: addAccountModel)

        // Assert
        XCTAssertEqual(httpClientSpy.data, data)
    }
}

// MARK: - RemoteAddAccountTests helpers

extension RemoteAddAccountTests {
    class HttpClientSpy: HttpPostClient {
        var url: URL?
        var data: Data?

        func post(to url: URL, with data: Data?) {
            self.url = url
            self.data = data
        }
    }
}
