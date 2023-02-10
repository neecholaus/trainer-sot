// todo - listener identifier so listener can remove self from listen queue
class EventPipeline {
    // Arrays of objects indexed by the event name. (First In First Out)
    events = {};

    // Arrays of Promise resolve functions indexed by the event name. (First In First Out)
    listeners = {};

    // Push data to a specific queue. Check for existing listener and resolve if found.
    push(to, data) {
        // Some integrity protections
        if (typeof(to) !== 'string' || typeof(data) !== 'object') {
            return false;
        }

        if (! this.events.hasOwnProperty(to)) {
            this.events[to] = [];
        }

        this.events[to].push(data);

        // If listener is waiting for event of same type, hydrate and resolve
        if (this.listeners.hasOwnProperty(to) && this.listeners[to].length) {
            // Resolve listener
            this.listeners[to][0](data);

            // Remove listener
            this.listeners[to].shift();

            // Remove event
            this.events[to].shift();
        }

        return true;
    }

    // Return promise, which resolves when the function inside the ListenerObject is called.
    // The ListenerObject is pushed onto the appropriate array in the listeners object.
    // When an event comes in, the first ListenerObject in that array is hydrated, and the
    // resolve func is called.
    async listen(eventName) {
        if (typeof(eventName) !== 'string') {
            return;
        }

        // Init listener key
        if (! this.listeners[eventName]) {
            this.listeners[eventName] = [];
        }

        return new Promise(resolve => {
            // Check for held event first, and resolve with that if available
            if (this.events[eventName] && this.events[eventName].length) {
                resolve(this.events[eventName].shift());
            }

            // Add self to queue of listeners for next available event of same type
            this.listeners[eventName].push(resolve);
        });
    }

    // todo - method to check state (how many unhandled events) (have them expire?)
}

window.EventPipeline = new EventPipeline();