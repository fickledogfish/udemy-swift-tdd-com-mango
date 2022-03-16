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
