import Koa from "koa";
import mount from "koa-mount";
import bodyparser from "koa-bodyparser";
import fs from "fs";
import { promisify } from "util";
import req from "request";

/**
 * Regex for a UUIDv4
 * @src https://www.regextester.com/99148
 */
const uuidRegex = new RegExp("([0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12})");

/**
 * Async fs.exists
 */
const exists = promisify(fs.exists);
/**
 * Async fs.stat
 */
const stat = promisify(fs.stat);

/**
 * Async fs.writeFile
 */
const writeFile = promisify(fs.writeFile);

/**
 * Async request
 */
const request = promisify(req);

/**
 * Includes the API Routes
 */
const app = new Koa();

/**
 * Extracts an UUID from a given string.
 * @param testee: String to extract UUID from.
 * @return the UUID
 * @throws if there is no valid UUID
 */
function extractUUID(testee: string): string {
    if (uuidRegex.test(testee)) {
        const regexResult = uuidRegex.exec(testee);
        if (regexResult && regexResult.length > 1) {
            return regexResult[1];
        }
    }
    throw ("No valid surveyId found");
}

/**
 * GET the Survey-Structure
 * @param surveyPath Path where the structure is located
 * @param ctx Request Context
 */
async function getSurveyStructure(surveyPath: string, ctx: Koa.Context): Promise<void> {
    const structurePath = surveyPath + "/structure.json";
    ctx.set("Content-Type", "application/json");
    ctx.set("Accept-Ranges", "bytes");
    ctx.set("Content-Length", (await stat(structurePath)).size.toString());
    ctx.body = fs.createReadStream(structurePath);
}

/**
 * POST save the filled-in data
 * @param surveyPath Path where the data should be stored
 * @param surveyId Id of the survey
 * @param ctx Request Context
 */
async function saveSurveyData(surveyPath: string, surveyId: string, ctx: Koa.Context): Promise<void> {
    const data = ctx.request.body;
    if (!data || !data.user || !data.survey || !data.pages || data.survey != surveyId) {
        ctx.body = "Invalid request";
        ctx.status = 400;
        return;
    }

    try {
        const userId = extractUUID(data.user);
        const dataPath = surveyPath + `/${userId}.json`;
        // Check if Data was already saved
        if (await exists(dataPath)) {
            ctx.body = "Already sent.";
            ctx.status = 409;
            return;
        }

        // Add Timestamp & write to File
        data.savedTime = Date.now();
        await writeFile(dataPath, JSON.stringify(data)).catch((err) => {
            console.log("Error on writing " + dataPath);
            console.log(err);
            ctx.status = 500;
            ctx.body = "Internal Server Error";
            return;
        });

        // Respond with success
        ctx.status = 200;
        ctx.body = "Success";
    } catch {
        ctx.body = "Invalid request";
        ctx.status = 400;
        return;
    }
}

/**
 * Validate the Captcha Response and sends the result back
 * @param captchaResponse Response of the Captcha
 * @param ctx Request-Context
 */
async function validateCaptcha(captchaResponse: string, ctx: Koa.Context): Promise<void>{
    const formData = {
        secret: process.env.RECAPTCHA_SECRET as string,
        response: captchaResponse
    };
    const response = await request({
        method: "POST",
        url: "https://www.google.com/recaptcha/api/siteverify",
        headers: {
            "Content-Type": "multipart/form-data"
        },
        formData
    });
    ctx.status = JSON.parse(response.body).success ? 200: 400;
}


/**
 * Tries to extract the surveyId from the given path
 * @param path: Path to extract surveyId from
 * @return the extracted surveyId
 * @throws if there is no valid UUID
 */
function extractSurveyId(path: string): string {
    return extractUUID(path.replace("/", ""));
}

/**
 * Checks wheter the survey exists
 * @param surveyPath Path to check
 */
async function checkExistense(surveyPath: string): Promise<boolean> {
    if (await exists(surveyPath)) {
        return await exists(surveyPath + "/structure.json");
    }
    return false;
}

/**
 * Register the Bodyparser
 */
app.use(bodyparser({
    enableTypes: ["json"],
    strict: true,
    onerror: (err, ctx) => {
        console.log(err);
        ctx.throw("Not a valid JSON", 400);
    }
}));

/**
 * Register the captcha endpoint
 */
app.use(mount("/captcha", async(ctx: Koa.Context) => {
    if (ctx.method != "POST"){
        ctx.status = 405;
        ctx.body = "Method not allowed";
        return;
    }

    const data = ctx.request.body;
    if (!data || !data.response) {
        ctx.body = "Invalid request";
        ctx.status = 400;
        return;
    }

    await validateCaptcha(data.response,ctx);
}));

/**
 * Register the survey endpoints
 */
app.use(mount("/", async (ctx: Koa.Context) => {
    // Check if the Method is GET or POST
    if (ctx.method != "GET" && ctx.method != "POST") {
        ctx.status = 405;
        ctx.body = "Method not allowed";
        return;
    }

    // Get the surveyId
    try {
        const surveyId = extractSurveyId(ctx.path);
        const surveyPath = (process.env.DATA_PATH as string) + surveyId;
        if (!await checkExistense(surveyPath)) {
            ctx.body = "This is not a valid SurveyId";
            ctx.status = 400;
            return;
        }

        // Pass on to the corresponding method
        if (ctx.method == "GET") {
            await getSurveyStructure(surveyPath, ctx);
        } else {
            await saveSurveyData(surveyPath, surveyId, ctx);
        }
    } catch (e) {
        console.log(e);
        ctx.body = "This is not a valid Survey-Id";
        ctx.status = 400;
        return;
    }

}));


export default app;
