#!/usr/bin/env node

import List from "./commands/list";
import AddShortcut from "./commands/add";
import RemoveShortcut from "./commands/remove";

let yargs = require('yargs');

yargs.version('0.0.1')
    .usage('$0 <cmd> [args]')
    .command('list', 'List all entries', () => {}, List.execute)
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
    }, AddShortcut.execute)
    .command('remove <name>', 'Remove a entry', () => {}, RemoveShortcut.execute)
    .demandCommand(1, 'You need at least one command before moving on')
    .help()
    .argv;