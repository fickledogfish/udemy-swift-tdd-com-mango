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
        sut.add(account: addAccountModel) { _ in }

        // Assert
        XCTAssertEqual(httpClientSpy.urls, [url])
    }

    func test_add_should_call_httpClient_with_correct_data() {
        // Arrange
        let (sut, httpClientSpy) = makeSut()
        let addAccountModel = makeAddAccountModel()

        // Act
        sut.add(account: addAccountModel) { _ in }

        // Assert
        XCTAssertEqual(httpClientSpy.data, addAccountModel.toData())
    }

    func test_add_should_complete_with_error_if_client_completes_with_error() {
        // Arrange
        let (sut, httpClientSpy) = makeSut()
        let addAccountModel = makeAddAccountModel()
        let exp = expectation(description: "waiting")

        // Act
        sut.add(account: addAccountModel) { result in
            switch result {
            case .failure(let error):
                // Assert
                XCTAssertEqual(error, .unexpected)

            default:
                XCTFail("Expected error, but got \(result) instead")
            }

            exp.fulfill()
        }

        httpClientSpy.completeWith(error: .noConnectivity)
        wait(for: [exp], timeout: 1)
    }

    func test_add_should_complete_with_account_if_client_completes_with_valid_data() {
        // Arrange
        let (sut, httpClientSpy) = makeSut()
        let addAccountModel = makeAddAccountModel()
        let exp = expectation(description: "waiting")
        let expectedAccount = makeAccountModel()

        // Act
        sut.add(account: addAccountModel) { result in
            switch result {
            case .success(let receivedAccount):
                // Assert
                XCTAssertEqual(receivedAccount, expectedAccount)

            default:
                XCTFail("Expected success, but got \(result) instead")
            }

            exp.fulfill()
        }

        httpClientSpy.completeWith(data: expectedAccount.toData()!)
        wait(for: [exp], timeout: 1)
    }

    func test_add_should_complete_with_error_if_client_completes_with_invalid_data() {
        // Arrange
        let (sut, httpClientSpy) = makeSut()
        let addAccountModel = makeAddAccountModel()
        let exp = expectation(description: "waiting")

        // Act
        sut.add(account: addAccountModel) { result in
            switch result {
            case .failure(let error):
                // Assert
                XCTAssertEqual(error, .unexpected)

            default:
                XCTFail("Expected error, but got \(result) instead")
            }

            exp.fulfill()
        }

        httpClientSpy.completeWith(data: Data("invalid_data".utf8))
        wait(for: [exp], timeout: 1)
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

    func makeAccountModel() -> AccountModel {
        AccountModel(
            id: "any_id",
            name: "Any Name",
            email: "any.name@example.com",
            password: "pass"
        )
    }

    class HttpClientSpy: HttpPostClient {
        var urls = [URL]()
        var data: Data?
        var completion: ((Result<Data, HttpError>) -> Void)?

        func post(
            to url: URL,
            with data: Data?,
            completion: @escaping (Result<Data, HttpError>) -> Void
        ) {
            self.urls.append(url)
            self.data = data
            self.completion = completion
        }

        func completeWith(error: HttpError) {
            completion?(.failure(error))
        }

        func completeWith(data: Data) {
            completion?(.success(data))
        }
    }
}
