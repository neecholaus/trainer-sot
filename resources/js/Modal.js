class Modal {
    elements = {
        modals: 'all:div.modal',
    }

    constructor() {
        ElementHydrator.hydrate(this);

        this.elements.modals.forEach(el => {
            // Don't close modal when clicking inside modal content
            el.querySelector('div.modal-content').addEventListener('click', event => {
                event.stopPropagation();
            });

            // Close/hide modal when clicking in the content container
            el.addEventListener('click', () => {
                el.style.display = 'none';
            });
        });
    }

}

// Initialize Modal class on page load (after all modal elements are in DOM)
document.addEventListener('DOMContentLoaded', () => {
    new Modal();
});
