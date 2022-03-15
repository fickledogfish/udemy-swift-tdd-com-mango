import Foundation

public struct AccountModel: Model {
    public let id: String
    public let name: String
    public let email: String
    public let password: String

    public init(
        id: String,
        name: String,
        email: String,
        password: String
    ) {
        self.id = id
        self.name = name
        self.email = email
        self.password = password
    }
}
