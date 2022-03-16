import XCTest

extension XCTestCase {
    func checkMemoryLeak(
        for instance: AnyObject,
        file: StaticString = #file,
        line: UInt = #line
    ) {
        addTeardownBlock { [weak instance] in
            XCTAssertNil(
                instance,
                "Memory leaked",
                file: file,
                line: line
            )
        }
    }
}
