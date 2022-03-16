import Foundation

public extension Data {
    func toModel<T: Decodable>() -> T? {
        try? JSONDecoder().decode(T.self, from: self)
    }

    func toJson() -> [String: Any]? {
        try? JSONSerialization.jsonObject(
            with: self,
            options: .allowFragments
        ) as? [String: Any]
    }
}
