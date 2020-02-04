import { UUID } from "./UUID";
import WidgetDescription from "./WidgetDescription";

/**
 * Describes the Structure of the Survey
 */
export default interface SurveyDescription {
  readonly user: UUID;
  readonly id: UUID;
  readonly pages: WidgetDescription[][];
}
