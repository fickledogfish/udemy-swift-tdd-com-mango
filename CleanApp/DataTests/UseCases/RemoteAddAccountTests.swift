import XCTest
import Data
import Domain

class RemoteAddAccountTests: XCTestCase {
    func test_add_should_call_httpClient_with_correct_url() {
        // Arrange
        let url = makeUrl()
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
        let (sut, httpClientSpy) = makeSut()

        expect(
            sut,
            completeWith: .failure(.unexpected),
            when: {
                httpClientSpy.completeWith(error: .noConnectivity)
            }
        )
    }

    func test_add_should_complete_with_account_if_client_completes_with_valid_data() {
        let (sut, httpClientSpy) = makeSut()
        let account = makeAccountModel()

        expect(
            sut,
            completeWith: .success(account),
            when: {
                httpClientSpy.completeWith(data: account.toData()!)
            }
        )
    }

    func test_add_should_complete_with_error_if_client_completes_with_invalid_data() {
        let (sut, httpClientSpy) = makeSut()

        expect(
            sut,
            completeWith: .failure(.unexpected),
            when: {
                httpClientSpy.completeWith(data: makeInvalidData())
            }
        )
    }

    func test_add_should_not_complete_if_sut_has_been_deallocated() {
        // Arrange
        let httpclientSpy = HttpClientSpy()
        var sut = Optional(RemoteAddAccount(
            url: makeUrl(),
            httpClient: httpclientSpy
        ))
        var result: Result<AccountModel, DomainError>?

        // Act
        sut?.add(account: makeAddAccountModel()) { result = $0 }

        sut = nil
        httpclientSpy.completeWith(error: .noConnectivity)

        // Assert
        XCTAssertNil(result)
    }
}

// MARK: - RemoteAddAccountTests helpers

extension RemoteAddAccountTests {
    func makeSut(
        url: URL = URL(string: "https://xkcd.com")!,
        file: StaticString = #file,
        line: UInt = #line
    ) -> (sut: RemoteAddAccount, httpClientSpy: HttpClientSpy) {
        let httpClientSpy = HttpClientSpy()
        let sut = RemoteAddAccount(url: url, httpClient: httpClientSpy)

        checkMemoryLeak(for: sut, file: file, line: line)
        checkMemoryLeak(for: httpClientSpy, file: file, line: line)

        return (sut, httpClientSpy)
    }

    func checkMemoryLeak(
        for instance: AnyObject,
        file: StaticString = #file,
        line: UInt = #line
    ) {
        addTeardownBlock { [weak instance] in
            XCTAssertNil(
                instance,
                "Memory leaked",
                file: file,
                line: line
            )
        }
    }

    func expect(
        _ sut: RemoteAddAccount,
        completeWith expectedResult: Result<AccountModel, DomainError>,
        when action: () -> Void,
        file: StaticString = #file,
        line: UInt = #line
    ){
        let exp = expectation(description: "waiting")

        sut.add(account: makeAddAccountModel()) { receivedResult in
            switch (expectedResult, receivedResult) {
            case (.failure(let expectedError), .failure(let receivedError)):
                XCTAssertEqual(
                    expectedError,
                    receivedError,
                    file: file,
                    line: line
                )

            case (.success(let expectedAccount), .success(let receivedAccount)):
                XCTAssertEqual(
                    expectedAccount,
                    receivedAccount,
                    file: file,
                    line: line
                )

            default:
                XCTFail(
                    "Expected \(expectedResult), but got \(receivedResult) instead",
                    file: file,
                    line: line
                )
            }

            exp.fulfill()
        }

        action()
        wait(for: [exp], timeout: 1)
    }

    func makeInvalidData() -> Data {
        Data("invalid_data".utf8)
    }

    func makeUrl() -> URL {
        URL(string: "https://curl.se/")!
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
