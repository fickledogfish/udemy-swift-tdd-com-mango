import Presentation

func makeSignUpViewModel(
    name: String? = "any_name",
    email: String? = "any_email@example.com",
    password: String? = "any_password",
    passwordConfirmation: String? = "any_password"
) -> SignUpViewModel {
    SignUpViewModel(
        name: name,
        email: email,
        password: password,
        passwordConfirmation: passwordConfirmation
    )
}

func makeRequiredAlertViewModel(fieldName: String) -> AlertViewModel {
    AlertViewModel(
        title: "Falha na validação",
        message: "O campo \(fieldName) é obrigatório"
    )
}

func makeInvalidAlertViewModel(fieldName: String) -> AlertViewModel {
    AlertViewModel(
        title: "Falha na validação",
        message: "O campo \(fieldName) é inválido"
    )
}

func makeErrorAlertViewModel(message: String) -> AlertViewModel {
    AlertViewModel(title: "Erro", message: message)
}

func makeSuccessAlertViewModel(message: String) -> AlertViewModel {
    AlertViewModel(title: "Sucesso", message: message)
}
