import Module from "../interfaces/Module";
import MouseData from "../interfaces/MouseData";

/**
 * Module which tracks the Mouse Movements
 */
export default abstract class MouseMovement implements Module {

    /**
     * MousePath
     */
    protected path: Array<MouseData> = new Array<MouseData>();

    /**
     * Handlers to remove the Listener
     */
    protected removeHandler: Array<() => void> = new Array<() => void>();

    /**
     * Name of the Module
     */
    public name = "Mouse";

    /**
     * If the mouse is currently being tracked
     */
    protected running = false;

    /**
     * will be called when the Captcha is rendered
     * does nothing
     */
    public start(): void { }

    /**
     * will be called when the Captcha is solved
     * does nothing
     */
    public stop(): void { }

    /**
     * returns the collected MousePath
     * @returns MouseMovement Data
     */
    public getCollectedData(): Array<MouseData> {
        return this.path;
    }

    /**
     * Registers the Listener to the Document
     * @param listener: Listener to call when the mouse is moved, where this is binded to the MouseMovement Class
     * @returns Method to remove the Listener
     */
    protected registerMouseListener(listener: (ev: Event) => any): () => void {
        const l = listener.bind(this);
        document.addEventListener("mousemove", l);
        return () => { document.removeEventListener("mousemove", l); };
    }

    /**
     * Executes and removes all RemoveHandlers
     */
    protected stopTracking() {
        while (this.removeHandler.length > 0) {
            const f = this.removeHandler.pop();
            if (f) {
                f();
            }
        }
    }

    /**
     * Adds a MousePoint from the given MouseEvent to the MousePath
     * @param ev MouseMoveEvent
     */
    protected onMouseMove(ev: Event): void {
        const e = ev as MouseEvent;
        this.path.push({
            x: e.screenX,
            y: e.screenY,
            time: new Date().getTime(),
        });
    }

}
