import frameListener from "../misc/FrameListener";
import MouseMovement from "./MouseMovement";

/**
 * Module which tracks the Mouse Movements until the Checkbox is clicked
 */
export default class MouseMovementCheckbox extends MouseMovement {

    /**
     * Name of the Module
     */
    public readonly name = "MouseCheckbox";

    /**
     * will be called when the Captcha is rendered
     * Starts tracking the MousePath and registers the listener for the Iframe Click.
     */
    public start(): void {
        super.start();
        this.removeHandler.push(frameListener.registerCaptchaFrameClickListener(this.stop, this));
        this.removeHandler.push(frameListener.registerCaptchaFrameMoveListener(this.onMouseMove, this));
        this.removeHandler.push(this.registerMouseListener(this.onMouseMove));
        this.running = true;
    }

    /**
     * Stops the Tracking
     */
    public stop(): void {
        if (this.running) {
            this.stopTracking();
        }
    }
}
