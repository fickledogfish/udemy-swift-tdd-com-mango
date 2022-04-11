import XCTest

import Domain
import Presentation

class SignUpPresenterTests: XCTestCase {
    func test_signUp_should_show_error_message_if_name_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { [weak self] viewModel in
            XCTAssertEqual(viewModel, self?.makeRequiredAlertViewModel(
                fieldName: "nome"
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel(name: nil))

        wait(for: [exp], timeout: 1)
    }

    func test_signUp_should_show_error_message_if_email_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { [weak self] viewModel in
            XCTAssertEqual(viewModel, self?.makeRequiredAlertViewModel(
                fieldName: "email"
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel(email: nil))

        wait(for: [exp], timeout: 1)
    }

    func test_signUp_should_show_error_message_if_password_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { [weak self] viewModel in
            XCTAssertEqual(viewModel, self?.makeRequiredAlertViewModel(
                fieldName: "senha"
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel(password: nil))

        wait(for: [exp], timeout: 1)
    }

    func test_signUp_should_show_error_message_if_passwordConfirmation_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { [weak self] viewModel in
            XCTAssertEqual(viewModel, self?.makeRequiredAlertViewModel(
                fieldName: "confirmação de senha"
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel(passwordConfirmation: nil))

        wait(for: [exp], timeout: 1)
    }

    func test_signUp_should_show_error_message_if_passwordConfirmation_does_not_match() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { [weak self] viewModel in
            XCTAssertEqual(viewModel, self?.makeInvalidAlertViewModel(
                fieldName: "confirmar senha"
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel(passwordConfirmation: "wrong_password"))

        wait(for: [exp], timeout: 1)
    }

    func test_signUp_should_show_error_message_if_invalid_email_is_provided() {
        let alertViewSpy = AlertViewSpy()
        let emailValidatorSpy = EmailValidatorSpy()
        let sut = makeSut(alertView: alertViewSpy, emailValidator: emailValidatorSpy)

        emailValidatorSpy.simulateInvalidEmail()

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { [weak self] viewModel in
            XCTAssertEqual(viewModel, self?.makeInvalidAlertViewModel(
                fieldName: "email"
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel())

        wait(for: [exp], timeout: 1)
    }

    func test_signUp_should_call_emailValidator_with_correct_email() {
        let emailValidatorSpy = EmailValidatorSpy()
        let sut = makeSut(emailValidator: emailValidatorSpy)
        let signUpViewModel = makeSignUpViewModel()

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(emailValidatorSpy.email, signUpViewModel.email)
    }

    func test_signUp_should_call_addAccount_with_correct_values() {
        let addAccountSpy = AddAccountSpy()
        let sut = makeSut(addAccount: addAccountSpy)

        sut.signUp(viewModel: makeSignUpViewModel())

        XCTAssertEqual(addAccountSpy.addAccountModel, makeAddAccountModel())
    }

    func test_signUp_should_show_error_message_if_addAccount_fails() {
        let alertViewSpy = AlertViewSpy()
        let addAccountSpy = AddAccountSpy()
        let sut = makeSut(alertView: alertViewSpy, addAccount: addAccountSpy)

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { [weak self] viewModel in
            XCTAssertEqual(viewModel, self?.makeErrorAlertViewModel(
                message: "Algo inesperado aconteceu, tente novamente em alguns instantes."
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel())
        addAccountSpy.completeWith(error: .unexpected)

        wait(for: [exp], timeout: 1)
    }
}

// MARK: - SignUpPresenterTests helpers

extension SignUpPresenterTests {
    func makeSut(
        alertView: AlertViewSpy = AlertViewSpy(),
        emailValidator: EmailValidatorSpy = EmailValidatorSpy(),
        addAccount: AddAccountSpy = AddAccountSpy()
    ) -> SignUpPresenter {
        SignUpPresenter(
            alertView: alertView,
            emailValidator: emailValidator,
            addAccount: addAccount
        )
    }

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

    class AlertViewSpy: AlertView {
        var emit: ((AlertViewModel) -> Void)? = nil

        func observe(completion: @escaping (AlertViewModel) -> Void) {
            self.emit = completion
        }

        func showMessage(viewModel: AlertViewModel) {
            self.emit?(viewModel)
        }
    }

    class EmailValidatorSpy: EmailValidator {
        var isValid = true
        var email: String?

        func isValid(email: String) -> Bool {
            self.email = email

            return isValid
        }

        func simulateInvalidEmail() {
            self.isValid = false
        }
    }

    class AddAccountSpy: AddAccount {
        var addAccountModel: AddAccountModel?
        var completion: ((Result<AccountModel, DomainError>) -> Void)? = nil

        func add(
            account: AddAccountModel,
            completion: @escaping (Result<AccountModel, DomainError>) -> Void
        ) {
            self.addAccountModel = account
            self.completion = completion
        }

        func completeWith(error: DomainError) {
            completion?(.failure(error))
        }
    }
}
