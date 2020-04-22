import { CellEvent, TaskData, TaskType } from "../interfaces/TaskData";
import { delay } from "../misc/Helper";

/**
 * Task represents a single Captcha Task
 */
export default class Task {

    /**
     * Name of the Candidate or "" if there is none
     */
    private candidate = "";

    /**
     * The payload of this task
     */
    private images: Array<string> = [];

    /**
     * Reference to TaskFrame
     */
    private frame: HTMLIFrameElement;

    /**
     * Reference to the Image Matrix
     */
    private table: HTMLTableElement | null = null;

    /**
     * Term to search for
     */
    private term = "";

    /**
     * Reference to Timer
     */
    private timer: number | null = null;

    /**
     * Type of the Task
     */
    public type: TaskType = TaskType.NO;

    /**
     * Wheter the Task failed
     */
    public failed = false;

    /**
     * Wheter this task was solved or reloaded/timed out
     */
    public solved = true;

    /**
     * Click Events on the Task
     */
    private events: Array<CellEvent> = [];

    /**
     * Which cells are currently selected?
     * Null for Dyn. Sim. Image
     */
    private selected: Array<Array<boolean>> | null = null;

    /**
     * Constructs the Task and starts the Analyzing
     * @param taskFrame Reference to the TaskFrame
     */
    public constructor(taskFrame: HTMLIFrameElement) {
        this.frame = taskFrame;
        this.getTerm();
        this.getCandidate();
        this.getTable()
            .then(() => delay(100))
            .then(this.getType.bind(this))
            .then(this.getImage.bind(this))
            .then(this.buildMatrix.bind(this))
            .then(this.startTimer.bind(this));
    }

    /**
     * Stops the Timer
     */
    public stop(): void {
        if(this.timer){
            window.clearInterval(this.timer);
        }
    }

    /**
     * gets the TaskData for this Task
     * @returns the Task Data for this Task
     */
    public getData(): TaskData {
        return {
            candidate: this.candidate,
            type: this.type,
            term: this.term,
            failed: this.failed,
            solved: this.solved,
            events: this.events,
            selected: this.type == TaskType.DYN ? null : this.selected,
            images: this.images,
        };
    }

    /**
     * Builds the selected Matrix
     */
    private async buildMatrix(): Promise<void>{
        let size = 3;
        if(this.type == TaskType.OBJ){
            size = 4;
        }

        this.selected = [];
        for (let i = 0; i < size; i++) {
            const intern: Array<boolean> = [];
            for(let j = 0; j < size; j++){
                intern.push(false);
            }
            this.selected.push(intern);
        }
    }

    /**
     * Saves the initial payload of the Task
     */
    private async getImage(): Promise<void> {
        if(this.frame.contentDocument){
            let imgEle;
            if(this.type == TaskType.SIM){
                imgEle = this.frame.contentDocument.getElementsByClassName("rc-image-tile-33");
            } else if(this.type == TaskType.OBJ){
                imgEle = this.frame.contentDocument.getElementsByClassName("rc-image-tile-44");
            }
            if (imgEle && imgEle.length > 0) {
                this.images.push(Task.toDataURL(imgEle[0] as HTMLImageElement));
            } else {
                await delay(100);
                this.getImage();
                return;
            }
        }
    }

    /**
     * Gets the Task Term
     */
    private async getTerm(): Promise<void> {
        if (this.frame.contentDocument) {
            let desc = this.frame.contentDocument.getElementsByClassName("rc-imageselect-desc-wrapper");
            while (desc.length == 0) {
                await delay(50);
                desc = this.frame.contentDocument.getElementsByClassName("rc-imageselect-desc-wrapper");
            }
            const termDiv = (desc[0] as HTMLDivElement).getElementsByTagName("strong");
            if (termDiv.length > 0 && termDiv[0].textContent) {
                this.term = termDiv[0].textContent;
            }
        }
    }

