import XCTest

import Data
import Infra
import Domain

class AddAccountIntegrationTests: XCTestCase {
    // https://fordevs.herokuapp.com/api
    private static let apiBaseUrl = "http://localhost:5050/api"

    func test_add_account() {
        let uuid = UUID()

        let alamofireAdapter = AlamofireAdapter()
        let url = URL(string: "\(Self.apiBaseUrl)/signup")!
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
                XCTAssertNotNil(account.id)
                XCTAssertEqual(account.name, addAccountModel.name)
                XCTAssertEqual(account.email, addAccountModel.email)
            }

            exp.fulfill()
        }

        wait(for: [exp], timeout: 10)
    }
}
