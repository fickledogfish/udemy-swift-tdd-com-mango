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
}

// MARK: - AlamofireAdapterTests helpers

extension AlamofireAdapterTests {
    func makeSut() -> AlamofireAdapter {
        let configuration = URLSessionConfiguration.default
        configuration.protocolClasses = [UrlProtocolStub.self]

        let session = Session(configuration: configuration)

        return AlamofireAdapter(session: session)
    }

    func testRequest(
        for url: URL = makeUrl(),
        with data: Data?,
        action: @escaping (URLRequest) -> Void
    ) {
        let sut = makeSut()
        let exp = expectation(description: "waiting")

        sut.post(to: url, with: data)

        UrlProtocolStub.observeRequest { request in
            action(request)
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
