class Validator {
    static asEmail(toValidate) {
        if (typeof toValidate !== 'string') {
            return false;
        }

        if (! toValidate.match(/^\w+@\w*.\w{2,}$/)) {
            return false;
        }

        return true;
    }
}
