import Foundation

public protocol Model: Encodable {}

public extension Model {
    func toData() -> Data? {
        try? JSONEncoder().encode(self)
    }
}
