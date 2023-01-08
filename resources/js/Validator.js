class Validator {
    static asEmail(toValidate) {
        if (typeof toValidate !== 'string') {
            return false;
        }

        // todo - actually validate

        return true;
    }
}
