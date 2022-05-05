import Foundation

import Data
import Infra
import Domain

final class UseCaseFactory {
    private static let httpClient = AlamofireAdapter()
    private static let apiBaseUrl = "https://fordevs.herokuapp.com/api"

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
