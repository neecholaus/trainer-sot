class SidebarActivator {
    /**
     * Note:
     * I am responsible for indicating which sidebar link is currently
     * being viewed in the trainer panel.
     */
    elements = {
        sidebarBtns: 'all:#sidebar a.sidebar-btn',
    }

    constructor() {
        ElementHydrator.hydrate(this);
    
        if (this.elements.sidebarBtns.length) {
            this.elements.sidebarBtns.forEach(btn => {
                if (btn.href === document.location.href) {
                    btn.classList.add('active');
                }
            });
        }
    }
}
