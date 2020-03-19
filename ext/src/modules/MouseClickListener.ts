import { delay } from "../misc/Helper";
import { MouseKey, ClickData } from "../interfaces/ClickData";
import frameListener from "../misc/FrameListener";
import MouseListener from "./MouseListener";

/**
 * Module which records the Mouseclicks
 */
export default class MouseClickListener extends MouseListener {

    /**
     * Click Events
     */
    private events: Array<ClickData> = [];

    /**
     * Name of the Module
     */
    public name = "MouseClicks";

    /**
     * will be called when the Captcha is rendered
     * registers the Listeners
     */
    public start(): void {
        const press = this.getClickHandler(false).bind(this);
        const release = this.getClickHandler(true).bind(this);

        this.removeHandler.push(this.registerClickListener("mousedown", press));
        this.removeHandler.push(this.registerClickListener("mouseup", release));

        this.registerOnFrames(press, release);
    }

    /**
    * Registers the Listener to the Document
    * @param listener: Listener to call
    * @returns Method to remove the Listener
    */
    private registerClickListener(event: "mousedown" | "mouseup", listener: (ev: MouseEvent) => any): () => void {
        document.addEventListener(event, listener);
        return (): void => { document.removeEventListener(event, listener); };
    }

    /**
     * Registers the Listeners on the Frames
     * @param press Listener on MouseDown
     * @param release Listener on MouseUp
     */
    private async registerOnFrames(press: (ev: MouseEvent) => any, release: (ev: MouseEvent) => any): Promise<void> {
        const t = true;

        while (t) {
            await delay(200);
            const captchaFrame = frameListener.getCaptchaFrame();
            if (captchaFrame && captchaFrame.contentDocument && captchaFrame.contentDocument.body) {
                captchaFrame.contentDocument.addEventListener("mousedown", press);
                captchaFrame.contentDocument.addEventListener("mouseup", release);
                this.removeHandler.push((): void => { if (captchaFrame && captchaFrame.contentDocument) { captchaFrame.contentDocument.removeEventListener("mousedown", press); } });
                this.removeHandler.push((): void => { if (captchaFrame && captchaFrame.contentDocument) { captchaFrame.contentDocument.removeEventListener("mouseup", release); } });
                break;
            }
        }

        while (t) {
            await delay(200);
            const taskFrame = frameListener.getTaskFrame();
            if (taskFrame && taskFrame.contentDocument && taskFrame.contentDocument.body) {
                taskFrame.contentDocument.addEventListener("mousedown", press);
                taskFrame.contentDocument.addEventListener("mouseup", release);
                this.removeHandler.push((): void => { if (taskFrame && taskFrame.contentDocument) { taskFrame.contentDocument.removeEventListener("mousedown", press); } });
                this.removeHandler.push((): void => { if (taskFrame && taskFrame.contentDocument) { taskFrame.contentDocument.removeEventListener("mouseup", release); } });
                break;
            }
        }
    }

    /**
     * will be called when the Captcha is solved
     * stops the tracking
     */
    public stop(): void {
        super.stopTracking();
    }

    /**
     * Returns a onClick Handler
     * @param release If the Button is released
     * @return onClick Handler
     */
    private getClickHandler(release: boolean): (ev: MouseEvent) => void {
        return (ev: MouseEvent): void => {
            this.onMouseClick(ev, release);
        };
    }

    /**
     * will be called when the mouse is clicked
     * stores the Event
     * @param ev Mouse Click Event
     * @param release If the Button is released
     */
    private onMouseClick(ev: MouseEvent, release: boolean): void {
        let key = MouseKey.LEFT;
        switch (ev.button) {
        case 1: {
            key = MouseKey.MIDDLE;
            break;
        }
        case 2: {
            key = MouseKey.RIGHT;
            break;
        }
        }

        this.events.push({
            x: ev.screenX,
            y: ev.screenY,
            time: new Date().getTime(),
            release: release,
            key: key,
        });
    }

    /**
     * returns the collected MouseClicks
     * @returns MouseClick Data
     */
    public getCollectedData(): Array<ClickData> {
        return this.events;
    }
    // eslint-disable-next-line semi
}