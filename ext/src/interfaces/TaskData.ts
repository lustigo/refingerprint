/**
 * Possible Task Types
 */
enum TaskType {
    NO = 0,
    SIM = 2,
    DYN = 3,
    OBJ = 4,
}

/**
 * Click Event on a Cell
 */
interface CellEvent {
    id: number;
    time: number;
    selected: boolean; //If the cell was selected or unselected
}

/**
 * Data which will be collected per task
 */
interface TaskData {
    candidate: string;
    term: string;
    type: TaskType;
    failed: boolean;
    solved: boolean;
    events: Array<CellEvent>;
    selected: Array<Array<boolean>> | null; // For Type Sim. Image or Obj. Ident. the selected Cells during Verify, for Dyn. Sim. Image this is null
}

export { CellEvent, TaskData, TaskType };
