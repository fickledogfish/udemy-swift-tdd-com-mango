import XCTest

import Presentation

class SignUpPresenterTests: XCTestCase {
    func test_signUp_should_show_error_message_if_name_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)
        let signUpViewModel = SignUpViewModel(
            name: nil,
            email: "any_email@example.com",
            password: "any_password",
            passwordConfirmation: "any_password"
        )

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "O campo nome é obrigatório"
        ))
    }

    func test_signUp_should_show_error_message_if_email_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)
        let signUpViewModel = SignUpViewModel(
            name: "any_name" ,
            email: nil,
            password: "any_password",
            passwordConfirmation: "any_password"
        )

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "O campo email é obrigatório"
        ))
    }

    func test_signUp_should_show_error_message_if_password_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)
        let signUpViewModel = SignUpViewModel(
            name: "any_name" ,
            email: "any_email@example.com",
            password: nil,
            passwordConfirmation: "any_password"
        )

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "O campo senha é obrigatório"
        ))
    }

    func test_signUp_should_show_error_message_if_passwordConfirmation_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)
        let signUpViewModel = SignUpViewModel(
            name: "any_name" ,
            email: "any_email@example.com",
            password: "any_password",
            passwordConfirmation: nil
        )

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "O campo confirmação de senha é obrigatório"
        ))
    }

    func test_signUp_should_show_error_message_if_passwordConfirmation_does_not_match() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)
        let signUpViewModel = SignUpViewModel(
            name: "any_name" ,
            email: "any_email@example.com",
            password: "any_password",
            passwordConfirmation: "wrong_password"
        )

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(alertViewSpy.viewModel, AlertViewModel(
            title: "Falha na validação",
            message: "Falha ao confirmar senha"
        ))
    }

    func test_signUp_should_call_emailValidator_with_correct_email() {
        let emailValidatorSpy = EmailValidatorSpy()
        let sut = makeSut(emailValidator: emailValidatorSpy)
        let signUpViewModel = SignUpViewModel(
            name: "any_name" ,
            email: "invalid_email@example.com",
            password: "any_password",
            passwordConfirmation: "any_password"
        )

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(emailValidatorSpy.email, signUpViewModel.email)
    }

    func test_signUp_should_show_error_message_if_invalid_email_is_provided() {
        let alertViewSpy = AlertViewSpy()
        let emailValidatorSpy = EmailValidatorSpy()
        let sut = makeSut(alertView: alertViewSpy, emailValidator: emailValidatorSpy)
        let signUpViewModel = SignUpViewModel(
            name: "any_name" ,
            email: "invalid_email@example.com",
            password: "any_password",
            passwordConfirmation: "any_password"
        )

        emailValidatorSpy.isValid = false

        sut.signUp(viewModel: signUpViewModel)

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
    }
}
