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
     * Frame of the ReCaptcha Task
     */
    protected taskFrame: HTMLIFrameElement | null = null;

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
     */
    public start(): void {
        this.getCaptchaFrame();
        this.getTaskFrame();
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
     * @returns Captcha Frame or null if it is not yet rendered
     */
    private getCaptchaFrame(): HTMLIFrameElement | null {
        if (this.captchaFrame) {
            return this.captchaFrame;
        }

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
        return this.captchaFrame;
    }

    /**
     * Stores the Task Frame
     * @returns Task Frame or null if it is not yet rendered
     */
    private getTaskFrame(): HTMLIFrameElement | null {
        if (this.taskFrame) {
            return this.taskFrame;
        }

        const iframes = document.getElementsByTagName("iframe");
        for (const iframe of iframes) {
            if (iframe.title.match(/reCAPTCHA/)) {
                this.taskFrame = iframe;
                return this.taskFrame;
            }
        }
        if (!this.taskFrame) {
            delay(10).then(() => this.getTaskFrame());
        }

        return this.taskFrame;
    }

    /**
     * Registers the Listener on the Recaptcha Iframe
     * @param listener Method to call, when the Iframe is clicked where this is binded to the MouseMovement Class
     * @returns Method to remove the Listener
     */
    protected registerCaptchaFrameClickListener(listener: (ev: Event) => any): () => void {
        return this.registerIframeListener(this.getCaptchaFrame.bind(this), "mousedown", listener.bind(this));
    }

    /**
     * Registers the Listener on the Recaptcha Iframe
     * @param listener Method to call, when the Mouse is moved in the Iframe where this is binded to the MouseMovement Class
     * @returns Method to remove the Listener
     */
    protected registerCaptchaFrameMoveListener(listener: (ev: Event) => any): () => void {
        return this.registerIframeListener(this.getCaptchaFrame.bind(this), "mousemove", listener.bind(this));
    }

    /**
     * Registers the Listener on the Task Iframe
     * @param listener Method to call, when the Iframe is clicked where this is binded to the MouseMovement Class
     * @returns Method to remove the Listener
     */
    protected registerTaskFrameClickListener(listener: (ev: Event) => any): () => void {
        return this.registerIframeListener(this.getTaskFrame.bind(this), "mousedown", listener.bind(this));
    }

    /**
     * Registers the Listener on the Task Iframe
     * @param listener Method to call, when the Mouse is moved in the Iframe where this is binded to the MouseMovement Class
     * @returns Method to remove the Listener
     */
    protected registerTaskFrameMoveListener(listener: (ev: Event) => any): () => void {
        return this.registerIframeListener(this.getTaskFrame.bind(this), "mousemove", listener.bind(this));
    }

    /**
    * Registers the Listener on the Recaptcha Iframe
    * @param getFrame Method that returns the Iframe or null if it is not yet rendered, if so, the registration will automatically retried after 10ms.
    * @param event Name of the Event to listen to
    * @param listener Method to call, when the event occurs
    * @returns Method to remove the Listener
    */
    private registerIframeListener(getFrame: () => HTMLIFrameElement | null, event: string, listener: (ev: Event) => any): () => void {
        const frame = getFrame();
        if (frame && frame.contentWindow) {
            frame.contentWindow.addEventListener(event, listener);
        } else {
            delay(10).then(() => this.registerIframeListener(getFrame, event, listener));
        }
        return () => { if (frame && frame.contentWindow) { frame.contentWindow.removeEventListener(event, listener) } };
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
