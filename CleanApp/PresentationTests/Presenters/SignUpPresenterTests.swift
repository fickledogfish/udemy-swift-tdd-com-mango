import XCTest

import Presentation

class SignUpPresenterTests: XCTestCase {
    func test_signUp_should_show_error_message_if_name_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(name: nil))

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "O campo nome é obrigatório"
        ))
    }

    func test_signUp_should_show_error_message_if_email_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(email: nil))

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "O campo email é obrigatório"
        ))
    }

    func test_signUp_should_show_error_message_if_password_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(password: nil))

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "O campo senha é obrigatório"
        ))
    }

    func test_signUp_should_show_error_message_if_passwordConfirmation_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(passwordConfirmation: nil))

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "O campo confirmação de senha é obrigatório"
        ))
    }

    func test_signUp_should_show_error_message_if_passwordConfirmation_does_not_match() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        sut.signUp(viewModel: makeSignUpViewModel(passwordConfirmation: "wrong_password"))

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "Falha ao confirmar senha"
        ))
    }

    func test_signUp_should_call_emailValidator_with_correct_email() {
        let emailValidatorSpy = EmailValidatorSpy()
        let sut = makeSut(emailValidator: emailValidatorSpy)
        let signUpViewModel = makeSignUpViewModel()

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(emailValidatorSpy.email, signUpViewModel.email)
    }

    func test_signUp_should_show_error_message_if_invalid_email_is_provided() {
        let alertViewSpy = AlertViewSpy()
        let emailValidatorSpy = EmailValidatorSpy()
        let sut = makeSut(alertView: alertViewSpy, emailValidator: emailValidatorSpy)

        emailValidatorSpy.simulateInvalidEmail()

        sut.signUp(viewModel: makeSignUpViewModel())

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "Email inválido"
        ))
    }
}

// MARK: - SignUpPresenterTests helpers

extension SignUpPresenterTests {
    func makeSut(
        alertView: AlertViewSpy = AlertViewSpy(),
        emailValidator: EmailValidatorSpy = EmailValidatorSpy()
    ) -> SignUpPresenter {
        SignUpPresenter(
            alertView: alertView,
            emailValidator: emailValidator
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
            isValid = false
        }
    }
}
