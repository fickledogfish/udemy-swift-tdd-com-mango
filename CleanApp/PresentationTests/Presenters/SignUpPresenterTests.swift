import XCTest

import Presentation

class SignUpPresenterTests: XCTestCase {
    func test_signUp_should_show_error_message_if_name_is_not_provided() {
        let (sut, alertViewSpy, _) = makeSut()
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
        let (sut, alertViewSpy, _) = makeSut()
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
        let (sut, alertViewSpy, _) = makeSut()
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
        let (sut, alertViewSpy, _) = makeSut()
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
        let (sut, alertViewSpy, _) = makeSut()
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
        let (sut, _, emailValidatorSpy) = makeSut()
        let signUpViewModel = SignUpViewModel(
            name: "any_name" ,
            email: "invalid_email@example.com",
            password: "any_password",
            passwordConfirmation: "any_password"
        )

        sut.signUp(viewModel: signUpViewModel)

        XCTAssertEqual(emailValidatorSpy.email, signUpViewModel.email)
    }
}

// MARK: - SignUpPresenterTests helpers

extension SignUpPresenterTests {
    func makeSut() -> (
        sut: SignUpPresenter,
        alertViewSpy: AlertViewSpy,
        emailValidatorSpy: EmailValidatorSpy
    ) {
        let alertViewSpy = AlertViewSpy()
        let emailValidatorSpy = EmailValidatorSpy()
        let sut = SignUpPresenter(
            alertView: alertViewSpy,
            emailValidator: emailValidatorSpy
        )

        return (sut, alertViewSpy, emailValidatorSpy)
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
