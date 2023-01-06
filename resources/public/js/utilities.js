class ElementHydrator {
    /**
     * Note:
     * The invoker's `elements` and `container` property must be public.
     *
     * @param invoker object
     */
    static hydrate(invoker) {

        // Hydrate container, if applicable, for more narrow element querying
        if (invoker.container !== undefined && typeof(invoker.container) == 'string') {
            invoker.container = this.#getElement(invoker.container, invoker);
        }

        this.#hydrateElements(invoker);
    }

    /**
     * Note:
     * `nestedMappings` allows for stepping through element groupings, such as
     * putting all form inputs inside an object keyed by `formInputs`. When absent,
     * the `invoker's` elements property is used.
     *
     * The invoker is passed on to #getElement in order to use the container for
     * narrow element querying.
     *
     * @param invoker object
     * @param nestedMappings ?object
     */
    static #hydrateElements(invoker, nestedMappings=null) {
        const mappings = nestedMappings || invoker.elements;

        for (let key in mappings) {
            if (typeof(mappings[key]) === 'object') {
                this.#hydrateElements(invoker, mappings[key]);
            } else {
                mappings[key] = this.#getElement(mappings[key], invoker);
            }
        }
    }

    /**
     *
     * @param selector string
     * @param invoker object
     * @return object
     */
    static #getElement(selector, invoker) {
        if (typeof(invoker.container) !== 'string') {
            return invoker.container.querySelector(selector);
        }
        return document.querySelector(selector);
    }
}