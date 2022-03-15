import Foundation

public protocol Model: Codable, Equatable {}

public extension Model {
    func toData() -> Data? {
        try? JSONEncoder().encode(self)
    }
}
