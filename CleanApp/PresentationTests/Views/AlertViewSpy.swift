import Presentation

class AlertViewSpy: AlertView {
    var emit: ((AlertViewModel) -> Void)? = nil

    func observe(completion: @escaping (AlertViewModel) -> Void) {
        self.emit = completion
    }

    func showMessage(viewModel: AlertViewModel) {
        self.emit?(viewModel)
    }
}

