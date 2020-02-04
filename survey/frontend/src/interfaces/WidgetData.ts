import { WidgetType } from "../enums/WidgetType";
/**
 * Describes the Data of a filled-in widget.
 */
export default interface WidgetData {
  readonly type: WidgetType;
  value: Object;
}
