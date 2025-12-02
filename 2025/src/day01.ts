import fs from "fs";
import path from "path";
import readline from "readline";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

function cycle(operation, a, b) {
    const mod = 100;
    let result;
    if (operation === "R") result = a + b;
    else if (operation === "L") result = a - b;
    return ((result % mod) + mod) % mod;
}

export async function day01() {
    const filepath = path.join(__dirname, "../inputs/day01.txt")
    const stream = fs.createReadStream(filepath);
    const rl = readline.createInterface({
        input: stream,
        crlfDelay: Infinity
    });

    const numbers = Array.from({ length: 100 }, (_, i) => i);
    let dial = 50;
    let points = 0;
    for await (const line of rl) {
        const direction = line.substr(0, 1);
        const value = Number(line.substr(1));
        dial = cycle(direction, dial, value);
        if (dial === 0) points++;
    }
    console.log("part one: ", points);
}
