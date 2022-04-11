import XCTest

import Domain
import Presentation

class SignUpPresenterTests: XCTestCase {
    func test_signUp_should_show_error_message_if_name_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(name: nil))

        XCTAssertEqual(
            alertViewSpy.viewModel,
            makeRequiredAlertViewModel(fieldName: "nome")
        )
    }

    func test_signUp_should_show_error_message_if_email_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(email: nil))

        XCTAssertEqual(
            alertViewSpy.viewModel,
            makeRequiredAlertViewModel(fieldName: "email")
        )
    }

    func test_signUp_should_show_error_message_if_password_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(password: nil))

        XCTAssertEqual(
            alertViewSpy.viewModel,
            makeRequiredAlertViewModel(fieldName: "senha")
        )
    }

    func test_signUp_should_show_error_message_if_passwordConfirmation_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(passwordConfirmation: nil))

        XCTAssertEqual(
            alertViewSpy.viewModel,
            makeRequiredAlertViewModel(fieldName: "confirmação de senha")
        )
    }

    func test_signUp_should_show_error_message_if_passwordConfirmation_does_not_match() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(passwordConfirmation: "wrong_password"))

        XCTAssertEqual(
            alertViewSpy.viewModel,
            makeInvalidAlertViewModel(fieldName: "confirmar senha")
        )
    }

    func test_signUp_should_show_error_message_if_invalid_email_is_provided() {
        let alertViewSpy = AlertViewSpy()
        let emailValidatorSpy = EmailValidatorSpy()
        let sut = makeSut(alertView: alertViewSpy, emailValidator: emailValidatorSpy)

        emailValidatorSpy.simulateInvalidEmail()

        sut.signUp(viewModel: makeSignUpViewModel())

        XCTAssertEqual(
            alertViewSpy.viewModel,
            makeInvalidAlertViewModel(fieldName: "email")
        )
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

    class AlertViewSpy: AlertView {
        var viewModel: AlertViewModel?

        func showMessage(viewModel: AlertViewModel) {
            self.viewModel =  viewModel
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

        func add(
            account: AddAccountModel,
            completion: @escaping (Result<AccountModel, DomainError>) -> Void
        ) {
            self.addAccountModel = account
        }
    }
}
