import Foundation

import Domain

public final class SignUpPresenter {
    private let alertView: AlertView
    private let emailValidator: EmailValidator
    private let addAccount: AddAccount
    private let loadingView: LoadingView

    public init(
        alertView: AlertView,
        emailValidator: EmailValidator,
        addAccount: AddAccount,
        loadingView: LoadingView
    ) {
        self.alertView = alertView
        self.emailValidator = emailValidator
        self.addAccount = addAccount
        self.loadingView = loadingView
    }

    public func signUp(viewModel: SignUpViewModel) {
        if let message = validate(viewModel: viewModel) {
            alertView.showMessage(viewModel: AlertViewModel(
                title: "Falha na validação",
                message: message
            ))
        } else {
            let addAccountModel = SignUpMapper.toAddAccountModel(viewModel: viewModel)

            loadingView.display(viewModel: LoadingViewModel(isLoading: true))

            addAccount.add(account: addAccountModel) { [weak self] result in
                guard let this = self else { return }

                switch result {
                case .failure:
                    this.alertView.showMessage(viewModel: AlertViewModel(
                        title: "Erro",
                        message: "Algo inesperado aconteceu, tente novamente em alguns instantes."
                    ))

                case .success:
                    this.alertView.showMessage(viewModel: AlertViewModel(
                        title: "Sucesso",
                        message: "Conta criada com sucesso."
                    ))
                }

                this.loadingView.display(viewModel: LoadingViewModel(isLoading: false))
            }
        }
    }

    private func validate(viewModel: SignUpViewModel) -> String? {
        if viewModel.name == nil || viewModel.name!.isEmpty {
            return "O campo nome é obrigatório"
        } else if viewModel.email == nil || viewModel.email!.isEmpty {
            return "O campo email é obrigatório"
        } else if viewModel.password == nil || viewModel.password!.isEmpty {
            return  "O campo senha é obrigatório"
        } else if viewModel.passwordConfirmation == nil || viewModel.passwordConfirmation!.isEmpty {
            return "O campo confirmação de senha é obrigatório"
        } else if viewModel.password != viewModel.passwordConfirmation {
            return "O campo confirmar senha é inválido"
        } else if !emailValidator.isValid(email: viewModel.email!) {
            return "O campo email é inválido"
        }

        return nil
    }
}
