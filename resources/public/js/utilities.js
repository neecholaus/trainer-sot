class ElementHydrator {
    /**
     * Note:
     * The invoker's `elements` and `container` property must be public.
     *
     * @param object invoker
     */
    static hydrate(invoker) {
        if (invoker.container !== undefined && typeof(invoker.container) == 'string') {
            invoker.container = this.#getElement(invoker.container, invoker);
        }

        this.#hydrateElements(invoker);
    }

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

    static #getElement(selector, invoker) {
        if (typeof(invoker.container) !== 'string') {
            return invoker.container.querySelector(selector);
        }
        return document.querySelector(selector);
    }
}