import XCTest

import Data
import Infra
import Domain

class AddAccountIntegrationTests: XCTestCase {
    // private static var apiBaseUrl: String = "http://localhost:5050"
    private static var apiBaseUrl: String = "https://fordevs.herokuapp.com"

    func test_add_account() {
        let uuid = UUID()

        let alamofireAdapter = AlamofireAdapter()
        let url = URL(string: "\(Self.apiBaseUrl)/api/signup")!
        let sut = RemoteAddAccount(url: url, httpClient: alamofireAdapter)
        let addAccountModel = AddAccountModel(
            name: "\(uuid)",
            email: "\(uuid)@example.com",
            password: "supersecretpassword",
            passwordConfirmation: "supersecretpassword"
        )

        let exp = expectation(description: "waiting")

        sut.add(account: addAccountModel) { result in
            switch result {
            case .failure:
                XCTFail("Expected success, but got \(result)")

            case .success(let account):
                XCTAssertNotNil(account.accessToken)
            }

            exp.fulfill()
        }

        wait(for: [exp], timeout: 10)
    }
}
