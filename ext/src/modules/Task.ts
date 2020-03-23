import { CellEvent, TaskData, TaskType } from "../interfaces/TaskData";

/**
 * Task represents a single Captcha Task
 */
export default class Task {

    /**
     * Term to search for
     */
    private term: string = "";

    /**
     * Reference to Timer
     */
    private timer: number | null = null;

    /**
     * Type of the Task
     */
    private type: TaskType = TaskType.NO;

    /**
     * Wheter the Task failed
     */
    public failed: boolean = false;

    /**
     * Wheter this task was solved or reloaded/timed out
     */
    public solved: boolean = true;

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
     * Constructs the Task and starts the Analyze
     * @param taskFrame Reference to the TaskFrame
     */
    public constructor(taskFrame: HTMLIFrameElement) {
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
            type: this.type,
            term: this.term,
            failed: this.failed,
            solved: this.solved,
            events: this.events,
            selected: this.type == TaskType.OBJ ? null : this.selected,
        };
    }

}
