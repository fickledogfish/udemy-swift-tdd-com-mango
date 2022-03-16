import Foundation
import Domain

func makeAccountModel() -> AccountModel {
    AccountModel(
        id: "any_id",
        name: "Any Name",
        email: "any.name@example.com",
        password: "pass"
    )
}

