import XCTest

class SignUpPresenter {
    private let alertView: AlertView

    init(alertView: AlertView) {
        self.alertView = alertView
    }

    func signUp(viewModel: SignUpViewModel) {
        if viewModel.name == nil || viewModel.name!.isEmpty {
            alertView.showMessage(viewModel: AlertViewModel(
                title: "Falha na validação",
                message: "O campo nome é obrigatório"
            ))
        }
    }
}

struct SignUpViewModel {
    var name: String?
    var email: String?
    var password: String?
    var passwordConfirmation: String?

    public init(
        name: String?,
        email: String?,
        password: String?,
        passwordConfirmation: String?
    ) {
        self.name = name
        self.email = email
        self.password = password
        self.passwordConfirmation = passwordConfirmation
    }
}

struct AlertViewModel: Equatable {
    var title: String
    var message: String
}

protocol AlertView {
    func showMessage(viewModel: AlertViewModel)
}

class SignUpPresenterTests: XCTestCase {
    func test_signUp_should_show_error_message_if_name_is_not_provided() {
        let alertViewSpy = AlertViewSpy()
        let sut = SignUpPresenter(alertView: alertViewSpy)
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
}

// MARK: - SignUpPresenterTests helpers

extension SignUpPresenterTests {
    class AlertViewSpy: AlertView {
        var viewModel: AlertViewModel?

        func showMessage(viewModel: AlertViewModel) {
            self.viewModel =  viewModel
        }
    }
}
