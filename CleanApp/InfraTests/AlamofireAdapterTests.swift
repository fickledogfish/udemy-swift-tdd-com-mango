import XCTest
import Alamofire
import Data

class AlamofireAdapter {
    private let session: Session

    init(session: Session = .default) {
        self.session = session
    }

    func post(
        to url: URL,
        with data: Data?,
        completion: @escaping (Result<Data, HttpError>) -> Void
    ) {
        session
            .request(
                url,
                method: .post,
                parameters: data?.toJson(),
                encoding: JSONEncoding.default
            )
            .responseData { dataResponse in
                guard let _ = dataResponse.response?.statusCode else {
                    return completion(.failure(.noConnectivity))
                }

                switch dataResponse.result {
                case .failure:
                    completion(.failure(.noConnectivity))

                case .success(let data):
                    completion(.success(data))
                }
            }
    }
}

class AlamofireAdapterTests: XCTestCase {
    func test_post_should_make_request_with_valid_url_and_method() {
        let url = makeUrl()

        testRequest(for: url, with: makeValidData()) { request in
            XCTAssertEqual(url, request.url)
            XCTAssertEqual("POST", request.httpMethod)
            XCTAssertNotNil(request.httpBodyStream)
        }
    }

    func test_post_should_make_request_with_no_data() {
        testRequest(with: nil) { request in
            XCTAssertNil(request.httpBodyStream)
        }
    }

    func test_post_should_complete_with_error_when_request_completes_with_error() {
        expect(
            .failure(.noConnectivity),
            when: (data: nil, response: nil, error: makeError())
        )
    }

    func test_post_should_complete_with_error_on_all_invalid_cases() {
        expect(.failure(.noConnectivity), when: (data: makeValidData(), response: makeHttpResponse(), error: makeError()))
        expect(.failure(.noConnectivity), when: (data: makeValidData(), response: nil, error: makeError()))
        expect(.failure(.noConnectivity), when: (data: makeValidData(), response: nil, error: nil))
        expect(.failure(.noConnectivity), when: (data: nil, response: makeHttpResponse(), error: makeError()))
        expect(.failure(.noConnectivity), when: (data: nil, response: makeHttpResponse(), error: nil))
        expect(.failure(.noConnectivity), when: (data: nil, response: nil, error: nil))
    }
}

// MARK: - AlamofireAdapterTests helpers

extension AlamofireAdapterTests {
    func makeSut(
        file: StaticString = #file,
        line: UInt = #line
    ) -> AlamofireAdapter {
        let configuration = URLSessionConfiguration.default
        configuration.protocolClasses = [UrlProtocolStub.self]

        let session = Session(configuration: configuration)
        let sut = AlamofireAdapter(session: session)

        checkMemoryLeak(for: sut, file: file, line: line)

        return sut
    }

    func testRequest(
        for url: URL = makeUrl(),
        with data: Data?,
        action: @escaping (URLRequest) -> Void
    ) {
        let sut = makeSut()
        var request: URLRequest?
        let exp = expectation(description: "waiting")

        sut.post(to: url, with: data) { _ in
            exp.fulfill()
        }

        UrlProtocolStub.observeRequest { request = $0 }

        wait(for: [exp], timeout: 1)
        action(request!)
    }

    func expect(
        _ expectedResult: Result<Data, HttpError>,
        when stub: (data: Data?, response: HTTPURLResponse?, error: Error?),
        file: StaticString = #file,
        line: UInt = #line
    ) {
        let sut = makeSut()
        let exp = expectation(description: "waiting")

        UrlProtocolStub.simulate(
            data: stub.data,
            response: stub.response,
            error: stub.error
        )

        sut.post(to: makeUrl(), with: makeValidData()) { receivedResult in
            switch (expectedResult, receivedResult) {
            case (.failure(let expectedError), .failure(let receivedError)):
                XCTAssertEqual(
                    expectedError, receivedError,
                    file: file,
                    line: line
                )

            case (.success(let expectedData), .success(let receivedData)):
                XCTAssertEqual(
                    expectedData, receivedData,
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

        wait(for: [exp], timeout: 1)
    }
}

class UrlProtocolStub: URLProtocol {
    static var emit: ((URLRequest) -> Void)?

    static var data: Data?
    static var response: HTTPURLResponse?
    static var error: Error?

    static func observeRequest(completion: @escaping (URLRequest) -> Void) {
        UrlProtocolStub.emit = completion
    }

    static func simulate(data: Data?, response: HTTPURLResponse?, error: Error?) {
        Self.data = data
        Self.response = response
        Self.error = error
    }

    open class override func canInit(with request: URLRequest) -> Bool {
        true
    }

    open class override func canonicalRequest(for request: URLRequest) -> URLRequest {
        request
    }

    override func startLoading() {
        UrlProtocolStub.emit?(request)

        if let data = UrlProtocolStub.data {
            client?.urlProtocol(self, didLoad: data)
        }

        if let response = UrlProtocolStub.response {
            client?.urlProtocol(
                self,
                didReceive: response,
                cacheStoragePolicy: .notAllowed
            )
        }

        if let error = UrlProtocolStub.error {
            client?.urlProtocol(self, didFailWithError: error)
        }

        client?.urlProtocolDidFinishLoading(self)
    }

    override func stopLoading() {
    }
}
