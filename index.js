#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const list_1 = require("./commands/list");
const add_1 = require("./commands/add");
const remove_1 = require("./commands/remove");
let yargs = require('yargs');
yargs.version('0.0.1')
    .usage('$0 <cmd> [args]')
    .command('list', 'List all entries', () => { }, list_1.default.execute)
    .command('add <name> <uri>', 'Add a new entry', (yargs) => {
    yargs.positional('name', {
        type: 'string',
        default: null,
        describe: 'Name'
    });
    yargs.positional('uri', {
        type: 'string',
        default: null,
        describe: 'Connection URI'
    });
}, add_1.default.execute)
    .command('remove <name>', 'Remove a entry', () => { }, remove_1.default.execute)
    .demandCommand(1, 'You need at least one command before moving on')
    .help()
    .argv;
