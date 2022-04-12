import Presentation

class LoadingViewSpy: LoadingView {
    var emit: ((LoadingViewModel) -> Void)? = nil

    func observe(completion: @escaping (LoadingViewModel) -> Void) {
        self.emit = completion
    }

    func display(viewModel: LoadingViewModel) {
        self.emit?(viewModel)
    }
}
