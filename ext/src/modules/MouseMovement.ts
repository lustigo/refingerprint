import MouseListener from "./MouseListener";
import MouseData from "../interfaces/MouseData";

/**
 * Module which tracks the Mouse Movements
 */
export default abstract class MouseMovement extends MouseListener {

    /**
     * MousePath
     */
    protected path: Array<MouseData> = new Array<MouseData>();

    /**
     * If the mouse is currently being tracked
     */
    protected running = false;

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
        return (): void => { document.removeEventListener("mousemove", l); };
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
