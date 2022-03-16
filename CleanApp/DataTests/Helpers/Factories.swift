import Foundation

func makeInvalidData() -> Data {
    Data("invalid_data".utf8)
}

func makeValidData() -> Data {
    Data("{\"name\":\"Fabian\"}".utf8)
}

func makeUrl() -> URL {
    URL(string: "https://curl.se/")!
}

func makeError() -> Error {
    NSError(domain: "any_error", code: 0)
}

func makeHttpResponse(statusCode: Int = 200) -> HTTPURLResponse {
    HTTPURLResponse(
        url: makeUrl(),
        statusCode: statusCode,
        httpVersion: nil,
        headerFields: nil
    )!
}
