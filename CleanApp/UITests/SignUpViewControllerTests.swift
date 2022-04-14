import XCTest

import Presentation
@testable import UI

class SignUpViewControllerTests: XCTestCase {
    func test_loading_is_hidden_on_start() {
        let sut = makeSut()

        XCTAssertEqual(sut.loadingIndicator?.isAnimating, false)
    }

    func test_sut_implements_loadingView() {
        _ = makeSut() as LoadingView
    }

    func test_sut_implements_alertView() {
        _ = makeSut() as AlertView
    }
}

extension SignUpViewControllerTests {
    func makeSut() -> SignUpViewController {
        let sb = UIStoryboard(
            name: "SignUp",
            bundle: Bundle(for: SignUpViewController.self)
        )
        let sut = sb.instantiateViewController(
            withIdentifier: "SignUpViewController"
        ) as! SignUpViewController

        sut.loadViewIfNeeded()

        return sut
    }
}
