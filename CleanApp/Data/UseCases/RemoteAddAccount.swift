import Foundation
import Domain

public final class RemoteAddAccount {
    private let url: URL
    private let httpClient: HttpPostClient

    public init(url: URL, httpClient: HttpPostClient) {
        self.url = url
        self.httpClient = httpClient
    }

    public func add(
        account: AddAccountModel,
        completion: @escaping (DomainError) -> Void
    ) {
        httpClient.post(to: url, with: account.toData()) { _ in
            completion(.unexpected)
        }
    }
}
