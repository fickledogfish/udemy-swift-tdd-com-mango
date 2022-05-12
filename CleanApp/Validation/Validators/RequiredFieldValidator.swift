import Presentation

public final class RequiredFieldValidation: Validation {
    private let fieldName: String
    private let fieldLabel: String

    public init(fieldName: String, fieldLabel: String) {
        self.fieldName = fieldName
        self.fieldLabel = fieldLabel
    }

    public func validate(data: [String: Any]?) -> String? {
        guard
            let fieldName = data?[fieldName] as? String,
            !fieldName.isEmpty
        else {
            return "O campo \(fieldLabel) é obrigatório"
        }

        return nil
    }
}

extension RequiredFieldValidation: Equatable {
    public static func == (
        lhs: RequiredFieldValidation,
        rhs: RequiredFieldValidation
    ) -> Bool {
        lhs.fieldLabel == rhs.fieldLabel &&
        lhs.fieldName == rhs.fieldName
    }
}
