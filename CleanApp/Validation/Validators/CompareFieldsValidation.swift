import Presentation

public final class CompareFieldsValidation: Validation {
    private let fieldName: String
    private let fieldNameToCompare: String
    private let fieldLabel: String

    public init(
        fieldName: String,
        fieldNameToCompare: String,
        fieldLabel: String
    ) {
        self.fieldName = fieldName
        self.fieldNameToCompare = fieldNameToCompare
        self.fieldLabel = fieldLabel
    }

    public func validate(data: [String: Any]?) -> String? {
        guard
            let fieldName = data?[fieldName] as? String,
            let fieldNameToCompare = data?[fieldNameToCompare] as? String,
            fieldName == fieldNameToCompare
        else {
            return "O campo \(fieldLabel) é inválido"
        }

        return nil
    }
}

extension CompareFieldsValidation: Equatable {
    public static func == (
        lhs: CompareFieldsValidation,
        rhs: CompareFieldsValidation
    ) -> Bool {
        lhs.fieldName == rhs.fieldName &&
        lhs.fieldNameToCompare == rhs.fieldNameToCompare &&
        lhs.fieldLabel == rhs.fieldLabel
    }
}
