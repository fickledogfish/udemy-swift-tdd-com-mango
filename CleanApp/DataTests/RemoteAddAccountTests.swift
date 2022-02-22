import XCTest
import Data
import Domain

class RemoteAddAccountTests: XCTestCase {
    func test_add_should_call_httpClient_with_correct_url() {
        // Arrange
        let url = URL(string: "https://curl.se/")!
        let (sut, httpClientSpy) = makeSut(url: url)
        let addAccountModel = makeAddAccountModel()

        // Act
        sut.add(account: addAccountModel)

        // Assert
        XCTAssertEqual(httpClientSpy.url, url)
    }

    func test_add_should_call_httpClient_with_correct_data() {
        // Arrange
        let (sut, httpClientSpy) = makeSut()
        let addAccountModel = makeAddAccountModel()

        // Act
        sut.add(account: addAccountModel)

        // Assert
        XCTAssertEqual(httpClientSpy.data, addAccountModel.toData())
    }
}

// MARK: - RemoteAddAccountTests helpers

extension RemoteAddAccountTests {
    func makeSut(
        url: URL = URL(string: "https://xkcd.com")!
    ) -> (sut: RemoteAddAccount, httpClientSpy: HttpClientSpy) {
        let httpClientSpy = HttpClientSpy()
        let sut = RemoteAddAccount(url: url, httpClient: httpClientSpy)

        return (sut, httpClientSpy)
    }

    func makeAddAccountModel() -> AddAccountModel {
        AddAccountModel(
            name: "Any Name",
            email: "any.name@example.com",
            password: "pass",
            passwordConfirmation: "pass"
        )
    }

    class HttpClientSpy: HttpPostClient {
        var url: URL?
        var data: Data?

        func post(to url: URL, with data: Data?) {
            self.url = url
            self.data = data
        }
    }
}
