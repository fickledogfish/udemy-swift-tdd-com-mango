import XCTest
import Alamofire

class AlamofireAdapter {
    private let session: Session

    init(session: Session = .default) {
        self.session = session
    }

    func post(to url: URL) {
        session
            .request(url, method: .post)
            .resume()
    }
}

class AlamofireAdapterTests: XCTestCase {
    func test_() {
        // Arrange
        let url = makeUrl()
        let configuration = URLSessionConfiguration.default
        configuration.protocolClasses = [UrlProtocolStub.self]
        let session = Session(configuration: configuration)

        let exp = expectation(description: "waiting")

        let sut = AlamofireAdapter(session: session)

        // Act
        sut.post(to: url)

        // Assert
        UrlProtocolStub.observeRequest { request in
            XCTAssertEqual(url, request.url)
            XCTAssertEqual("POST", request.httpMethod)
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
