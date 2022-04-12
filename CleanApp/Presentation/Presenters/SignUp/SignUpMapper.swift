import Foundation

import Domain

public final class SignUpMapper {
    static func toAddAccountModel(
        viewModel: SignUpViewModel
    ) -> AddAccountModel {
        AddAccountModel(
            name: viewModel.name!,
            email: viewModel.email!,
            password: viewModel.password!,
            passwordConfirmation: viewModel.passwordConfirmation!
        )
    }
}
