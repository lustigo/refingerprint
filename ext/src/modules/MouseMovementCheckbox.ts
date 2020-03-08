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
        this.removeHandler.push(this.registerCaptchaFrameClickListener(this.stop));
        this.removeHandler.push(this.registerCaptchaFrameMoveListener(this.onMouseMove));
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
