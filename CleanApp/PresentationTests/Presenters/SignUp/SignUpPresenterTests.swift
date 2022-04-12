import XCTest

import Domain
import Presentation

class SignUpPresenterTests: XCTestCase {
    func test_signUp_should_show_error_message_if_name_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = makeSut(alertView: alertViewSpy)

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, makeRequiredAlertViewModel(
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

        alertViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, makeRequiredAlertViewModel(
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

        alertViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, makeRequiredAlertViewModel(
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

        alertViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, makeRequiredAlertViewModel(
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

        alertViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, makeInvalidAlertViewModel(
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

        alertViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, makeInvalidAlertViewModel(
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

        alertViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, makeErrorAlertViewModel(
                message: "Algo inesperado aconteceu, tente novamente em alguns instantes."
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel())
        addAccountSpy.completeWith(error: .unexpected)

        wait(for: [exp], timeout: 1)
    }

    func test_signUp_should_show_success_message_if_addAccount_succeeds() {
        let alertViewSpy = AlertViewSpy()
        let addAccountSpy = AddAccountSpy()
        let sut = makeSut(alertView: alertViewSpy, addAccount: addAccountSpy)

        let exp = expectation(description: "waiting")

        alertViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, makeSuccessAlertViewModel(
                message: "Conta criada com sucesso."
            ))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel())
        addAccountSpy.completeWith(account: makeAccountModel())

        wait(for: [exp], timeout: 1)
    }

    func test_signUp_should_show_loading_before_and_after_addAccount() {
        let addAccountSpy = AddAccountSpy()
        let loadingViewSpy = LoadingViewSpy()
        let sut = makeSut(
            addAccount: addAccountSpy,
            loadingView: loadingViewSpy
        )

        let exp = expectation(description: "waiting")

        loadingViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, LoadingViewModel(isLoading: true))

            exp.fulfill()
        }

        sut.signUp(viewModel: makeSignUpViewModel())

        wait(for: [exp], timeout: 1)

        // ---

        let exp2 = expectation(description: "waiting")

        loadingViewSpy.observe { viewModel in
            XCTAssertEqual(viewModel, LoadingViewModel(isLoading: false))

            exp2.fulfill()
        }

        addAccountSpy.completeWith(error: .unexpected)

        wait(for: [exp2], timeout: 1)
    }
}

// MARK: - SignUpPresenterTests helpers

extension SignUpPresenterTests {
    func makeSut(
        alertView: AlertViewSpy = AlertViewSpy(),
        emailValidator: EmailValidatorSpy = EmailValidatorSpy(),
        addAccount: AddAccountSpy = AddAccountSpy(),
        loadingView: LoadingViewSpy = LoadingViewSpy(),
        file: StaticString = #file,
        line: UInt = #line
    ) -> SignUpPresenter {
        let sut = SignUpPresenter(
            alertView: alertView,
            emailValidator: emailValidator,
            addAccount: addAccount,
            loadingView: loadingView
        )

        checkMemoryLeak(for: sut, file: file, line: line)

        return sut
    }
}
