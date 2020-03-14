import MouseMovement from "./MouseMovement";
import frameListener from "../misc/FrameListener";

/**
 * Module which tracks the Mouse Movements after the Checkbox is clicked
 */
export default class MouseMovementRest extends MouseMovement {

    /**
     * Name of the Module
     */
    public readonly name = "MouseRest";

    /**
     * will be called when the Captcha is rendered
     * Registers the Listener for the Iframe Click.
     */
    public start(): void {
        super.start();
        frameListener.registerCaptchaFrameClickListener(this.startTracking, this);
    }

    /**
     * will be called when the Captcha is solved
     * Stops tracking the MousePath.
     */
    public stop(): void {
        this.stopTracking();
        this.running = false;
    }

    /**
     * Starts the Tracking
     */
    private startTracking(): void {
        if (!this.running) {
            this.running = true;
            this.stopTracking();
            this.removeHandler.push(this.registerMouseListener(this.onMouseMove));
            this.removeHandler.push(frameListener.registerCaptchaFrameMoveListener(this.onMouseMove, this));
            this.removeHandler.push(frameListener.registerTaskFrameMoveListener(this.onMouseMove, this));
        }
    }

}
