import Foundation

public struct AccountModel: Model {
    public let accessToken: String

    public init(accessToken: String) {
        self.accessToken = accessToken
    }
}
