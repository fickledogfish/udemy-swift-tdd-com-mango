import Domain
import UI

public final class SignUpComposer {
    public static func composeControllerWith(
        addAccount: AddAccount
    ) -> SignUpViewController {
        ControllerFactory.makeSignUp(addAccount: addAccount)
    }
}
