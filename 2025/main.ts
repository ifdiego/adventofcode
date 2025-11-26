function main(day: number) {
    switch (day) {
        default:
            console.log(`day ${day} not yet implemented`)
            break;
    }
}

const day = parseInt(process.argv[2], 10);
isNaN(day)
    ? console.error("usage: node main.ts <day>")
    : main(day);
