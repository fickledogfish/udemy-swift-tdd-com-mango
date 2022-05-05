import XCTest

import Main
@testable import PresentationTests

class SignUpIntegrationTests: XCTestCase {
    func test_ui_presentation_integration() {
        let sut = SignUpComposer.composeControllerWith(
            addAccount: AddAccountSpy()
        )

        checkMemoryLeak(for: sut)
    }
}