    /**
     * Checks if the Description has a Candidate and stores the name
     */
    private async getCandidate(): Promise<void> {
        if (this.frame.contentDocument) {
            let desc = this.frame.contentDocument.getElementsByClassName("rc-imageselect-desc-wrapper");
            while (desc.length == 0) {
                await delay(50);
                desc = this.frame.contentDocument.getElementsByClassName("rc-imageselect-desc-wrapper");
            }
            const candidatesDiv = desc[0].getElementsByClassName("rc-imageselect-candidates");
            if (candidatesDiv.length > 0) {
                const candidate = (candidatesDiv[0] as HTMLDivElement).children;
                if (candidate.length > 0 && candidate[0].tagName == "DIV") {
                    // remove 'rc-canonical-'
                    this.candidate = candidate[0].className.substr(13);
                }
            } 
        }
    }

    /**
     * Stores the Reference to the Table
     */
    private async getTable(): Promise<void> {
        if (this.frame.contentDocument){
            let tableDiv = this.frame.contentDocument.getElementsByClassName("rc-imageselect-target");
            while (tableDiv.length == 0) {
                await delay(10);
                tableDiv = this.frame.contentDocument.getElementsByClassName("rc-imageselect-target");
            }

            const table = tableDiv[0].getElementsByTagName("TABLE");
            if(table.length > 0){
                this.table = table[0] as HTMLTableElement;
            }
        }
    }

    /**
     * Determins the Type of the Task
     */
    private async getType(): Promise<void>{
        if(this.table){
            const splittedClass = this.table.className.split("-");
            if(splittedClass[splittedClass.length -1] == "33"){
            // 3x3 Matrix
                this.type = TaskType.SIM;
            } else {
                for(const splitted of splittedClass){
                    if(splitted.match(/44/)){
                        // 4x4 Matrix
                        this.type = TaskType.OBJ;
                        return;
                    }
                }
            }
        }
    }

    /**
     * Will be called for each Timer Tick
     * Checks if a new cell was selected
     */
    private onTick(): void {
        if(this.table && this.selected){
            const rows = this.table.getElementsByTagName("TR");
            for(let i = 0; i< rows.length; i++){
                const row = rows[i];
                const cells = row.getElementsByTagName("TD");
                for (let j = 0; j < cells.length; j++){
                    const cell = cells[j] as HTMLTableDataCellElement;
                    if(cell.className.match(/rc-imageselect-dynamic-selected/)){
                        this.type = TaskType.DYN;
                        if(!this.selected[i][j]){
                            this.events.push({
                                id: i*3+j,
                                selected: true,
                                time: new Date().getTime(),
                            });
                            this.selected[i][j] = true;
                        }
                    } else  if (this.type != TaskType.DYN){
                        const marked = this.selected[i][j];
                        const sel = cell.className.match(/rc-imageselect-tileselected/);

                        if (!(marked ||sel) || sel && marked) {
                            // nothing new
                            continue;
                        } else {
                            this.selected[i][j] = sel != null;
                            this.events.push({
                                id: i * this.type==TaskType.OBJ? 4: 3 + j,
                                selected: true,
                                time: new Date().getTime(),
                            });
                        }
                    } else {
                        if(this.type == TaskType.DYN && this.selected[i][j]){
                            const imgEle = cell.getElementsByClassName("rc-image-tile-11");
                            if(imgEle && imgEle.length > 0 ){
                                this.images.push(Task.toDataURL(imgEle[0] as HTMLImageElement));
                            }
                        }
                        this.selected[i][j] = false;
                    }
                }
            }
        }
    }

    /**
     * Starts the Timer
     */
    private async startTimer(): Promise<void> {
        this.timer = window.setInterval(this.onTick.bind(this),100);
    }

    /**
     * Converts the Image to a Data URI
     * @param img Image Element to convert
     * @return Data URI of that image
     * @src https://www.tutorialspoint.com/convert-image-to-data-uri-with-javascript
     */
    private static toDataURL(img: HTMLImageElement): string {
        const canvas = document.createElement("canvas");
        const ctx = canvas.getContext("2d");
        canvas.width = img.width;
        canvas.height = img.height;
        if(ctx){
            ctx.drawImage(img, 0, 0);
        }
        return canvas.toDataURL("image/jpeg");
    }
}

