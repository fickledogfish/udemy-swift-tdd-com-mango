import Foundation

import Domain

public final class SignUpPresenter {
    private let alertView: AlertView
    private let addAccount: AddAccount
    private let loadingView: LoadingView
    private let validation: Validation

    public init(
        alertView: AlertView,
        addAccount: AddAccount,
        loadingView: LoadingView,
        validation: Validation
    ) {
        self.alertView = alertView
        self.addAccount = addAccount
        self.loadingView = loadingView
        self.validation = validation
    }

    public func signUp(viewModel: SignUpViewModel) {
        if let message = validation.validate(data: viewModel.toJson()) {
            alertView.showMessage(viewModel: AlertViewModel(
                title: "Falha na validação",
                message: message
            ))
        } else {
            let addAccountModel = viewModel.toAddAccountModel()

            loadingView.display(viewModel: LoadingViewModel(isLoading: true))

            addAccount.add(account: addAccountModel) { [weak self] result in
                guard let this = self else { return }

                this.loadingView.display(viewModel: LoadingViewModel(isLoading: false))

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
            }
        }
    }
}
