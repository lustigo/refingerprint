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
     * If the mouse is currently being tracked
     */
    private running = false;

    /**
     * will be called when the Captcha is rendered
     * Starts tracking the MousePath and registers the listener for the Iframe Click.
     */
    public start(): void {
        super.start();
        this.registerIframeClickListener(this.stop);
        this.registerIframeMoveListener(this.getOnMouseMove(this));
        this.registerMouseListener(this.getOnMouseMove(this));
        this.running = true;
    }

    /**
     * Stops the Tracking
     */
    public stop(): void {
        if (this.running) {
            this.removeMouseListener(this.getOnMouseMove(this));
            this.removeIframeClickListener(this.stop);
            this.removeIframeMoveListener(this.getOnMouseMove(this));
        }
    }
}
