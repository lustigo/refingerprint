import { delay } from "../misc/Helper";
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
     * Frame of the Captcha
     */
    protected captchaFrame: HTMLIFrameElement | null = null;

    /**
     * Name of the Module
     */
    public name = "Mouse";

    /**
     * will be called when the Captcha is rendered
     */
    public start(): void {
        this.getCaptchaFrame();
    }


    /**
     * will be called when the Captcha is solved
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
     * Stores the Captcha Frame
     */
    private getCaptchaFrame() {
        const gotcha = document.getElementById("gotcha")
        if (gotcha) {
            const childs = gotcha.getElementsByTagName("iframe");
            if (childs.length > 0) {
                this.captchaFrame = childs[0];
            }
        }
        if (!this.captchaFrame) {
            delay(10).then(() => this.getCaptchaFrame());
        }
    }

    /**
     * Registers the Listener on the Recaptcha Iframe
     * @param listener Method to call, when the Iframe is clicked
     */
    protected registerIframeClickListener(listener: (ev: Event) => any): void {
        this.registerIframeListener("mousedown", listener);
    }

    /**
     * Registers the Listener on the Recaptcha Iframe
     * @param listener Method to call, when the Mouse is moved in the Iframe
     */
    protected registerIframeMoveListener(listener: (ev: Event) => any): void {
        this.registerIframeListener("mousemove", listener);
    }

    /**
    * Registers the Listener on the Recaptcha Iframe
    * @param event Name of the Event to listen to
    * @param listener Method to call, when the event occurs
    */
    private registerIframeListener(event: string, listener: (ev: Event) => any): void {
        if (this.captchaFrame && this.captchaFrame.contentWindow) {
            this.captchaFrame.contentWindow.addEventListener(event, listener);
        } else {
            delay(10).then(() => this.registerIframeListener(event, listener));
        }
    }

    /**
    * Removes the Listener from the Recaptcha Iframe
    * @param listener Listener to remove
    */
    protected removeIframeClickListener(listener: (ev: Event) => any): void {
        this.removeIframeListener("mousedown", listener);
    }

    /**
    * Removes the Listener from the Recaptcha Iframe
    * @param listener Listener to remove
    */
    protected removeIframeMoveListener(listener: (ev: Event) => any): void {
        this.removeIframeListener("mousemove", listener);
    }

    /**
    * Removes the Listener from the Recaptcha Iframe
    * @param event Name of the Event that is listened to
    * @param listener Listener to remove
    */
    private removeIframeListener(event: string, listener: (ev: Event) => any): void {
        if (this.captchaFrame && this.captchaFrame.contentWindow) {
            this.captchaFrame.contentWindow.removeEventListener(event, listener);
        }
    }

    /**
     * Registers the Listener to the Document
     * @param listener: Listener to call when the mouse is moved
     */
    protected registerMouseListener(listener: (ev: Event) => any): void {
        document.addEventListener("mousemove", listener);
    }

    /**
    * Removes the Listener from the Document
    * @param listener: Listener to remove
    */
    protected removeMouseListener(listener: (ev: Event) => any): void {
        document.removeEventListener("mousemove", listener);
    }

    /**
     * Builds a listener, so that onMouseMove is called and pointer to the Class is being kept
     * @param mm MouseMovement Class
     * @returns  Listener, that calls onMouseMove
     */
    protected getOnMouseMove(mm: MouseMovement): (ev: Event) => any {
        return (ev) => {
            mm.onMouseMove(ev);
        };
    }

    /**
     * Adds a MousePoint from the given MouseEvent to the MousePath
     * @param ev MouseMoveEvent
     */
    private onMouseMove(ev: Event): void {
        const e = ev as MouseEvent;
        this.path.push({
            x: e.screenX,
            y: e.screenY,
            time: new Date().getTime(),
        });
    }

}
