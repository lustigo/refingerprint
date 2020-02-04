import { UUID } from "./UUID";
import WidgetData from "./WidgetData";
/**
 * Describes the structure of the filled-in survey, which will be sent to the server.
 */
export default interface SurveyData {
  readonly user: UUID;
  readonly id: UUID;
  pages: WidgetData[][];
}
