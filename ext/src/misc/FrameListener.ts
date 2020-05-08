import { delay } from "./Helper";

/**
 * Helper which monitors the ReCaptcha Frames
 */
class FrameListener {
    /**
     * Frame of the Captcha
     */
    private captchaFrame: HTMLIFrameElement | null = null;

    /**
     * Frame of the ReCaptcha Task
     */
    private taskFrame: HTMLIFrameElement | null = null;

    /**
     * Stores the Captcha Frame
     * @returns Captcha Frame or null if it is not yet rendered
     */
    public getCaptchaFrame(): HTMLIFrameElement | null {
        if (this.captchaFrame) {
            return this.captchaFrame;
        }

        const gotcha = document.getElementById("gotcha");
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
    public getTaskFrame(): HTMLIFrameElement | null {
        if (this.taskFrame) {
            return this.taskFrame;
        }

        const iframes = document.getElementsByTagName("iframe");
        for (const iframe of iframes) {
            if (iframe.title.match(/recaptcha challenge/)) {
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
     * @param listener Method to call, when the Iframe is clicked, this is binded to t
     * @param t Instance
     * @returns Method to remove the Listener
     */
    public registerCaptchaFrameClickListener(listener: (ev: Event) => any, t: any): () => void {
        return this.registerIframeListener(this.getCaptchaFrame.bind(this), "mouseup", listener.bind(t));
    }

    /**
     * Registers the Listener on the Recaptcha Iframe
     * @param listener Method to call, when the Mouse is moved in the Iframe, this is binded to t
     * @param t Instance
     * @returns Method to remove the Listener
     */
    public registerCaptchaFrameMoveListener(listener: (ev: Event) => any, t: any): () => void {
        return this.registerIframeListener(this.getCaptchaFrame.bind(this), "mousemove", listener.bind(t));
    }

    /**
     * Registers the Listener on the Task Iframe
     * @param listener Method to call, when the Iframe is clicked, this is binded to t
     * @param t Instance
     * @returns Method to remove the Listener
     */
    public registerTaskFrameClickListener(listener: (ev: Event) => any, t: any): () => void {
        return this.registerIframeListener(this.getTaskFrame.bind(this), "mouseup", listener.bind(t));
    }

    /**
     * Registers the Listener on the Task Iframe
     * @param listener Method to call, when the Mouse is moved in the Iframe, this is binded to t
     * @param t Instance
     * @returns Method to remove the Listener
     */
    public registerTaskFrameMoveListener(listener: (ev: Event) => any, t: any): () => void {
        return this.registerIframeListener(this.getTaskFrame.bind(this), "mousemove", listener.bind(t));
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
        return (): void => { if (frame && frame.contentWindow) { frame.contentWindow.removeEventListener(event, listener); } };
    }

}

// frameListener gives access to the Captcha Frames
const frameListener = new FrameListener();

export default frameListener;
