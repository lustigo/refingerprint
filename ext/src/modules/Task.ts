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
        this.getTable();
        this.getType();
    }

    /**
     * Stops the Timer
     */
    public stop(): void { }

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
            selected: this.type == TaskType.OBJ ? null : this.selected,
        };
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
        while (!(this.frame && this.frame.contentDocument)) {
            await delay(10);
        }
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

    /**
     * Determins the Type of the Task
     */
    private async getType(): Promise<void>{
        while(this.table == null){
            await delay(10);
        }
        const splittedClass = this.table.className.split("-");
        if(splittedClass[splittedClass.length -1] == "33"){
            // 3x3 Matrix
            this.type = TaskType.SIM;
        } else if (splittedClass[splittedClass.length - 1] == "44"){
            // 4x4 Matrix
            this.type = TaskType.OBJ;
        }
    }
}

