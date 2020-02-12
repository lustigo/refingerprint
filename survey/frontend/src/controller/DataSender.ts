import CheckboxDescription from "../interfaces/CheckboxDescription";
import DropdownDescription from "../interfaces/DropdownDescription";
import MatrixDescription from "../interfaces/MatrixDescription";
import SurveyData from "../interfaces/SurveyData";
import SurveyDescription from "../interfaces/SurveyDescription";
import TextboxDescription from "../interfaces/TextboxDescription";
import TextFieldDescription from "../interfaces/TextFieldDescription";
import { UUID } from "../interfaces/UUID";
import VectorDescription from "../interfaces/VectorDescription";
import WidgetData from "../interfaces/WidgetData";
import WidgetDescription from "../interfaces/WidgetDescription";
import { WidgetType } from "../enums/WidgetType";


/**
 * Class that handles the sending of the Data to the server.
 */
export default class DataSender {

    /**
     * Data Object to send to the server
     */
    private data: SurveyData;

    /**
     * Constructs a new Datasender with the given Data
     * @param data All Widgets with their respective data
     * @param userId Unique ID of the User
     */
    public constructor(data: SurveyDescription, userId: UUID) {
        this.data = {
            survey: data.survey,
            user: userId,
            pages: new Array<WidgetData[]>(data.pages.length),
        };
        this.convertData(data);
    }

    /**
     * Converts the data enclosed in the widget descriptions to
     * the Widgetdata-shape and saves it in data.pages
     * @param data All Widgets with their respective data
     */
    private convertData(data: SurveyDescription) {
        for (let i = 0; i < data.pages.length; i++) {
            this.data.pages[i] = new Array<WidgetData>();
            for (const widget of data.pages[i]) {
                this.data.pages[i].push(this.getDataFromWidget(widget));
            }
        }
    }

    /**
     * Extracts the data from a WidgetDescription
     * @param widget Widget to extract the data from
     * @return extracted Data
     */
    private getDataFromWidget(widget: WidgetDescription): WidgetData {
        let value = {};

        switch (widget.type) {
        case WidgetType.CHECKBOX: {
            const choosen = new Array<string>();
            const checkbox = widget.structure as CheckboxDescription;
            let bitMap = 0;
            for (let i = 0; i < checkbox.checked.length; i++) {
                if (checkbox.checked[i]) {
                    choosen.push(checkbox.selection[i]);
                    bitMap += Math.pow(2, i);
                }
            }
            value = { choosen, bitMap };
            break;
        }
        case WidgetType.DROPDOWN: {
            const dropdown = widget.structure as DropdownDescription;
            value = { selected: dropdown.selected };
            break;
        }
        case WidgetType.MATRIX: {
            const matrix = widget.structure as MatrixDescription;
            value = { answers: matrix.answers };
            break;
        }
        case WidgetType.TEXTBOX: {
            const textbox = widget.structure as TextboxDescription;
            value = { answer: textbox.answer };
            break;
        }
        case WidgetType.TEXTFIELD: {
            const textfield = widget.structure as TextFieldDescription;
            value = { answer: textfield.content };
            break;
        }
        case WidgetType.VECTOR: {
            const vector = widget.structure as VectorDescription;
            value = { answer: vector.answer };
            break;
        }
        }

        return {
            type: widget.type,
            value
        };
    }

    /**
     * Sends the prepared data
     * @return if it was successful
     */
    public async send(): Promise<boolean> {
        try {
            const resp = await fetch(`/api/${this.data.survey}`, {
                method: "POST",
                body: JSON.stringify(this.data),
                headers: { "content-type": "application/json" }
            });
            return resp.status == 200;
        } catch{
            return false;
        }
    }

    // eslint-disable-next-line semi
}