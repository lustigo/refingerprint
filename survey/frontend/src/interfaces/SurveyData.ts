import { UUID } from "./UUID";
import WidgetData from "./WidgetData";
/**
 * Describes the structure of the filled-in survey, which will be sent to the server.
 */
export default interface SurveyData {
    readonly survey: UUID;
    readonly session: UUID;
    pages: WidgetData[][];
    // eslint-disable-next-line semi
}
