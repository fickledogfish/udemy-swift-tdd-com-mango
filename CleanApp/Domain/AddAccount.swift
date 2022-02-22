import Foundation

protocol AddAccount {
    func add(
        addAccountModel: AddAccountModel,
        completion: @escaping (Result<AccountModel, Error>) -> Void
    )
}

struct AddAccountModel {
    let name: String
    let email: String
    let password: String
    let passwordConfirmation: String
}
