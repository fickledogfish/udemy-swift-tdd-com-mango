import Foundation

import Data
import Infra
import Domain

final class UseCaseFactory {
    private static let httpClient = AlamofireAdapter()
    private static let apiBaseUrl = Environment.variable(.apiBaseUrl)

    private static func makeUrl(path: String) -> URL {
        URL(string: "\(apiBaseUrl)/\(path)")!
    }

    static func makeRemoteAddAccount() -> AddAccount {
        return RemoteAddAccount(
            url: makeUrl(path: "signup"),
            httpClient: httpClient
        )
    }
}
