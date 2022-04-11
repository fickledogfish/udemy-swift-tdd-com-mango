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

func makeAddAccountModel() -> AddAccountModel {
    AddAccountModel(
        name: "any_name",
        email: "any_email@example.com",
        password: "any_password",
        passwordConfirmation: "any_password"
    )
}
