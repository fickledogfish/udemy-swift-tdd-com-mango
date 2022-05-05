import XCTest

import Main
import PresentationTests

class SignUpComposerTests: XCTestCase {
    func test_ui_presentation_integration() {
        let sut = SignUpComposer.composeControllerWith(
            addAccount: AddAccountSpy()
        )

        checkMemoryLeak(for: sut)
    }
}
