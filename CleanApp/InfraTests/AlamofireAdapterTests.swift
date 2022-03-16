import XCTest
import Alamofire

class AlamofireAdapter {
    private let session: Session

    init(session: Session = .default) {
        self.session = session
    }

    func post(to url: URL, with data: Data?) {
        let json = data == nil ? nil : try? JSONSerialization.jsonObject(
                with: data!,
                options: .allowFragments
            ) as? [String: Any]

        session
            .request(
                url,
                method: .post,
                parameters: json,
                encoding: JSONEncoding.default
            )
            .resume()
    }
}

class AlamofireAdapterTests: XCTestCase {
    func test_post_should_make_request_with_valid_url_and_method() {
        // Arrange
        let url = makeUrl()
        let configuration = URLSessionConfiguration.default
        configuration.protocolClasses = [UrlProtocolStub.self]
        let session = Session(configuration: configuration)

        let exp = expectation(description: "waiting")

        let sut = AlamofireAdapter(session: session)

        // Act
        sut.post(to: url, with: makeValidData())

        // Assert
        UrlProtocolStub.observeRequest { request in
            XCTAssertEqual(url, request.url)
            XCTAssertEqual("POST", request.httpMethod)
            XCTAssertNotNil(request.httpBodyStream)

            exp.fulfill()
        }

        wait(for: [exp], timeout: 1)
    }

    func test_post_should_make_request_with_no_data() {
        // Arrange
        let url = makeUrl()
        let configuration = URLSessionConfiguration.default
        configuration.protocolClasses = [UrlProtocolStub.self]
        let session = Session(configuration: configuration)

        let exp = expectation(description: "waiting")

        let sut = AlamofireAdapter(session: session)

        // Act
        sut.post(to: url, with: nil)

        // Assert
        UrlProtocolStub.observeRequest { request in
            XCTAssertNil(request.httpBodyStream)

            exp.fulfill()
        }

        wait(for: [exp], timeout: 1)
    }
}

class UrlProtocolStub: URLProtocol {
    static var emit: ((URLRequest) -> Void)?

    static func observeRequest(completion: @escaping (URLRequest) -> Void) {
        UrlProtocolStub.emit = completion
    }

    open class override func canInit(with request: URLRequest) -> Bool {
        true
    }

    open class override func canonicalRequest(for request: URLRequest) -> URLRequest {
        request
    }

    override func startLoading() {
        UrlProtocolStub.emit?(request)
    }

    override func stopLoading() {
    }
}
